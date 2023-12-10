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

// checks if the RequestStatusClientModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestStatusClientModel{}

// RequestStatusClientModel struct for RequestStatusClientModel
type RequestStatusClientModel struct {
	Status      AccessStatus      `json:"status"`
	Description NullableString    `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata"`
}

type _RequestStatusClientModel RequestStatusClientModel

// NewRequestStatusClientModel instantiates a new RequestStatusClientModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestStatusClientModel(status AccessStatus, metadata map[string]string) *RequestStatusClientModel {
	this := RequestStatusClientModel{}
	this.Status = status
	this.Metadata = metadata
	return &this
}

// NewRequestStatusClientModelWithDefaults instantiates a new RequestStatusClientModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestStatusClientModelWithDefaults() *RequestStatusClientModel {
	this := RequestStatusClientModel{}
	return &this
}

// GetStatus returns the Status field value
func (o *RequestStatusClientModel) GetStatus() AccessStatus {
	if o == nil {
		var ret AccessStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestStatusClientModel) GetStatusOk() (*AccessStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestStatusClientModel) SetStatus(v AccessStatus) {
	o.Status = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestStatusClientModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestStatusClientModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestStatusClientModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RequestStatusClientModel) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RequestStatusClientModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RequestStatusClientModel) UnsetDescription() {
	o.Description.Unset()
}

// GetMetadata returns the Metadata field value
func (o *RequestStatusClientModel) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *RequestStatusClientModel) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *RequestStatusClientModel) SetMetadata(v map[string]string) {
	o.Metadata = v
}

func (o RequestStatusClientModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestStatusClientModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *RequestStatusClientModel) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"metadata",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRequestStatusClientModel := _RequestStatusClientModel{}

	err = json.Unmarshal(bytes, &varRequestStatusClientModel)

	if err != nil {
		return err
	}

	*o = RequestStatusClientModel(varRequestStatusClientModel)

	return err
}

type NullableRequestStatusClientModel struct {
	value *RequestStatusClientModel
	isSet bool
}

func (v NullableRequestStatusClientModel) Get() *RequestStatusClientModel {
	return v.value
}

func (v *NullableRequestStatusClientModel) Set(val *RequestStatusClientModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestStatusClientModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestStatusClientModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestStatusClientModel(val *RequestStatusClientModel) *NullableRequestStatusClientModel {
	return &NullableRequestStatusClientModel{value: val, isSet: true}
}

func (v NullableRequestStatusClientModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestStatusClientModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
