package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minuchi/go-auth-api/controllers"
)

func SetupRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	authV1 := r.Group("/api/auth/v1")
	controllers.LoadAuth(authV1)

	return r
}
