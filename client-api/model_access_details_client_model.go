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

// checks if the AccessDetailsClientModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessDetailsClientModel{}

// AccessDetailsClientModel struct for AccessDetailsClientModel
type AccessDetailsClientModel struct {
	Instructions        string                               `json:"instructions"`
	InstructionsMd      string                               `json:"instructions_md"`
	CanResetCredentials bool                                 `json:"can_reset_credentials"`
	Credentials         map[string]interface{}               `json:"credentials,omitempty"`
	Cli                 NullableString                       `json:"cli,omitempty"`
	Link                NullableAccessDetailsClientModelLink `json:"link,omitempty"`
	Plain               string                               `json:"plain"`
	Markdown            string                               `json:"markdown"`
}

type _AccessDetailsClientModel AccessDetailsClientModel

// NewAccessDetailsClientModel instantiates a new AccessDetailsClientModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessDetailsClientModel(instructions string, instructionsMd string, canResetCredentials bool, plain string, markdown string) *AccessDetailsClientModel {
	this := AccessDetailsClientModel{}
	this.Instructions = instructions
	this.InstructionsMd = instructionsMd
	this.CanResetCredentials = canResetCredentials
	this.Plain = plain
	this.Markdown = markdown
	return &this
}

// NewAccessDetailsClientModelWithDefaults instantiates a new AccessDetailsClientModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessDetailsClientModelWithDefaults() *AccessDetailsClientModel {
	this := AccessDetailsClientModel{}
	return &this
}

// GetInstructions returns the Instructions field value
func (o *AccessDetailsClientModel) GetInstructions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value
// and a boolean to check if the value has been set.
func (o *AccessDetailsClientModel) GetInstructionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instructions, true
}

// SetInstructions sets field value
func (o *AccessDetailsClientModel) SetInstructions(v string) {
	o.Instructions = v
}

// GetInstructionsMd returns the InstructionsMd field value
func (o *AccessDetailsClientModel) GetInstructionsMd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstructionsMd
}

// GetInstructionsMdOk returns a tuple with the InstructionsMd field value
// and a boolean to check if the value has been set.
func (o *AccessDetailsClientModel) GetInstructionsMdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstructionsMd, true
}

// SetInstructionsMd sets field value
func (o *AccessDetailsClientModel) SetInstructionsMd(v string) {
	o.InstructionsMd = v
}

// GetCanResetCredentials returns the CanResetCredentials field value
func (o *AccessDetailsClientModel) GetCanResetCredentials() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanResetCredentials
}

// GetCanResetCredentialsOk returns a tuple with the CanResetCredentials field value
// and a boolean to check if the value has been set.
func (o *AccessDetailsClientModel) GetCanResetCredentialsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanResetCredentials, true
}

// SetCanResetCredentials sets field value
func (o *AccessDetailsClientModel) SetCanResetCredentials(v bool) {
	o.CanResetCredentials = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessDetailsClientModel) GetCredentials() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessDetailsClientModel) GetCredentialsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Credentials) {
		return map[string]interface{}{}, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *AccessDetailsClientModel) HasCredentials() bool {
	if o != nil && IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given map[string]interface{} and assigns it to the Credentials field.
func (o *AccessDetailsClientModel) SetCredentials(v map[string]interface{}) {
	o.Credentials = v
}

// GetCli returns the Cli field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessDetailsClientModel) GetCli() string {
	if o == nil || IsNil(o.Cli.Get()) {
		var ret string
		return ret
	}
	return *o.Cli.Get()
}

// GetCliOk returns a tuple with the Cli field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessDetailsClientModel) GetCliOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cli.Get(), o.Cli.IsSet()
}

// HasCli returns a boolean if a field has been set.
func (o *AccessDetailsClientModel) HasCli() bool {
	if o != nil && o.Cli.IsSet() {
		return true
	}

	return false
}

// SetCli gets a reference to the given NullableString and assigns it to the Cli field.
func (o *AccessDetailsClientModel) SetCli(v string) {
	o.Cli.Set(&v)
}

// SetCliNil sets the value for Cli to be an explicit nil
func (o *AccessDetailsClientModel) SetCliNil() {
	o.Cli.Set(nil)
}

// UnsetCli ensures that no value is present for Cli, not even an explicit nil
func (o *AccessDetailsClientModel) UnsetCli() {
	o.Cli.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessDetailsClientModel) GetLink() AccessDetailsClientModelLink {
	if o == nil || IsNil(o.Link.Get()) {
		var ret AccessDetailsClientModelLink
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessDetailsClientModel) GetLinkOk() (*AccessDetailsClientModelLink, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *AccessDetailsClientModel) HasLink() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableAccessDetailsClientModelLink and assigns it to the Link field.
func (o *AccessDetailsClientModel) SetLink(v AccessDetailsClientModelLink) {
	o.Link.Set(&v)
}

// SetLinkNil sets the value for Link to be an explicit nil
func (o *AccessDetailsClientModel) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *AccessDetailsClientModel) UnsetLink() {
	o.Link.Unset()
}

// GetPlain returns the Plain field value
func (o *AccessDetailsClientModel) GetPlain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plain
}

// GetPlainOk returns a tuple with the Plain field value
// and a boolean to check if the value has been set.
func (o *AccessDetailsClientModel) GetPlainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plain, true
}

// SetPlain sets field value
func (o *AccessDetailsClientModel) SetPlain(v string) {
	o.Plain = v
}

// GetMarkdown returns the Markdown field value
func (o *AccessDetailsClientModel) GetMarkdown() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Markdown
}

// GetMarkdownOk returns a tuple with the Markdown field value
// and a boolean to check if the value has been set.
func (o *AccessDetailsClientModel) GetMarkdownOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Markdown, true
}

// SetMarkdown sets field value
func (o *AccessDetailsClientModel) SetMarkdown(v string) {
	o.Markdown = v
}

func (o AccessDetailsClientModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessDetailsClientModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instructions"] = o.Instructions
	toSerialize["instructions_md"] = o.InstructionsMd
	toSerialize["can_reset_credentials"] = o.CanResetCredentials
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Cli.IsSet() {
		toSerialize["cli"] = o.Cli.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	toSerialize["plain"] = o.Plain
	toSerialize["markdown"] = o.Markdown
	return toSerialize, nil
}

func (o *AccessDetailsClientModel) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instructions",
		"instructions_md",
		"can_reset_credentials",
		"plain",
		"markdown",
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

	varAccessDetailsClientModel := _AccessDetailsClientModel{}

	err = json.Unmarshal(bytes, &varAccessDetailsClientModel)

	if err != nil {
		return err
	}

	*o = AccessDetailsClientModel(varAccessDetailsClientModel)

	return err
}

type NullableAccessDetailsClientModel struct {
	value *AccessDetailsClientModel
	isSet bool
}

func (v NullableAccessDetailsClientModel) Get() *AccessDetailsClientModel {
	return v.value
}

func (v *NullableAccessDetailsClientModel) Set(val *AccessDetailsClientModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDetailsClientModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDetailsClientModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDetailsClientModel(val *AccessDetailsClientModel) *NullableAccessDetailsClientModel {
	return &NullableAccessDetailsClientModel{value: val, isSet: true}
}

func (v NullableAccessDetailsClientModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessDetailsClientModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
