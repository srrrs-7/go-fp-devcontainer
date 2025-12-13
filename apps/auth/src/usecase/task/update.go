package task

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"utils/types"
)

// UpdateTask updates an existing task with the given parameters.
func UpdateTask(
	id model.TaskID,
	title model.TaskTitle,
	description model.TaskDescription,
	completed model.TaskCompleted,
) types.Result[model.Task, model.AppError] {
	return task_repository.UpdateTask(id, title, description, completed)
}
