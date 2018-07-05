// Code generated by protoc-gen-validate
// source: proto/tessellate.proto
// DO NOT EDIT!!!

package server

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

// Validate checks the field values on GetWorkspaceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetWorkspaceRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return GetWorkspaceRequestValidationError{
			Field:  "Id",
			Reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetWorkspaceRequestValidationError is the validation error returned by
// GetWorkspaceRequest.Validate if the designated constraints aren't met.
type GetWorkspaceRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GetWorkspaceRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetWorkspaceRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GetWorkspaceRequestValidationError{}

// Validate checks the field values on Workspace with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Workspace) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Vars

	return nil
}

// WorkspaceValidationError is the validation error returned by
// Workspace.Validate if the designated constraints aren't met.
type WorkspaceValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e WorkspaceValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkspace.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = WorkspaceValidationError{}

// Validate checks the field values on Layouts with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Layouts) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetLayouts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LayoutsValidationError{
					Field:  fmt.Sprintf("Layouts[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// LayoutsValidationError is the validation error returned by Layouts.Validate
// if the designated constraints aren't met.
type LayoutsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e LayoutsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLayouts.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = LayoutsValidationError{}

// Validate checks the field values on Layout with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Layout) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Plan

	// no validation rules for Vars

	// no validation rules for Status

	return nil
}

// LayoutValidationError is the validation error returned by Layout.Validate if
// the designated constraints aren't met.
type LayoutValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e LayoutValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLayout.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = LayoutValidationError{}

// Validate checks the field values on SaveWorkspaceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SaveWorkspaceRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return SaveWorkspaceRequestValidationError{
			Field:  "Id",
			Reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Vars

	return nil
}

// SaveWorkspaceRequestValidationError is the validation error returned by
// SaveWorkspaceRequest.Validate if the designated constraints aren't met.
type SaveWorkspaceRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e SaveWorkspaceRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveWorkspaceRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = SaveWorkspaceRequestValidationError{}

// Validate checks the field values on GetWorkspaceLayoutsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetWorkspaceLayoutsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return GetWorkspaceLayoutsRequestValidationError{
			Field:  "Id",
			Reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetWorkspaceLayoutsRequestValidationError is the validation error returned
// by GetWorkspaceLayoutsRequest.Validate if the designated constraints aren't met.
type GetWorkspaceLayoutsRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GetWorkspaceLayoutsRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetWorkspaceLayoutsRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GetWorkspaceLayoutsRequestValidationError{}

// Validate checks the field values on JobStatus with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JobStatus) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Status

	return nil
}

// JobStatusValidationError is the validation error returned by
// JobStatus.Validate if the designated constraints aren't met.
type JobStatusValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e JobStatusValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobStatus.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = JobStatusValidationError{}

// Validate checks the field values on Vars with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Vars) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Vars

	return nil
}

// VarsValidationError is the validation error returned by Vars.Validate if the
// designated constraints aren't met.
type VarsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e VarsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVars.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = VarsValidationError{}

// Validate checks the field values on JobRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JobRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return JobRequestValidationError{
			Field:  "Id",
			Reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// JobRequestValidationError is the validation error returned by
// JobRequest.Validate if the designated constraints aren't met.
type JobRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e JobRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJobRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = JobRequestValidationError{}

// Validate checks the field values on Ok with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Ok) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// OkValidationError is the validation error returned by Ok.Validate if the
// designated constraints aren't met.
type OkValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e OkValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOk.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = OkValidationError{}

// Validate checks the field values on LayoutRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LayoutRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetWorkspaceId()) < 1 {
		return LayoutRequestValidationError{
			Field:  "WorkspaceId",
			Reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return LayoutRequestValidationError{
			Field:  "Id",
			Reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// LayoutRequestValidationError is the validation error returned by
// LayoutRequest.Validate if the designated constraints aren't met.
type LayoutRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e LayoutRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLayoutRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = LayoutRequestValidationError{}

// Validate checks the field values on SaveLayoutRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SaveLayoutRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetWorkspaceId()) < 1 {
		return SaveLayoutRequestValidationError{
			Field:  "WorkspaceId",
			Reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return SaveLayoutRequestValidationError{
			Field:  "Id",
			Reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetPlan()) < 1 {
		return SaveLayoutRequestValidationError{
			Field:  "Plan",
			Reason: "value must contain at least 1 pair(s)",
		}
	}

	// no validation rules for Vars

	return nil
}

// SaveLayoutRequestValidationError is the validation error returned by
// SaveLayoutRequest.Validate if the designated constraints aren't met.
type SaveLayoutRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e SaveLayoutRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveLayoutRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = SaveLayoutRequestValidationError{}

// Validate checks the field values on ApplyLayoutRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplyLayoutRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetWorkspaceId()) < 1 {
		return ApplyLayoutRequestValidationError{
			Field:  "WorkspaceId",
			Reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return ApplyLayoutRequestValidationError{
			Field:  "Id",
			Reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetVars()) < 1 {
		return ApplyLayoutRequestValidationError{
			Field:  "Vars",
			Reason: "value must contain at least 1 pair(s)",
		}
	}

	// no validation rules for Dry

	return nil
}

// ApplyLayoutRequestValidationError is the validation error returned by
// ApplyLayoutRequest.Validate if the designated constraints aren't met.
type ApplyLayoutRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ApplyLayoutRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplyLayoutRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ApplyLayoutRequestValidationError{}

// Validate checks the field values on StartWatchRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StartWatchRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetWorkspaceId()) < 1 {
		return StartWatchRequestValidationError{
			Field:  "WorkspaceId",
			Reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return StartWatchRequestValidationError{
			Field:  "Id",
			Reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for SuccessCallback

	// no validation rules for FailureCallback

	return nil
}

// StartWatchRequestValidationError is the validation error returned by
// StartWatchRequest.Validate if the designated constraints aren't met.
type StartWatchRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StartWatchRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartWatchRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StartWatchRequestValidationError{}

// Validate checks the field values on StopWatchRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StopWatchRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetWorkspaceId()) < 1 {
		return StopWatchRequestValidationError{
			Field:  "WorkspaceId",
			Reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return StopWatchRequestValidationError{
			Field:  "Id",
			Reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// StopWatchRequestValidationError is the validation error returned by
// StopWatchRequest.Validate if the designated constraints aren't met.
type StopWatchRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StopWatchRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStopWatchRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StopWatchRequestValidationError{}
