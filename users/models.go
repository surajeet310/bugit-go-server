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

type UserDetails struct {
	User_id uuid.UUID `json:"user_id"`
	Fname   string    `json:"fname"`
	Lname   string    `json:"lname"`
}

type BasicUser struct {
	Email string `json:"email"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Alert int    `json:"alert"`
}

type UserPwd struct {
	User_id  uuid.UUID `json:"user_id"`
	Password string    `json:"password"`
}
