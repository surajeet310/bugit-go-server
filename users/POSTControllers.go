package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
	"golang.org/x/crypto/bcrypt"
)

func ChangeUserFname(c *gin.Context) {
	var user UserDetails
	if err := c.ShouldBindJSON(&user); err != nil {
		handleRequestError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	query := "UPDATE users SET fname = $1 WHERE user_id = $2"
	_, err := db.Query(query, user.Fname, user.User_id)
	if err != nil {
		handleRequestError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func ChangeUserLname(c *gin.Context) {
	var user UserDetails
	if err := c.ShouldBindJSON(&user); err != nil {
		handleRequestError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	query := "UPDATE users SET lname = $1 WHERE user_id = $2"
	_, err := db.Query(query, user.Lname, user.User_id)
	if err != nil {
		handleRequestError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func ChangePwd(c *gin.Context) {
	var user UserPwd
	if err := c.ShouldBindJSON(&user); err != nil {
		handleRequestError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	query := "UPDATE users SET pwd = $1 WHERE user_id = $2"
	user.Password = getHashedPwd(user.Password)
	_, err := db.Query(query, user.Password, user.User_id)
	if err != nil {
		handleRequestError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}

func CheckOldPwd(c *gin.Context) {
	var actualPass string
	var userPwd UserPwd
	if err := c.ShouldBindJSON(&userPwd); err != nil {
		handleRequestError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	err := db.QueryRow("SELECT pwd FROM users WHERE user_id = $1", userPwd.User_id).Scan(&actualPass)
	if err != nil {
		handleRequestError(c)
		return
	}
	if err = checkPassword(actualPass, userPwd.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		handleRequestError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   nil,
	})
}
