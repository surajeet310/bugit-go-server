package tasks

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
)

func handleError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"response": "error",
		"result":   nil,
	})
}

func getUUID() uuid.UUID {
	return uuid.NewRandom()
}

func changeTaskCount(db *sql.DB, p_id uuid.UUID, option string) error {
	var taskCount int
	query := "SELECT task_count FROM projects WHERE p_id = $1"
	err := db.QueryRow(query, p_id).Scan(&taskCount)
	if err != nil {
		return err
	}
	if option == "add" {
		taskCount++
	} else {
		taskCount--
	}
	query = "UPDATE projects SET task_count = $1 WHERE p_id = $2"
	_, err = db.Query(query, taskCount, p_id)
	return err
}

func AddTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		handleError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	query := "INSERT INTO tasks (t_id,p_id,name,descp,assignee,createdat,deadline,tech) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	task.T_id = getUUID()
	_, err := db.Query(query, task.T_id, task.P_id, task.Name, task.Descp, task.Assignee, task.CreatedAt, task.Deadline, task.Tech)
	if err != nil {
		handleError(c)
		return
	}
	err = changeTaskCount(db, task.P_id, "add")
	if err != nil {
		handleError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func AssignTask(c *gin.Context) {
	var count int
	var task TaskAssign
	if err := c.ShouldBindJSON(&task); err != nil {
		handleError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	_ = db.QueryRow("SELECT count(*) FROM task_members WHERE t_id = $1", task.T_id).Scan(&count)
	if count == 0 {
		query := "INSERT INTO task_members (t_id,assignedto) VALUES ($1,$2)"
		_, err := db.Query(query, task.T_id, task.User_id)
		if err != nil {
			handleError(c)
			return
		}
	} else {
		query := "UPDATE task_members SET assignedto = $1 WHERE t_id = $2"
		_, err := db.Query(query, task.User_id, task.T_id)
		if err != nil {
			handleError(c)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func AddComment(c *gin.Context) {
	var comment TaskComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		handleError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	query := "INSERT INTO task_comments (tc_id,t_id,comment,user_id,createdat) VALUES ($1,$2,$3,$4,$5)"
	comment.Tc_id = getUUID()
	_, err := db.Query(query, comment.Tc_id, comment.T_id, comment.Comment, comment.User_id, comment.CreatedAt)
	if err != nil {
		handleError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}
