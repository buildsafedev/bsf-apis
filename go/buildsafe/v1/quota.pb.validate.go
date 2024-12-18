// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: buildsafe/v1/quota.proto

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

// Validate checks the field values on ViewQuotaRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ViewQuotaRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ViewQuotaRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ViewQuotaRequestMultiError, or nil if none found.
func (m *ViewQuotaRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ViewQuotaRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectId

	if len(errors) > 0 {
		return ViewQuotaRequestMultiError(errors)
	}

	return nil
}

// ViewQuotaRequestMultiError is an error wrapping multiple validation errors
// returned by ViewQuotaRequest.ValidateAll() if the designated constraints
// aren't met.
type ViewQuotaRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ViewQuotaRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ViewQuotaRequestMultiError) AllErrors() []error { return m }

// ViewQuotaRequestValidationError is the validation error returned by
// ViewQuotaRequest.Validate if the designated constraints aren't met.
type ViewQuotaRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ViewQuotaRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ViewQuotaRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ViewQuotaRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ViewQuotaRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ViewQuotaRequestValidationError) ErrorName() string { return "ViewQuotaRequestValidationError" }

// Error satisfies the builtin error interface
func (e ViewQuotaRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sViewQuotaRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ViewQuotaRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ViewQuotaRequestValidationError{}

// Validate checks the field values on ViewQuotaResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ViewQuotaResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ViewQuotaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ViewQuotaResponseMultiError, or nil if none found.
func (m *ViewQuotaResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ViewQuotaResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PullRequestsCount

	// no validation rules for PullRequestsLimit

	if len(errors) > 0 {
		return ViewQuotaResponseMultiError(errors)
	}

	return nil
}

// ViewQuotaResponseMultiError is an error wrapping multiple validation errors
// returned by ViewQuotaResponse.ValidateAll() if the designated constraints
// aren't met.
type ViewQuotaResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ViewQuotaResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ViewQuotaResponseMultiError) AllErrors() []error { return m }

// ViewQuotaResponseValidationError is the validation error returned by
// ViewQuotaResponse.Validate if the designated constraints aren't met.
type ViewQuotaResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ViewQuotaResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ViewQuotaResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ViewQuotaResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ViewQuotaResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ViewQuotaResponseValidationError) ErrorName() string {
	return "ViewQuotaResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ViewQuotaResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sViewQuotaResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ViewQuotaResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ViewQuotaResponseValidationError{}
