package tinyblog

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/aiservice"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	"github.com/thoseJanes/tinyblog/internal/pkg/errno"
	"github.com/thoseJanes/tinyblog/internal/pkg/log"
	"github.com/thoseJanes/tinyblog/internal/pkg/middleware"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/controller/ai"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/controller/post"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/controller/user"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/store"
	"github.com/thoseJanes/tinyblog/pkg/auth"
)

func installRouters(g *gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	g.GET("/healthz", func(c *gin.Context){
		log.Infow("headlth function called")
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	pprof.Register(g)

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)
	pc := post.New(store.S)
	ac := ai.New(aiservice.Client)
	
	
	api := g.Group("/api")
	api.POST("/login", uc.Login)
	v1 := api.Group("/v1")
	{
		v1User := v1.Group("/users")
		{
			v1User.POST("", uc.Create)
			v1User.PUT(":name/change-password", uc.ChangePassword)

			v1User.Use(middleware.Authn(), middleware.Authz(authz))
			v1User.GET(":name", uc.Get)
			v1User.GET("", uc.List)
			v1User.DELETE(":name", uc.Delete)
			v1User.PUT(":name", uc.Update)
		}
		v1Post := v1.Group("/posts", middleware.Authn())
		{
			v1Post.POST("", pc.Create)
			v1Post.GET(":postId", pc.Get)
			v1Post.GET("", pc.List)
			v1Post.DELETE(":postId", pc.Delete)
			v1Post.PUT(":postId", pc.Update)

			v1Post.GET("/search", pc.Search)
			v1Post.GET("/aisearch", pc.AiSearch)
		}
		v1Ai := v1.Group("/ai", middleware.Authn())
		{
			v1Ai.GET("/polish-content", ac.PolishContent)
			v1Ai.GET("/generate-title", ac.GenerateTitle)
		}
	}

	return nil
}
