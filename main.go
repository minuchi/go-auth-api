package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minuchi/go-auth-api/lib"
)

func main() {
	r := lib.SetupRouter(gin.DebugMode)

	r.Run()
}
