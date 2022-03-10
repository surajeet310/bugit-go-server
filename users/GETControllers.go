package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
)

func handleRequestError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"response": "error",
		"result":   nil,
	})
}

func GetUserFromId(c *gin.Context) {
	var user BasicUser
	var fname, lname, email string
	var alert int
	db := databaseHandler.OpenDbConnectionLocal()
	user_id := c.Query("user_id")
	query := "SELECT email,fname,lname,alert FROM users WHERE user_id = $1"
	err := db.QueryRow(query, user_id).Scan(&email, &fname, &lname, &alert)
	if err != nil {
		handleRequestError(c)
		return
	}
	user.Alert = alert
	user.Fname = fname
	user.Lname = lname
	user.Email = email

	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   user,
	})
}

func DeleteUser(c *gin.Context) {
	user_id := c.Query("user_id")
	db := databaseHandler.OpenDbConnectionLocal()
	query := "DELETE FROM users WHERE user_id = $1"
	_, err := db.Query(query, user_id)
	if err != nil {
		handleRequestError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}
