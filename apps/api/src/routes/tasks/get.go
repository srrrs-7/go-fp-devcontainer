package tasks

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"api/src/routes/response"
	"net/http"
	"utils/types"
)

type getResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	res := types.Map(
		types.FlatMap(
			newGetRequest(r).validate(),
			func(req getRequest) types.Result[model.Task, model.AppError] {
				return task_repository.FindTaskByID(model.NewTaskID(req.ID))
			},
		),
		func(task model.Task) getResponse {
			return getResponse{
				ID:          task.ID.String(),
				Title:       task.Title.String(),
				Description: task.Description.String(),
				Completed:   task.Completed.Bool(),
			}
		},
	)

	res.Match(
		func(resp getResponse) {
			response.OK(w, resp)
		},
		func(e model.AppError) {
			response.HandleAppError(w, e)
		},
	)
}
