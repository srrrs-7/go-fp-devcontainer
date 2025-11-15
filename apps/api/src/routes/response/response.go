package response

import (
	"api/src/domain/model"
	"net/http"
)

func OK(w http.ResponseWriter) {}

func Created(w http.ResponseWriter) {}

func NoContent(w http.ResponseWriter) {}

// handleAppError - AppErrorを網羅的に処理し、適切なHTTPレスポンスを返す
func HandleAppError(w http.ResponseWriter, err model.AppError) {
	errName := err.ErrorName()

	switch errName {
	case model.ValidationErrorName:
		badRequest(w)
	case model.NotFoundErrorName:
		notFound(w)
	case model.UnauthorizedErrorName:
		unauthorized(w)
	case model.ForbiddenErrorName:
		forbidden(w)
	case model.BadRequestErrorName:
		badRequest(w)
	case model.ConflictErrorName:
		conflict(w)
	case model.DatabaseErrorName:
		internalError(w)
	case model.InternalServerErrorName:
		internalError(w)
	default:
		unexpectedError(w)
	}
}

func badRequest(w http.ResponseWriter) {}

func notFound(w http.ResponseWriter) {}

func unauthorized(w http.ResponseWriter) {}

func internalError(w http.ResponseWriter) {}

func forbidden(w http.ResponseWriter) {}

func conflict(w http.ResponseWriter) {}

func unexpectedError(w http.ResponseWriter) {}
