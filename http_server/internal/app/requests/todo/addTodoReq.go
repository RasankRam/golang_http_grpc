package todo_requests

import (
	"io"
	"net/http"
	"todo-list/internal/errors_constants"
	"todo-list/internal/utils"
)

type AddTodoReq struct {
	Title string `json:"title"`
	Dsc   string `json:"dsc"`
}

func CreateAddTodoReq(body io.ReadCloser, headers *http.Header) (*AddTodoReq, error) {
	var addTodoReq AddTodoReq
	err := utils.DecodeJSONBodyDisallow(body, headers, &addTodoReq)
	if err != nil {
		return nil, errors_constants.ErrInvalidCreds
	}

	return &addTodoReq, nil
}
