package tasks

import "github.com/pborman/uuid"

type HomeTasks struct {
	T_id uuid.UUID `json:"t_id"`
	Name string    `json:"name"`
}
