package endpoints

import (
	"svc-discord/config"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func GetUsersCount(c *gin.Context) {
	session := c.MustGet("discord").(*discordgo.Session)

	guild, err := session.GuildWithCounts(config.GetGuildID())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	if guild == nil {
		c.JSON(500, gin.H{"error": "guild_is_nil"})
	}

	c.JSON(200, gin.H{"status": "ok", "count_members": guild.ApproximateMemberCount})
}
