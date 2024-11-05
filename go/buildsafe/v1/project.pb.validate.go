// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: buildsafe/v1/project.proto

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

// Validate checks the field values on CreateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProjectRequestMultiError, or nil if none found.
func (m *CreateProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectName

	if len(errors) > 0 {
		return CreateProjectRequestMultiError(errors)
	}

	return nil
}

// CreateProjectRequestMultiError is an error wrapping multiple validation
// errors returned by CreateProjectRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProjectRequestMultiError) AllErrors() []error { return m }

// CreateProjectRequestValidationError is the validation error returned by
// CreateProjectRequest.Validate if the designated constraints aren't met.
type CreateProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProjectRequestValidationError) ErrorName() string {
	return "CreateProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProjectRequestValidationError{}

// Validate checks the field values on CreateProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProjectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProjectResponseMultiError, or nil if none found.
func (m *CreateProjectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProjectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectId

	if len(errors) > 0 {
		return CreateProjectResponseMultiError(errors)
	}

	return nil
}

// CreateProjectResponseMultiError is an error wrapping multiple validation
// errors returned by CreateProjectResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateProjectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProjectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProjectResponseMultiError) AllErrors() []error { return m }

// CreateProjectResponseValidationError is the validation error returned by
// CreateProjectResponse.Validate if the designated constraints aren't met.
type CreateProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProjectResponseValidationError) ErrorName() string {
	return "CreateProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProjectResponseValidationError{}

// Validate checks the field values on GetProjectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProjectRequestMultiError, or nil if none found.
func (m *GetProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectId

	if len(errors) > 0 {
		return GetProjectRequestMultiError(errors)
	}

	return nil
}

// GetProjectRequestMultiError is an error wrapping multiple validation errors
// returned by GetProjectRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProjectRequestMultiError) AllErrors() []error { return m }

// GetProjectRequestValidationError is the validation error returned by
// GetProjectRequest.Validate if the designated constraints aren't met.
type GetProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectRequestValidationError) ErrorName() string {
	return "GetProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectRequestValidationError{}

// Validate checks the field values on GetProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProjectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProjectResponseMultiError, or nil if none found.
func (m *GetProjectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProjectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProjectResponseValidationError{
					field:  "Project",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProjectResponseValidationError{
					field:  "Project",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProjectResponseValidationError{
				field:  "Project",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetProjectResponseMultiError(errors)
	}

	return nil
}

// GetProjectResponseMultiError is an error wrapping multiple validation errors
// returned by GetProjectResponse.ValidateAll() if the designated constraints
// aren't met.
type GetProjectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProjectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProjectResponseMultiError) AllErrors() []error { return m }

// GetProjectResponseValidationError is the validation error returned by
// GetProjectResponse.Validate if the designated constraints aren't met.
type GetProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectResponseValidationError) ErrorName() string {
	return "GetProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectResponseValidationError{}

// Validate checks the field values on ListProjectsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProjectsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProjectsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProjectsRequestMultiError, or nil if none found.
func (m *ListProjectsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProjectsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListProjectsRequestMultiError(errors)
	}

	return nil
}

// ListProjectsRequestMultiError is an error wrapping multiple validation
// errors returned by ListProjectsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListProjectsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProjectsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProjectsRequestMultiError) AllErrors() []error { return m }

// ListProjectsRequestValidationError is the validation error returned by
// ListProjectsRequest.Validate if the designated constraints aren't met.
type ListProjectsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectsRequestValidationError) ErrorName() string {
	return "ListProjectsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProjectsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectsRequestValidationError{}

// Validate checks the field values on Project with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Project) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Project with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProjectMultiError, or nil if none found.
func (m *Project) ValidateAll() error {
	return m.validate(true)
}

func (m *Project) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectId

	// no validation rules for ProjectName

	if len(errors) > 0 {
		return ProjectMultiError(errors)
	}

	return nil
}

// ProjectMultiError is an error wrapping multiple validation errors returned
// by Project.ValidateAll() if the designated constraints aren't met.
type ProjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProjectMultiError) AllErrors() []error { return m }

