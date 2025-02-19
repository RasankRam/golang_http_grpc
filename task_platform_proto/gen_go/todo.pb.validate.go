// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: todo.proto

package gen

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on TodoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TodoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TodoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TodoRequestMultiError, or
// nil if none found.
func (m *TodoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TodoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 6 {
		err := TodoRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPrice() <= 0 {
		err := TodoRequestValidationError{
			field:  "Price",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return TodoRequestMultiError(errors)
	}

	return nil
}

// TodoRequestMultiError is an error wrapping multiple validation errors
// returned by TodoRequest.ValidateAll() if the designated constraints aren't met.
type TodoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TodoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TodoRequestMultiError) AllErrors() []error { return m }

// TodoRequestValidationError is the validation error returned by
// TodoRequest.Validate if the designated constraints aren't met.
type TodoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TodoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TodoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TodoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TodoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TodoRequestValidationError) ErrorName() string { return "TodoRequestValidationError" }

// Error satisfies the builtin error interface
func (e TodoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTodoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TodoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TodoRequestValidationError{}

// Validate checks the field values on TodoResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TodoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TodoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TodoResponseMultiError, or
// nil if none found.
func (m *TodoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TodoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DiscountedPrice

	if len(errors) > 0 {
		return TodoResponseMultiError(errors)
	}

	return nil
}

// TodoResponseMultiError is an error wrapping multiple validation errors
// returned by TodoResponse.ValidateAll() if the designated constraints aren't met.
type TodoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TodoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TodoResponseMultiError) AllErrors() []error { return m }

// TodoResponseValidationError is the validation error returned by
// TodoResponse.Validate if the designated constraints aren't met.
type TodoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TodoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TodoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TodoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TodoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TodoResponseValidationError) ErrorName() string { return "TodoResponseValidationError" }

// Error satisfies the builtin error interface
func (e TodoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTodoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TodoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TodoResponseValidationError{}
