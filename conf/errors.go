package conf

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	// ErrInternalServerError msg
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound msg
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict msg
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput ,sg
	ErrBadParamInput = errors.New("Given Param is not valid")
)

// GetStatusCode funcction
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
