package tasks

import (
	"api/src/domain/model"
	"api/src/routes/response"
	usecase "api/src/usecase/task"
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
	res := types.Pipe2(
		newGetRequest(r).validate(),
		func(req getRequest) types.Result[model.Task, model.AppError] {
			return usecase.GetTask(model.NewTaskID(req.ID))
		},
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
