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

// checks if the UserModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModel{}

// UserModel struct for UserModel
type UserModel struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// NewUserModel instantiates a new UserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModel(id string, email string, firstName string, lastName string) *UserModel {
	this := UserModel{}
	this.Id = id
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewUserModelWithDefaults instantiates a new UserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelWithDefaults() *UserModel {
	this := UserModel{}
	return &this
}

// GetId returns the Id field value
func (o *UserModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserModel) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *UserModel) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserModel) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *UserModel) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UserModel) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UserModel) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserModel) SetLastName(v string) {
	o.LastName = v
}

func (o UserModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["email"] = o.Email
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	return toSerialize, nil
}

type NullableUserModel struct {
	value *UserModel
	isSet bool
}

func (v NullableUserModel) Get() *UserModel {
	return v.value
}

func (v *NullableUserModel) Set(val *UserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModel(val *UserModel) *NullableUserModel {
	return &NullableUserModel{value: val, isSet: true}
}

func (v NullableUserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
