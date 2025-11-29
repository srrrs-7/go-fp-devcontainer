package response

import (
	"api/src/domain/model"
	"encoding/json"
	"net/http"
)

func OK(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Created(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Accepted(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func NoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// handleAppError - AppErrorを網羅的に処理し、適切なHTTPレスポンスを返す
func HandleAppError(w http.ResponseWriter, err model.AppError) {
	errName := err.ErrorName()

	switch errName {
	case model.ValidationErrorName:
		badRequest(w, err)
	case model.NotFoundErrorName:
		notFound(w, err)
	case model.UnauthorizedErrorName:
		unauthorized(w, err)
	case model.ForbiddenErrorName:
		forbidden(w, err)
	case model.BadRequestErrorName:
		badRequest(w, err)
	case model.ConflictErrorName:
		conflict(w, err)
	case model.DatabaseErrorName:
		internalError(w, err)
	case model.InternalServerErrorName:
		internalError(w, err)
	default:
		unexpectedError(w, err)
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Domain  string `json:"domain"`
}

func newErrorResponse(err model.AppError) ErrorResponse {
	return ErrorResponse{
		Message: err.Error(),
		Type:    err.ErrorName(),
		Domain:  err.DomainName(),
	}
}

func writeError(w http.ResponseWriter, status int, err model.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(newErrorResponse(err)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func badRequest(w http.ResponseWriter, err model.AppError) {
	writeError(w, http.StatusBadRequest, err)
}

func notFound(w http.ResponseWriter, err model.AppError) {
	writeError(w, http.StatusNotFound, err)
}

func unauthorized(w http.ResponseWriter, err model.AppError) {
	writeError(w, http.StatusUnauthorized, err)
}

func internalError(w http.ResponseWriter, err model.AppError) {
	writeError(w, http.StatusInternalServerError, err)
}

func forbidden(w http.ResponseWriter, err model.AppError) {
	writeError(w, http.StatusForbidden, err)
}

func conflict(w http.ResponseWriter, err model.AppError) {
	writeError(w, http.StatusConflict, err)
}

func unexpectedError(w http.ResponseWriter, err model.AppError) {
	writeError(w, http.StatusInternalServerError, err)
}
