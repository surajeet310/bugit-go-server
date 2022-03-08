package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
	"github.com/surajeet310/bugit-go-server/tasks"
)

func SingleProjectList(c *gin.Context) {
	var project SingleProject
	var task tasks.HomeTasks
	var tasks []tasks.HomeTasks
	p_id := c.Query("project_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT p_id,name,descp,task_count,member_count,createdat,deadline,tech FROM projects WHERE p_id = $1"
	err := db.QueryRow(query, p_id).Scan(&project.P_id, &project.Name, &project.Descp, &project.TaskCount, &project.MemberCount, &project.CreatedAt, &project.Deadline, &project.Tech)
	if err != nil {
		handleBadReqError(c)
		return
	}
	query = "SELECT t_id,name FROM tasks WHERE p_id = $1"
	taskList, err := db.Query(query, project.P_id)
	if err != nil {
		handleBadReqError(c)
		return
	}
	for taskList.Next() {
		taskList.Scan(&task.T_id, &task.Name)
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result": gin.H{
			"project": project,
			"tasks":   tasks,
		},
	})
}

func GetProjectMembers(c *gin.Context) {
	var projectMember ProjectMember
	var projectMembers []ProjectMember
	var fname, lname string
	var assignedTo uuid.UUID
	p_id := c.Query("project_id")
	t_id := c.Query("task_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT user_id,is_admin FROM project_members WHERE p_id = $1"
	users, err := db.Query(query, p_id)
	if err != nil {
		handleBadReqError(c)
		return
	}
	for users.Next() {
		users.Scan(&projectMember.User_id, &projectMember.IsAdmin)
		err = db.QueryRow("SELECT fname,lname FROM users WHERE user_id = $1", projectMember.User_id).Scan(&fname, &lname)
		if err != nil {
			handleBadReqError(c)
			return
		}
		err = db.QueryRow("SELECT assignedto FROM tasks WHERE t_id = $1", t_id).Scan(&assignedTo)
		if err != nil {
			handleBadReqError(c)
			return
		}
		if assignedTo.String() == projectMember.User_id.String() {
			projectMember.IsAssigned = true
		} else {
			projectMember.IsAssigned = false
		}
		projectMember.UserName = fname + " " + lname
		projectMembers = append(projectMembers, projectMember)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   projectMembers,
	})
}

func GetAllProjectMembers(c *gin.Context) {
	var projectMember ProjectMember
	var projectMembers []ProjectMember
	var fname, lname string
	p_id := c.Query("project_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT user_id,is_admin FROM project_members WHERE p_id = $1"
	users, err := db.Query(query, p_id)
	if err != nil {
		handleBadReqError(c)
		return
	}
	for users.Next() {
		users.Scan(&projectMember.User_id, &projectMember.IsAdmin)
		err = db.QueryRow("SELECT fname,lname FROM users WHERE user_id = $1", projectMember.User_id).Scan(&fname, &lname)
		if err != nil {
			handleBadReqError(c)
			return
		}
		projectMember.IsAssigned = false
		projectMember.UserName = fname + " " + lname
		projectMembers = append(projectMembers, projectMember)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   projectMembers,
	})
}

func DeleteProject(c *gin.Context) {
	var w_id uuid.UUID
	p_id := c.Query("project_id")
	db := databaseHandler.OpenDbConnectionLocal()
	err := db.QueryRow("SELECT w_id FROM projects WHERE p_id = $1", p_id).Scan(&w_id)
	if err != nil {
		handleBadReqError(c)
		return
	}
	query := "DELETE FROM projects WHERE p_id = $1"
	_, err = db.Query(query, p_id)
	if err != nil {
		handleBadReqError(c)
		return
	}
	err = changeProjectCount(db, w_id, "sub")
	if err != nil {
		handleBadReqError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func RemoveProjectMember(c *gin.Context) {
	var t_id, assignedTo uuid.UUID
	user_id := c.Query("user_id")
	p_id := c.Query("project_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "DELETE FROM project_members WHERE user_id = $1 AND p_id = $2"
	_, err := db.Query(query, user_id, p_id)
	if err != nil {
		handleBadReqError(c)
		return
	}
	query = "SELECT t_id FROM tasks WHERE p_id = $1"
	tasks, err := db.Query(query, p_id)
	if err != nil {
		handleBadReqError(c)
		return
	}
	for tasks.Next() {
		tasks.Scan(&t_id)
		_ = db.QueryRow("SELECT assignedto FROM tasks WHERE t_id = $1", t_id).Scan(&assignedTo)
		if assignedTo.String() == user_id {
			_, err := db.Query("UPDATE tasks SET assignedto = $1 WHERE t_id = $2", nil, t_id)
			if err != nil {
				handleBadReqError(c)
				return
			}
		}
	}
	err = ChangeProjectMemberCount(db, uuid.Parse(p_id), "sub")
	if err != nil {
		handleBadReqError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}
