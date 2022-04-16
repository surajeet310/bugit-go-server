package projects

import "github.com/pborman/uuid"

type Project struct {
	P_id        uuid.UUID `json:"p_id"`
	W_id        uuid.UUID `json:"w_id"`
	User_id     uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Descp       string    `json:"descp"`
	TaskCount   int       `json:"task_count"`
	MemberCount int       `json:"member_count"`
	CreatedAt   string    `json:"created_at"`
	Deadline    string    `json:"deadline"`
	Tech        string    `json:"tech"`
	IsAdmin     bool      `json:"is_admin"`
}

type HomeProjects struct {
	P_id        uuid.UUID `json:"p_id"`
	W_id        uuid.UUID `json:"w_id"`
	Name        string    `json:"name"`
	TaskCount   int       `json:"task_count"`
	MemberCount int       `json:"member_count"`
}

type ProjectUser struct {
	P_id    uuid.UUID `json:"p_id"`
	User_id uuid.UUID `json:"user_id"`
	IsAdmin bool      `json:"is_admin"`
}

type SingleProject struct {
	P_id        uuid.UUID `json:"p_id"`
	Name        string    `json:"name"`
	Descp       string    `json:"descp"`
	TaskCount   int       `json:"task_count"`
	MemberCount int       `json:"member_count"`
	IsAdmin     bool      `json:"is_admin"`
	CreatedAt   string    `json:"created_at"`
	Deadline    string    `json:"deadline"`
	Tech        string    `json:"tech"`
}

type ProjectMember struct {
	UserName   string    `json:"user_name"`
	User_id    uuid.UUID `json:"user_id"`
	IsAdmin    bool      `json:"is_admin"`
	IsAssigned bool      `json:"is_assigned"`
}
