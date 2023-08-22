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

// checks if the UpdateResourceUserTagsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateResourceUserTagsRequest{}

// UpdateResourceUserTagsRequest struct for UpdateResourceUserTagsRequest
type UpdateResourceUserTagsRequest struct {
	Tags map[string]string `json:"tags"`
}

// NewUpdateResourceUserTagsRequest instantiates a new UpdateResourceUserTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateResourceUserTagsRequest(tags map[string]string) *UpdateResourceUserTagsRequest {
	this := UpdateResourceUserTagsRequest{}
	this.Tags = tags
	return &this
}

// NewUpdateResourceUserTagsRequestWithDefaults instantiates a new UpdateResourceUserTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateResourceUserTagsRequestWithDefaults() *UpdateResourceUserTagsRequest {
	this := UpdateResourceUserTagsRequest{}
	return &this
}

// GetTags returns the Tags field value
func (o *UpdateResourceUserTagsRequest) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceUserTagsRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *UpdateResourceUserTagsRequest) SetTags(v map[string]string) {
	o.Tags = v
}

func (o UpdateResourceUserTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateResourceUserTagsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

type NullableUpdateResourceUserTagsRequest struct {
	value *UpdateResourceUserTagsRequest
	isSet bool
}

func (v NullableUpdateResourceUserTagsRequest) Get() *UpdateResourceUserTagsRequest {
	return v.value
}

func (v *NullableUpdateResourceUserTagsRequest) Set(val *UpdateResourceUserTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateResourceUserTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateResourceUserTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateResourceUserTagsRequest(val *UpdateResourceUserTagsRequest) *NullableUpdateResourceUserTagsRequest {
	return &NullableUpdateResourceUserTagsRequest{value: val, isSet: true}
}

func (v NullableUpdateResourceUserTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateResourceUserTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}