package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/minuchi/go-auth-api/controllers/auth"
)

func LoadAuth(r *gin.RouterGroup) {
	r.GET("/time", auth.GetTime)
	r.POST("/login", auth.Login)
	r.POST("/signup", auth.SignUp)
	r.GET("/token", auth.IssueAccessToken)
	r.GET("/verify", auth.Verify)
}
