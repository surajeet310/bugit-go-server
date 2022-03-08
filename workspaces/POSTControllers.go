package workspaces

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
)

func generateUUID() uuid.UUID {
	return uuid.NewRandom()
}

func handleError(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"response": err,
		"result":   nil,
	})
}

func AddWorkspace(c *gin.Context) {
	var workspace AddWorkspaceStruct
	var workspaceMember WorkspaceMembers
	if err := c.ShouldBindJSON(&workspace); err != nil {
		handleError(c, "error")
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	query := "INSERT INTO workspaces (w_id,name,descp,project_count,member_count,createdat) VALUES ($1,$2,$3,$4,$5,$6)"
	workspace.W_id = generateUUID()
	workspace.MemberCount = 1
	workspace.ProjectCount = 0
	_, err := db.Query(query, workspace.W_id, workspace.Name, workspace.Descp, workspace.ProjectCount, workspace.MemberCount, workspace.CreatedAt)
	if err != nil {
		handleError(c, "error")
		return
	}

	workspaceMember.UserId = workspace.UserId
	workspaceMember.W_id = workspace.W_id
	workspaceMember.IsAdmin = true

	query = "INSERT INTO workspace_members (w_id,user_id,is_admin) VALUES ($1,$2,$3)"
	_, err = db.Query(query, workspaceMember.W_id, workspaceMember.UserId, workspaceMember.IsAdmin)
	if err != nil {
		handleError(c, "error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func MakeWorkspaceMemberAdmin(c *gin.Context) {
	var workspaceMember WorkspaceMembers
	if err := c.ShouldBindJSON(&workspaceMember); err != nil {
		handleError(c, "error")
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	workspaceMember.IsAdmin = true
	query := "UPDATE workspace_members SET is_admin = $1 WHERE user_id = $2 AND w_id = $3"
	_, err := db.Query(query, workspaceMember.IsAdmin, workspaceMember.UserId, workspaceMember.W_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func changeAlertColumn(db *sql.DB, option string, user_id uuid.UUID) error {
	var alert int
	query := "SELECT alert FROM users WHERE user_id = $1"
	err := db.QueryRow(query, user_id).Scan(&alert)
	if err != nil {
		return err
	}
	query = "UPDATE users SET alert = $1 WHERE user_id = $2"
	if option == "add" {
		alert += 1
	} else {
		alert -= 1
	}
	_, err = db.Query(query, alert, user_id)
	return err
}

func changeMemberCount(db *sql.DB, option string, w_id uuid.UUID) error {
	var memberCount int
	query := "SELECT member_count FROM workspaces WHERE w_id = $1"
	err := db.QueryRow(query, w_id).Scan(&memberCount)
	if err != nil {
		return err
	}
	if option == "add" {
		memberCount += 1
	} else {
		memberCount -= 1
	}
	query = "UPDATE workspaces SET member_count = $1 WHERE w_id = $2"
	_, err = db.Query(query, memberCount, w_id)
	return err
}

func checkExistingMember(c *gin.Context, db *sql.DB, w_id, user_id uuid.UUID) (int, error) {
	var count int
	query := "SELECT count(*) FROM workspace_members WHERE user_id = $1 AND w_id = $2"
	err := db.QueryRow(query, user_id, w_id).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, err
}

func AddWorkspaceMemberRequest(c *gin.Context) {
	var workspaceMemberReq RequestAddMember
	var user_id uuid.UUID
	var fname, lname string
	if err := c.ShouldBindJSON(&workspaceMemberReq); err != nil {
		handleError(c, "error")
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT user_id FROM users WHERE email = $1"
	err := db.QueryRow(query, workspaceMemberReq.Email).Scan(&user_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	count, err := checkExistingMember(c, db, workspaceMemberReq.W_id, user_id)
	if err != nil {
		handleError(c, "error")
		return
	}

	if count == 0 {
		err = db.QueryRow("SELECT fname,lname FROM users WHERE user_id = $1", workspaceMemberReq.RequesteeId).Scan(&fname, &lname)
		if err != nil {
			handleError(c, "error")
			return
		}
		fullName := fname + " " + lname
		query = "INSERT INTO requests (req_id,w_id,user_id,requestee) VALUES ($1,$2,$3,$4)"
		req_id := generateUUID()
		_, err = db.Query(query, req_id, workspaceMemberReq.W_id, user_id, fullName)
		if err != nil {
			handleError(c, "error")
			return
		}
		err = changeAlertColumn(db, "add", user_id)
		if err != nil {
			handleError(c, "error")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"response": "success",
			"result":   nil,
		})
		return
	}
	handleError(c, "exists")
}

func AddWorkspaceMember(c *gin.Context) {
	var workspaceMember WorkspaceMembers
	var w_id, user_id uuid.UUID
	var req_id RequestIdStruct
	if err := c.ShouldBindJSON(&req_id); err != nil {
		handleError(c, "error")
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	err := db.QueryRow("SELECT user_id,w_id FROM requests WHERE req_id = $1", req_id.RequestId).Scan(&user_id, &w_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	workspaceMember.IsAdmin = false
	workspaceMember.W_id = w_id
	workspaceMember.UserId = user_id
	query := "INSERT INTO workspace_members (w_id,user_id,is_admin) VALUES ($1,$2,$3)"
	_, err = db.Query(query, workspaceMember.W_id, workspaceMember.UserId, workspaceMember.IsAdmin)
	if err != nil {
		handleError(c, "error")
		return
	}
	query = "DELETE FROM requests WHERE req_id = $1"
	_, err = db.Query(query, req_id.RequestId)
	if err != nil {
		handleError(c, "error")
		return
	}
	err = changeAlertColumn(db, "sub", workspaceMember.UserId)
	if err != nil {
		handleError(c, "error")
		return
	}
	err = changeMemberCount(db, "add", workspaceMember.W_id)
	if err != nil {
		handleError(c, "error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}
