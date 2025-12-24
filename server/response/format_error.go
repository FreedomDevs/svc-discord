package response

import (
	"svc-discord/utils"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(code ErrorResponseCode, details any, c *gin.Context) {
	traceID := c.Request.Header.Get("X-Trace-Id")
	c.Header("X-Trace-Id", traceID)
	c.JSON(code.Status, gin.H{
		"error": gin.H{
			"message": code.Message,
			"code":    code.Code,
			"details": details,
		},
		"meta": gin.H{
			"traceId":   traceID,
			"timestamp": utils.GetCurrentTimestampString(),
		},
	})
}

type ErrorResponseCode struct {
	Code    string
	Message string
	Status  int
}
