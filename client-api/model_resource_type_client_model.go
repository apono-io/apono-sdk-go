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

// checks if the ResourceTypeClientModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceTypeClientModel{}

// ResourceTypeClientModel struct for ResourceTypeClientModel
type ResourceTypeClientModel struct {
	Id                       string                 `json:"id"`
	Name                     string                 `json:"name"`
	DisplayPath              string                 `json:"display_path"`
	AllowMultiplePermissions bool                   `json:"allow_multiple_permissions"`
	Icons                    IconsConfigClientModel `json:"icons"`
}

type _ResourceTypeClientModel ResourceTypeClientModel

// NewResourceTypeClientModel instantiates a new ResourceTypeClientModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTypeClientModel(id string, name string, displayPath string, allowMultiplePermissions bool, icons IconsConfigClientModel) *ResourceTypeClientModel {
	this := ResourceTypeClientModel{}
	this.Id = id
	this.Name = name
	this.DisplayPath = displayPath
	this.AllowMultiplePermissions = allowMultiplePermissions
	this.Icons = icons
	return &this
}

// NewResourceTypeClientModelWithDefaults instantiates a new ResourceTypeClientModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeClientModelWithDefaults() *ResourceTypeClientModel {
	this := ResourceTypeClientModel{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceTypeClientModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceTypeClientModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceTypeClientModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResourceTypeClientModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceTypeClientModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceTypeClientModel) SetName(v string) {
	o.Name = v
}

// GetDisplayPath returns the DisplayPath field value
func (o *ResourceTypeClientModel) GetDisplayPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayPath
}

// GetDisplayPathOk returns a tuple with the DisplayPath field value
// and a boolean to check if the value has been set.
func (o *ResourceTypeClientModel) GetDisplayPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayPath, true
}

// SetDisplayPath sets field value
func (o *ResourceTypeClientModel) SetDisplayPath(v string) {
	o.DisplayPath = v
}

// GetAllowMultiplePermissions returns the AllowMultiplePermissions field value
func (o *ResourceTypeClientModel) GetAllowMultiplePermissions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowMultiplePermissions
}

// GetAllowMultiplePermissionsOk returns a tuple with the AllowMultiplePermissions field value
// and a boolean to check if the value has been set.
func (o *ResourceTypeClientModel) GetAllowMultiplePermissionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowMultiplePermissions, true
}

// SetAllowMultiplePermissions sets field value
func (o *ResourceTypeClientModel) SetAllowMultiplePermissions(v bool) {
	o.AllowMultiplePermissions = v
}

// GetIcons returns the Icons field value
func (o *ResourceTypeClientModel) GetIcons() IconsConfigClientModel {
	if o == nil {
		var ret IconsConfigClientModel
		return ret
	}

	return o.Icons
}

// GetIconsOk returns a tuple with the Icons field value
// and a boolean to check if the value has been set.
func (o *ResourceTypeClientModel) GetIconsOk() (*IconsConfigClientModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icons, true
}

// SetIcons sets field value
func (o *ResourceTypeClientModel) SetIcons(v IconsConfigClientModel) {
	o.Icons = v
}

func (o ResourceTypeClientModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceTypeClientModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["display_path"] = o.DisplayPath
	toSerialize["allow_multiple_permissions"] = o.AllowMultiplePermissions
	toSerialize["icons"] = o.Icons
	return toSerialize, nil
}

func (o *ResourceTypeClientModel) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"display_path",
		"allow_multiple_permissions",
		"icons",
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

	varResourceTypeClientModel := _ResourceTypeClientModel{}

	err = json.Unmarshal(bytes, &varResourceTypeClientModel)

	if err != nil {
		return err
	}

	*o = ResourceTypeClientModel(varResourceTypeClientModel)

	return err
}

type NullableResourceTypeClientModel struct {
	value *ResourceTypeClientModel
	isSet bool
}

func (v NullableResourceTypeClientModel) Get() *ResourceTypeClientModel {
	return v.value
}

func (v *NullableResourceTypeClientModel) Set(val *ResourceTypeClientModel) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeClientModel) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeClientModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeClientModel(val *ResourceTypeClientModel) *NullableResourceTypeClientModel {
	return &NullableResourceTypeClientModel{value: val, isSet: true}
}

func (v NullableResourceTypeClientModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeClientModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
