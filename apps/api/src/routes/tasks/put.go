package tasks

import (
	"api/src/domain/model"
	"api/src/routes/response"
	usecase "api/src/usecase/task"
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
			func(req putRequest) types.Result[model.Task, model.AppError] {
				input := usecase.UpdateInput{
					ID:          model.NewTaskID(req.ID),
					Title:       model.TaskTitle(req.Title),
					Description: model.TaskDescription(req.Description),
					Status:      model.TaskStatus(req.Status),
				}
				return usecase.UpdateTask(
					q,
					r.Context(),
					input,
				)
			},
			func(task model.Task) putResponse {
				return putResponse{
					ID:          task.ID.String(),
					Title:       task.Title.String(),
					Description: task.Description.String(),
					Status:      task.Status.String(),
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
}
