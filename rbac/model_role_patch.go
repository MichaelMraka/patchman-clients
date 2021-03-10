/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac

import (
	"encoding/json"
)

// RolePatch struct for RolePatch
type RolePatch struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewRolePatch instantiates a new RolePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePatch() *RolePatch {
	this := RolePatch{}
	return &this
}

// NewRolePatchWithDefaults instantiates a new RolePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePatchWithDefaults() *RolePatch {
	this := RolePatch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RolePatch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePatch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RolePatch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RolePatch) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *RolePatch) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePatch) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RolePatch) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RolePatch) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RolePatch) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePatch) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RolePatch) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RolePatch) SetDescription(v string) {
	o.Description = &v
}

func (o RolePatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableRolePatch struct {
	value *RolePatch
	isSet bool
}

func (v NullableRolePatch) Get() *RolePatch {
	return v.value
}

func (v *NullableRolePatch) Set(val *RolePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePatch(val *RolePatch) *NullableRolePatch {
	return &NullableRolePatch{value: val, isSet: true}
}

func (v NullableRolePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


