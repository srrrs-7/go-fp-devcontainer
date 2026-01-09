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

type postResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewPostHandler(q db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := types.Pipe2(
			newPostRequest(r).validate(),
			func(req postRequest) types.Result[task.Task, apperror.AppError] {
				cmd := task.TaskCmd{
					Title:       task.TaskTitle(req.Title),
					Description: task.TaskDescription(req.Description),
					Status:      task.TaskStatusPending,
				}
				return task_repository.CreateTask(q, r.Context(), cmd)
			},
			func(t task.Task) postResponse {
				return postResponse{
					ID:          t.ID.String(),
					Title:       t.Title.String(),
					Description: t.Description.String(),
					Status:      t.Status.String(),
				}
			},
		)

		res.Match(
			func(resp postResponse) {
				response.Created(w, resp)
			},
			func(e apperror.AppError) {
				response.HandleAppError(w, e)
			},
		)
	}
}
