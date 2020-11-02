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

// PkgtreeResponse struct for PkgtreeResponse
type PkgtreeResponse struct {
	LastChange interface{} `json:"last_change,omitempty"`
	PackageNameList *map[string][]map[string]interface{} `json:"package_name_list,omitempty"`
}

// NewPkgtreeResponse instantiates a new PkgtreeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgtreeResponse() *PkgtreeResponse {
	this := PkgtreeResponse{}
	return &this
}

// NewPkgtreeResponseWithDefaults instantiates a new PkgtreeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgtreeResponseWithDefaults() *PkgtreeResponse {
	this := PkgtreeResponse{}
	return &this
}

// GetLastChange returns the LastChange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkgtreeResponse) GetLastChange() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkgtreeResponse) GetLastChangeOk() (*interface{}, bool) {
	if o == nil || o.LastChange == nil {
		return nil, false
	}
	return &o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *PkgtreeResponse) HasLastChange() bool {
	if o != nil && o.LastChange != nil {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given interface{} and assigns it to the LastChange field.
func (o *PkgtreeResponse) SetLastChange(v interface{}) {
	o.LastChange = v
}

// GetPackageNameList returns the PackageNameList field value if set, zero value otherwise.
func (o *PkgtreeResponse) GetPackageNameList() map[string][]map[string]interface{} {
	if o == nil || o.PackageNameList == nil {
		var ret map[string][]map[string]interface{}
		return ret
	}
	return *o.PackageNameList
}

// GetPackageNameListOk returns a tuple with the PackageNameList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgtreeResponse) GetPackageNameListOk() (*map[string][]map[string]interface{}, bool) {
	if o == nil || o.PackageNameList == nil {
		return nil, false
	}
	return o.PackageNameList, true
}

// HasPackageNameList returns a boolean if a field has been set.
func (o *PkgtreeResponse) HasPackageNameList() bool {
	if o != nil && o.PackageNameList != nil {
		return true
	}

	return false
}

// SetPackageNameList gets a reference to the given map[string][]map[string]interface{} and assigns it to the PackageNameList field.
func (o *PkgtreeResponse) SetPackageNameList(v map[string][]map[string]interface{}) {
	o.PackageNameList = &v
}

func (o PkgtreeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastChange != nil {
		toSerialize["last_change"] = o.LastChange
	}
	if o.PackageNameList != nil {
		toSerialize["package_name_list"] = o.PackageNameList
	}
	return json.Marshal(toSerialize)
}

type NullablePkgtreeResponse struct {
	value *PkgtreeResponse
	isSet bool
}

func (v NullablePkgtreeResponse) Get() *PkgtreeResponse {
	return v.value
}

func (v *NullablePkgtreeResponse) Set(val *PkgtreeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgtreeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgtreeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgtreeResponse(val *PkgtreeResponse) *NullablePkgtreeResponse {
	return &NullablePkgtreeResponse{value: val, isSet: true}
}

func (v NullablePkgtreeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgtreeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


