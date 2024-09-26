package auth_requests

import (
	"io"
	"net/http"
	"todo-list/internal/utils"
)

type LoginReq struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CreateLoginReq(body io.ReadCloser, headers *http.Header) (*LoginReq, error) {

	var loginReq LoginReq
	err := utils.DecodeJSONBodyDisallow(body, headers, &loginReq)
	if err != nil {
		return nil, err
	}

	return &loginReq, nil
}