// ProjectValidationError is the validation error returned by Project.Validate
// if the designated constraints aren't met.
type ProjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectValidationError) ErrorName() string { return "ProjectValidationError" }

// Error satisfies the builtin error interface
func (e ProjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectValidationError{}

// Validate checks the field values on ListProjectsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProjectsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProjectsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProjectsResponseMultiError, or nil if none found.
func (m *ListProjectsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProjectsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListProjectsResponseValidationError{
						field:  fmt.Sprintf("Projects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListProjectsResponseValidationError{
						field:  fmt.Sprintf("Projects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProjectsResponseValidationError{
					field:  fmt.Sprintf("Projects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListProjectsResponseMultiError(errors)
	}

	return nil
}

// ListProjectsResponseMultiError is an error wrapping multiple validation
// errors returned by ListProjectsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListProjectsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProjectsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProjectsResponseMultiError) AllErrors() []error { return m }

// ListProjectsResponseValidationError is the validation error returned by
// ListProjectsResponse.Validate if the designated constraints aren't met.
type ListProjectsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectsResponseValidationError) ErrorName() string {
	return "ListProjectsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListProjectsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectsResponseValidationError{}

// Validate checks the field values on RenameProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RenameProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RenameProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RenameProjectRequestMultiError, or nil if none found.
func (m *RenameProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RenameProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectId

	// no validation rules for ProjectName

	if len(errors) > 0 {
		return RenameProjectRequestMultiError(errors)
	}

	return nil
}

// RenameProjectRequestMultiError is an error wrapping multiple validation
// errors returned by RenameProjectRequest.ValidateAll() if the designated
// constraints aren't met.
type RenameProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RenameProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RenameProjectRequestMultiError) AllErrors() []error { return m }

// RenameProjectRequestValidationError is the validation error returned by
// RenameProjectRequest.Validate if the designated constraints aren't met.
type RenameProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RenameProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RenameProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RenameProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RenameProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RenameProjectRequestValidationError) ErrorName() string {
	return "RenameProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RenameProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRenameProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RenameProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RenameProjectRequestValidationError{}

// Validate checks the field values on RenameProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RenameProjectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RenameProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RenameProjectResponseMultiError, or nil if none found.
func (m *RenameProjectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RenameProjectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RenameProjectResponseMultiError(errors)
	}

	return nil
}

// RenameProjectResponseMultiError is an error wrapping multiple validation
// errors returned by RenameProjectResponse.ValidateAll() if the designated
// constraints aren't met.
type RenameProjectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RenameProjectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RenameProjectResponseMultiError) AllErrors() []error { return m }

// RenameProjectResponseValidationError is the validation error returned by
// RenameProjectResponse.Validate if the designated constraints aren't met.
type RenameProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RenameProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RenameProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RenameProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RenameProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RenameProjectResponseValidationError) ErrorName() string {
	return "RenameProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RenameProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRenameProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RenameProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RenameProjectResponseValidationError{}

// Validate checks the field values on DeleteProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProjectRequestMultiError, or nil if none found.
func (m *DeleteProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectId

	if len(errors) > 0 {
		return DeleteProjectRequestMultiError(errors)
	}

	return nil
}

// DeleteProjectRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteProjectRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProjectRequestMultiError) AllErrors() []error { return m }

// DeleteProjectRequestValidationError is the validation error returned by
// DeleteProjectRequest.Validate if the designated constraints aren't met.
type DeleteProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProjectRequestValidationError) ErrorName() string {
	return "DeleteProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProjectRequestValidationError{}

// Validate checks the field values on DeleteProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProjectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProjectResponseMultiError, or nil if none found.
func (m *DeleteProjectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProjectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteProjectResponseMultiError(errors)
	}

	return nil
}

// DeleteProjectResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteProjectResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteProjectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProjectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProjectResponseMultiError) AllErrors() []error { return m }

// DeleteProjectResponseValidationError is the validation error returned by
// DeleteProjectResponse.Validate if the designated constraints aren't met.
type DeleteProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProjectResponseValidationError) ErrorName() string {
	return "DeleteProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProjectResponseValidationError{}
