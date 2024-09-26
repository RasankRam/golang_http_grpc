package auth_requests

import (
	"io"
	"net/http"
	"todo-list/internal/errors_constants"
	"todo-list/internal/utils"
)

type LogoutReq struct {
	Token string
}

func CreateLogoutReq(body io.ReadCloser, headers *http.Header) (*LogoutReq, error) {

	var logReq LogoutReq
	err := utils.DecodeJSONBody(body, headers, &logReq)
	if err != nil {
		return nil, errors_constants.ErrInvalidCreds
	}

	return &logReq, nil
}
