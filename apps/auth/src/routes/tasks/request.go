package tasks

import (
	"api/src/domain/model"
	"net/http"
	"utils/types"

	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
)

type getRequest struct {
	ID string `json:"id" validate:"required,uuid4"`
}

func newGetRequest(r *http.Request) getRequest {
	return getRequest{
		ID: r.URL.Query().Get("id"),
	}
}

func (r getRequest) validate() types.Result[getRequest, model.AppError] {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[getRequest, model.AppError](
			model.NewValidationError(err, "GetRequest"),
		)
	}
	return types.Ok[getRequest, model.AppError](r)
}

type listRequest struct {
	ID          string `json:"id" validate:"required,uuid4"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
	Completed   bool   `json:"completed"`
}

func newListRequest(r *http.Request) listRequest {
	return listRequest{
		ID:          r.URL.Query().Get("id"),
		Title:       r.URL.Query().Get("title"),
		Description: r.URL.Query().Get("description"),
		Completed:   r.URL.Query().Get("completed") == "true",
	}
}

func (r listRequest) validate() types.Result[listRequest, model.AppError] {
	sanitize := bluemonday.StrictPolicy()
	r.Title = sanitize.Sanitize(r.Title)
	r.Description = sanitize.Sanitize(r.Description)

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[listRequest, model.AppError](
			model.NewValidationError(err, "listRequest"),
		)
	}
	return types.Ok[listRequest, model.AppError](r)
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

func (r postRequest) validate() types.Result[postRequest, model.AppError] {
	sanitize := bluemonday.StrictPolicy()
	r.Title = sanitize.Sanitize(r.Title)
	r.Description = sanitize.Sanitize(r.Description)

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[postRequest, model.AppError](
			model.NewValidationError(err, "postRequest"),
		)
	}
	return types.Ok[postRequest, model.AppError](r)
}

type putRequest struct {
	ID          string `json:"id" validate:"required,uuid4"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
	Completed   bool   `json:"completed"`
}

func newPutRequest(r *http.Request) putRequest {
	return putRequest{
		ID:          r.FormValue("id"),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Completed:   r.FormValue("completed") == "true",
	}
}

func (r putRequest) validate() types.Result[putRequest, model.AppError] {
	sanitize := bluemonday.StrictPolicy()
	r.Title = sanitize.Sanitize(r.Title)
	r.Description = sanitize.Sanitize(r.Description)

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[putRequest, model.AppError](
			model.NewValidationError(err, "putRequest"),
		)
	}
	return types.Ok[putRequest, model.AppError](r)
}
