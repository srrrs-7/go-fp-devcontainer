package model

const (
	NotFoundErrorName       = "NotFoundError"
	ValidationErrorName     = "ValidationError"
	UnauthorizedErrorName   = "UnauthorizedError"
	InternalServerErrorName = "InternalServerError"
	BadRequestErrorName     = "BadRequestError"
	ConflictErrorName       = "ConflictError"
	ForbiddenErrorName      = "ForbiddenError"
	DatabaseErrorName       = "DatabaseError"
)

type baseErr struct {
	errName    string
	domainName string
	err        error
}

func (e baseErr) ErrorName() string {
	return e.errName
}

func (e baseErr) DomainName() string {
	return e.domainName
}

func (e baseErr) Error() error {
	return e.err
}

// NotFoundError
type NotFoundError struct {
	baseErr
}

func NewNotFoundError(err error, dName string) NotFoundError {
	return NotFoundError{
		baseErr: baseErr{
			errName:    NotFoundErrorName,
			domainName: dName,
			err:        err,
		},
	}
}

// ValidationError
type ValidationError struct {
	baseErr
}

func NewValidationError(err error, dName string) ValidationError {
	return ValidationError{
		baseErr: baseErr{
			errName:    ValidationErrorName,
			domainName: dName,
			err:        err,
		},
	}
}

// UnauthorizedError
type UnauthorizedError struct {
	baseErr
}

func NewUnauthorizedError(err error, dName string) UnauthorizedError {
	return UnauthorizedError{
		baseErr: baseErr{
			errName:    UnauthorizedErrorName,
			domainName: dName,
			err:        err,
		},
	}
}

// InternalServerError
type InternalServerError struct {
	baseErr
}

func NewInternalServerError(err error, dName string) InternalServerError {
	return InternalServerError{
		baseErr: baseErr{
			errName:    InternalServerErrorName,
			domainName: dName,
			err:        err,
		},
	}
}

// BadRequestError
type BadRequestError struct {
	baseErr
}

func NewBadRequestError(err error, dName string) BadRequestError {
	return BadRequestError{
		baseErr: baseErr{
			errName:    BadRequestErrorName,
			domainName: dName,
			err:        err,
		},
	}
}

// ConflictError
type ConflictError struct {
	baseErr
}

func NewConflictError(err error, dName string) ConflictError {
	return ConflictError{
		baseErr: baseErr{
			errName:    ConflictErrorName,
			domainName: dName,
			err:        err,
		},
	}
}

// ForbiddenError
type ForbiddenError struct {
	baseErr
}

func NewForbiddenError(err error, dName string) ForbiddenError {
	return ForbiddenError{
		baseErr: baseErr{
			errName:    ForbiddenErrorName,
			domainName: dName,
			err:        err,
		},
	}
}

type DatabaseError struct {
	baseErr
}

func NewDatabaseError(err error, dName string) DatabaseError {
	return DatabaseError{
		baseErr: baseErr{
			errName:    DatabaseErrorName,
			domainName: dName,
			err:        err,
		},
	}
}
