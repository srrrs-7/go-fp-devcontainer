package tasks

import (
	"api/src/domain/model"
	"api/src/routes/response"
	usecase "api/src/usecase/task"
	"net/http"
	"utils/db/db"
	"utils/types"
)

type getResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewGetHandler(q db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := types.Pipe2(
			newGetRequest(r).validate(),
			func(req getRequest) types.Result[model.Task, model.AppError] {
				input := usecase.GetInput{
					ID: model.NewTaskID(req.ID),
				}
				return usecase.GetTask(q, r.Context(), input)
			},
			func(task model.Task) getResponse {
				return getResponse{
					ID:          task.ID.String(),
					Title:       task.Title.String(),
					Description: task.Description.String(),
					Status:      task.Status.String(),
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
}
