package tasks

import (
	"api/src/domain/apperror"
	"api/src/domain/task"
	"api/src/infra/rds/task_repository"
	"api/src/routes/response"
	"net/http"
	"utils/db/db"
	"utils/types"
)

type putResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewPutHandler(q db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := types.Pipe2(
			newPutRequest(r).validate(),
			func(req putRequest) types.Result[task.Task, apperror.AppError] {
				cmd := task.TaskCmd{
					Title:       task.TaskTitle(req.Title),
					Description: task.TaskDescription(req.Description),
					Status:      task.TaskStatus(req.Status),
				}
				return task_repository.UpdateTask(q, r.Context(), task.NewTaskID(req.ID), cmd)
			},
			func(t task.Task) putResponse {
				return putResponse{
					ID:          t.ID.String(),
					Title:       t.Title.String(),
					Description: t.Description.String(),
					Status:      t.Status.String(),
				}
			},
		)

		res.Match(
			func(resp putResponse) {
				response.OK(w, resp)
			},
			func(e apperror.AppError) {
				response.HandleAppError(w, e)
			},
		)
	}
}
