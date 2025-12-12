package task

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"utils/types"
)

// CreateTask creates a new task with the given title and description.
func CreateTask(title model.TaskTitle, description model.TaskDescription) types.Result[model.Task, model.AppError] {
	return task_repository.CreateTask(title, description)
}
