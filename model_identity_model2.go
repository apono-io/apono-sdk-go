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

// checks if the IdentityModel2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityModel2{}

// IdentityModel2 struct for IdentityModel2
type IdentityModel2 struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Id   string `json:"id"`
}

// NewIdentityModel2 instantiates a new IdentityModel2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityModel2(type_ string, name string, id string) *IdentityModel2 {
	this := IdentityModel2{}
	this.Type = type_
	this.Name = name
	this.Id = id
	return &this
}

// NewIdentityModel2WithDefaults instantiates a new IdentityModel2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityModel2WithDefaults() *IdentityModel2 {
	this := IdentityModel2{}
	return &this
}

// GetType returns the Type field value
func (o *IdentityModel2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IdentityModel2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IdentityModel2) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *IdentityModel2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityModel2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityModel2) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *IdentityModel2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdentityModel2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdentityModel2) SetId(v string) {
	o.Id = v
}

func (o IdentityModel2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityModel2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableIdentityModel2 struct {
	value *IdentityModel2
	isSet bool
}

func (v NullableIdentityModel2) Get() *IdentityModel2 {
	return v.value
}

func (v *NullableIdentityModel2) Set(val *IdentityModel2) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityModel2) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityModel2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityModel2(val *IdentityModel2) *NullableIdentityModel2 {
	return &NullableIdentityModel2{value: val, isSet: true}
}

func (v NullableIdentityModel2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityModel2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
