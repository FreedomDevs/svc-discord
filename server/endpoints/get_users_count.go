package endpoints

import (
	"svc-discord/config"
	"svc-discord/server/codes"

	"github.com/FreedomDevs/svcLibs/go/svcLibs"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func GetUsersCountHandler(c *gin.Context) {
	session := c.MustGet("discord").(*discordgo.Session)

	guild, err := session.GuildWithCounts(config.GetGuildID())
	if err != nil {
		svcLibs.SendErrorResponse(svcLibs.ErrInternalError(err), c)
		return
	}
	if guild == nil {
		svcLibs.SendErrorResponse(codes.ErrGuildIsNull, c)
		return
	}

	svcLibs.SendSuccessResponse(codes.SuccessUsersCountOK, gin.H{"count_members": guild.ApproximateMemberCount}, c)
}
