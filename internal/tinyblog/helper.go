package tinyblog

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thoseJanes/tinyblog/internal/pkg/log"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/store"
	"github.com/thoseJanes/tinyblog/pkg/db"
)

const(
	recommandedHomeDir = ".tinyblog"
	defaultConfigName = "tinyblog"
)

func initConfig(){
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}else{
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.SetConfigType("yaml")
		viper.AddConfigPath(filepath.Join(home, recommandedHomeDir))
		viper.AddConfigPath(".")
		viper.AddConfigPath("configs")
		viper.SetConfigName(defaultConfigName)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("TINYBLOG")

	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read in config.", "err", err)
	}

	log.Infow("Using config file", "file", viper.ConfigFileUsed())
}

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller: viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		LogLevel: viper.GetString("log.log-level"),
		Format: viper.GetString("log.format"),
		OutputPaths: viper.GetStringSlice("log.output-paths"),
		ErrOutputPaths: viper.GetStringSlice("log.err-output-paths"),
	}
}

func initStore() error {
	sqlOptions := &db.MySQLOptions{
		MaxIdleConnections: viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections: viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel: viper.GetInt("db.log-level"),
		MySQLDSN: db.MySQLDSN{
			User: viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			Database: viper.GetString("db.database"),
			Host: viper.GetString("db.host"),
		},
	}

	db, err := db.NewMySQL(sqlOptions)
	if err != nil {
		return err
	}

	store.InitDataStore(db)
	return nil
}
