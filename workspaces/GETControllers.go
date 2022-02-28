package workspaces

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
	"github.com/surajeet310/bugit-go-server/projects"
)

func handleError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"response": "",
	})
}

func ListOfWorkspaces(c *gin.Context) {
	var workspace HomeWorkspaces
	var workspaceId uuid.UUID
	var home []HomeWorkspaces
	user_id := c.Query("user_id")
	db := databaseHandler.OpenDbConnectionLocal()
	workspaceIds, err := db.Query("SELECT w_id FROM workspace_members WHERE user_id = $1", user_id)
	if err != nil {
		log.Println(err)
		handleError(c)
		return
	}
	for workspaceIds.Next() {
		workspaceIds.Scan(&workspaceId)
		query := "SELECT w_id,name,project_count,member_count FROM workspaces WHERE w_id = $1"
		err := db.QueryRow(query, workspaceId).Scan(&workspace.W_id, &workspace.Name, &workspace.ProjectCount, &workspace.MemberCount)
		if err != nil {
			log.Println(err)
			handleError(c)
			return
		}
		home = append(home, workspace)
	}

	c.JSON(http.StatusOK, gin.H{
		"response": home,
	})
}

func SingleWorkspace(c *gin.Context) {
	var workspace Workspace
	var project projects.HomeProjects
	var projects []projects.HomeProjects

	var workspace_id = c.Query("workspace_id")
	db := databaseHandler.OpenDbConnectionLocal()
	err := db.QueryRow("SELECT * FROM workspaces WHERE w_id = $1", workspace_id).Scan(&workspace.W_id, &workspace.Name, &workspace.Descp, &workspace.ProjectCount, &workspace.MemberCount, &workspace.CreatedAt)
	if err != nil {
		handleError(c)
		return
	}
	projectList, err := db.Query("SELECT p_id,name,task_count,member_count FROM projects WHERE w_id = $1", workspace_id)
	if err != nil {
		handleError(c)
		return
	}
	for projectList.Next() {
		projectList.Scan(&project)
		projects = append(projects, project)
	}
	c.JSON(http.StatusOK, gin.H{
		"workspace": workspace,
		"projects":  projects,
	})
}
