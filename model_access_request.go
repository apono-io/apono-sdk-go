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

// checks if the AccessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequest{}

// AccessRequest struct for AccessRequest
type AccessRequest struct {
	RequestId         string            `json:"request_id"`
	FriendlyRequestId string            `json:"friendly_request_id"`
	UserId            string            `json:"user_id"`
	Status            AccessStatusModel `json:"status"`
	IntegrationId     string            `json:"integration_id"`
	ResourceIds       []string          `json:"resource_ids"`
	Permissions       []string          `json:"permissions"`
	Justification     string            `json:"justification"`
}

// NewAccessRequest instantiates a new AccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequest(requestId string, friendlyRequestId string, userId string, status AccessStatusModel, integrationId string, resourceIds []string, permissions []string, justification string) *AccessRequest {
	this := AccessRequest{}
	this.RequestId = requestId
	this.FriendlyRequestId = friendlyRequestId
	this.UserId = userId
	this.Status = status
	this.IntegrationId = integrationId
	this.ResourceIds = resourceIds
	this.Permissions = permissions
	this.Justification = justification
	return &this
}

// NewAccessRequestWithDefaults instantiates a new AccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestWithDefaults() *AccessRequest {
	this := AccessRequest{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *AccessRequest) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AccessRequest) SetRequestId(v string) {
	o.RequestId = v
}

// GetFriendlyRequestId returns the FriendlyRequestId field value
func (o *AccessRequest) GetFriendlyRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FriendlyRequestId
}

// GetFriendlyRequestIdOk returns a tuple with the FriendlyRequestId field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetFriendlyRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FriendlyRequestId, true
}

// SetFriendlyRequestId sets field value
func (o *AccessRequest) SetFriendlyRequestId(v string) {
	o.FriendlyRequestId = v
}

// GetUserId returns the UserId field value
func (o *AccessRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AccessRequest) SetUserId(v string) {
	o.UserId = v
}

// GetStatus returns the Status field value
func (o *AccessRequest) GetStatus() AccessStatusModel {
	if o == nil {
		var ret AccessStatusModel
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetStatusOk() (*AccessStatusModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccessRequest) SetStatus(v AccessStatusModel) {
	o.Status = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *AccessRequest) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *AccessRequest) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetResourceIds returns the ResourceIds field value
func (o *AccessRequest) GetResourceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetResourceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceIds, true
}

// SetResourceIds sets field value
func (o *AccessRequest) SetResourceIds(v []string) {
	o.ResourceIds = v
}

// GetPermissions returns the Permissions field value
func (o *AccessRequest) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *AccessRequest) SetPermissions(v []string) {
	o.Permissions = v
}

// GetJustification returns the Justification field value
func (o *AccessRequest) GetJustification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Justification
}

// GetJustificationOk returns a tuple with the Justification field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetJustificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Justification, true
}

// SetJustification sets field value
func (o *AccessRequest) SetJustification(v string) {
	o.Justification = v
}

func (o AccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["friendly_request_id"] = o.FriendlyRequestId
	toSerialize["user_id"] = o.UserId
	toSerialize["status"] = o.Status
	toSerialize["integration_id"] = o.IntegrationId
	toSerialize["resource_ids"] = o.ResourceIds
	toSerialize["permissions"] = o.Permissions
	toSerialize["justification"] = o.Justification
	return toSerialize, nil
}

type NullableAccessRequest struct {
	value *AccessRequest
	isSet bool
}

func (v NullableAccessRequest) Get() *AccessRequest {
	return v.value
}

func (v *NullableAccessRequest) Set(val *AccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequest(val *AccessRequest) *NullableAccessRequest {
	return &NullableAccessRequest{value: val, isSet: true}
}

func (v NullableAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
