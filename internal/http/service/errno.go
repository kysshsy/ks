package service

import (
	"net/http"

	"ks/pkg/errors"
)

var (
	UserNotFoundErr = errors.NewErr(http.StatusOK, 1, "user not found")
)
