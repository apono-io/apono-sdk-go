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

// checks if the PaginatedAccessFlowV3SearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedAccessFlowV3SearchResponse{}

// PaginatedAccessFlowV3SearchResponse struct for PaginatedAccessFlowV3SearchResponse
type PaginatedAccessFlowV3SearchResponse struct {
	Data         []AccessFlowModelV3      `json:"data"`
	FilterResult []AccessFlowFilterResult `json:"filter_result,omitempty"`
}

// NewPaginatedAccessFlowV3SearchResponse instantiates a new PaginatedAccessFlowV3SearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAccessFlowV3SearchResponse(data []AccessFlowModelV3) *PaginatedAccessFlowV3SearchResponse {
	this := PaginatedAccessFlowV3SearchResponse{}
	this.Data = data
	return &this
}

// NewPaginatedAccessFlowV3SearchResponseWithDefaults instantiates a new PaginatedAccessFlowV3SearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAccessFlowV3SearchResponseWithDefaults() *PaginatedAccessFlowV3SearchResponse {
	this := PaginatedAccessFlowV3SearchResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PaginatedAccessFlowV3SearchResponse) GetData() []AccessFlowModelV3 {
	if o == nil {
		var ret []AccessFlowModelV3
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedAccessFlowV3SearchResponse) GetDataOk() ([]AccessFlowModelV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedAccessFlowV3SearchResponse) SetData(v []AccessFlowModelV3) {
	o.Data = v
}

// GetFilterResult returns the FilterResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedAccessFlowV3SearchResponse) GetFilterResult() []AccessFlowFilterResult {
	if o == nil {
		var ret []AccessFlowFilterResult
		return ret
	}
	return o.FilterResult
}

// GetFilterResultOk returns a tuple with the FilterResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedAccessFlowV3SearchResponse) GetFilterResultOk() ([]AccessFlowFilterResult, bool) {
	if o == nil || IsNil(o.FilterResult) {
		return nil, false
	}
	return o.FilterResult, true
}

// HasFilterResult returns a boolean if a field has been set.
func (o *PaginatedAccessFlowV3SearchResponse) HasFilterResult() bool {
	if o != nil && IsNil(o.FilterResult) {
		return true
	}

	return false
}

// SetFilterResult gets a reference to the given []AccessFlowFilterResult and assigns it to the FilterResult field.
func (o *PaginatedAccessFlowV3SearchResponse) SetFilterResult(v []AccessFlowFilterResult) {
	o.FilterResult = v
}

func (o PaginatedAccessFlowV3SearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedAccessFlowV3SearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if o.FilterResult != nil {
		toSerialize["filter_result"] = o.FilterResult
	}
	return toSerialize, nil
}

type NullablePaginatedAccessFlowV3SearchResponse struct {
	value *PaginatedAccessFlowV3SearchResponse
	isSet bool
}

func (v NullablePaginatedAccessFlowV3SearchResponse) Get() *PaginatedAccessFlowV3SearchResponse {
	return v.value
}

func (v *NullablePaginatedAccessFlowV3SearchResponse) Set(val *PaginatedAccessFlowV3SearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAccessFlowV3SearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAccessFlowV3SearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAccessFlowV3SearchResponse(val *PaginatedAccessFlowV3SearchResponse) *NullablePaginatedAccessFlowV3SearchResponse {
	return &NullablePaginatedAccessFlowV3SearchResponse{value: val, isSet: true}
}

func (v NullablePaginatedAccessFlowV3SearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAccessFlowV3SearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}