package server

import (
	"svc-discord/config"
	"svc-discord/server/endpoints"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func Init(session *discordgo.Session) *gin.Engine {
	if config.GetIsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	var r *gin.Engine = gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(func(c *gin.Context) {
		c.Set("discord", session)
		c.Next()
	})

	r.SetTrustedProxies([]string{"127.0.0.1", "172.16.0.0/12"})

	r.GET("/count_members", endpoints.GetUsersCountHandler)
	r.GET("/live", endpoints.LiveHandler)

	return r
}
