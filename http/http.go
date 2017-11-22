package http

import (
	"github.com/gin-gonic/gin"
	"github.com/axetroy/go-upload/config"
	"github.com/axetroy/gin-uploader"
	"github.com/itsjamie/gin-cors"
	"time"
)

var (
	Router *gin.Engine
)

func RunServer() (err error) {
	gin.SetMode(config.Config.Mode)
	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	Router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	if upload, download, err := uploader.New(Router, config.Config.Upload); err != nil {
		return err
	} else {
		if err := uploader.Resolve(upload, download); err != nil {
			return err
		}
	}

	return Router.Run(config.Http.Host + ":" + config.Http.Port)
}
