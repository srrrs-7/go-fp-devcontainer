package tasks

import (
	apperror "api/src/domain/error"
	"net/http"
	"utils/types"

	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
)

type getRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

func newGetRequest(r *http.Request) getRequest {
	return getRequest{
		ID: r.URL.Query().Get("id"),
	}
}

func (r getRequest) validate() types.Result[getRequest, apperror.AppError] {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[getRequest, apperror.AppError](
			apperror.NewValidationError(err, "GetRequest"),
		)
	}
	return types.Ok[getRequest, apperror.AppError](r)
}

type listRequest struct {
	ID          string `json:"id" validate:"omitempty,uuid"`
	Title       string `json:"title" validate:"omitempty,min=3,max=100"`
	Description string `json:"description" validate:"omitempty,max=500"`
	Status      string `json:"status" validate:"omitempty"`
}

func newListRequest(r *http.Request) listRequest {
	return listRequest{
		ID:          r.URL.Query().Get("id"),
		Title:       r.URL.Query().Get("title"),
		Description: r.URL.Query().Get("description"),
		Status:      r.URL.Query().Get("status"),
	}
}

func (r listRequest) validate() types.Result[listRequest, apperror.AppError] {
	sanitize := bluemonday.StrictPolicy()
	r.Title = sanitize.Sanitize(r.Title)
	r.Description = sanitize.Sanitize(r.Description)

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[listRequest, apperror.AppError](
			apperror.NewValidationError(err, "listRequest"),
		)
	}
	return types.Ok[listRequest, apperror.AppError](r)
}

type postRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
}

func newPostRequest(r *http.Request) postRequest {
	return postRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}
}

func (r postRequest) validate() types.Result[postRequest, apperror.AppError] {
	sanitize := bluemonday.StrictPolicy()
	r.Title = sanitize.Sanitize(r.Title)
	r.Description = sanitize.Sanitize(r.Description)

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[postRequest, apperror.AppError](
			apperror.NewValidationError(err, "postRequest"),
		)
	}
	return types.Ok[postRequest, apperror.AppError](r)
}

type putRequest struct {
	ID          string `json:"id" validate:"required,uuid"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
	Status      string `json:"status"`
}

func newPutRequest(r *http.Request) putRequest {
	return putRequest{
		ID:          r.FormValue("id"),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Status:      r.FormValue("status"),
	}
}

func (r putRequest) validate() types.Result[putRequest, apperror.AppError] {
	sanitize := bluemonday.StrictPolicy()
	r.Title = sanitize.Sanitize(r.Title)
	r.Description = sanitize.Sanitize(r.Description)

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[putRequest, apperror.AppError](
			apperror.NewValidationError(err, "putRequest"),
		)
	}
	return types.Ok[putRequest, apperror.AppError](r)
}
