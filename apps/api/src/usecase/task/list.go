package task

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"context"
	"utils/db/db"
	"utils/types"
)

// ListTasks retrieves all tasks from the repository.
func ListTasks(q db.Querier, ctx context.Context) types.Result[[]model.Task, model.AppError] {
	return task_repository.FindAllTasks(q, ctx)
}
