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

func CreateTask(q db.Querier, ctx context.Context, cmd task.TaskCmd) types.Result[task.Task, apperror.AppError] {
	return types.Map(
		types.MapErr(
			types.FromPair(q.CreateTask(ctx, db.CreateTaskParams{
				Title:       cmd.Title.String(),
				Description: sql.NullString{String: cmd.Description.String(), Valid: true},
				Status:      task.TaskStatusPending.String(), // Default status
				Priority:    "medium",                        // Default priority
			})),
			func(e error) apperror.AppError {
				return apperror.NewDatabaseError(e, "TaskRepository")
			},
		),
		func(t db.Task) task.Task {
			return task.Task{
				ID:          task.TaskID(t.ID),
				Title:       task.TaskTitle(t.Title),
				Description: task.TaskDescription(t.Description.String),
				Status:      task.TaskStatus(t.Status),
			}
		},
	)
}

func UpdateTask(q db.Querier, ctx context.Context, id task.TaskID, cmd task.TaskCmd) types.Result[task.Task, apperror.AppError] {
	// Status logic: use cmd.Status if provided, else pending? Or default logic?
	// Assuming cmd.Status is populated correctly by caller
	status := cmd.Status.String()
	if status == "" {
		status = task.TaskStatusPending.String()
	}

	// Note: priority and due_date are not yet in domain model, setting defaults or ignoring
	return types.Map(
		types.MapErr(
			types.FromPair(q.UpdateTask(ctx, db.UpdateTaskParams{
				ID:          uuid.UUID(id),
				Title:       sql.NullString{String: cmd.Title.String(), Valid: true},
				Description: sql.NullString{String: cmd.Description.String(), Valid: true},
				Status:      sql.NullString{String: status, Valid: true},
				Priority:    sql.NullString{String: "medium", Valid: true}, // Default
			})),
			func(e error) apperror.AppError {
				if errors.Is(e, sql.ErrNoRows) {
					return apperror.NewNotFoundError(e, "Task")
				}
				return apperror.NewDatabaseError(e, "TaskRepository")
			},
		),
		func(t db.Task) task.Task {
			return task.Task{
				ID:          task.TaskID(t.ID),
				Title:       task.TaskTitle(t.Title),
				Description: task.TaskDescription(t.Description.String),
				Status:      task.TaskStatus(t.Status),
			}
		},
	)
}
