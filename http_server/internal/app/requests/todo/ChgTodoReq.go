package todo_requests

import (
	"io"
	"net/http"
	"todo-list/internal/errors_constants"
	"todo-list/internal/utils"
)

type ChgTodoReq struct {
	Title string `json:"title"`
	Dsc   string `json:"dsc"`
}

func CreateChgTodoReq(body io.ReadCloser, headers *http.Header) (*ChgTodoReq, error) {
	var chgTodoReq ChgTodoReq
	err := utils.DecodeJSONBodyDisallow(body, headers, &chgTodoReq)
	if err != nil {
		return nil, errors_constants.ErrInvalidCreds
	}

	return &chgTodoReq, nil
}
