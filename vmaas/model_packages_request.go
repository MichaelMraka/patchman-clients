/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// PackagesRequest struct for PackagesRequest
type PackagesRequest struct {
	PackageList []string `json:"package_list"`
	// Include content from \"third party\" repositories into the response, disabled by default.
	ThirdParty *bool `json:"third_party,omitempty"`
}

// NewPackagesRequest instantiates a new PackagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesRequest(packageList []string, ) *PackagesRequest {
	this := PackagesRequest{}
	this.PackageList = packageList
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	return &this
}

// NewPackagesRequestWithDefaults instantiates a new PackagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesRequestWithDefaults() *PackagesRequest {
	this := PackagesRequest{}
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	return &this
}

// GetPackageList returns the PackageList field value
func (o *PackagesRequest) GetPackageList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.PackageList
}

// GetPackageListOk returns a tuple with the PackageList field value
// and a boolean to check if the value has been set.
func (o *PackagesRequest) GetPackageListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PackageList, true
}

// SetPackageList sets field value
func (o *PackagesRequest) SetPackageList(v []string) {
	o.PackageList = v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *PackagesRequest) GetThirdParty() bool {
	if o == nil || o.ThirdParty == nil {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesRequest) GetThirdPartyOk() (*bool, bool) {
	if o == nil || o.ThirdParty == nil {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *PackagesRequest) HasThirdParty() bool {
	if o != nil && o.ThirdParty != nil {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *PackagesRequest) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

func (o PackagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["package_list"] = o.PackageList
	}
	if o.ThirdParty != nil {
		toSerialize["third_party"] = o.ThirdParty
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesRequest struct {
	value *PackagesRequest
	isSet bool
}

func (v NullablePackagesRequest) Get() *PackagesRequest {
	return v.value
}

func (v *NullablePackagesRequest) Set(val *PackagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesRequest(val *PackagesRequest) *NullablePackagesRequest {
	return &NullablePackagesRequest{value: val, isSet: true}
}

func (v NullablePackagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


