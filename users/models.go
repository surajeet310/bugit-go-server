package users

import (
	"github.com/pborman/uuid"
)

type User struct {
	User_id  uuid.UUID `json:"user_id"`
	Fname    string    `json:"fname"`
	Lname    string    `json:"lname"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Alert    int       `json:"alert"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
