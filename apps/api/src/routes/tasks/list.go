package tasks

import (
	"api/src/domain/model"
	"api/src/routes/response"
	usecase "api/src/usecase/task"
	"net/http"
	"utils/db/db"
	"utils/types"
)

type listResponse struct {
	Tasks []taskItem `json:"tasks"`
}

type taskItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewListHandler(q db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := types.Pipe2(
			newListRequest(r).validate(),
			func(req listRequest) types.Result[[]model.Task, model.AppError] {
				input := usecase.ListInput{}
				return usecase.ListTasks(q, r.Context(), input)
			},
			func(tasks []model.Task) listResponse {
				items := make([]taskItem, len(tasks))
				for i, task := range tasks {
					items[i] = taskItem{
						ID:          task.ID.String(),
						Title:       task.Title.String(),
						Description: task.Description.String(),
						Status:      task.Status.String(),
					}
				}
				return listResponse{Tasks: items}
			},
		)

		res.Match(
			func(resp listResponse) {
				response.OK(w, resp)
			},
			func(e model.AppError) {
				response.HandleAppError(w, e)
			},
		)
	}
}
