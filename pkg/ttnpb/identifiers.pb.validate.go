// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _identifiers_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on ApplicationIdentifiers with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplicationIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ApplicationIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "application_id":

			if utf8.RuneCountInString(m.GetApplicationId()) > 36 {
				return ApplicationIdentifiersValidationError{
					field:  "application_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_ApplicationIdentifiers_ApplicationId_Pattern.MatchString(m.GetApplicationId()) {
				return ApplicationIdentifiersValidationError{
					field:  "application_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
				}
			}

		default:
			return ApplicationIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ApplicationIdentifiersValidationError is the validation error returned by
// ApplicationIdentifiers.ValidateFields if the designated constraints aren't met.
type ApplicationIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationIdentifiersValidationError) ErrorName() string {
	return "ApplicationIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationIdentifiersValidationError{}

var _ApplicationIdentifiers_ApplicationId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on ClientIdentifiers with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClientIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ClientIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "client_id":

			if utf8.RuneCountInString(m.GetClientId()) > 36 {
				return ClientIdentifiersValidationError{
					field:  "client_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_ClientIdentifiers_ClientId_Pattern.MatchString(m.GetClientId()) {
				return ClientIdentifiersValidationError{
					field:  "client_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
				}
			}

		default:
			return ClientIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ClientIdentifiersValidationError is the validation error returned by
// ClientIdentifiers.ValidateFields if the designated constraints aren't met.
type ClientIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientIdentifiersValidationError) ErrorName() string {
	return "ClientIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e ClientIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientIdentifiersValidationError{}

var _ClientIdentifiers_ClientId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on EndDeviceIdentifiers with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EndDeviceIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = EndDeviceIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "device_id":

			if utf8.RuneCountInString(m.GetDeviceId()) > 36 {
				return EndDeviceIdentifiersValidationError{
					field:  "device_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_EndDeviceIdentifiers_DeviceId_Pattern.MatchString(m.GetDeviceId()) {
				return EndDeviceIdentifiersValidationError{
					field:  "device_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
				}
			}

		case "application_ids":

			if m.GetApplicationIds() == nil {
				return EndDeviceIdentifiersValidationError{
					field:  "application_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetApplicationIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return EndDeviceIdentifiersValidationError{
						field:  "application_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "dev_eui":
			// no validation rules for DevEui
		case "join_eui":
			// no validation rules for JoinEui
		case "dev_addr":
			// no validation rules for DevAddr
		default:
			return EndDeviceIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// EndDeviceIdentifiersValidationError is the validation error returned by
// EndDeviceIdentifiers.ValidateFields if the designated constraints aren't met.
type EndDeviceIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndDeviceIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndDeviceIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndDeviceIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndDeviceIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndDeviceIdentifiersValidationError) ErrorName() string {
	return "EndDeviceIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e EndDeviceIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndDeviceIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndDeviceIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndDeviceIdentifiersValidationError{}

var _EndDeviceIdentifiers_DeviceId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on GatewayIdentifiers with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GatewayIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GatewayIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "gateway_id":

			if utf8.RuneCountInString(m.GetGatewayId()) > 36 {
				return GatewayIdentifiersValidationError{
					field:  "gateway_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_GatewayIdentifiers_GatewayId_Pattern.MatchString(m.GetGatewayId()) {
				return GatewayIdentifiersValidationError{
					field:  "gateway_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
				}
			}

		case "eui":
			// no validation rules for Eui
		default:
			return GatewayIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GatewayIdentifiersValidationError is the validation error returned by
// GatewayIdentifiers.ValidateFields if the designated constraints aren't met.
type GatewayIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GatewayIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GatewayIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GatewayIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GatewayIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GatewayIdentifiersValidationError) ErrorName() string {
	return "GatewayIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e GatewayIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGatewayIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GatewayIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GatewayIdentifiersValidationError{}

var _GatewayIdentifiers_GatewayId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on OrganizationIdentifiers with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OrganizationIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = OrganizationIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "organization_id":

			if utf8.RuneCountInString(m.GetOrganizationId()) > 36 {
				return OrganizationIdentifiersValidationError{
					field:  "organization_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_OrganizationIdentifiers_OrganizationId_Pattern.MatchString(m.GetOrganizationId()) {
				return OrganizationIdentifiersValidationError{
					field:  "organization_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
				}
			}

		default:
			return OrganizationIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// OrganizationIdentifiersValidationError is the validation error returned by
// OrganizationIdentifiers.ValidateFields if the designated constraints aren't met.
type OrganizationIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrganizationIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrganizationIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrganizationIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrganizationIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrganizationIdentifiersValidationError) ErrorName() string {
	return "OrganizationIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e OrganizationIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrganizationIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrganizationIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrganizationIdentifiersValidationError{}

var _OrganizationIdentifiers_OrganizationId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on UserIdentifiers with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UserIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = UserIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "user_id":

			if utf8.RuneCountInString(m.GetUserId()) > 36 {
				return UserIdentifiersValidationError{
					field:  "user_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_UserIdentifiers_UserId_Pattern.MatchString(m.GetUserId()) {
				return UserIdentifiersValidationError{
					field:  "user_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){1,}$\"",
				}
			}

		case "email":
			// no validation rules for Email
		default:
			return UserIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// UserIdentifiersValidationError is the validation error returned by
// UserIdentifiers.ValidateFields if the designated constraints aren't met.
type UserIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserIdentifiersValidationError) ErrorName() string { return "UserIdentifiersValidationError" }

// Error satisfies the builtin error interface
func (e UserIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserIdentifiersValidationError{}

var _UserIdentifiers_UserId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){1,}$")

// ValidateFields checks the field values on OrganizationOrUserIdentifiers with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *OrganizationOrUserIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = OrganizationOrUserIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":
			if m.Ids == nil {
				return OrganizationOrUserIdentifiersValidationError{
					field:  "ids",
					reason: "value is required",
				}
			}
			if len(subs) == 0 {
				subs = []string{
					"organization_ids", "user_ids",
				}
			}
			for name, subs := range _processPaths(subs) {
				_ = subs
				switch name {
				case "organization_ids":
					w, ok := m.Ids.(*OrganizationOrUserIdentifiers_OrganizationIds)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetOrganizationIds()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return OrganizationOrUserIdentifiersValidationError{
								field:  "organization_ids",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "user_ids":
					w, ok := m.Ids.(*OrganizationOrUserIdentifiers_UserIds)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetUserIds()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return OrganizationOrUserIdentifiersValidationError{
								field:  "user_ids",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				}
			}
		default:
			return OrganizationOrUserIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// OrganizationOrUserIdentifiersValidationError is the validation error
// returned by OrganizationOrUserIdentifiers.ValidateFields if the designated
// constraints aren't met.
type OrganizationOrUserIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrganizationOrUserIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrganizationOrUserIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrganizationOrUserIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrganizationOrUserIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrganizationOrUserIdentifiersValidationError) ErrorName() string {
	return "OrganizationOrUserIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e OrganizationOrUserIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrganizationOrUserIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrganizationOrUserIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrganizationOrUserIdentifiersValidationError{}

// ValidateFields checks the field values on EntityIdentifiers with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EntityIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = EntityIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":
			if m.Ids == nil {
				return EntityIdentifiersValidationError{
					field:  "ids",
					reason: "value is required",
				}
			}
			if len(subs) == 0 {
				subs = []string{
					"application_ids", "client_ids", "device_ids", "gateway_ids", "organization_ids", "user_ids",
				}
			}
			for name, subs := range _processPaths(subs) {
				_ = subs
				switch name {
				case "application_ids":
					w, ok := m.Ids.(*EntityIdentifiers_ApplicationIds)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetApplicationIds()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return EntityIdentifiersValidationError{
								field:  "application_ids",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "client_ids":
					w, ok := m.Ids.(*EntityIdentifiers_ClientIds)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetClientIds()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return EntityIdentifiersValidationError{
								field:  "client_ids",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "device_ids":
					w, ok := m.Ids.(*EntityIdentifiers_DeviceIds)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetDeviceIds()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return EntityIdentifiersValidationError{
								field:  "device_ids",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "gateway_ids":
					w, ok := m.Ids.(*EntityIdentifiers_GatewayIds)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetGatewayIds()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return EntityIdentifiersValidationError{
								field:  "gateway_ids",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "organization_ids":
					w, ok := m.Ids.(*EntityIdentifiers_OrganizationIds)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetOrganizationIds()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return EntityIdentifiersValidationError{
								field:  "organization_ids",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "user_ids":
					w, ok := m.Ids.(*EntityIdentifiers_UserIds)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetUserIds()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return EntityIdentifiersValidationError{
								field:  "user_ids",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				}
			}
		default:
			return EntityIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// EntityIdentifiersValidationError is the validation error returned by
// EntityIdentifiers.ValidateFields if the designated constraints aren't met.
type EntityIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityIdentifiersValidationError) ErrorName() string {
	return "EntityIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e EntityIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntityIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityIdentifiersValidationError{}

// ValidateFields checks the field values on EndDeviceVersionIdentifiers with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *EndDeviceVersionIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = EndDeviceVersionIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "brand_id":

			if m.GetBrandId() != "" {

				if utf8.RuneCountInString(m.GetBrandId()) > 36 {
					return EndDeviceVersionIdentifiersValidationError{
						field:  "brand_id",
						reason: "value length must be at most 36 runes",
					}
				}

				if !_EndDeviceVersionIdentifiers_BrandId_Pattern.MatchString(m.GetBrandId()) {
					return EndDeviceVersionIdentifiersValidationError{
						field:  "brand_id",
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

			}

		case "model_id":

			if m.GetModelId() != "" {

				if utf8.RuneCountInString(m.GetModelId()) > 36 {
					return EndDeviceVersionIdentifiersValidationError{
						field:  "model_id",
						reason: "value length must be at most 36 runes",
					}
				}

				if !_EndDeviceVersionIdentifiers_ModelId_Pattern.MatchString(m.GetModelId()) {
					return EndDeviceVersionIdentifiersValidationError{
						field:  "model_id",
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

			}

		case "hardware_version":

			if utf8.RuneCountInString(m.GetHardwareVersion()) > 32 {
				return EndDeviceVersionIdentifiersValidationError{
					field:  "hardware_version",
					reason: "value length must be at most 32 runes",
				}
			}

		case "firmware_version":

			if utf8.RuneCountInString(m.GetFirmwareVersion()) > 32 {
				return EndDeviceVersionIdentifiersValidationError{
					field:  "firmware_version",
					reason: "value length must be at most 32 runes",
				}
			}

		case "band_id":

			if utf8.RuneCountInString(m.GetBandId()) > 32 {
				return EndDeviceVersionIdentifiersValidationError{
					field:  "band_id",
					reason: "value length must be at most 32 runes",
				}
			}

		default:
			return EndDeviceVersionIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// EndDeviceVersionIdentifiersValidationError is the validation error returned
// by EndDeviceVersionIdentifiers.ValidateFields if the designated constraints
// aren't met.
type EndDeviceVersionIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndDeviceVersionIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndDeviceVersionIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndDeviceVersionIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndDeviceVersionIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndDeviceVersionIdentifiersValidationError) ErrorName() string {
	return "EndDeviceVersionIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e EndDeviceVersionIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndDeviceVersionIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndDeviceVersionIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndDeviceVersionIdentifiersValidationError{}

var _EndDeviceVersionIdentifiers_BrandId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

var _EndDeviceVersionIdentifiers_ModelId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on NetworkIdentifiers with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NetworkIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = NetworkIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "net_id":
			// no validation rules for NetId
		case "tenant_id":

			if utf8.RuneCountInString(m.GetTenantId()) > 36 {
				return NetworkIdentifiersValidationError{
					field:  "tenant_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_NetworkIdentifiers_TenantId_Pattern.MatchString(m.GetTenantId()) {
				return NetworkIdentifiersValidationError{
					field:  "tenant_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$\"",
				}
			}

		case "cluster_id":

			if utf8.RuneCountInString(m.GetClusterId()) > 64 {
				return NetworkIdentifiersValidationError{
					field:  "cluster_id",
					reason: "value length must be at most 64 runes",
				}
			}

		case "cluster_address":

			if utf8.RuneCountInString(m.GetClusterAddress()) > 256 {
				return NetworkIdentifiersValidationError{
					field:  "cluster_address",
					reason: "value length must be at most 256 runes",
				}
			}

		case "tenant_address":

			if utf8.RuneCountInString(m.GetTenantAddress()) > 256 {
				return NetworkIdentifiersValidationError{
					field:  "tenant_address",
					reason: "value length must be at most 256 runes",
				}
			}

		default:
			return NetworkIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// NetworkIdentifiersValidationError is the validation error returned by
// NetworkIdentifiers.ValidateFields if the designated constraints aren't met.
type NetworkIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NetworkIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NetworkIdentifiersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NetworkIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NetworkIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NetworkIdentifiersValidationError) ErrorName() string {
	return "NetworkIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e NetworkIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNetworkIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NetworkIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NetworkIdentifiersValidationError{}

var _NetworkIdentifiers_TenantId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")
