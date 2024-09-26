package auth_requests

import (
	"io"
	"net/http"
	"todo-list/internal/utils"
)

type RegisterReq struct {
	Login    string `json:"login" validate:"required,gte=3,lte=255"`
	Password string `json:"password" validate:"required,gte=6,lte=255"`
	Role     string `json:"role" validate:"required"`
}

func CreateRegisterReq(body io.ReadCloser, headers *http.Header) (*RegisterReq, error) {

	var regReq RegisterReq
	err := utils.DecodeJSONBodyDisallow(body, headers, &regReq)
	if err != nil {
		return nil, err
	}

	return &regReq, nil
}
