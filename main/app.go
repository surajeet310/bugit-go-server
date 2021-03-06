package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/surajeet310/bugit-go-server/databaseHandler"
	"github.com/surajeet310/bugit-go-server/middlewares"
	"github.com/surajeet310/bugit-go-server/projects"
	"github.com/surajeet310/bugit-go-server/tasks"
	"github.com/surajeet310/bugit-go-server/users"
	"github.com/surajeet310/bugit-go-server/workspaces"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func getPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$Port not set")
	}
	return ":" + port, nil
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	urlRouter := gin.New()
	db := databaseHandler.OpenDbConnection()
	defer db.Close()
	publicRouter := urlRouter.Group("/open")
	{
		publicRouter.POST("/register", users.RegisterUser)
		publicRouter.POST("/login", users.LoginUser)
	}

	privateRouter := urlRouter.Group("/auth")
	privateRouter.Use(middlewares.AuthMiddleware())
	{
		//users
		privateRouter.GET("/user", users.GetUserFromId)
		privateRouter.GET("/getUserId", users.GetUser)
		privateRouter.POST("/checkPwd", users.CheckOldPwd)
		privateRouter.PATCH("/changePwd", users.ChangePwd)
		privateRouter.PATCH("/changeFname", users.ChangeUserFname)
		privateRouter.PATCH("/changeLname", users.ChangeUserLname)
		privateRouter.DELETE("/deleteUser", users.DeleteUser)
		//workspaces
		privateRouter.GET("/home", workspaces.ListOfWorkspaces)
		privateRouter.GET("/home/workspace", workspaces.SingleWorkspace)
		privateRouter.GET("/workspaceMembers", workspaces.GetWorkspaceMembers)
		privateRouter.GET("/allWorkspaceMembers", workspaces.GetAllWorkspaceMembers)
		privateRouter.GET("/requests", workspaces.GetRequests)
		privateRouter.POST("/ignoreRequest", workspaces.IgnoreWorkspaceRequest)
		privateRouter.POST("/addWorkspace", workspaces.AddWorkspace)
		privateRouter.POST("/makeUserAdmin", workspaces.MakeWorkspaceMemberAdmin)
		privateRouter.POST("/addWorkspaceMemberReq", workspaces.AddWorkspaceMemberRequest)
		privateRouter.POST("/addWorkspaceMember", workspaces.AddWorkspaceMember)
		privateRouter.DELETE("/removeWorkspaceMember", workspaces.RemoveWorkspaceMember)
		privateRouter.DELETE("/deleteWorkspace", workspaces.DeleteWorkspace)
		//projects
		privateRouter.GET("/project", projects.SingleProjectList)
		privateRouter.GET("/projectMembers", projects.GetProjectMembers)
		privateRouter.GET("/allProjectMembers", projects.GetAllProjectMembers)
		privateRouter.POST("/addProject", projects.AddProject)
		privateRouter.POST("/makeProjectUserAdmin", projects.MakeProjectMemberAdmin)
		privateRouter.POST("/addProjectMember", projects.AddProjectMember)
		privateRouter.DELETE("/deleteProject", projects.DeleteProject)
		privateRouter.DELETE("/removeProjectMember", projects.RemoveProjectMember)
		//tasks
		privateRouter.GET("/task", tasks.GetTask)
		privateRouter.POST("/addTask", tasks.AddTask)
		privateRouter.POST("/assignTask", tasks.AssignTask)
		privateRouter.POST("/addComment", tasks.AddComment)
		privateRouter.DELETE("/deleteTask", tasks.DeleteTask)

	}
	urlRouter.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"response": "404",
			"result":   nil,
		})
	})

	port, err := getPort()
	if err != nil {
		log.Println(err)
		return
	}
	err = urlRouter.Run(port)
	if err != nil {
		return
	}
}
