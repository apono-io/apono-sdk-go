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

// checks if the PaginatedResponseAccessBundleV1Model type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponseAccessBundleV1Model{}

// PaginatedResponseAccessBundleV1Model struct for PaginatedResponseAccessBundleV1Model
type PaginatedResponseAccessBundleV1Model struct {
	Data       []AccessBundleV1 `json:"data"`
	Pagination PaginationInfo   `json:"pagination"`
}

// NewPaginatedResponseAccessBundleV1Model instantiates a new PaginatedResponseAccessBundleV1Model object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseAccessBundleV1Model(data []AccessBundleV1, pagination PaginationInfo) *PaginatedResponseAccessBundleV1Model {
	this := PaginatedResponseAccessBundleV1Model{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewPaginatedResponseAccessBundleV1ModelWithDefaults instantiates a new PaginatedResponseAccessBundleV1Model object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseAccessBundleV1ModelWithDefaults() *PaginatedResponseAccessBundleV1Model {
	this := PaginatedResponseAccessBundleV1Model{}
	return &this
}

// GetData returns the Data field value
func (o *PaginatedResponseAccessBundleV1Model) GetData() []AccessBundleV1 {
	if o == nil {
		var ret []AccessBundleV1
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponseAccessBundleV1Model) GetDataOk() ([]AccessBundleV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedResponseAccessBundleV1Model) SetData(v []AccessBundleV1) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *PaginatedResponseAccessBundleV1Model) GetPagination() PaginationInfo {
	if o == nil {
		var ret PaginationInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponseAccessBundleV1Model) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedResponseAccessBundleV1Model) SetPagination(v PaginationInfo) {
	o.Pagination = v
}

func (o PaginatedResponseAccessBundleV1Model) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponseAccessBundleV1Model) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullablePaginatedResponseAccessBundleV1Model struct {
	value *PaginatedResponseAccessBundleV1Model
	isSet bool
}

func (v NullablePaginatedResponseAccessBundleV1Model) Get() *PaginatedResponseAccessBundleV1Model {
	return v.value
}

func (v *NullablePaginatedResponseAccessBundleV1Model) Set(val *PaginatedResponseAccessBundleV1Model) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseAccessBundleV1Model) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseAccessBundleV1Model) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseAccessBundleV1Model(val *PaginatedResponseAccessBundleV1Model) *NullablePaginatedResponseAccessBundleV1Model {
	return &NullablePaginatedResponseAccessBundleV1Model{value: val, isSet: true}
}

func (v NullablePaginatedResponseAccessBundleV1Model) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseAccessBundleV1Model) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}