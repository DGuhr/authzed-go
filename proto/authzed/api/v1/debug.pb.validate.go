// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authzed/api/v1/debug.proto

package v1

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

// Validate checks the field values on DebugInformation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DebugInformation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DebugInformation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DebugInformationMultiError, or nil if none found.
func (m *DebugInformation) ValidateAll() error {
	return m.validate(true)
}

func (m *DebugInformation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCheck()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DebugInformationValidationError{
					field:  "Check",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DebugInformationValidationError{
					field:  "Check",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCheck()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DebugInformationValidationError{
				field:  "Check",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SchemaUsed

	if len(errors) > 0 {
		return DebugInformationMultiError(errors)
	}

	return nil
}

// DebugInformationMultiError is an error wrapping multiple validation errors
// returned by DebugInformation.ValidateAll() if the designated constraints
// aren't met.
type DebugInformationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DebugInformationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DebugInformationMultiError) AllErrors() []error { return m }

// DebugInformationValidationError is the validation error returned by
// DebugInformation.Validate if the designated constraints aren't met.
type DebugInformationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DebugInformationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DebugInformationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DebugInformationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DebugInformationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DebugInformationValidationError) ErrorName() string { return "DebugInformationValidationError" }

// Error satisfies the builtin error interface
func (e DebugInformationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDebugInformation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DebugInformationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DebugInformationValidationError{}

// Validate checks the field values on CheckDebugTrace with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CheckDebugTrace) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckDebugTrace with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckDebugTraceMultiError, or nil if none found.
func (m *CheckDebugTrace) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckDebugTrace) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetResource() == nil {
		err := CheckDebugTraceValidationError{
			field:  "Resource",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckDebugTraceValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckDebugTraceValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckDebugTraceValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Permission

	if _, ok := _CheckDebugTrace_PermissionType_NotInLookup[m.GetPermissionType()]; ok {
		err := CheckDebugTraceValidationError{
			field:  "PermissionType",
			reason: "value must not be in list [0]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := CheckDebugTrace_PermissionType_name[int32(m.GetPermissionType())]; !ok {
		err := CheckDebugTraceValidationError{
			field:  "PermissionType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSubject() == nil {
		err := CheckDebugTraceValidationError{
			field:  "Subject",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckDebugTraceValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckDebugTraceValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckDebugTraceValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := _CheckDebugTrace_Result_NotInLookup[m.GetResult()]; ok {
		err := CheckDebugTraceValidationError{
			field:  "Result",
			reason: "value must not be in list [0]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := CheckDebugTrace_Permissionship_name[int32(m.GetResult())]; !ok {
		err := CheckDebugTraceValidationError{
			field:  "Result",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCaveatEvaluationInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckDebugTraceValidationError{
					field:  "CaveatEvaluationInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckDebugTraceValidationError{
					field:  "CaveatEvaluationInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCaveatEvaluationInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckDebugTraceValidationError{
				field:  "CaveatEvaluationInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Resolution.(type) {

	case *CheckDebugTrace_WasCachedResult:
		// no validation rules for WasCachedResult

	case *CheckDebugTrace_SubProblems_:

		if all {
			switch v := interface{}(m.GetSubProblems()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CheckDebugTraceValidationError{
						field:  "SubProblems",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckDebugTraceValidationError{
						field:  "SubProblems",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSubProblems()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckDebugTraceValidationError{
					field:  "SubProblems",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := CheckDebugTraceValidationError{
			field:  "Resolution",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return CheckDebugTraceMultiError(errors)
	}

	return nil
}

// CheckDebugTraceMultiError is an error wrapping multiple validation errors
// returned by CheckDebugTrace.ValidateAll() if the designated constraints
// aren't met.
type CheckDebugTraceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckDebugTraceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckDebugTraceMultiError) AllErrors() []error { return m }

// CheckDebugTraceValidationError is the validation error returned by
// CheckDebugTrace.Validate if the designated constraints aren't met.
type CheckDebugTraceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckDebugTraceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckDebugTraceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckDebugTraceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckDebugTraceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckDebugTraceValidationError) ErrorName() string { return "CheckDebugTraceValidationError" }

// Error satisfies the builtin error interface
func (e CheckDebugTraceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckDebugTrace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckDebugTraceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckDebugTraceValidationError{}

var _CheckDebugTrace_PermissionType_NotInLookup = map[CheckDebugTrace_PermissionType]struct{}{
	0: {},
}

var _CheckDebugTrace_Result_NotInLookup = map[CheckDebugTrace_Permissionship]struct{}{
	0: {},
}

// Validate checks the field values on CaveatEvalInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CaveatEvalInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CaveatEvalInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CaveatEvalInfoMultiError,
// or nil if none found.
func (m *CaveatEvalInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *CaveatEvalInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Expression

	// no validation rules for Result

	if all {
		switch v := interface{}(m.GetContext()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CaveatEvalInfoValidationError{
					field:  "Context",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CaveatEvalInfoValidationError{
					field:  "Context",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CaveatEvalInfoValidationError{
				field:  "Context",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPartialCaveatInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CaveatEvalInfoValidationError{
					field:  "PartialCaveatInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CaveatEvalInfoValidationError{
					field:  "PartialCaveatInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPartialCaveatInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CaveatEvalInfoValidationError{
				field:  "PartialCaveatInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CaveatEvalInfoMultiError(errors)
	}

	return nil
}

// CaveatEvalInfoMultiError is an error wrapping multiple validation errors
// returned by CaveatEvalInfo.ValidateAll() if the designated constraints
// aren't met.
type CaveatEvalInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CaveatEvalInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CaveatEvalInfoMultiError) AllErrors() []error { return m }

// CaveatEvalInfoValidationError is the validation error returned by
// CaveatEvalInfo.Validate if the designated constraints aren't met.
type CaveatEvalInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CaveatEvalInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CaveatEvalInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CaveatEvalInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CaveatEvalInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CaveatEvalInfoValidationError) ErrorName() string { return "CaveatEvalInfoValidationError" }

// Error satisfies the builtin error interface
func (e CaveatEvalInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCaveatEvalInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CaveatEvalInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CaveatEvalInfoValidationError{}

// Validate checks the field values on CheckDebugTrace_SubProblems with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckDebugTrace_SubProblems) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckDebugTrace_SubProblems with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckDebugTrace_SubProblemsMultiError, or nil if none found.
func (m *CheckDebugTrace_SubProblems) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckDebugTrace_SubProblems) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTraces() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CheckDebugTrace_SubProblemsValidationError{
						field:  fmt.Sprintf("Traces[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckDebugTrace_SubProblemsValidationError{
						field:  fmt.Sprintf("Traces[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckDebugTrace_SubProblemsValidationError{
					field:  fmt.Sprintf("Traces[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CheckDebugTrace_SubProblemsMultiError(errors)
	}

	return nil
}

// CheckDebugTrace_SubProblemsMultiError is an error wrapping multiple
// validation errors returned by CheckDebugTrace_SubProblems.ValidateAll() if
// the designated constraints aren't met.
type CheckDebugTrace_SubProblemsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckDebugTrace_SubProblemsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckDebugTrace_SubProblemsMultiError) AllErrors() []error { return m }

// CheckDebugTrace_SubProblemsValidationError is the validation error returned
// by CheckDebugTrace_SubProblems.Validate if the designated constraints
// aren't met.
type CheckDebugTrace_SubProblemsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckDebugTrace_SubProblemsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckDebugTrace_SubProblemsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckDebugTrace_SubProblemsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckDebugTrace_SubProblemsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckDebugTrace_SubProblemsValidationError) ErrorName() string {
	return "CheckDebugTrace_SubProblemsValidationError"
}

// Error satisfies the builtin error interface
func (e CheckDebugTrace_SubProblemsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckDebugTrace_SubProblems.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckDebugTrace_SubProblemsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckDebugTrace_SubProblemsValidationError{}
