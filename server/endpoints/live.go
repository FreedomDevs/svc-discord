package endpoints

import (
	"github.com/FreedomDevs/svcLibs/go/svcLibs"
	"github.com/gin-gonic/gin"
)

func LiveHandler(c *gin.Context) {
	svcLibs.SendSuccessResponse(svcLibs.SuccessLiveOK, gin.H{"alive": true}, c)
}
