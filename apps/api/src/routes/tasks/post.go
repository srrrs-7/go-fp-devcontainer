package tasks

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"api/src/routes/response"
	"net/http"
	"utils/types"
)

type postResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	res := types.Pipe2(
		newPostRequest(r).validate(),
		func(req postRequest) types.Result[model.Task, model.AppError] {
			return task_repository.CreateTask(
				model.TaskTitle(req.Title),
				model.TaskDescription(req.Description),
			)
		},
		func(task model.Task) postResponse {
			return postResponse{
				ID:          task.ID.String(),
				Title:       task.Title.String(),
				Description: task.Description.String(),
				Completed:   task.Completed.Bool(),
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
