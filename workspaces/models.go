package workspaces

import (
	"github.com/pborman/uuid"
)

type Workspace struct {
	W_id         uuid.UUID `json:"w_id"`
	Name         string    `json:"name"`
	Descp        string    `json:"descp"`
	IsAdmin      string    `json:"is_admin"`
	ProjectCount int       `json:"project_count"`
	MemberCount  int       `json:"member_count"`
	CreatedAt    string    `json:"created_at"`
}

type HomeWorkspaces struct {
	W_id         uuid.UUID `json:"w_id"`
	Name         string    `json:"name"`
	ProjectCount int       `json:"project_count"`
	MemberCount  int       `json:"member_count"`
}

type WorkspaceMembers struct {
	W_id    uuid.UUID `json:"w_id"`
	UserId  uuid.UUID `json:"user_id"`
	IsAdmin bool      `json:"is_admin"`
}

type AddWorkspaceStruct struct {
	W_id         uuid.UUID `json:"w_id"`
	UserId       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	Descp        string    `json:"descp"`
	ProjectCount int       `json:"project_count"`
	MemberCount  int       `json:"member_count"`
	CreatedAt    string    `json:"created_at"`
}

type RequestAddMember struct {
	W_id        uuid.UUID `json:"w_id"`
	RequesteeId uuid.UUID `json:"requestee_id"`
	Email       string    `json:"email"`
}

type Request struct {
	Req_id    uuid.UUID `json:"req_id"`
	W_id      uuid.UUID `json:"w_id"`
	UserId    uuid.UUID `json:"user_id"`
	Requestee string    `json:"requestee"`
	Priority  int       `json:"priority"`
}

type GetWorkspaceMemberStruct struct {
	UserId   uuid.UUID `json:"user_id"`
	UserName string    `json:"user_name"`
	IsAdmin  bool      `json:"is_admin"`
	IsTaken  bool      `json:"is_taken"`
}

type RequestIdStruct struct {
	RequestId uuid.UUID `json:"req_id"`
}
