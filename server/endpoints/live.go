package endpoints

import (
	"svc-discord/server/response"
	"svc-discord/server/response/codes"

	"github.com/gin-gonic/gin"
)

func LiveHandler(c *gin.Context) {
	response.SendSuccessResponse(codes.SuccessLiveOK, gin.H{"alive": true}, c)
}
