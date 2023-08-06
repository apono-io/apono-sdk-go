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

// checks if the SelectablePermissionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectablePermissionsResponse{}

// SelectablePermissionsResponse struct for SelectablePermissionsResponse
type SelectablePermissionsResponse struct {
	Data          []string       `json:"data"`
	Pagination    PaginationInfo `json:"pagination"`
	AllowMultiple bool           `json:"allow_multiple"`
}

// NewSelectablePermissionsResponse instantiates a new SelectablePermissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectablePermissionsResponse(data []string, pagination PaginationInfo, allowMultiple bool) *SelectablePermissionsResponse {
	this := SelectablePermissionsResponse{}
	this.Data = data
	this.Pagination = pagination
	this.AllowMultiple = allowMultiple
	return &this
}

// NewSelectablePermissionsResponseWithDefaults instantiates a new SelectablePermissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectablePermissionsResponseWithDefaults() *SelectablePermissionsResponse {
	this := SelectablePermissionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SelectablePermissionsResponse) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SelectablePermissionsResponse) GetDataOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SelectablePermissionsResponse) SetData(v []string) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *SelectablePermissionsResponse) GetPagination() PaginationInfo {
	if o == nil {
		var ret PaginationInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *SelectablePermissionsResponse) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *SelectablePermissionsResponse) SetPagination(v PaginationInfo) {
	o.Pagination = v
}

// GetAllowMultiple returns the AllowMultiple field value
func (o *SelectablePermissionsResponse) GetAllowMultiple() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowMultiple
}

// GetAllowMultipleOk returns a tuple with the AllowMultiple field value
// and a boolean to check if the value has been set.
func (o *SelectablePermissionsResponse) GetAllowMultipleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowMultiple, true
}

// SetAllowMultiple sets field value
func (o *SelectablePermissionsResponse) SetAllowMultiple(v bool) {
	o.AllowMultiple = v
}

func (o SelectablePermissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectablePermissionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination
	toSerialize["allow_multiple"] = o.AllowMultiple
	return toSerialize, nil
}

type NullableSelectablePermissionsResponse struct {
	value *SelectablePermissionsResponse
	isSet bool
}

func (v NullableSelectablePermissionsResponse) Get() *SelectablePermissionsResponse {
	return v.value
}

func (v *NullableSelectablePermissionsResponse) Set(val *SelectablePermissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectablePermissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectablePermissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectablePermissionsResponse(val *SelectablePermissionsResponse) *NullableSelectablePermissionsResponse {
	return &NullableSelectablePermissionsResponse{value: val, isSet: true}
}

func (v NullableSelectablePermissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectablePermissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}