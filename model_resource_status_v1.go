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

// ResourceStatusV1 the model 'ResourceStatusV1'
type ResourceStatusV1 string

// List of ResourceStatusV1
const (
	RESOURCESTATUSV1_ACTIVE  ResourceStatusV1 = "Active"
	RESOURCESTATUSV1_ERROR   ResourceStatusV1 = "Error"
	RESOURCESTATUSV1_DELETED ResourceStatusV1 = "Deleted"
)

// All allowed values of ResourceStatusV1 enum
var AllowedResourceStatusV1EnumValues = []ResourceStatusV1{
	"Active",
	"Error",
	"Deleted",
}

func (v *ResourceStatusV1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceStatusV1(value)
	for _, existing := range AllowedResourceStatusV1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceStatusV1", value)
}

// NewResourceStatusV1FromValue returns a pointer to a valid ResourceStatusV1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceStatusV1FromValue(v string) (*ResourceStatusV1, error) {
	ev := ResourceStatusV1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceStatusV1: valid values are %v", v, AllowedResourceStatusV1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceStatusV1) IsValid() bool {
	for _, existing := range AllowedResourceStatusV1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResourceStatusV1 value
func (v ResourceStatusV1) Ptr() *ResourceStatusV1 {
	return &v
}

type NullableResourceStatusV1 struct {
	value *ResourceStatusV1
	isSet bool
}

func (v NullableResourceStatusV1) Get() *ResourceStatusV1 {
	return v.value
}

func (v *NullableResourceStatusV1) Set(val *ResourceStatusV1) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceStatusV1) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceStatusV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceStatusV1(val *ResourceStatusV1) *NullableResourceStatusV1 {
	return &NullableResourceStatusV1{value: val, isSet: true}
}

func (v NullableResourceStatusV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceStatusV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}