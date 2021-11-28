package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WelcomeHandler returns an HTTP handler to show welcome info.
func WelcomeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//ctx.JSON(http.StatusOK, struct{ data string }{data: })
		ctx.JSON(
			http.StatusOK,
			gin.H{"data": "!~ Go(lang) powered API ~!"},
		)
	}
}
