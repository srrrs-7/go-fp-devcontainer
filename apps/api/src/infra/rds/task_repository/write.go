package task_repository

import (
	"api/src/domain/model"
	"utils/types"

	"github.com/google/uuid"
)

func CreateTask(title model.TaskTitle, description model.TaskDescription) types.Result[model.Task, model.AppError] {
	newID := model.NewTaskID(uuid.New().String())
	task := model.Task{
		ID:          newID,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	return types.Ok[model.Task, model.AppError](task)
}

func UpdateTask(id model.TaskID, title model.TaskTitle, description model.TaskDescription, completed model.TaskCompleted) types.Result[model.Task, model.AppError] {
	task := model.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
	return types.Ok[model.Task, model.AppError](task)
}
