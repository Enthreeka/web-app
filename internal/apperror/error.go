package apperror

import "encoding/json"

var (
	ErrNotFound = NewAppError(nil, "not found", "", "us-000003")
	NoAuthErr   = NewAppError(nil, "", "", "")
)

type AppError struct {
	Err              error  `json:"error"`
	Massage          string `json:"massage"`
	DeveloperMassage string `json:"developerMassage"`
	Code             string `json:"code"`
}

func (e *AppError) Error() string {
	return e.Massage
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, Massage, DeveloperMassage, Code string) *AppError {
	return &AppError{
		Err:              err,
		Massage:          Massage,
		DeveloperMassage: DeveloperMassage,
		Code:             Code,
	}
}

func systemError(err error) *AppError {
	return NewAppError(err, "internal system error", err.Error(), "us-000000")
}
