package projects

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
)

func handleBadReqError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"response": "error",
		"result":   nil,
	})
}

func generateUUID() uuid.UUID {
	return uuid.NewRandom()
}

func changeProjectCount(db *sql.DB, w_id uuid.UUID, option string) error {
	var projectCount = 0
	query := "SELECT project_count FROM workspaces WHERE w_id = $1"
	err := db.QueryRow(query, w_id).Scan(&projectCount)
	if err != nil {
		return err
	}
	if option == "add" {
		projectCount++
	} else {
		projectCount--
	}
	query = "UPDATE workspaces SET project_count = $1 WHERE w_id = $2"
	_, err = db.Query(query, projectCount, w_id)
	return err
}

func changeProjectMemberCount(db *sql.DB, p_id uuid.UUID, option string) error {
	var memberCount = 0
	query := "SELECT member_count FROM projects WHERE p_id = $1"
	err := db.QueryRow(query, p_id).Scan(&memberCount)
	if err != nil {
		return err
	}
	if option == "add" {
		memberCount++
	} else {
		memberCount--
	}
	query = "UPDATE projects SET member_count = $1 WHERE p_id = $2"
	_, err = db.Query(query, memberCount, p_id)
	return err
}

func AddProject(c *gin.Context) {
	var project Project
	if err := c.ShouldBindJSON(&project); err != nil {
		handleBadReqError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	project.P_id = generateUUID()
	project.MemberCount = 1
	project.TaskCount = 0
	query := "INSERT INTO projects (p_id,w_id,name,descp,task_count,member_count,createdat,deadline,tech) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	_, err := db.Query(query, project.P_id, project.W_id, project.Name, project.Descp, project.TaskCount, project.MemberCount, project.CreatedAt, project.Deadline, project.Tech)
	if err != nil {
		handleBadReqError(c)
		return
	}
	query = "INSERT INTO project_members (p_id,user_id,is_admin) VALUES ($1,$2,$3)"
	project.IsAdmin = true
	_, err = db.Query(query, project.P_id, project.User_id, project.IsAdmin)
	if err != nil {
		handleBadReqError(c)
		return
	}
	err = changeProjectCount(db, project.W_id, "add")
	if err != nil {
		handleBadReqError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func MakeProjectMemberAdmin(c *gin.Context) {
	var projectMember ProjectUser
	if err := c.ShouldBindJSON(&projectMember); err != nil {
		handleBadReqError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	query := "UPDATE project_members SET is_admin = $1 WHERE p_id = $2 AND user_id = $3"
	projectMember.IsAdmin = true
	_, err := db.Query(query, projectMember.IsAdmin, projectMember.P_id, projectMember.User_id)
	if err != nil {
		handleBadReqError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func AddProjectMember(c *gin.Context) {
	var projectMember ProjectUser
	if err := c.ShouldBindJSON(&projectMember); err != nil {
		handleBadReqError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	projectMember.IsAdmin = false
	query := "INSERT INTO project_members (p_id,user_id,is_admin) VALUES ($1,$2,$3)"
	_, err := db.Query(query, projectMember.P_id, projectMember.User_id, projectMember.IsAdmin)
	if err != nil {
		handleBadReqError(c)
		return
	}
	err = changeProjectMemberCount(db, projectMember.P_id, "add")
	if err != nil {
		handleBadReqError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}
