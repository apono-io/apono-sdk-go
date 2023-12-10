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

// checks if the AccessRequestClientModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestClientModel{}

// AccessRequestClientModel struct for AccessRequestClientModel
type AccessRequestClientModel struct {
	Id             string                                    `json:"id"`
	Requestor      UserClientModel                           `json:"requestor"`
	CreationTime   Instant                                   `json:"creation_time"`
	RevocationTime NullableInstant                           `json:"revocation_time,omitempty"`
	Status         RequestStatusClientModel                  `json:"status"`
	Justification  NullableString                            `json:"justification,omitempty"`
	AccessGroups   []AccessGroupClientModel                  `json:"access_groups"`
	Bundle         NullableAccessRequestClientModelBundle    `json:"bundle,omitempty"`
	Challenge      NullableAccessRequestClientModelChallenge `json:"challenge,omitempty"`
}

type _AccessRequestClientModel AccessRequestClientModel

// NewAccessRequestClientModel instantiates a new AccessRequestClientModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestClientModel(id string, requestor UserClientModel, creationTime Instant, status RequestStatusClientModel, accessGroups []AccessGroupClientModel) *AccessRequestClientModel {
	this := AccessRequestClientModel{}
	this.Id = id
	this.Requestor = requestor
	this.CreationTime = creationTime
	this.Status = status
	this.AccessGroups = accessGroups
	return &this
}

// NewAccessRequestClientModelWithDefaults instantiates a new AccessRequestClientModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestClientModelWithDefaults() *AccessRequestClientModel {
	this := AccessRequestClientModel{}
	return &this
}

// GetId returns the Id field value
func (o *AccessRequestClientModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessRequestClientModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessRequestClientModel) SetId(v string) {
	o.Id = v
}

// GetRequestor returns the Requestor field value
func (o *AccessRequestClientModel) GetRequestor() UserClientModel {
	if o == nil {
		var ret UserClientModel
		return ret
	}

	return o.Requestor
}

// GetRequestorOk returns a tuple with the Requestor field value
// and a boolean to check if the value has been set.
func (o *AccessRequestClientModel) GetRequestorOk() (*UserClientModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requestor, true
}

// SetRequestor sets field value
func (o *AccessRequestClientModel) SetRequestor(v UserClientModel) {
	o.Requestor = v
}

// GetCreationTime returns the CreationTime field value
func (o *AccessRequestClientModel) GetCreationTime() Instant {
	if o == nil {
		var ret Instant
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *AccessRequestClientModel) GetCreationTimeOk() (*Instant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *AccessRequestClientModel) SetCreationTime(v Instant) {
	o.CreationTime = v
}

// GetRevocationTime returns the RevocationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestClientModel) GetRevocationTime() Instant {
	if o == nil || IsNil(o.RevocationTime.Get()) {
		var ret Instant
		return ret
	}
	return *o.RevocationTime.Get()
}

// GetRevocationTimeOk returns a tuple with the RevocationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestClientModel) GetRevocationTimeOk() (*Instant, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevocationTime.Get(), o.RevocationTime.IsSet()
}

// HasRevocationTime returns a boolean if a field has been set.
func (o *AccessRequestClientModel) HasRevocationTime() bool {
	if o != nil && o.RevocationTime.IsSet() {
		return true
	}

	return false
}

// SetRevocationTime gets a reference to the given NullableInstant and assigns it to the RevocationTime field.
func (o *AccessRequestClientModel) SetRevocationTime(v Instant) {
	o.RevocationTime.Set(&v)
}

// SetRevocationTimeNil sets the value for RevocationTime to be an explicit nil
func (o *AccessRequestClientModel) SetRevocationTimeNil() {
	o.RevocationTime.Set(nil)
}

// UnsetRevocationTime ensures that no value is present for RevocationTime, not even an explicit nil
func (o *AccessRequestClientModel) UnsetRevocationTime() {
	o.RevocationTime.Unset()
}

// GetStatus returns the Status field value
func (o *AccessRequestClientModel) GetStatus() RequestStatusClientModel {
	if o == nil {
		var ret RequestStatusClientModel
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccessRequestClientModel) GetStatusOk() (*RequestStatusClientModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccessRequestClientModel) SetStatus(v RequestStatusClientModel) {
	o.Status = v
}

// GetJustification returns the Justification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestClientModel) GetJustification() string {
	if o == nil || IsNil(o.Justification.Get()) {
		var ret string
		return ret
	}
	return *o.Justification.Get()
}

// GetJustificationOk returns a tuple with the Justification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestClientModel) GetJustificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Justification.Get(), o.Justification.IsSet()
}

// HasJustification returns a boolean if a field has been set.
func (o *AccessRequestClientModel) HasJustification() bool {
	if o != nil && o.Justification.IsSet() {
		return true
	}

	return false
}

// SetJustification gets a reference to the given NullableString and assigns it to the Justification field.
func (o *AccessRequestClientModel) SetJustification(v string) {
	o.Justification.Set(&v)
}

