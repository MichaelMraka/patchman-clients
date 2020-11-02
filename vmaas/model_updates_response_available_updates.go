/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.20.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// UpdatesResponseAvailableUpdates struct for UpdatesResponseAvailableUpdates
type UpdatesResponseAvailableUpdates struct {
	Repository *string `json:"repository,omitempty"`
	Releasever *string `json:"releasever,omitempty"`
	Basearch *string `json:"basearch,omitempty"`
	Erratum *string `json:"erratum,omitempty"`
	Package *string `json:"package,omitempty"`
}

// NewUpdatesResponseAvailableUpdates instantiates a new UpdatesResponseAvailableUpdates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatesResponseAvailableUpdates() *UpdatesResponseAvailableUpdates {
	this := UpdatesResponseAvailableUpdates{}
	return &this
}

// NewUpdatesResponseAvailableUpdatesWithDefaults instantiates a new UpdatesResponseAvailableUpdates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatesResponseAvailableUpdatesWithDefaults() *UpdatesResponseAvailableUpdates {
	this := UpdatesResponseAvailableUpdates{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *UpdatesResponseAvailableUpdates) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponseAvailableUpdates) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *UpdatesResponseAvailableUpdates) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *UpdatesResponseAvailableUpdates) SetRepository(v string) {
	o.Repository = &v
}

// GetReleasever returns the Releasever field value if set, zero value otherwise.
func (o *UpdatesResponseAvailableUpdates) GetReleasever() string {
	if o == nil || o.Releasever == nil {
		var ret string
		return ret
	}
	return *o.Releasever
}

// GetReleaseverOk returns a tuple with the Releasever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponseAvailableUpdates) GetReleaseverOk() (*string, bool) {
	if o == nil || o.Releasever == nil {
		return nil, false
	}
	return o.Releasever, true
}

// HasReleasever returns a boolean if a field has been set.
func (o *UpdatesResponseAvailableUpdates) HasReleasever() bool {
	if o != nil && o.Releasever != nil {
		return true
	}

	return false
}

// SetReleasever gets a reference to the given string and assigns it to the Releasever field.
func (o *UpdatesResponseAvailableUpdates) SetReleasever(v string) {
	o.Releasever = &v
}

// GetBasearch returns the Basearch field value if set, zero value otherwise.
func (o *UpdatesResponseAvailableUpdates) GetBasearch() string {
	if o == nil || o.Basearch == nil {
		var ret string
		return ret
	}
	return *o.Basearch
}

// GetBasearchOk returns a tuple with the Basearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponseAvailableUpdates) GetBasearchOk() (*string, bool) {
	if o == nil || o.Basearch == nil {
		return nil, false
	}
	return o.Basearch, true
}

// HasBasearch returns a boolean if a field has been set.
func (o *UpdatesResponseAvailableUpdates) HasBasearch() bool {
	if o != nil && o.Basearch != nil {
		return true
	}

	return false
}

// SetBasearch gets a reference to the given string and assigns it to the Basearch field.
func (o *UpdatesResponseAvailableUpdates) SetBasearch(v string) {
	o.Basearch = &v
}

// GetErratum returns the Erratum field value if set, zero value otherwise.
func (o *UpdatesResponseAvailableUpdates) GetErratum() string {
	if o == nil || o.Erratum == nil {
		var ret string
		return ret
	}
	return *o.Erratum
}

// GetErratumOk returns a tuple with the Erratum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponseAvailableUpdates) GetErratumOk() (*string, bool) {
	if o == nil || o.Erratum == nil {
		return nil, false
	}
	return o.Erratum, true
}

// HasErratum returns a boolean if a field has been set.
func (o *UpdatesResponseAvailableUpdates) HasErratum() bool {
	if o != nil && o.Erratum != nil {
		return true
	}

	return false
}

// SetErratum gets a reference to the given string and assigns it to the Erratum field.
func (o *UpdatesResponseAvailableUpdates) SetErratum(v string) {
	o.Erratum = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *UpdatesResponseAvailableUpdates) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponseAvailableUpdates) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *UpdatesResponseAvailableUpdates) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *UpdatesResponseAvailableUpdates) SetPackage(v string) {
	o.Package = &v
}

func (o UpdatesResponseAvailableUpdates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Releasever != nil {
		toSerialize["releasever"] = o.Releasever
	}
	if o.Basearch != nil {
		toSerialize["basearch"] = o.Basearch
	}
	if o.Erratum != nil {
		toSerialize["erratum"] = o.Erratum
	}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatesResponseAvailableUpdates struct {
	value *UpdatesResponseAvailableUpdates
	isSet bool
}

func (v NullableUpdatesResponseAvailableUpdates) Get() *UpdatesResponseAvailableUpdates {
	return v.value
}

func (v *NullableUpdatesResponseAvailableUpdates) Set(val *UpdatesResponseAvailableUpdates) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatesResponseAvailableUpdates) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatesResponseAvailableUpdates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatesResponseAvailableUpdates(val *UpdatesResponseAvailableUpdates) *NullableUpdatesResponseAvailableUpdates {
	return &NullableUpdatesResponseAvailableUpdates{value: val, isSet: true}
}

func (v NullableUpdatesResponseAvailableUpdates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatesResponseAvailableUpdates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


