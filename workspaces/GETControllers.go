package workspaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
	"github.com/surajeet310/bugit-go-server/projects"
)

func ListOfWorkspaces(c *gin.Context) {
	var workspace HomeWorkspaces
	var workspaceId uuid.UUID
	var home []HomeWorkspaces
	user_id := c.Query("user_id")
	db := databaseHandler.OpenDbConnectionLocal()
	workspaceIds, err := db.Query("SELECT w_id FROM workspace_members WHERE user_id = $1", user_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	for workspaceIds.Next() {
		workspaceIds.Scan(&workspaceId)
		query := "SELECT w_id,name,project_count,member_count FROM workspaces WHERE w_id = $1"
		err := db.QueryRow(query, workspaceId).Scan(&workspace.W_id, &workspace.Name, &workspace.ProjectCount, &workspace.MemberCount)
		if err != nil {
			handleError(c, "error")
			return
		}
		home = append(home, workspace)
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   home,
	})
}

func SingleWorkspace(c *gin.Context) {
	var workspace Workspace
	var project projects.HomeProjects
	var projects []projects.HomeProjects

	var workspace_id = c.Query("workspace_id")
	var user_id = c.Query("user_id")
	db := databaseHandler.OpenDbConnectionLocal()
	err := db.QueryRow("SELECT * FROM workspaces WHERE w_id = $1", workspace_id).Scan(&workspace.W_id, &workspace.Name, &workspace.Descp, &workspace.ProjectCount, &workspace.MemberCount, &workspace.CreatedAt)
	if err != nil {
		handleError(c, "error")
		return
	}
	err = db.QueryRow("SELECT is_admin FROM workspace_members WHERE w_id = $1 AND user_id = $2", workspace_id, user_id).Scan(&workspace.IsAdmin)
	if err != nil {
		handleError(c, "error")
		return
	}
	projectList, err := db.Query("SELECT p_id,name,task_count,member_count FROM projects WHERE w_id = $1", workspace_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	for projectList.Next() {
		projectList.Scan(&project.P_id, &project.Name, &project.TaskCount, &project.MemberCount)
		projects = append(projects, project)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result": gin.H{
			"workspace": workspace,
			"projects":  projects,
		},
	})
}

func DeleteWorkspace(c *gin.Context) {
	db := databaseHandler.OpenDbConnectionLocal()
	workspace_id := c.Query("workspace_id")
	query := "DELETE FROM workspaces WHERE w_id = $1"
	_, err := db.Query(query, workspace_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"results":  nil,
	})
}

func RemoveWorkspaceMember(c *gin.Context) {
	var count int
	var p_id uuid.UUID
	user_id := c.Query("user_id")
	w_id := c.Query("workspace_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "DELETE FROM workspace_members WHERE user_id = $1 AND w_id = $2"
	_, err := db.Query(query, user_id, w_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	query = "SELECT p_id FROM projects WHERE w_id = $1"
	projectList, err := db.Query(query, w_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	for projectList.Next() {
		projectList.Scan(&p_id)
		_ = db.QueryRow("SELECT count(*) FROM project_members WHERE user_id = $1 AND p_id = $2", user_id, p_id).Scan(&count)
		if count != 0 {
			_, err = db.Query("DELETE FROM project_members WHERE user_id = $1 AND p_id = $2", user_id, p_id)
			if err != nil {
				handleError(c, "error")
				return
			}
			err = projects.ChangeProjectMemberCount(db, p_id, "sub")
			if err != nil {
				handleError(c, "error")
				return
			}
		}
	}
	err = changeMemberCount(db, "sub", uuid.Parse(w_id))
	if err != nil {
		handleError(c, "error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func GetRequests(c *gin.Context) {
	var request Request
	var requests []Request
	user_id := c.Query("user_id")
	query := "SELECT * FROM requests WHERE user_id = $1 ORDER BY priority DESC"
	db := databaseHandler.OpenDbConnectionLocal()
	reqs, err := db.Query(query, user_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	for reqs.Next() {
		reqs.Scan(&request.Req_id, &request.W_id, &request.UserId, &request.Priority, &request.Requestee)
		requests = append(requests, request)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   requests,
	})
}

func GetWorkspaceMembers(c *gin.Context) {
	var count int
	var fname, lname string
	var workspaceMember GetWorkspaceMemberStruct
	var workspaceMembers []GetWorkspaceMemberStruct
	w_id := c.Query("workspace_id")
	p_id := c.Query("project_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT user_id,is_admin FROM workspace_members WHERE w_id = $1"
	members, err := db.Query(query, w_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	for members.Next() {
		members.Scan(&workspaceMember.UserId, &workspaceMember.IsAdmin)
		err = db.QueryRow("SELECT fname,lname FROM users WHERE user_id = $1", workspaceMember.UserId).Scan(&fname, &lname)
		if err != nil {
			handleError(c, "error")
			return
		}
		_ = db.QueryRow("SELECT count(*) FROM project_members WHERE p_id = $1 AND user_id = $2", p_id, workspaceMember.UserId).Scan(&count)
		if count == 0 {
			workspaceMember.IsTaken = false
		} else {
			workspaceMember.IsTaken = true
		}
		workspaceMember.UserName = fname + " " + lname
		workspaceMembers = append(workspaceMembers, workspaceMember)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   workspaceMembers,
	})
}

func GetAllWorkspaceMembers(c *gin.Context) {
	var fname, lname string
	var workspaceMember GetWorkspaceMemberStruct
	var workspaceMembers []GetWorkspaceMemberStruct
	w_id := c.Query("workspace_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT user_id,is_admin FROM workspace_members WHERE w_id = $1"
	members, err := db.Query(query, w_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	for members.Next() {
		members.Scan(&workspaceMember.UserId, &workspaceMember.IsAdmin)
		err = db.QueryRow("SELECT fname,lname FROM users WHERE user_id = $1", workspaceMember.UserId).Scan(&fname, &lname)
		if err != nil {
			handleError(c, "error")
			return
		}
		workspaceMember.IsTaken = false
		workspaceMember.UserName = fname + " " + lname
		workspaceMembers = append(workspaceMembers, workspaceMember)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   workspaceMembers,
	})
}
