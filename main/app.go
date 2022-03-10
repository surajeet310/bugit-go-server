package main

import (
	"github.com/surajeet310/bugit-go-server/projects"
	"github.com/surajeet310/bugit-go-server/tasks"
	"github.com/surajeet310/bugit-go-server/users"
	"github.com/surajeet310/bugit-go-server/workspaces"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	urlRouter := gin.New()
	//users
	urlRouter.POST("/register", users.RegisterUser)
	urlRouter.POST("/login", users.LoginUser)
	urlRouter.GET("/user", users.GetUserFromId)
	urlRouter.POST("/checkPwd", users.CheckOldPwd)
	urlRouter.PATCH("/changePwd", users.ChangePwd)
	urlRouter.PATCH("/changeFname", users.ChangeUserFname)
	urlRouter.PATCH("/changeLname", users.ChangeUserLname)
	urlRouter.DELETE("/deleteUser", users.DeleteUser)
	//workspaces
	urlRouter.GET("/home", workspaces.ListOfWorkspaces)
	urlRouter.GET("/home/workspace", workspaces.SingleWorkspace)
	urlRouter.GET("/workspaceMembers", workspaces.GetWorkspaceMembers)
	urlRouter.GET("/allWorkspaceMembers", workspaces.GetAllWorkspaceMembers)
	urlRouter.GET("/requests", workspaces.GetRequests)
	urlRouter.POST("/addWorkspace", workspaces.AddWorkspace)
	urlRouter.POST("/makeUserAdmin", workspaces.MakeWorkspaceMemberAdmin)
	urlRouter.POST("/addWorkspaceMemberReq", workspaces.AddWorkspaceMemberRequest)
	urlRouter.POST("/addWorkspaceMember", workspaces.AddWorkspaceMember)
	urlRouter.DELETE("/removeWorkspaceMember", workspaces.RemoveWorkspaceMember)
	urlRouter.DELETE("/deleteWorkspace", workspaces.DeleteWorkspace)
	//projects
	urlRouter.POST("/addProject", projects.AddProject)
	urlRouter.POST("/makeProjectUserAdmin", projects.MakeProjectMemberAdmin)
	urlRouter.POST("/addProjectMember", projects.AddProjectMember)
	urlRouter.GET("/project", projects.SingleProjectList)
	urlRouter.GET("/projectMembers", projects.GetProjectMembers)
	urlRouter.GET("/allProjectMembers", projects.GetAllProjectMembers)
	urlRouter.DELETE("/deleteProject", projects.DeleteProject)
	urlRouter.DELETE("/removeProjectMember", projects.RemoveProjectMember)
	//tasks
	urlRouter.GET("/task", tasks.GetTask)
	urlRouter.POST("/addTask", tasks.AddTask)
	urlRouter.POST("/assignTask", tasks.AssignTask)
	urlRouter.POST("/addComment", tasks.AddComment)
	urlRouter.DELETE("/deleteTask", tasks.DeleteTask)
	err := urlRouter.Run()
	if err != nil {
		return
	}
}
