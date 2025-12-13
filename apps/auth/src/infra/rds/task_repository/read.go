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

func FindAllTasks() types.Result[[]model.Task, model.AppError] {
	tasks := []model.Task{
		{
			ID:          model.NewTaskID("00000000-0000-0000-0000-000000000001"),
			Title:       "Sample Task 1",
			Description: "This is the first sample task.",
			Completed:   false,
		},
		{
			ID:          model.NewTaskID("00000000-0000-0000-0000-000000000002"),
			Title:       "Sample Task 2",
			Description: "This is the second sample task.",
			Completed:   true,
		},
	}
	return types.Ok[[]model.Task, model.AppError](tasks)
}
