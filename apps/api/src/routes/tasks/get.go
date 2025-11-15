package tasks

import (
	"api/src/domain/model"
	"api/src/infra/rds_repository"
	"api/src/routes"
	"utils/types"
)

type response struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func Handler() {
	res := types.Map(
		rds_repository.FindTaskByID(
			model.NewTaskID("550e8400-e29b-41d4-a716-446655440000"),
		),
		func(t model.Task) response {
			return response{
				ID:          t.ID.String(),
				Title:       t.Title.String(),
				Description: t.Description.String(),
				Completed:   t.Completed.Bool(),
			}
		},
	)

	res.Match(
		func(t response) {
			routes.ResponseOK()
		},
		func(e model.DatabaseError) {
			errName := e.ErrorName()
			switch errName {
			case model.BadRequestErrorName:
				routes.ResponseBadRequest()
			case model.NotFoundErrorName:
				routes.ResponseNotFound()
			case model.UnauthorizedErrorName:
				routes.ResponseUnauthorized()
			default:
				routes.ResponseInternalError()
			}
		},
	)
}
