package task_repository

import (
	"api/src/domain/model"
	"context"
	"database/sql"
	"utils/db/db"
	"utils/types"

	"github.com/google/uuid"
)

func CreateTask(q db.Querier, ctx context.Context, cmd model.TaskCmd) types.Result[model.Task, model.AppError] {
	return types.Map(
		types.MapErr(
			types.FromPair(q.CreateTask(ctx, db.CreateTaskParams{
				Title:       cmd.Title.String(),
				Description: sql.NullString{String: cmd.Description.String(), Valid: true},
				Status:      model.TaskStatusPending.String(), // Default status
				Priority:    "medium",                         // Default priority
			})),
			func(e error) model.AppError {
				return model.NewDatabaseError(e, "TaskRepository")
			},
		),
		func(t db.Task) model.Task {
			return model.Task{
				ID:          model.TaskID(t.ID),
				Title:       model.TaskTitle(t.Title),
				Description: model.TaskDescription(t.Description.String),
				Status:      model.TaskStatus(t.Status),
			}
		},
	)
}

func UpdateTask(q db.Querier, ctx context.Context, id model.TaskID, cmd model.TaskCmd) types.Result[model.Task, model.AppError] {
	// Status logic: use cmd.Status if provided, else pending? Or default logic?
	// Assuming cmd.Status is populated correctly by caller
	status := cmd.Status.String()
	if status == "" {
		status = model.TaskStatusPending.String()
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
			func(e error) model.AppError {
				return model.NewDatabaseError(e, "TaskRepository")
			},
		),
		func(t db.Task) model.Task {
			return model.Task{
				ID:          model.TaskID(t.ID),
				Title:       model.TaskTitle(t.Title),
				Description: model.TaskDescription(t.Description.String),
				Status:      model.TaskStatus(t.Status),
			}
		},
	)
}
