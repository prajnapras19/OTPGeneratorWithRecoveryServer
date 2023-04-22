package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prajnapras19/otpgeneratorwithrecoveryserver/client/mysql"
	"github.com/prajnapras19/otpgeneratorwithrecoveryserver/config"
	"github.com/prajnapras19/otpgeneratorwithrecoveryserver/sharedsecret"
	"net/http"
)

func main() {
	// config
	cfg := config.Get()

	// database
	mySQLService := mysql.NewService(cfg.MySQLConfig)

	// auto migrate
	mySQLService.GetDB().AutoMigrate(&sharedsecret.SharedSecret{})

	// repository
	sharedSecretRepository := sharedsecret.NewRepository(mySQLService.GetDB())

	// handler
	sharedSecretHandler := sharedsecret.NewHandler(sharedSecretRepository)

	r := gin.Default()
	r.POST("/insert", sharedSecretHandler.Insert)
	r.GET("/get/:client_id", sharedSecretHandler.Get)
	r.GET("/_healthcheck", func(gc *gin.Context) {
		gc.String(http.StatusOK, "OK")
	})

	r.Run(":" + cfg.RESTPort)
}
