package main

import (
	"github.com/surajeet310/bugit-go-server/users"
	"github.com/surajeet310/bugit-go-server/workspaces"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	urlRouter := gin.New()
	urlRouter.POST("/register", users.RegisterUser)
	urlRouter.POST("/login", users.LoginUser)
	urlRouter.GET("/home", workspaces.ListOfWorkspaces)
	urlRouter.GET("/home/workspace", workspaces.SingleWorkspace)
	urlRouter.POST("/addWorkspace", workspaces.AddWorkspace)
	err := urlRouter.Run()
	if err != nil {
		return
	}
}