// SetJustificationNil sets the value for Justification to be an explicit nil
func (o *AccessRequestClientModel) SetJustificationNil() {
	o.Justification.Set(nil)
}

// UnsetJustification ensures that no value is present for Justification, not even an explicit nil
func (o *AccessRequestClientModel) UnsetJustification() {
	o.Justification.Unset()
}

// GetAccessGroups returns the AccessGroups field value
func (o *AccessRequestClientModel) GetAccessGroups() []AccessGroupClientModel {
	if o == nil {
		var ret []AccessGroupClientModel
		return ret
	}

	return o.AccessGroups
}

// GetAccessGroupsOk returns a tuple with the AccessGroups field value
// and a boolean to check if the value has been set.
func (o *AccessRequestClientModel) GetAccessGroupsOk() ([]AccessGroupClientModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessGroups, true
}

// SetAccessGroups sets field value
func (o *AccessRequestClientModel) SetAccessGroups(v []AccessGroupClientModel) {
	o.AccessGroups = v
}

// GetBundle returns the Bundle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestClientModel) GetBundle() AccessRequestClientModelBundle {
	if o == nil || IsNil(o.Bundle.Get()) {
		var ret AccessRequestClientModelBundle
		return ret
	}
	return *o.Bundle.Get()
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestClientModel) GetBundleOk() (*AccessRequestClientModelBundle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bundle.Get(), o.Bundle.IsSet()
}

// HasBundle returns a boolean if a field has been set.
func (o *AccessRequestClientModel) HasBundle() bool {
	if o != nil && o.Bundle.IsSet() {
		return true
	}

	return false
}

// SetBundle gets a reference to the given NullableAccessRequestClientModelBundle and assigns it to the Bundle field.
func (o *AccessRequestClientModel) SetBundle(v AccessRequestClientModelBundle) {
	o.Bundle.Set(&v)
}

// SetBundleNil sets the value for Bundle to be an explicit nil
func (o *AccessRequestClientModel) SetBundleNil() {
	o.Bundle.Set(nil)
}

// UnsetBundle ensures that no value is present for Bundle, not even an explicit nil
func (o *AccessRequestClientModel) UnsetBundle() {
	o.Bundle.Unset()
}

// GetChallenge returns the Challenge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequestClientModel) GetChallenge() AccessRequestClientModelChallenge {
	if o == nil || IsNil(o.Challenge.Get()) {
		var ret AccessRequestClientModelChallenge
		return ret
	}
	return *o.Challenge.Get()
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequestClientModel) GetChallengeOk() (*AccessRequestClientModelChallenge, bool) {
	if o == nil {
		return nil, false
	}
	return o.Challenge.Get(), o.Challenge.IsSet()
}

// HasChallenge returns a boolean if a field has been set.
func (o *AccessRequestClientModel) HasChallenge() bool {
	if o != nil && o.Challenge.IsSet() {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given NullableAccessRequestClientModelChallenge and assigns it to the Challenge field.
func (o *AccessRequestClientModel) SetChallenge(v AccessRequestClientModelChallenge) {
	o.Challenge.Set(&v)
}

// SetChallengeNil sets the value for Challenge to be an explicit nil
func (o *AccessRequestClientModel) SetChallengeNil() {
	o.Challenge.Set(nil)
}

// UnsetChallenge ensures that no value is present for Challenge, not even an explicit nil
func (o *AccessRequestClientModel) UnsetChallenge() {
	o.Challenge.Unset()
}

func (o AccessRequestClientModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestClientModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["requestor"] = o.Requestor
	toSerialize["creation_time"] = o.CreationTime
	if o.RevocationTime.IsSet() {
		toSerialize["revocation_time"] = o.RevocationTime.Get()
	}
	toSerialize["status"] = o.Status
	if o.Justification.IsSet() {
		toSerialize["justification"] = o.Justification.Get()
	}
	toSerialize["access_groups"] = o.AccessGroups
	if o.Bundle.IsSet() {
		toSerialize["bundle"] = o.Bundle.Get()
	}
	if o.Challenge.IsSet() {
		toSerialize["challenge"] = o.Challenge.Get()
	}
	return toSerialize, nil
}

func (o *AccessRequestClientModel) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"requestor",
		"creation_time",
		"status",
		"access_groups",
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

	varAccessRequestClientModel := _AccessRequestClientModel{}

	err = json.Unmarshal(bytes, &varAccessRequestClientModel)

	if err != nil {
		return err
	}

	*o = AccessRequestClientModel(varAccessRequestClientModel)

	return err
}

type NullableAccessRequestClientModel struct {
	value *AccessRequestClientModel
	isSet bool
}

func (v NullableAccessRequestClientModel) Get() *AccessRequestClientModel {
	return v.value
}

func (v *NullableAccessRequestClientModel) Set(val *AccessRequestClientModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestClientModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestClientModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestClientModel(val *AccessRequestClientModel) *NullableAccessRequestClientModel {
	return &NullableAccessRequestClientModel{value: val, isSet: true}
}

func (v NullableAccessRequestClientModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestClientModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
