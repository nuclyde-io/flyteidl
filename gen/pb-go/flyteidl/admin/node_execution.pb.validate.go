// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/node_execution.proto

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

// Validate checks the field values on NodeExecutionIdentifier with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionIdentifier) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NodeId

	// no validation rules for ExecutionId

	// no validation rules for RetryAttempt

	return nil
}

// NodeExecutionIdentifierValidationError is the validation error returned by
// NodeExecutionIdentifier.Validate if the designated constraints aren't met.
type NodeExecutionIdentifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionIdentifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionIdentifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionIdentifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionIdentifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionIdentifierValidationError) ErrorName() string {
	return "NodeExecutionIdentifierValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionIdentifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionIdentifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionIdentifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionIdentifierValidationError{}

// Validate checks the field values on NodeExecutionGetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionGetRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NodeExecutionId

	return nil
}

// NodeExecutionGetRequestValidationError is the validation error returned by
// NodeExecutionGetRequest.Validate if the designated constraints aren't met.
type NodeExecutionGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionGetRequestValidationError) ErrorName() string {
	return "NodeExecutionGetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionGetRequestValidationError{}

// Validate checks the field values on NodeExecutionListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionListRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	// no validation rules for Filters

	return nil
}

// NodeExecutionListRequestValidationError is the validation error returned by
// NodeExecutionListRequest.Validate if the designated constraints aren't met.
type NodeExecutionListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionListRequestValidationError) ErrorName() string {
	return "NodeExecutionListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionListRequestValidationError{}

// Validate checks the field values on NodeExecution with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *NodeExecution) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for NodeExecutionId

	// no validation rules for InputUri

	if v, ok := interface{}(m.GetClosure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionValidationError{
				field:  "Closure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// NodeExecutionValidationError is the validation error returned by
// NodeExecution.Validate if the designated constraints aren't met.
type NodeExecutionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionValidationError) ErrorName() string { return "NodeExecutionValidationError" }

// Error satisfies the builtin error interface
func (e NodeExecutionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecution.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionValidationError{}

// Validate checks the field values on NodeExecutionList with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *NodeExecutionList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetNodeExecutions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionListValidationError{
					field:  fmt.Sprintf("NodeExecutions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// NodeExecutionListValidationError is the validation error returned by
// NodeExecutionList.Validate if the designated constraints aren't met.
type NodeExecutionListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionListValidationError) ErrorName() string {
	return "NodeExecutionListValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionListValidationError{}

// Validate checks the field values on NodeExecutionClosure with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionClosure) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Phase

	switch m.OutputResult.(type) {

	case *NodeExecutionClosure_OutputUri:
		// no validation rules for OutputUri

	case *NodeExecutionClosure_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionClosureValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// NodeExecutionClosureValidationError is the validation error returned by
// NodeExecutionClosure.Validate if the designated constraints aren't met.
type NodeExecutionClosureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionClosureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionClosureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionClosureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionClosureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionClosureValidationError) ErrorName() string {
	return "NodeExecutionClosureValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionClosureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionClosure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionClosureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionClosureValidationError{}