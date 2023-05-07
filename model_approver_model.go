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

// checks if the ApproverModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApproverModel{}

// ApproverModel struct for ApproverModel
type ApproverModel struct {
	Id   string            `json:"id"`
	Type ApproverTypeModel `json:"type"`
}

// NewApproverModel instantiates a new ApproverModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproverModel(id string, type_ ApproverTypeModel) *ApproverModel {
	this := ApproverModel{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewApproverModelWithDefaults instantiates a new ApproverModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproverModelWithDefaults() *ApproverModel {
	this := ApproverModel{}
	return &this
}

// GetId returns the Id field value
func (o *ApproverModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApproverModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApproverModel) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ApproverModel) GetType() ApproverTypeModel {
	if o == nil {
		var ret ApproverTypeModel
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApproverModel) GetTypeOk() (*ApproverTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApproverModel) SetType(v ApproverTypeModel) {
	o.Type = v
}

func (o ApproverModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApproverModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableApproverModel struct {
	value *ApproverModel
	isSet bool
}

func (v NullableApproverModel) Get() *ApproverModel {
	return v.value
}

func (v *NullableApproverModel) Set(val *ApproverModel) {
	v.value = val
	v.isSet = true
}

func (v NullableApproverModel) IsSet() bool {
	return v.isSet
}

func (v *NullableApproverModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproverModel(val *ApproverModel) *NullableApproverModel {
	return &NullableApproverModel{value: val, isSet: true}
}

func (v NullableApproverModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproverModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
