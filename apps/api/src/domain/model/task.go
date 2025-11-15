package model

import "github.com/google/uuid"

type TaskID uuid.UUID

func NewTaskID(id string) TaskID {
	return TaskID(uuid.MustParse(id))
}

func (t TaskID) String() string {
	return uuid.UUID(t).String()
}

type TaskTitle string

func (t TaskTitle) String() string {
	return string(t)
}

type TaskDescription string

func (t TaskDescription) String() string {
	return string(t)
}

type TaskCompleted bool

func (t TaskCompleted) Bool() bool {
	return bool(t)
}

type Task struct {
	ID          TaskID
	Title       TaskTitle
	Description TaskDescription
	Completed   TaskCompleted
}

type TaskCmd struct {
	Title       TaskTitle
	Description TaskDescription
}
