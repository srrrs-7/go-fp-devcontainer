package tasks

import (
	"api/src/domain/model"
	"api/src/infra/rds/task_repository"
	"api/src/routes/response"
	"net/http"
	"utils/types"
)

type listResponse struct {
	Tasks []taskItem `json:"tasks"`
}

type taskItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	res := types.Pipe2(
		newListRequest(r).validate(),
		func(req listRequest) types.Result[[]model.Task, model.AppError] {
			return task_repository.FindAllTasks()
		},
		func(tasks []model.Task) listResponse {
			items := make([]taskItem, len(tasks))
			for i, task := range tasks {
				items[i] = taskItem{
					ID:          task.ID.String(),
					Title:       task.Title.String(),
					Description: task.Description.String(),
					Completed:   task.Completed.Bool(),
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
