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

// checks if the TimeframeV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeframeV1{}

// TimeframeV1 struct for TimeframeV1
type TimeframeV1 struct {
	StartOfDayTimeInSeconds int64       `json:"start_of_day_time_in_seconds"`
	EndOfDayTimeInSeconds   int64       `json:"end_of_day_time_in_seconds"`
	DaysInWeek              []DayOfWeek `json:"days_in_week"`
	TimeZone                string      `json:"time_zone"`
}

// NewTimeframeV1 instantiates a new TimeframeV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeframeV1(startOfDayTimeInSeconds int64, endOfDayTimeInSeconds int64, daysInWeek []DayOfWeek, timeZone string) *TimeframeV1 {
	this := TimeframeV1{}
	this.StartOfDayTimeInSeconds = startOfDayTimeInSeconds
	this.EndOfDayTimeInSeconds = endOfDayTimeInSeconds
	this.DaysInWeek = daysInWeek
	this.TimeZone = timeZone
	return &this
}

// NewTimeframeV1WithDefaults instantiates a new TimeframeV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeframeV1WithDefaults() *TimeframeV1 {
	this := TimeframeV1{}
	return &this
}

// GetStartOfDayTimeInSeconds returns the StartOfDayTimeInSeconds field value
func (o *TimeframeV1) GetStartOfDayTimeInSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartOfDayTimeInSeconds
}

// GetStartOfDayTimeInSecondsOk returns a tuple with the StartOfDayTimeInSeconds field value
// and a boolean to check if the value has been set.
func (o *TimeframeV1) GetStartOfDayTimeInSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartOfDayTimeInSeconds, true
}

// SetStartOfDayTimeInSeconds sets field value
func (o *TimeframeV1) SetStartOfDayTimeInSeconds(v int64) {
	o.StartOfDayTimeInSeconds = v
}

// GetEndOfDayTimeInSeconds returns the EndOfDayTimeInSeconds field value
func (o *TimeframeV1) GetEndOfDayTimeInSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndOfDayTimeInSeconds
}

// GetEndOfDayTimeInSecondsOk returns a tuple with the EndOfDayTimeInSeconds field value
// and a boolean to check if the value has been set.
func (o *TimeframeV1) GetEndOfDayTimeInSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndOfDayTimeInSeconds, true
}

// SetEndOfDayTimeInSeconds sets field value
func (o *TimeframeV1) SetEndOfDayTimeInSeconds(v int64) {
	o.EndOfDayTimeInSeconds = v
}

// GetDaysInWeek returns the DaysInWeek field value
func (o *TimeframeV1) GetDaysInWeek() []DayOfWeek {
	if o == nil {
		var ret []DayOfWeek
		return ret
	}

	return o.DaysInWeek
}

// GetDaysInWeekOk returns a tuple with the DaysInWeek field value
// and a boolean to check if the value has been set.
func (o *TimeframeV1) GetDaysInWeekOk() ([]DayOfWeek, bool) {
	if o == nil {
		return nil, false
	}
	return o.DaysInWeek, true
}

// SetDaysInWeek sets field value
func (o *TimeframeV1) SetDaysInWeek(v []DayOfWeek) {
	o.DaysInWeek = v
}

// GetTimeZone returns the TimeZone field value
func (o *TimeframeV1) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *TimeframeV1) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *TimeframeV1) SetTimeZone(v string) {
	o.TimeZone = v
}

func (o TimeframeV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeframeV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start_of_day_time_in_seconds"] = o.StartOfDayTimeInSeconds
	toSerialize["end_of_day_time_in_seconds"] = o.EndOfDayTimeInSeconds
	toSerialize["days_in_week"] = o.DaysInWeek
	toSerialize["time_zone"] = o.TimeZone
	return toSerialize, nil
}

type NullableTimeframeV1 struct {
	value *TimeframeV1
	isSet bool
}

func (v NullableTimeframeV1) Get() *TimeframeV1 {
	return v.value
}

func (v *NullableTimeframeV1) Set(val *TimeframeV1) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeframeV1) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeframeV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeframeV1(val *TimeframeV1) *NullableTimeframeV1 {
	return &NullableTimeframeV1{value: val, isSet: true}
}

func (v NullableTimeframeV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeframeV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
