package task

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"context"
	"utils/db/db"
	"utils/types"
)

// GetInput represents the input data for retrieving a task.
type GetInput struct {
	ID model.TaskID
}

// GetTask retrieves a task by its ID.
func GetTask(q db.Querier, ctx context.Context, input GetInput) types.Result[model.Task, model.AppError] {
	return task_repository.FindTaskByID(q, ctx, input.ID)
}
