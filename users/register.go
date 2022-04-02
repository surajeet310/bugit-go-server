package users

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
	"golang.org/x/crypto/bcrypt"
)

func generateUUID() uuid.UUID {
	return uuid.NewRandom()
}

func getHashedPwd(pass string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return pass
	}
	return string(hashedPass)
}

func isUser(email string, db *sql.DB) bool {
	var count int
	query := "SELECT count(*) FROM users WHERE email = $1"
	db.QueryRow(query, email).Scan(&count)
	return count == 0
}

func RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		handleRequestError(c)
		return
	}
	db := databaseHandler.OpenDbConnectionLocal()
	checkUser := isUser(user.Email, db)
	if checkUser {
		query := "INSERT INTO users (user_id,fname,lname,email,pwd,alert) VALUES ($1,$2,$3,$4,$5,$6)"
		user.User_id = generateUUID()
		user.Password = getHashedPwd(user.Password)
		user.Alert = 0
		_, err := db.Query(query, user.User_id, user.Fname, user.Lname, user.Email, user.Password, user.Alert)
		if err != nil {
			handleRequestError(c)
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"response": "success",
			"result":   nil,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"response": "exists",
			"result":   nil,
		})
		return
	}
}
