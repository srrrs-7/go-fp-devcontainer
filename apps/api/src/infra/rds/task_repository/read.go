package task_repository

import (
	"api/src/domain/apperror"
	"api/src/domain/task"
	"context"
	"database/sql"
	"errors"
	"utils/db/db"
	"utils/types"

	"github.com/google/uuid"
)

func FindTaskByID(q db.Querier, ctx context.Context, id task.TaskID) types.Result[task.Task, apperror.AppError] {
	return types.Map(
		types.MapErr(
			types.FromPair(q.GetTask(ctx, uuid.UUID(id))),
			func(e error) apperror.AppError {
				if errors.Is(e, sql.ErrNoRows) {
					return apperror.NewNotFoundError(e, "Task")
				}
				return apperror.NewDatabaseError(e, "TaskRepository")
			},
		),
		func(t db.Task) task.Task {
			return task.NewTask(
				task.TaskID(t.ID),
				task.TaskTitle(t.Title),
				task.TaskDescription(t.Description.String),
				task.TaskStatus(t.Status),
			)
		},
	)
}

func FindAllTasks(q db.Querier, ctx context.Context) types.Result[[]task.Task, apperror.AppError] {
	return types.Map(
		types.MapErr(
			types.FromPair(q.ListTasks(ctx)),
			func(e error) apperror.AppError {
				return apperror.NewDatabaseError(e, "TaskRepository")
			},
		),
		func(tasks []db.Task) []task.Task {
			var result []task.Task
			for _, t := range tasks {
				result = append(result, task.NewTask(
					task.TaskID(t.ID),
					task.TaskTitle(t.Title),
					task.TaskDescription(t.Description.String),
					task.TaskStatus(t.Status),
				))
			}
			return result
		},
	)
}
