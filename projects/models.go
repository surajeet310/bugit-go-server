package projects

import "github.com/pborman/uuid"

type Project struct {
	P_id        uuid.UUID `json:"p_id"`
	W_id        uuid.UUID `json:"w_id"`
	Name        string    `json:"name"`
	Descp       string    `json:"descp"`
	TaskCount   int       `json:"task_count"`
	MemberCount int       `json:"member_count"`
	CreatedAt   string    `json:"created_at"`
	Deadline    string    `json:"deadline"`
	Tech        string    `json:"tech"`
}

type HomeProjects struct {
	P_id        uuid.UUID `json:"p_id"`
	W_id        uuid.UUID `json:"w_id"`
	Name        string    `json:"name"`
	TaskCount   int       `json:"task_count"`
	MemberCount int       `json:"member_count"`
}
