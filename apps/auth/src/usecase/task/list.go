package task

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"utils/types"
)

// ListTasks retrieves all tasks from the repository.
func ListTasks() types.Result[[]model.Task, model.AppError] {
	return task_repository.FindAllTasks()
}
