/*
Apono

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apono

import (
	"encoding/json"
	"fmt"
)

// ReportField the model 'ReportField'
type ReportField string

// List of ReportField
const (
	REPORTFIELD_REQUEST_ID          ReportField = "request_id"
	REPORTFIELD_REQUEST_DATE        ReportField = "request_date"
	REPORTFIELD_REQUEST_GRANT_DATE  ReportField = "request_grant_date"
	REPORTFIELD_REQUEST_REVOKE_DATE ReportField = "request_revoke_date"
	REPORTFIELD_REQUESTOR_NAME      ReportField = "requestor_name"
	REPORTFIELD_REQUESTOR_EMAIL     ReportField = "requestor_email"
	REPORTFIELD_INTEGRATION         ReportField = "integration"
	REPORTFIELD_RESOURCES           ReportField = "resources"
	REPORTFIELD_RESOURCE_TYPE       ReportField = "resource_type"
	REPORTFIELD_PERMISSIONS         ReportField = "permissions"
	REPORTFIELD_APPROVER_NAMES      ReportField = "approver_names"
	REPORTFIELD_APPROVER_EMAILS     ReportField = "approver_emails"
	REPORTFIELD_APPROVER_TYPES      ReportField = "approver_types"
	REPORTFIELD_JUSTIFICATION       ReportField = "justification"
	REPORTFIELD_STATUS              ReportField = "status"
	REPORTFIELD_RESOURCES_STATUS    ReportField = "resources_status"
	REPORTFIELD_TRIGGER_TYPE        ReportField = "trigger_type"
	REPORTFIELD_ACCESS_FLOW         ReportField = "access_flow"
	REPORTFIELD_BUNDLE_NAME         ReportField = "bundle_name"
)

// All allowed values of ReportField enum
var AllowedReportFieldEnumValues = []ReportField{
	"request_id",
	"request_date",
	"request_grant_date",
	"request_revoke_date",
	"requestor_name",
	"requestor_email",
	"integration",
	"resources",
	"resource_type",
	"permissions",
	"approver_names",
	"approver_emails",
	"approver_types",
	"justification",
	"status",
	"resources_status",
	"trigger_type",
	"access_flow",
	"bundle_name",
}

func (v *ReportField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportField(value)
	for _, existing := range AllowedReportFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportField", value)
}

// NewReportFieldFromValue returns a pointer to a valid ReportField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportFieldFromValue(v string) (*ReportField, error) {
	ev := ReportField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportField: valid values are %v", v, AllowedReportFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportField) IsValid() bool {
	for _, existing := range AllowedReportFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportField value
func (v ReportField) Ptr() *ReportField {
	return &v
}

type NullableReportField struct {
	value *ReportField
	isSet bool
}

func (v NullableReportField) Get() *ReportField {
	return v.value
}

func (v *NullableReportField) Set(val *ReportField) {
	v.value = val
	v.isSet = true
}

func (v NullableReportField) IsSet() bool {
	return v.isSet
}

func (v *NullableReportField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportField(val *ReportField) *NullableReportField {
	return &NullableReportField{value: val, isSet: true}
}

func (v NullableReportField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
