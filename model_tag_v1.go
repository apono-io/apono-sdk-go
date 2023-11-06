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

// checks if the TagV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagV1{}

// TagV1 struct for TagV1
type TagV1 struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NewTagV1 instantiates a new TagV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagV1(name string, value string) *TagV1 {
	this := TagV1{}
	this.Name = name
	this.Value = value
	return &this
}

// NewTagV1WithDefaults instantiates a new TagV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagV1WithDefaults() *TagV1 {
	this := TagV1{}
	return &this
}

// GetName returns the Name field value
func (o *TagV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagV1) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *TagV1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TagV1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TagV1) SetValue(v string) {
	o.Value = v
}

func (o TagV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableTagV1 struct {
	value *TagV1
	isSet bool
}

func (v NullableTagV1) Get() *TagV1 {
	return v.value
}

func (v *NullableTagV1) Set(val *TagV1) {
	v.value = val
	v.isSet = true
}

func (v NullableTagV1) IsSet() bool {
	return v.isSet
}

func (v *NullableTagV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagV1(val *TagV1) *NullableTagV1 {
	return &NullableTagV1{value: val, isSet: true}
}

func (v NullableTagV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}