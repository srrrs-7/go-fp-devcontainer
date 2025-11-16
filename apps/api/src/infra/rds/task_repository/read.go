package task_repository

import (
	"api/src/domain/model"
	"utils/types"
)

func FindTaskByID(id model.TaskID) types.Result[model.Task, model.AppError] {
	return types.Ok[model.Task, model.AppError](model.Task{
		ID:          id,
		Title:       "Sample Task",
		Description: "This is a sample task description.",
		Completed:   false,
	})
}
