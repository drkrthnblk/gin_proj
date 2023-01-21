package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
	"github.com/spf13/viper"

	// "gin_proj/configs"
	"gin_proj/api/rest"
	"gin_proj/pkg/common/logging"
	"gin_proj/pkg/dbConns/postgresConn"
)

func setupConfig() {
	configName := os.Getenv("HOST_TYPE")
	if configName == "" {
		configName = "local"
	}
	viper.SetConfigName(configName)
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error readin config file ", err)
	}

	fmt.Printf("Running in %s environment", configName)

	// configs.LoadConfigFromParameterStore()

}

func setupLogger() {
	logging.SetFormat(viper.GetString("logging.format"))
	logging.SetFormat(viper.GetString("logging.level"))
}

func init() {
	fmt.Println("loading config...........")
	setupConfig()

	fmt.Println("loading logger...........")
	setupLogger()

	logrus.Info("initializing postgres db connection.........")
	if postgresError := postgresConn.NewPostgresClient().InitDB(); postgresError != nil {
		logrus.Panicf("cannot connect to postgres db: %s", postgresError.Error())
	}
}

func main() {
	logrus.Info("starting server.........")
	server := rest.BuildServer()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to pages",
		})
	})
	server.GET("/health_check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})
	if err := server.Run("0.0.0.0:8080"); err != nil {
		logrus.Fatalf("cannot start the server %s", err.Error())
	}
}
