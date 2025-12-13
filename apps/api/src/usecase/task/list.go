package task

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"context"
	"utils/db/db"
	"utils/types"
)

// ListInput represents the input data for listing tasks.
type ListInput struct {
	// Future filters can be added here
}

// ListTasks retrieves all tasks from the repository.
func ListTasks(q db.Querier, ctx context.Context, input ListInput) types.Result[[]model.Task, model.AppError] {
	return task_repository.FindAllTasks(q, ctx)
}
