package tasks

import "github.com/pborman/uuid"

type HomeTasks struct {
	T_id uuid.UUID `json:"t_id"`
	Name string    `json:"name"`
}

type Task struct {
	T_id       uuid.UUID `json:"t_id"`
	P_id       uuid.UUID `json:"p_id"`
	Name       string    `json:"name"`
	Descp      string    `json:"descp"`
	Assignee   uuid.UUID `json:"assignee"`
	AssignedTo uuid.UUID `json:"assigned_to"`
	CreatedAt  string    `json:"created_at"`
	Deadline   string    `json:"deadline"`
	Tech       string    `json:"tech"`
}

type GetTaskStruct struct {
	T_id           uuid.UUID `json:"t_id"`
	P_id           uuid.UUID `json:"p_id"`
	Name           string    `json:"name"`
	Descp          string    `json:"descp"`
	Assignee       uuid.UUID `json:"assignee"`
	AssigneeName   string    `json:"assignee_name"`
	AssignedTo     uuid.UUID `json:"assigned_to"`
	AssignedToName string    `json:"assigned_to_name"`
	CreatedAt      string    `json:"created_at"`
	Deadline       string    `json:"deadline"`
	Tech           string    `json:"tech"`
}

type TaskAssign struct {
	T_id    uuid.UUID `json:"t_id"`
	User_id uuid.UUID `json:"user_id"`
}

type TaskComment struct {
	Tc_id     uuid.UUID `json:"tc_id"`
	T_id      uuid.UUID `json:"t_id"`
	User_id   uuid.UUID `json:"user_id"`
	UserName  string    `json:"user_name"`
	Comment   string    `json:"comment"`
	CreatedAt string    `json:"created_at"`
}
