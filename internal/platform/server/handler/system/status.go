package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusHandler returns an HTTP handler to perform status checks.
func StatusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Alive!", "details": "Connectivity ok"})
	}
}
