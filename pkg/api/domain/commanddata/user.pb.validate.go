// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/domain/commanddata/user.proto

package commanddata

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

// define the regex for a UUID once up-front
var _user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateUserCommandData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUserCommandData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserCommandData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserCommandDataMultiError, or nil if none found.
func (m *CreateUserCommandData) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserCommandData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = CreateUserCommandDataValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetName()) > 150 {
		err := CreateUserCommandDataValidationError{
			field:  "Name",
			reason: "value length must be at most 150 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateUserCommandDataMultiError(errors)
	}
	return nil
}

func (m *CreateUserCommandData) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreateUserCommandData) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreateUserCommandDataMultiError is an error wrapping multiple validation
// errors returned by CreateUserCommandData.ValidateAll() if the designated
// constraints aren't met.
type CreateUserCommandDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserCommandDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserCommandDataMultiError) AllErrors() []error { return m }

// CreateUserCommandDataValidationError is the validation error returned by
// CreateUserCommandData.Validate if the designated constraints aren't met.
type CreateUserCommandDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserCommandDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserCommandDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserCommandDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserCommandDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserCommandDataValidationError) ErrorName() string {
	return "CreateUserCommandDataValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserCommandDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserCommandData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserCommandDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserCommandDataValidationError{}

// Validate checks the field values on CreateUserRoleBindingCommandData with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CreateUserRoleBindingCommandData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRoleBindingCommandData with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreateUserRoleBindingCommandDataMultiError, or nil if none found.
func (m *CreateUserRoleBindingCommandData) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRoleBindingCommandData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUserId()); err != nil {
		err = CreateUserRoleBindingCommandDataValidationError{
			field:  "UserId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetRole()) > 60 {
		err := CreateUserRoleBindingCommandDataValidationError{
			field:  "Role",
			reason: "value length must be at most 60 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CreateUserRoleBindingCommandData_Role_Pattern.MatchString(m.GetRole()) {
		err := CreateUserRoleBindingCommandDataValidationError{
			field:  "Role",
			reason: "value does not match regex pattern \"^[a-z]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetScope()) > 60 {
		err := CreateUserRoleBindingCommandDataValidationError{
			field:  "Scope",
			reason: "value length must be at most 60 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CreateUserRoleBindingCommandData_Scope_Pattern.MatchString(m.GetScope()) {
		err := CreateUserRoleBindingCommandDataValidationError{
			field:  "Scope",
			reason: "value does not match regex pattern \"^[a-z]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateUuid(m.GetResource()); err != nil {
		err = CreateUserRoleBindingCommandDataValidationError{
			field:  "Resource",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateUserRoleBindingCommandDataMultiError(errors)
	}
	return nil
}

func (m *CreateUserRoleBindingCommandData) _validateUuid(uuid string) error {
	if matched := _user_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateUserRoleBindingCommandDataMultiError is an error wrapping multiple
// validation errors returned by
// CreateUserRoleBindingCommandData.ValidateAll() if the designated
// constraints aren't met.
type CreateUserRoleBindingCommandDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRoleBindingCommandDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRoleBindingCommandDataMultiError) AllErrors() []error { return m }

// CreateUserRoleBindingCommandDataValidationError is the validation error
// returned by CreateUserRoleBindingCommandData.Validate if the designated
// constraints aren't met.
type CreateUserRoleBindingCommandDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRoleBindingCommandDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRoleBindingCommandDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRoleBindingCommandDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRoleBindingCommandDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRoleBindingCommandDataValidationError) ErrorName() string {
	return "CreateUserRoleBindingCommandDataValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRoleBindingCommandDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRoleBindingCommandData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRoleBindingCommandDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRoleBindingCommandDataValidationError{}

var _CreateUserRoleBindingCommandData_Role_Pattern = regexp.MustCompile("^[a-z]+$")

var _CreateUserRoleBindingCommandData_Scope_Pattern = regexp.MustCompile("^[a-z]+$")