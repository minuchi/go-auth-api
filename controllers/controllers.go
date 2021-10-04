package controllers

import "github.com/gin-gonic/gin"

func LoadAuth(r *gin.RouterGroup) {
	r.GET("/time", GetTime)
	r.POST("/login", Login)
	r.POST("/signup", SignUp)
	r.GET("/token", IssueAccessToken)
	r.GET("/verify", Verify)
}
