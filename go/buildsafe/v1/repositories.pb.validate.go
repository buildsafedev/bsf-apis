// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: buildsafe/v1/repositories.proto

package bsfv1

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

// Validate checks the field values on ListRepositoriesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListRepositoriesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRepositoriesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRepositoriesRequestMultiError, or nil if none found.
func (m *ListRepositoriesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRepositoriesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for InstallationId

	if len(errors) > 0 {
		return ListRepositoriesRequestMultiError(errors)
	}

	return nil
}

// ListRepositoriesRequestMultiError is an error wrapping multiple validation
// errors returned by ListRepositoriesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListRepositoriesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRepositoriesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRepositoriesRequestMultiError) AllErrors() []error { return m }

// ListRepositoriesRequestValidationError is the validation error returned by
// ListRepositoriesRequest.Validate if the designated constraints aren't met.
type ListRepositoriesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRepositoriesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRepositoriesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRepositoriesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRepositoriesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRepositoriesRequestValidationError) ErrorName() string {
	return "ListRepositoriesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRepositoriesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRepositoriesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRepositoriesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRepositoriesRequestValidationError{}

// Validate checks the field values on ListRepositoriesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListRepositoriesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRepositoriesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRepositoriesResponseMultiError, or nil if none found.
func (m *ListRepositoriesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRepositoriesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRepositories() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListRepositoriesResponseValidationError{
						field:  fmt.Sprintf("Repositories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListRepositoriesResponseValidationError{
						field:  fmt.Sprintf("Repositories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRepositoriesResponseValidationError{
					field:  fmt.Sprintf("Repositories[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListRepositoriesResponseMultiError(errors)
	}

	return nil
}

// ListRepositoriesResponseMultiError is an error wrapping multiple validation
// errors returned by ListRepositoriesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListRepositoriesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRepositoriesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRepositoriesResponseMultiError) AllErrors() []error { return m }

// ListRepositoriesResponseValidationError is the validation error returned by
// ListRepositoriesResponse.Validate if the designated constraints aren't met.
type ListRepositoriesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRepositoriesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRepositoriesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRepositoriesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRepositoriesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRepositoriesResponseValidationError) ErrorName() string {
	return "ListRepositoriesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRepositoriesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRepositoriesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRepositoriesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRepositoriesResponseValidationError{}

// Validate checks the field values on Repository with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Repository) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Repository with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RepositoryMultiError, or
// nil if none found.
func (m *Repository) ValidateAll() error {
	return m.validate(true)
}

func (m *Repository) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Url

	if len(errors) > 0 {
		return RepositoryMultiError(errors)
	}

	return nil
}

// RepositoryMultiError is an error wrapping multiple validation errors
// returned by Repository.ValidateAll() if the designated constraints aren't met.
type RepositoryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RepositoryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RepositoryMultiError) AllErrors() []error { return m }

// RepositoryValidationError is the validation error returned by
// Repository.Validate if the designated constraints aren't met.
type RepositoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RepositoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RepositoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RepositoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RepositoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RepositoryValidationError) ErrorName() string { return "RepositoryValidationError" }

// Error satisfies the builtin error interface
func (e RepositoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRepository.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RepositoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RepositoryValidationError{}
