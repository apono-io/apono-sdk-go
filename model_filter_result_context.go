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

// FilterResultContext the model 'FilterResultContext'
type FilterResultContext string

// List of FilterResultContext
const (
	FILTERRESULTCONTEXT_APPROVERS        FilterResultContext = "Approvers"
	FILTERRESULTCONTEXT_ACCESS_TARGET    FilterResultContext = "AccessTarget"
	FILTERRESULTCONTEXT_GRANTEES         FilterResultContext = "Grantees"
	FILTERRESULTCONTEXT_ACCESS_FLOW_NAME FilterResultContext = "AccessFlowName"
	FILTERRESULTCONTEXT_TRIGGER_TYPE     FilterResultContext = "TriggerType"
)

// All allowed values of FilterResultContext enum
var AllowedFilterResultContextEnumValues = []FilterResultContext{
	"Approvers",
	"AccessTarget",
	"Grantees",
	"AccessFlowName",
	"TriggerType",
}

func (v *FilterResultContext) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FilterResultContext(value)
	for _, existing := range AllowedFilterResultContextEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FilterResultContext", value)
}

// NewFilterResultContextFromValue returns a pointer to a valid FilterResultContext
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilterResultContextFromValue(v string) (*FilterResultContext, error) {
	ev := FilterResultContext(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilterResultContext: valid values are %v", v, AllowedFilterResultContextEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilterResultContext) IsValid() bool {
	for _, existing := range AllowedFilterResultContextEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilterResultContext value
func (v FilterResultContext) Ptr() *FilterResultContext {
	return &v
}

type NullableFilterResultContext struct {
	value *FilterResultContext
	isSet bool
}

func (v NullableFilterResultContext) Get() *FilterResultContext {
	return v.value
}

func (v *NullableFilterResultContext) Set(val *FilterResultContext) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterResultContext) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterResultContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterResultContext(val *FilterResultContext) *NullableFilterResultContext {
	return &NullableFilterResultContext{value: val, isSet: true}
}

func (v NullableFilterResultContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterResultContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}