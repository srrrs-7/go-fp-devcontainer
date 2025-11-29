package response

import (
	"api/src/domain/model"
	"net/http"
)

func OK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func Created(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created"))
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("NoContent"))
}

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

func badRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("BadRequest"))
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("NotFound"))
}

func unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Unauthorized"))
}

func internalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("InternalError"))
}

func forbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Forbidden"))
}

func conflict(w http.ResponseWriter) {
	w.WriteHeader(http.StatusConflict)
	w.Write([]byte("Conflict"))
}

func unexpectedError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("UnexpectedError"))
}
