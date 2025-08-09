package tinyblog

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/gogo/protobuf/protoc-gen-gogo/grpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	"github.com/thoseJanes/tinyblog/internal/pkg/log"
	"github.com/thoseJanes/tinyblog/internal/pkg/middleware"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/controller/user"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/store"
	pb "github.com/thoseJanes/tinyblog/pkg/proto/tinyblog/v1"
	"github.com/thoseJanes/tinyblog/pkg/token"
	"github.com/thoseJanes/tinyblog/pkg/version/verflag"
	"google.golang.org/grpc"
)

var cfgFile string



func NewTinyBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tinyblog",
		Short: "A practice project following marmotedu's miniblog",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			verflag.HandleFlag()
			log.Init(logOptions())
			defer log.Sync()

			if err := initStore(); err != nil {
				return err
			}

			if err := initAiClient(); err != nil {
				return err
			}

			return run()
		},
		Args: cobra.NoArgs,
	}

	cobra.OnInitialize(initConfig)
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the configuration to init log and database.")

	verflag.AddToFlagSet(cmd.PersistentFlags())
	return cmd
}


func run() error {
	token.Init(viper.GetString("jwt-secret"), core.XUsernameKey)
	gin.SetMode(viper.GetString("runmode"))

	mws := []gin.HandlerFunc{gin.Recovery(), middleware.NoCache, middleware.CORS, middleware.RequestId}
	g := gin.New()
	g.Use(mws...)//apply to all routers after

	if err := installRouters(g); err != nil {
		return err
	}

	httpSrv := startHttpServer(g)
	httpsSrv := startHttpsServer(g)
	// gRpcSrv := startGRPCServer(g)


	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infow("Shutting down the server...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Errorw("Http server forced to shutdown", "err", err)
		return err
	}
	if err := httpsSrv.Shutdown(ctx); err != nil {
		log.Errorw("Https server forced to shutdown", "err", err)
		return err
	}
	//gRpcSrv.GracefulStop()

	log.Infow("Server exited")

	return nil
}


func startHttpServer(g *gin.Engine) *http.Server {
	srv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	log.Infow("start listening address on http server", "addr", viper.GetString("addr"))
	go func(){
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed){
			log.Fatalw(err.Error())
		}
	}()

	return srv
}

func startHttpsServer(g *gin.Engine) *http.Server {
	srv := &http.Server{Addr: viper.GetString("tls.addr"), Handler: g}

	log.Infow("start listening address on https server", "addr", viper.GetString("tls.addr"))
	cert, key := viper.GetString("tls.cert"), viper.GetString("tls.key")
	go func(){
		if err := srv.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	return srv
}

func startGRPCServer(g *gin.Engine) *grpc.Server {
	lis, err := net.Listen("tcp", viper.GetString("grpc.addr"))
	if err != nil {
		log.Fatalw("failed to listen address on grpc server", "addr", viper.GetString("grpc.addr"), "err", err.Error())
	}

	srv := grpc.NewServer()
	pb.RegisterTinyBlogServer(srv, user.New(store.S, nil))
	log.Infow("start listening address " + viper.GetString("grpc.addr") + " on grpc server")
	go func(){
		if err := srv.Serve(lis); err != nil {
			log.Fatalw(err.Error())
		}
	}()
	return srv
}
