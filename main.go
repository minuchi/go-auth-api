package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minuchi/go-auth-api/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	authV1 := r.Group("/api/auth/v1")
	controllers.LoadAuth(authV1)

	r.Run()
}
