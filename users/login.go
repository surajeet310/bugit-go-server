package users

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/pborman/uuid"
	"github.com/surajeet310/bugit-go-server/databaseHandler"
	"golang.org/x/crypto/bcrypt"
)

func checkPassword(actualPass, givenPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(actualPass), []byte(givenPass))
}

func getToken(uid uuid.UUID) (string, error) {
	err := godotenv.Load("../local.env")
	if err != nil {
		log.Println(err)
		return "", err
	}
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = uid
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET_KEY")))
}

func LoginUser(c *gin.Context) {
	var loginUser UserLogin
	var pass string
	var uid uuid.UUID
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "",
		})
		log.Print(err)
		return
	}

	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT pwd,user_id FROM users WHERE email = $1"
	err := db.QueryRow(query, loginUser.Email).Scan(&pass, &uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"response": "",
		})
		return
	}
	if passVerify := checkPassword(pass, loginUser.Password); passVerify != nil && passVerify == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"response": "",
		})
		return
	}
	token, err := getToken(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": token,
	})
}
