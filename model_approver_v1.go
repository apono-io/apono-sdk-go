/*
Apono

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apono

import (
	"encoding/json"
)

// checks if the ApproverV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApproverV1{}

// ApproverV1 struct for ApproverV1
type ApproverV1 struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

// NewApproverV1 instantiates a new ApproverV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproverV1(id string, type_ string) *ApproverV1 {
	this := ApproverV1{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewApproverV1WithDefaults instantiates a new ApproverV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproverV1WithDefaults() *ApproverV1 {
	this := ApproverV1{}
	return &this
}

// GetId returns the Id field value
func (o *ApproverV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApproverV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApproverV1) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ApproverV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApproverV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApproverV1) SetType(v string) {
	o.Type = v
}

func (o ApproverV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApproverV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableApproverV1 struct {
	value *ApproverV1
	isSet bool
}

func (v NullableApproverV1) Get() *ApproverV1 {
	return v.value
}

func (v *NullableApproverV1) Set(val *ApproverV1) {
	v.value = val
	v.isSet = true
}

func (v NullableApproverV1) IsSet() bool {
	return v.isSet
}

func (v *NullableApproverV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproverV1(val *ApproverV1) *NullableApproverV1 {
	return &NullableApproverV1{value: val, isSet: true}
}

func (v NullableApproverV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproverV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
