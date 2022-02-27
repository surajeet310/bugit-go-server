package workspaces

import (
	"github.com/google/uuid"
)

type Workspace struct {
	W_id         uuid.UUID `json:"w_id"`
	Name         string    `json:"name"`
	Descp        string    `json:"descp"`
	ProjectCount int       `json:"project_count"`
	MemberCount  int       `json:"member_count"`
	CreatedAt    string    `json:"created_at"`
	IsActive     bool      `json:"is_active"`
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
