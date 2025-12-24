package endpoints

import (
	"svc-discord/config"
	"svc-discord/server/response"
	"svc-discord/server/response/codes"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func GetUsersCountHandler(c *gin.Context) {
	session := c.MustGet("discord").(*discordgo.Session)

	guild, err := session.GuildWithCounts(config.GetGuildID())
	if err != nil {
		response.SendErrorResponse(codes.ErrInternalError(err), nil, c)
		return
	}
	if guild == nil {
		response.SendErrorResponse(codes.ErrGuildIsNull, nil, c)
		return
	}

	response.SendSuccessResponse(codes.SuccessUsersCountOK, gin.H{"count_members": guild.ApproximateMemberCount}, c)
}
