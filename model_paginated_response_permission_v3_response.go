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

// checks if the PaginatedResponsePermissionV3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponsePermissionV3Response{}

// PaginatedResponsePermissionV3Response struct for PaginatedResponsePermissionV3Response
type PaginatedResponsePermissionV3Response struct {
	Data       []PermissionV3 `json:"data"`
	Pagination PaginationInfo `json:"pagination"`
}

// NewPaginatedResponsePermissionV3Response instantiates a new PaginatedResponsePermissionV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponsePermissionV3Response(data []PermissionV3, pagination PaginationInfo) *PaginatedResponsePermissionV3Response {
	this := PaginatedResponsePermissionV3Response{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewPaginatedResponsePermissionV3ResponseWithDefaults instantiates a new PaginatedResponsePermissionV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponsePermissionV3ResponseWithDefaults() *PaginatedResponsePermissionV3Response {
	this := PaginatedResponsePermissionV3Response{}
	return &this
}

// GetData returns the Data field value
func (o *PaginatedResponsePermissionV3Response) GetData() []PermissionV3 {
	if o == nil {
		var ret []PermissionV3
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponsePermissionV3Response) GetDataOk() ([]PermissionV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedResponsePermissionV3Response) SetData(v []PermissionV3) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *PaginatedResponsePermissionV3Response) GetPagination() PaginationInfo {
	if o == nil {
		var ret PaginationInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponsePermissionV3Response) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedResponsePermissionV3Response) SetPagination(v PaginationInfo) {
	o.Pagination = v
}

func (o PaginatedResponsePermissionV3Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponsePermissionV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullablePaginatedResponsePermissionV3Response struct {
	value *PaginatedResponsePermissionV3Response
	isSet bool
}

func (v NullablePaginatedResponsePermissionV3Response) Get() *PaginatedResponsePermissionV3Response {
	return v.value
}

func (v *NullablePaginatedResponsePermissionV3Response) Set(val *PaginatedResponsePermissionV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponsePermissionV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponsePermissionV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponsePermissionV3Response(val *PaginatedResponsePermissionV3Response) *NullablePaginatedResponsePermissionV3Response {
	return &NullablePaginatedResponsePermissionV3Response{value: val, isSet: true}
}

func (v NullablePaginatedResponsePermissionV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponsePermissionV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
