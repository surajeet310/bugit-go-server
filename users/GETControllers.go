package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
	"golang.org/x/crypto/bcrypt"
)

func handleRequestError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"response": "",
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
		"response": user,
	})
}

func CheckOldPwd(c *gin.Context) {
	var actualPass string
	oldPass := c.Query("pwd")
	user_id := c.Query("user_id")
	db := databaseHandler.OpenDbConnectionLocal()
	err := db.QueryRow("SELECT pwd FROM users WHERE user_id = $1", user_id).Scan(&actualPass)
	if err != nil {
		handleRequestError(c)
		return
	}
	if err = checkPassword(actualPass, oldPass); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusNotFound, gin.H{
			"response": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
	})
}
