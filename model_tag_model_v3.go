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

// checks if the TagModelV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagModelV3{}

// TagModelV3 struct for TagModelV3
type TagModelV3 struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NewTagModelV3 instantiates a new TagModelV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagModelV3(name string, value string) *TagModelV3 {
	this := TagModelV3{}
	this.Name = name
	this.Value = value
	return &this
}

// NewTagModelV3WithDefaults instantiates a new TagModelV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagModelV3WithDefaults() *TagModelV3 {
	this := TagModelV3{}
	return &this
}

// GetName returns the Name field value
func (o *TagModelV3) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagModelV3) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagModelV3) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *TagModelV3) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TagModelV3) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TagModelV3) SetValue(v string) {
	o.Value = v
}

func (o TagModelV3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagModelV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableTagModelV3 struct {
	value *TagModelV3
	isSet bool
}

func (v NullableTagModelV3) Get() *TagModelV3 {
	return v.value
}

func (v *NullableTagModelV3) Set(val *TagModelV3) {
	v.value = val
	v.isSet = true
}

func (v NullableTagModelV3) IsSet() bool {
	return v.isSet
}

func (v *NullableTagModelV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagModelV3(val *TagModelV3) *NullableTagModelV3 {
	return &NullableTagModelV3{value: val, isSet: true}
}

func (v NullableTagModelV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagModelV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}