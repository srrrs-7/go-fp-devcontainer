package tasks

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"api/src/routes/response"
	"net/http"
	"utils/types"
)

type putResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	res := types.Map(
		types.FlatMap(
			newPutRequest(r).validate(),
			func(req putRequest) types.Result[model.Task, model.AppError] {
				return task_repository.UpdateTask(
					model.NewTaskID(req.ID),
					model.TaskTitle(req.Title),
					model.TaskDescription(req.Description),
					model.TaskCompleted(req.Completed),
				)
			},
		),
		func(task model.Task) putResponse {
			return putResponse{
				ID:          task.ID.String(),
				Title:       task.Title.String(),
				Description: task.Description.String(),
				Completed:   task.Completed.Bool(),
			}
		},
	)

	res.Match(
		func(resp putResponse) {
			response.OK(w, resp)
		},
		func(e model.AppError) {
			response.HandleAppError(w, e)
		},
	)
}
