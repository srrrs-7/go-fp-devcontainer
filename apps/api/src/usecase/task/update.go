package task

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"context"
	"utils/db/db"
	"utils/types"
)

// UpdateInput represents the input data for updating a task.
type UpdateInput struct {
	ID          model.TaskID
	Title       model.TaskTitle
	Description model.TaskDescription
	Completed   model.TaskCompleted
}

// UpdateTask updates an existing task with the given parameters.
func UpdateTask(
	q db.Querier,
	ctx context.Context,
	input UpdateInput,
) types.Result[model.Task, model.AppError] {
	cmd := model.TaskCmd{
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
	}
	return task_repository.UpdateTask(q, ctx, input.ID, cmd)
}
