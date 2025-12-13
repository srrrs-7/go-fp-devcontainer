package task

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"utils/types"
)

// GetTask retrieves a task by its ID.
func GetTask(id model.TaskID) types.Result[model.Task, model.AppError] {
	return task_repository.FindTaskByID(id)
}
