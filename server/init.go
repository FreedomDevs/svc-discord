package server

import (
	"svc-discord/config"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	if config.GetIsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	var r *gin.Engine = gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.SetTrustedProxies([]string{"127.0.0.1", "172.16.0.0/12"})

	return r
}
