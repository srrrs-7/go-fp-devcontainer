package task

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"context"
	"utils/db/db"
	"utils/types"
)

// CreateInput represents the input data for creating a task.
type CreateInput struct {
	Title       model.TaskTitle
	Description model.TaskDescription
}

// CreateTask creates a new task with the given title and description.
func CreateTask(q db.Querier, ctx context.Context, input CreateInput) types.Result[model.Task, model.AppError] {
	cmd := model.TaskCmd{
		Title:       input.Title,
		Description: input.Description,
		Completed:   false, // Default value
	}
	return task_repository.CreateTask(q, ctx, cmd)
}
