package main

import (
	"github.com/surajeet310/bugit-go-server/users"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	urlRouter := gin.New()
	urlRouter.POST("/register", users.RegisterUser)
	urlRouter.POST("/login", users.LoginUser)
	err := urlRouter.Run()
	if err != nil {
		return
	}
}
