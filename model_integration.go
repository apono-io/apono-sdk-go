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

// checks if the Integration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Integration{}

// Integration struct for Integration
type Integration struct {
	Id            string                 `json:"id"`
	Name          string                 `json:"name"`
	Type          string                 `json:"type"`
	Status        IntegrationStatus      `json:"status"`
	Details       NullableString         `json:"details,omitempty"`
	ProvisionerId NullableString         `json:"provisioner_id,omitempty"`
	Connection    map[string]interface{} `json:"connection,omitempty"`
	LastSyncTime  NullableInstant        `json:"last_sync_time,omitempty"`
	Metadata      map[string]interface{} `json:"metadata"`
	SecretConfig  map[string]interface{} `json:"secret_config,omitempty"`
}

// NewIntegration instantiates a new Integration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegration(id string, name string, type_ string, status IntegrationStatus, metadata map[string]interface{}) *Integration {
	this := Integration{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Status = status
	this.Metadata = metadata
	return &this
}

// NewIntegrationWithDefaults instantiates a new Integration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationWithDefaults() *Integration {
	this := Integration{}
	return &this
}

// GetId returns the Id field value
func (o *Integration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Integration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Integration) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Integration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Integration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Integration) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Integration) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Integration) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Integration) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *Integration) GetStatus() IntegrationStatus {
	if o == nil {
		var ret IntegrationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Integration) GetStatusOk() (*IntegrationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Integration) SetStatus(v IntegrationStatus) {
	o.Status = v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Integration) GetDetails() string {
	if o == nil || IsNil(o.Details.Get()) {
		var ret string
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Integration) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *Integration) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableString and assigns it to the Details field.
func (o *Integration) SetDetails(v string) {
	o.Details.Set(&v)
}

// SetDetailsNil sets the value for Details to be an explicit nil
func (o *Integration) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *Integration) UnsetDetails() {
	o.Details.Unset()
}

// GetProvisionerId returns the ProvisionerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Integration) GetProvisionerId() string {
	if o == nil || IsNil(o.ProvisionerId.Get()) {
		var ret string
		return ret
	}
	return *o.ProvisionerId.Get()
}

// GetProvisionerIdOk returns a tuple with the ProvisionerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Integration) GetProvisionerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvisionerId.Get(), o.ProvisionerId.IsSet()
}

// HasProvisionerId returns a boolean if a field has been set.
func (o *Integration) HasProvisionerId() bool {
	if o != nil && o.ProvisionerId.IsSet() {
		return true
	}

	return false
}

// SetProvisionerId gets a reference to the given NullableString and assigns it to the ProvisionerId field.
func (o *Integration) SetProvisionerId(v string) {
	o.ProvisionerId.Set(&v)
}

// SetProvisionerIdNil sets the value for ProvisionerId to be an explicit nil
func (o *Integration) SetProvisionerIdNil() {
	o.ProvisionerId.Set(nil)
}

// UnsetProvisionerId ensures that no value is present for ProvisionerId, not even an explicit nil
func (o *Integration) UnsetProvisionerId() {
	o.ProvisionerId.Unset()
}

// GetConnection returns the Connection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Integration) GetConnection() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Integration) GetConnectionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Connection) {
		return map[string]interface{}{}, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *Integration) HasConnection() bool {
	if o != nil && IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given map[string]interface{} and assigns it to the Connection field.
func (o *Integration) SetConnection(v map[string]interface{}) {
	o.Connection = v
}

// GetLastSyncTime returns the LastSyncTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Integration) GetLastSyncTime() Instant {
	if o == nil || IsNil(o.LastSyncTime.Get()) {
		var ret Instant
		return ret
	}
	return *o.LastSyncTime.Get()
}

// GetLastSyncTimeOk returns a tuple with the LastSyncTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Integration) GetLastSyncTimeOk() (*Instant, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSyncTime.Get(), o.LastSyncTime.IsSet()
}

// HasLastSyncTime returns a boolean if a field has been set.
func (o *Integration) HasLastSyncTime() bool {
	if o != nil && o.LastSyncTime.IsSet() {
		return true
	}

	return false
}

// SetLastSyncTime gets a reference to the given NullableInstant and assigns it to the LastSyncTime field.
func (o *Integration) SetLastSyncTime(v Instant) {
	o.LastSyncTime.Set(&v)
}

// SetLastSyncTimeNil sets the value for LastSyncTime to be an explicit nil
func (o *Integration) SetLastSyncTimeNil() {
	o.LastSyncTime.Set(nil)
}

// UnsetLastSyncTime ensures that no value is present for LastSyncTime, not even an explicit nil
func (o *Integration) UnsetLastSyncTime() {
	o.LastSyncTime.Unset()
}

// GetMetadata returns the Metadata field value
func (o *Integration) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Integration) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Integration) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetSecretConfig returns the SecretConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Integration) GetSecretConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SecretConfig
}

// GetSecretConfigOk returns a tuple with the SecretConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Integration) GetSecretConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SecretConfig) {
		return map[string]interface{}{}, false
	}
	return o.SecretConfig, true
}

// HasSecretConfig returns a boolean if a field has been set.
func (o *Integration) HasSecretConfig() bool {
	if o != nil && IsNil(o.SecretConfig) {
		return true
	}

	return false
}

// SetSecretConfig gets a reference to the given map[string]interface{} and assigns it to the SecretConfig field.
func (o *Integration) SetSecretConfig(v map[string]interface{}) {
	o.SecretConfig = v
}

func (o Integration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Integration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.ProvisionerId.IsSet() {
		toSerialize["provisioner_id"] = o.ProvisionerId.Get()
	}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	if o.LastSyncTime.IsSet() {
		toSerialize["last_sync_time"] = o.LastSyncTime.Get()
	}
	toSerialize["metadata"] = o.Metadata
	if o.SecretConfig != nil {
		toSerialize["secret_config"] = o.SecretConfig
	}
	return toSerialize, nil
}

type NullableIntegration struct {
	value *Integration
	isSet bool
}

func (v NullableIntegration) Get() *Integration {
	return v.value
}

func (v *NullableIntegration) Set(val *Integration) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegration(val *Integration) *NullableIntegration {
	return &NullableIntegration{value: val, isSet: true}
}

func (v NullableIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
