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

// checks if the IntegrationConfigParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationConfigParam{}

// IntegrationConfigParam struct for IntegrationConfigParam
type IntegrationConfigParam struct {
	Id       string   `json:"id"`
	Label    string   `json:"label"`
	Values   []string `json:"values"`
	Default  string   `json:"default"`
	Optional bool     `json:"optional"`
}

// NewIntegrationConfigParam instantiates a new IntegrationConfigParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationConfigParam(id string, label string, values []string, default_ string, optional bool) *IntegrationConfigParam {
	this := IntegrationConfigParam{}
	this.Id = id
	this.Label = label
	this.Values = values
	this.Default = default_
	this.Optional = optional
	return &this
}

// NewIntegrationConfigParamWithDefaults instantiates a new IntegrationConfigParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationConfigParamWithDefaults() *IntegrationConfigParam {
	this := IntegrationConfigParam{}
	return &this
}

// GetId returns the Id field value
func (o *IntegrationConfigParam) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigParam) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IntegrationConfigParam) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *IntegrationConfigParam) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigParam) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *IntegrationConfigParam) SetLabel(v string) {
	o.Label = v
}

// GetValues returns the Values field value
func (o *IntegrationConfigParam) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigParam) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *IntegrationConfigParam) SetValues(v []string) {
	o.Values = v
}

// GetDefault returns the Default field value
func (o *IntegrationConfigParam) GetDefault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigParam) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *IntegrationConfigParam) SetDefault(v string) {
	o.Default = v
}

// GetOptional returns the Optional field value
func (o *IntegrationConfigParam) GetOptional() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigParam) GetOptionalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Optional, true
}

// SetOptional sets field value
func (o *IntegrationConfigParam) SetOptional(v bool) {
	o.Optional = v
}

func (o IntegrationConfigParam) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationConfigParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	toSerialize["values"] = o.Values
	toSerialize["default"] = o.Default
	toSerialize["optional"] = o.Optional
	return toSerialize, nil
}

type NullableIntegrationConfigParam struct {
	value *IntegrationConfigParam
	isSet bool
}

func (v NullableIntegrationConfigParam) Get() *IntegrationConfigParam {
	return v.value
}

func (v *NullableIntegrationConfigParam) Set(val *IntegrationConfigParam) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationConfigParam) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationConfigParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationConfigParam(val *IntegrationConfigParam) *NullableIntegrationConfigParam {
	return &NullableIntegrationConfigParam{value: val, isSet: true}
}

func (v NullableIntegrationConfigParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationConfigParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
