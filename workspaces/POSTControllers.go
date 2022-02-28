package workspaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
)

func generateUUID() uuid.UUID {
	return uuid.NewRandom()
}

func handleInsertError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"response": "",
	})
}

func AddWorkspace(c *gin.Context) {
	var workspace AddWorkspaceStruct
	var workspaceMember WorkspaceMembers
	if err := c.ShouldBindJSON(&workspace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "",
		})
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	query := "INSERT INTO workspaces (w_id,name,descp,project_count,member_count,createdat) VALUES ($1,$2,$3,$4,$5,$6)"
	workspace.W_id = generateUUID()
	_, err := db.Query(query, workspace.W_id, workspace.Name, workspace.Descp, workspace.ProjectCount, workspace.MemberCount, workspace.CreatedAt)
	if err != nil {
		handleInsertError(c)
		return
	}

	workspaceMember.UserId = workspace.UserId
	workspaceMember.W_id = workspace.W_id
	workspaceMember.IsAdmin = true

	query = "INSERT INTO workspace_members (w_id,user_id,is_admin) VALUES ($1,$2,$3)"
	_, err = db.Query(query, workspaceMember.W_id, workspaceMember.UserId, workspaceMember.IsAdmin)
	if err != nil {
		handleInsertError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"workspace":        workspace,
		"workspace-member": workspaceMember,
	})
}
