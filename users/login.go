package users

import (
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
		handleRequestError(c)
		return
	}

	db := databaseHandler.OpenDbConnectionLocal()
	query := "SELECT pwd,user_id FROM users WHERE email = $1"
	err := db.QueryRow(query, loginUser.Email).Scan(&pass, &uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"response": "Doesn't exist",
			"result":   nil,
		})
		return
	}
	if passVerify := checkPassword(pass, loginUser.Password); passVerify != nil && passVerify == bcrypt.ErrMismatchedHashAndPassword {
		handleRequestError(c)
		return
	}
	token, err := getToken(uid)
	if err != nil {
		handleRequestError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   token,
	})
}

func GetUser(c *gin.Context) {
	user_id, err := ExtractIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "error",
			"result":   nil,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"result":   user_id,
	})
}
