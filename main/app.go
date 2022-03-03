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
	urlRouter.GET("/user", users.GetUserFromId)
	urlRouter.GET("/checkPwd", users.CheckOldPwd)
	urlRouter.PATCH("/changePwd", users.ChangePwd)
	urlRouter.PATCH("/changeFname", users.ChangeUserFname)
	urlRouter.PATCH("/changeLname", users.ChangeUserLname)
	urlRouter.GET("/home", workspaces.ListOfWorkspaces)
	urlRouter.GET("/home/workspace", workspaces.SingleWorkspace)
	urlRouter.GET("/workspaceMembers", workspaces.GetWorkspaceMembers)
	urlRouter.GET("/requests", workspaces.GetRequests)
	urlRouter.POST("/addWorkspace", workspaces.AddWorkspace)
	urlRouter.POST("/makeUserAdmin", workspaces.MakeWorkspaceMemberAdmin)
	urlRouter.POST("/addWorkspaceMemberReq", workspaces.AddWorkspaceMemberRequest)
	urlRouter.POST("/addWorkspaceMember", workspaces.AddWorkspaceMember)
	urlRouter.DELETE("/removeWorkspaceMember", workspaces.RemoveWorkspaceMember)
	urlRouter.DELETE("/deleteWorkspace", workspaces.DeleteWorkspace)
	err := urlRouter.Run()
	if err != nil {
		return
	}
}
