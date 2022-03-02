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

func DeleteWorkspace(c *gin.Context) {
	db := databaseHandler.OpenDbConnectionLocal()
	workspace_id := c.Query("workspace_id")
	query := "DELETE FROM workspaces WHERE w_id = $1"
	_, err := db.Query(query, workspace_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
	})
}

func RemoveWorkspaceMember(c *gin.Context) {
	var w_id uuid.UUID
	user_id := c.Query("user_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT w_id FROM workspace_members WHERE user_id = $1"
	widList, err := db.Query(query, user_id)
	if err != nil {
		handleServerError(c)
		return
	}
	for widList.Next() {
		widList.Scan(&w_id)
		err = changeMemberCount(c, db, "sub", w_id)
		if err != nil {
			handleServerError(c)
			return
		}
	}
	query = "DELETE FROM workspace_members WHERE user_id = $1"
	_, err = db.Query(query, user_id)
	if err != nil {
		handleServerError(c)
		return
	}
	query = "DELETE FROM project_members WHERE user_id = $1"
	_, err = db.Query(query, user_id)
	if err != nil {
		handleServerError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
	})
}

func GetRequests(c *gin.Context) {
	var request Request
	var requests []Request
	user_id := c.Query("user_id")
	query := "SELECT * FROM requests WHERE user_id = $1"
	db := databaseHandler.OpenDbConnectionLocal()
	reqs, err := db.Query(query, user_id)
	if err != nil {
		handleError(c)
		return
	}
	for reqs.Next() {
		reqs.Scan(&request)
		requests = append(requests, request)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": requests,
	})
}

func GetWorkspaceMembers(c *gin.Context) {
	var fname, lname string
	var user_id uuid.UUID
	var is_admin bool
	var workspaceMember GetWorkspaceMemberStruct
	var workspaceMembers []GetWorkspaceMemberStruct
	w_id := c.Query("workspace_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT user_id,is_admin FROM workspace_members WHERE w_id = $1"
	members, err := db.Query(query, w_id)
	if err != nil {
		handleServerError(c)
		return
	}
	for members.Next() {
		members.Scan(&user_id, &is_admin)
		err = db.QueryRow("SELECT fname,lname FROM users WHERE user_id = $1", user_id).Scan(&fname, &lname)
		if err != nil {
			handleError(c)
			return
		}
		workspaceMember.UserId = user_id
		workspaceMember.IsAdmin = is_admin
		workspaceMember.UserName = fname + " " + lname
		workspaceMembers = append(workspaceMembers, workspaceMember)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": workspaceMembers,
	})
}
