package error

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	Error() string

	ToJson() string
	Wrap(error) Error
}

var _ Error = (*Err)(nil)

type Err struct {
	HttpCode     int
	BusinessCode int
	Message      string
	CauseErr     error
}

func NewErr(httpCode int, code int, msg string) Error {
	return &Err{
		HttpCode:     httpCode,
		BusinessCode: code,
		Message:      msg,
	}
}

func (e *Err) Error() string {
	return fmt.Sprintf("%s :%s", e.Message, e.CauseErr.Error())
}

func (e *Err) ToJson() string {
	err := &struct {
		HttpCode     int    `json:"http_code"`
		BusinessCode int    `json:"business_code"`
		Message      string `json:"message"`
	}{
		HttpCode:     e.HttpCode,
		BusinessCode: e.BusinessCode,
		Message:      e.Error(),
	}

	raw, _ := json.Marshal(err)

	return string(raw)
}

func (e *Err) Wrap(err error) Error {
	e.CauseErr = err
	return e
}
