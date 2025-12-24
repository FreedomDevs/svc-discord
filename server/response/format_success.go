package response

import (
	"svc-discord/utils"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(code SuccessResponseCode, data any, c *gin.Context) {
	traceID := c.Request.Header.Get("traceId")
	c.Header("traceId", traceID)
	c.JSON(code.Status, gin.H{
		"data":    data,
		"message": code.Message,
		"meta": gin.H{
			"code":      code.Code,
			"traceId":   traceID,
			"timestamp": utils.GetCurrentTimestampString(),
		},
	})
}

type SuccessResponseCode struct {
	Code    string
	Message string
	Status  int
}
