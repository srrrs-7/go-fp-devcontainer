package task_repository

import (
	"api/src/domain/model"
	"context"
	"utils/db/db"
	"utils/types"

	"github.com/google/uuid"
)

func FindTaskByID(q db.Querier, ctx context.Context, id model.TaskID) types.Result[model.Task, model.AppError] {
	return types.Map(
		types.MapErr(
			types.FromPair(q.GetTask(ctx, uuid.UUID(id))),
			func(e error) model.AppError {
				return model.NewDatabaseError(e, "TaskRepository")
			},
		),
		func(t db.Task) model.Task {
			return model.Task{
				ID:          model.TaskID(t.ID),
				Title:       model.TaskTitle(t.Title),
				Description: model.TaskDescription(t.Description.String),
				Completed:   model.TaskCompleted(t.Status == "completed"),
			}
		},
	)
}

func FindAllTasks(q db.Querier, ctx context.Context) types.Result[[]model.Task, model.AppError] {
	return types.Map(
		types.MapErr(
			types.FromPair(q.ListTasks(ctx)),
			func(e error) model.AppError {
				return model.NewDatabaseError(e, "TaskRepository")
			},
		),
		func(tasks []db.Task) []model.Task {
			var result []model.Task
			for _, task := range tasks {
				result = append(result, model.Task{
					ID:          model.TaskID(task.ID),
					Title:       model.TaskTitle(task.Title),
					Description: model.TaskDescription(task.Description.String),
					Completed:   model.TaskCompleted(task.Status == "completed"),
				})
			}
			return result
		},
	)
}
