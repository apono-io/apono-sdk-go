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

// checks if the ResourceStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceStatusResponse{}

// ResourceStatusResponse struct for ResourceStatusResponse
type ResourceStatusResponse struct {
	Status  ResourceStatusV1 `json:"status"`
	Message NullableString   `json:"message,omitempty"`
}

// NewResourceStatusResponse instantiates a new ResourceStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceStatusResponse(status ResourceStatusV1) *ResourceStatusResponse {
	this := ResourceStatusResponse{}
	this.Status = status
	return &this
}

// NewResourceStatusResponseWithDefaults instantiates a new ResourceStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceStatusResponseWithDefaults() *ResourceStatusResponse {
	this := ResourceStatusResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *ResourceStatusResponse) GetStatus() ResourceStatusV1 {
	if o == nil {
		var ret ResourceStatusV1
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResourceStatusResponse) GetStatusOk() (*ResourceStatusV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResourceStatusResponse) SetStatus(v ResourceStatusV1) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceStatusResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceStatusResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ResourceStatusResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ResourceStatusResponse) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *ResourceStatusResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ResourceStatusResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o ResourceStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableResourceStatusResponse struct {
	value *ResourceStatusResponse
	isSet bool
}

func (v NullableResourceStatusResponse) Get() *ResourceStatusResponse {
	return v.value
}

func (v *NullableResourceStatusResponse) Set(val *ResourceStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceStatusResponse(val *ResourceStatusResponse) *NullableResourceStatusResponse {
	return &NullableResourceStatusResponse{value: val, isSet: true}
}

func (v NullableResourceStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
