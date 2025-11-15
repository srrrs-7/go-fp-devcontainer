package tasks

import (
	"api/src/domain/model"
	"utils/types"

	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
)

type GetRequest struct {
	ID string `json:"id" validate:"required,uuid4"`
}

func (r GetRequest) Validate() types.Result[GetRequest, model.ValidationError] {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[GetRequest, model.ValidationError](
			model.NewValidationError(err, "GetRequest"),
		)
	}
	return types.Ok[GetRequest, model.ValidationError](r)
}

type ListRequest struct {
	ID          string `json:"id" validate:"required,uuid4"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
	Completed   bool   `json:"completed"`
}

func (r ListRequest) Validate() types.Result[ListRequest, model.ValidationError] {
	sanitize := bluemonday.StrictPolicy()
	r.Title = sanitize.Sanitize(r.Title)
	r.Description = sanitize.Sanitize(r.Description)

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[ListRequest, model.ValidationError](
			model.NewValidationError(err, "ListRequest"),
		)
	}
	return types.Ok[ListRequest, model.ValidationError](r)
}

type PostRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
}

func (r PostRequest) Validate() types.Result[PostRequest, model.ValidationError] {
	sanitize := bluemonday.StrictPolicy()
	r.Title = sanitize.Sanitize(r.Title)
	r.Description = sanitize.Sanitize(r.Description)

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[PostRequest, model.ValidationError](
			model.NewValidationError(err, "PostRequest"),
		)
	}
	return types.Ok[PostRequest, model.ValidationError](r)
}

type PutRequest struct {
	ID          string `json:"id" validate:"required,uuid4"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=500"`
	Completed   bool   `json:"completed"`
}

func (r PutRequest) Validate() types.Result[PutRequest, model.ValidationError] {
	sanitize := bluemonday.StrictPolicy()
	r.Title = sanitize.Sanitize(r.Title)
	r.Description = sanitize.Sanitize(r.Description)

	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return types.Err[PutRequest, model.ValidationError](
			model.NewValidationError(err, "PutRequest"),
		)
	}
	return types.Ok[PutRequest, model.ValidationError](r)
}
