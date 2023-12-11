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

// checks if the CredentialsClientModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsClientModel{}

// CredentialsClientModel struct for CredentialsClientModel
type CredentialsClientModel struct {
	Id       string `json:"id"`
	Status   string `json:"status"`
	CanReset bool   `json:"can_reset"`
}

type _CredentialsClientModel CredentialsClientModel

// NewCredentialsClientModel instantiates a new CredentialsClientModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsClientModel(id string, status string, canReset bool) *CredentialsClientModel {
	this := CredentialsClientModel{}
	this.Id = id
	this.Status = status
	this.CanReset = canReset
	return &this
}

// NewCredentialsClientModelWithDefaults instantiates a new CredentialsClientModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsClientModelWithDefaults() *CredentialsClientModel {
	this := CredentialsClientModel{}
	return &this
}

// GetId returns the Id field value
func (o *CredentialsClientModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CredentialsClientModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CredentialsClientModel) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *CredentialsClientModel) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CredentialsClientModel) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CredentialsClientModel) SetStatus(v string) {
	o.Status = v
}

// GetCanReset returns the CanReset field value
func (o *CredentialsClientModel) GetCanReset() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanReset
}

// GetCanResetOk returns a tuple with the CanReset field value
// and a boolean to check if the value has been set.
func (o *CredentialsClientModel) GetCanResetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanReset, true
}

// SetCanReset sets field value
func (o *CredentialsClientModel) SetCanReset(v bool) {
	o.CanReset = v
}

func (o CredentialsClientModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsClientModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["can_reset"] = o.CanReset
	return toSerialize, nil
}

func (o *CredentialsClientModel) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"can_reset",
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

	varCredentialsClientModel := _CredentialsClientModel{}

	err = json.Unmarshal(bytes, &varCredentialsClientModel)

	if err != nil {
		return err
	}

	*o = CredentialsClientModel(varCredentialsClientModel)

	return err
}

type NullableCredentialsClientModel struct {
	value *CredentialsClientModel
	isSet bool
}

func (v NullableCredentialsClientModel) Get() *CredentialsClientModel {
	return v.value
}

func (v *NullableCredentialsClientModel) Set(val *CredentialsClientModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsClientModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsClientModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsClientModel(val *CredentialsClientModel) *NullableCredentialsClientModel {
	return &NullableCredentialsClientModel{value: val, isSet: true}
}

func (v NullableCredentialsClientModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsClientModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}