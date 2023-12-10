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

// AccessRequestsScopeModel the model 'AccessRequestsScopeModel'
type AccessRequestsScopeModel string

// List of AccessRequestsScopeModel
const (
	ACCESSREQUESTSSCOPEMODEL_MY_REQUESTS AccessRequestsScopeModel = "MyRequests"
	ACCESSREQUESTSSCOPEMODEL_MY_TASKS    AccessRequestsScopeModel = "MyTasks"
)

// All allowed values of AccessRequestsScopeModel enum
var AllowedAccessRequestsScopeModelEnumValues = []AccessRequestsScopeModel{
	"MyRequests",
	"MyTasks",
}

func (v *AccessRequestsScopeModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessRequestsScopeModel(value)
	for _, existing := range AllowedAccessRequestsScopeModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessRequestsScopeModel", value)
}

// NewAccessRequestsScopeModelFromValue returns a pointer to a valid AccessRequestsScopeModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessRequestsScopeModelFromValue(v string) (*AccessRequestsScopeModel, error) {
	ev := AccessRequestsScopeModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessRequestsScopeModel: valid values are %v", v, AllowedAccessRequestsScopeModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessRequestsScopeModel) IsValid() bool {
	for _, existing := range AllowedAccessRequestsScopeModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessRequestsScopeModel value
func (v AccessRequestsScopeModel) Ptr() *AccessRequestsScopeModel {
	return &v
}

type NullableAccessRequestsScopeModel struct {
	value *AccessRequestsScopeModel
	isSet bool
}

func (v NullableAccessRequestsScopeModel) Get() *AccessRequestsScopeModel {
	return v.value
}

func (v *NullableAccessRequestsScopeModel) Set(val *AccessRequestsScopeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestsScopeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestsScopeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestsScopeModel(val *AccessRequestsScopeModel) *NullableAccessRequestsScopeModel {
	return &NullableAccessRequestsScopeModel{value: val, isSet: true}
}

func (v NullableAccessRequestsScopeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestsScopeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
