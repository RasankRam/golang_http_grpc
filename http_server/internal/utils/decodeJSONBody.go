package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strings"
	"todo-list/internal/validate"
)

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

type validatorErrors struct {
	errorsList []string
}

func (mr *validatorErrors) Error() string {
	return strings.Join(mr.errorsList, ";")
}

func DecodeJSONBody[T any](reqBody io.ReadCloser, headers *http.Header, dst *T) error {
	return decodeJSONBodyFn(reqBody, headers, dst, false, true)
}

func DecodeJSONBodyDisallow[T any](reqBody io.ReadCloser, headers *http.Header, dst *T) error {
	return decodeJSONBodyFn(reqBody, headers, dst, true, true)
}

func decodeJSONBodyFn[T any](reqBody io.ReadCloser, headers *http.Header, dst *T, disallowUnknownFields bool, allowEmptyBody bool) error {
	ct := headers.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}

	dec := json.NewDecoder(reqBody)
	fmt.Println("disallowUnknownFields", disallowUnknownFields)
	if disallowUnknownFields {
		dec.DisallowUnknownFields()
	}

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			if allowEmptyBody {
				return nil
			} else {
				msg := "Request body must not be empty"
				return &malformedRequest{status: http.StatusBadRequest, msg: msg}
			}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	if err := validate.Validator.Struct(dst); err != nil {
		var errorsList []string
		for _, e := range err.(validator.ValidationErrors) {
			errorsList = append(errorsList, e.Translate(validate.Trans))
		}
		return &validatorErrors{errorsList: errorsList}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	fmt.Println("dst", dst)

	return nil
}
