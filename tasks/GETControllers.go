package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
)

func GetTask(c *gin.Context) {
	var task GetTaskStruct
	var comment TaskComment
	var comments []TaskComment
	var fname, lname string
	t_id := c.Query("task_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT * FROM tasks WHERE t_id = $1"
	taskList, err := db.Query(query, t_id)
	if err != nil {
		handleError(c)
		return
	}
	for taskList.Next() {
		taskList.Scan(&task.T_id, &task.P_id, &task.Name, &task.Descp, &task.Assignee, &task.CreatedAt, &task.Deadline, &task.Tech)
	}
	query = "SELECT assignedto FROM task_members WHERE t_id = $1"
	_ = db.QueryRow(query, t_id).Scan(&task.AssignedTo)
	_ = db.QueryRow("SELECT fname,lname FROM users WHERE user_id = $1", task.Assignee).Scan(&fname, &lname)
	task.AssigneeName = fname + " " + lname
	_ = db.QueryRow("SELECT fname,lname FROM users WHERE user_id = $1", task.AssignedTo).Scan(&fname, &lname)
	task.AssignedToName = fname + " " + lname
	query = "SELECT * FROM task_comments WHERE t_id = $1"
	commentList, err := db.Query(query, t_id)
	if err != nil {
		handleError(c)
		return
	}
	for commentList.Next() {
		commentList.Scan(&comment.Tc_id, &comment.T_id, &comment.Comment, &comment.User_id, &comment.CreatedAt)
		_ = db.QueryRow("SELECT fname,lname FROM users WHERE user_id = $1", comment.User_id).Scan(&fname, &lname)
		comment.UserName = fname + " " + lname
		comments = append(comments, comment)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result": gin.H{
			"task":     task,
			"comments": comments,
		},
	})
}

func DeleteTask(c *gin.Context) {
	var p_id uuid.UUID
	t_id := c.Query("task_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT p_id FROM tasks WHERE t_id = $1"
	err := db.QueryRow(query, t_id).Scan(&p_id)
	if err != nil {
		handleError(c)
		return
	}
	query = "DELETE FROM tasks WHERE t_id = $1"
	_, err = db.Query(query, t_id)
	if err != nil {
		handleError(c)
		return
	}
	err = changeTaskCount(db, p_id, "sub")
	if err != nil {
		handleError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}
