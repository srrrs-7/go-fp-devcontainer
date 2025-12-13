package tasks

import (
	"api/src/domain/model"
	"api/src/routes/response"
	usecase "api/src/usecase/task"
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
			func(req postRequest) types.Result[model.Task, model.AppError] {
				input := usecase.CreateInput{
					Title:       model.TaskTitle(req.Title),
					Description: model.TaskDescription(req.Description),
				}
				return usecase.CreateTask(q, r.Context(), input)
			},
			func(task model.Task) postResponse {
				return postResponse{
					ID:          task.ID.String(),
					Title:       task.Title.String(),
					Description: task.Description.String(),
					Status:      task.Status.String(),
				}
			},
		)

		res.Match(
			func(resp postResponse) {
				response.Created(w, resp)
			},
			func(e model.AppError) {
				response.HandleAppError(w, e)
			},
		)
	}
}
