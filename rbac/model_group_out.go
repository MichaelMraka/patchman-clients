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
	"time"
)

// GroupOut struct for GroupOut
type GroupOut struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Uuid string `json:"uuid"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	PrincipalCount *int32 `json:"principalCount,omitempty"`
	RoleCount *int32 `json:"roleCount,omitempty"`
	System *bool `json:"system,omitempty"`
	PlatformDefault *bool `json:"platform_default,omitempty"`
}

// NewGroupOut instantiates a new GroupOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupOut(name string, uuid string, created time.Time, modified time.Time, ) *GroupOut {
	this := GroupOut{}
	this.Name = name
	this.Uuid = uuid
	this.Created = created
	this.Modified = modified
	var system bool = false
	this.System = &system
	var platformDefault bool = false
	this.PlatformDefault = &platformDefault
	return &this
}

// NewGroupOutWithDefaults instantiates a new GroupOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupOutWithDefaults() *GroupOut {
	this := GroupOut{}
	var system bool = false
	this.System = &system
	var platformDefault bool = false
	this.PlatformDefault = &platformDefault
	return &this
}

// GetName returns the Name field value
func (o *GroupOut) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupOut) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupOut) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupOut) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOut) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupOut) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupOut) SetDescription(v string) {
	o.Description = &v
}

// GetUuid returns the Uuid field value
func (o *GroupOut) GetUuid() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *GroupOut) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *GroupOut) SetUuid(v string) {
	o.Uuid = v
}

// GetCreated returns the Created field value
func (o *GroupOut) GetCreated() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GroupOut) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GroupOut) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *GroupOut) GetModified() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *GroupOut) GetModifiedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *GroupOut) SetModified(v time.Time) {
	o.Modified = v
}

// GetPrincipalCount returns the PrincipalCount field value if set, zero value otherwise.
func (o *GroupOut) GetPrincipalCount() int32 {
	if o == nil || o.PrincipalCount == nil {
		var ret int32
		return ret
	}
	return *o.PrincipalCount
}

// GetPrincipalCountOk returns a tuple with the PrincipalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOut) GetPrincipalCountOk() (*int32, bool) {
	if o == nil || o.PrincipalCount == nil {
		return nil, false
	}
	return o.PrincipalCount, true
}

// HasPrincipalCount returns a boolean if a field has been set.
func (o *GroupOut) HasPrincipalCount() bool {
	if o != nil && o.PrincipalCount != nil {
		return true
	}

	return false
}

// SetPrincipalCount gets a reference to the given int32 and assigns it to the PrincipalCount field.
func (o *GroupOut) SetPrincipalCount(v int32) {
	o.PrincipalCount = &v
}

// GetRoleCount returns the RoleCount field value if set, zero value otherwise.
func (o *GroupOut) GetRoleCount() int32 {
	if o == nil || o.RoleCount == nil {
		var ret int32
		return ret
	}
	return *o.RoleCount
}

// GetRoleCountOk returns a tuple with the RoleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOut) GetRoleCountOk() (*int32, bool) {
	if o == nil || o.RoleCount == nil {
		return nil, false
	}
	return o.RoleCount, true
}

// HasRoleCount returns a boolean if a field has been set.
func (o *GroupOut) HasRoleCount() bool {
	if o != nil && o.RoleCount != nil {
		return true
	}

	return false
}

// SetRoleCount gets a reference to the given int32 and assigns it to the RoleCount field.
func (o *GroupOut) SetRoleCount(v int32) {
	o.RoleCount = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *GroupOut) GetSystem() bool {
	if o == nil || o.System == nil {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOut) GetSystemOk() (*bool, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *GroupOut) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *GroupOut) SetSystem(v bool) {
	o.System = &v
}

// GetPlatformDefault returns the PlatformDefault field value if set, zero value otherwise.
func (o *GroupOut) GetPlatformDefault() bool {
	if o == nil || o.PlatformDefault == nil {
		var ret bool
		return ret
	}
	return *o.PlatformDefault
}

// GetPlatformDefaultOk returns a tuple with the PlatformDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOut) GetPlatformDefaultOk() (*bool, bool) {
	if o == nil || o.PlatformDefault == nil {
		return nil, false
	}
	return o.PlatformDefault, true
}

// HasPlatformDefault returns a boolean if a field has been set.
func (o *GroupOut) HasPlatformDefault() bool {
	if o != nil && o.PlatformDefault != nil {
		return true
	}

	return false
}

// SetPlatformDefault gets a reference to the given bool and assigns it to the PlatformDefault field.
func (o *GroupOut) SetPlatformDefault(v bool) {
	o.PlatformDefault = &v
}

func (o GroupOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	if o.PrincipalCount != nil {
		toSerialize["principalCount"] = o.PrincipalCount
	}
	if o.RoleCount != nil {
		toSerialize["roleCount"] = o.RoleCount
	}
	if o.System != nil {
		toSerialize["system"] = o.System
	}
	if o.PlatformDefault != nil {
		toSerialize["platform_default"] = o.PlatformDefault
	}
	return json.Marshal(toSerialize)
}

type NullableGroupOut struct {
	value *GroupOut
	isSet bool
}

func (v NullableGroupOut) Get() *GroupOut {
	return v.value
}

func (v *NullableGroupOut) Set(val *GroupOut) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupOut) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupOut(val *GroupOut) *NullableGroupOut {
	return &NullableGroupOut{value: val, isSet: true}
}

func (v NullableGroupOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


