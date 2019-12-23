// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/project_domain_attributes.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _project_domain_attributes_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ProjectDomainAttributes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProjectDomainAttributes) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Project

	// no validation rules for Domain

	if v, ok := interface{}(m.GetMatchingAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectDomainAttributesValidationError{
				field:  "MatchingAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProjectDomainAttributesValidationError is the validation error returned by
// ProjectDomainAttributes.Validate if the designated constraints aren't met.
type ProjectDomainAttributesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesValidationError) ErrorName() string {
	return "ProjectDomainAttributesValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesValidationError{}

// Validate checks the field values on ProjectDomainAttributesUpdateRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ProjectDomainAttributesUpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectDomainAttributesUpdateRequestValidationError{
				field:  "Attributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProjectDomainAttributesUpdateRequestValidationError is the validation error
// returned by ProjectDomainAttributesUpdateRequest.Validate if the designated
// constraints aren't met.
type ProjectDomainAttributesUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesUpdateRequestValidationError) ErrorName() string {
	return "ProjectDomainAttributesUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributesUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesUpdateRequestValidationError{}

// Validate checks the field values on ProjectDomainAttributesUpdateResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ProjectDomainAttributesUpdateResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ProjectDomainAttributesUpdateResponseValidationError is the validation error
// returned by ProjectDomainAttributesUpdateResponse.Validate if the
// designated constraints aren't met.
type ProjectDomainAttributesUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesUpdateResponseValidationError) ErrorName() string {
	return "ProjectDomainAttributesUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributesUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesUpdateResponseValidationError{}
