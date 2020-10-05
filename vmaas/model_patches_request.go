/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.20.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// PatchesRequest struct for PatchesRequest
type PatchesRequest struct {
	PackageList []string `json:"package_list"`
	RepositoryList *[]string `json:"repository_list,omitempty"`
	ModulesList *[]UpdatesRequestModulesList `json:"modules_list,omitempty"`
	Releasever *string `json:"releasever,omitempty"`
	Basearch *string `json:"basearch,omitempty"`
}

// NewPatchesRequest instantiates a new PatchesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchesRequest(packageList []string, ) *PatchesRequest {
	this := PatchesRequest{}
	this.PackageList = packageList
	return &this
}

// NewPatchesRequestWithDefaults instantiates a new PatchesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchesRequestWithDefaults() *PatchesRequest {
	this := PatchesRequest{}
	return &this
}

// GetPackageList returns the PackageList field value
func (o *PatchesRequest) GetPackageList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.PackageList
}

// GetPackageListOk returns a tuple with the PackageList field value
// and a boolean to check if the value has been set.
func (o *PatchesRequest) GetPackageListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PackageList, true
}

// SetPackageList sets field value
func (o *PatchesRequest) SetPackageList(v []string) {
	o.PackageList = v
}

// GetRepositoryList returns the RepositoryList field value if set, zero value otherwise.
func (o *PatchesRequest) GetRepositoryList() []string {
	if o == nil || o.RepositoryList == nil {
		var ret []string
		return ret
	}
	return *o.RepositoryList
}

// GetRepositoryListOk returns a tuple with the RepositoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchesRequest) GetRepositoryListOk() (*[]string, bool) {
	if o == nil || o.RepositoryList == nil {
		return nil, false
	}
	return o.RepositoryList, true
}

// HasRepositoryList returns a boolean if a field has been set.
func (o *PatchesRequest) HasRepositoryList() bool {
	if o != nil && o.RepositoryList != nil {
		return true
	}

	return false
}

// SetRepositoryList gets a reference to the given []string and assigns it to the RepositoryList field.
func (o *PatchesRequest) SetRepositoryList(v []string) {
	o.RepositoryList = &v
}

// GetModulesList returns the ModulesList field value if set, zero value otherwise.
func (o *PatchesRequest) GetModulesList() []UpdatesRequestModulesList {
	if o == nil || o.ModulesList == nil {
		var ret []UpdatesRequestModulesList
		return ret
	}
	return *o.ModulesList
}

// GetModulesListOk returns a tuple with the ModulesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchesRequest) GetModulesListOk() (*[]UpdatesRequestModulesList, bool) {
	if o == nil || o.ModulesList == nil {
		return nil, false
	}
	return o.ModulesList, true
}

// HasModulesList returns a boolean if a field has been set.
func (o *PatchesRequest) HasModulesList() bool {
	if o != nil && o.ModulesList != nil {
		return true
	}

	return false
}

// SetModulesList gets a reference to the given []UpdatesRequestModulesList and assigns it to the ModulesList field.
func (o *PatchesRequest) SetModulesList(v []UpdatesRequestModulesList) {
	o.ModulesList = &v
}

// GetReleasever returns the Releasever field value if set, zero value otherwise.
func (o *PatchesRequest) GetReleasever() string {
	if o == nil || o.Releasever == nil {
		var ret string
		return ret
	}
	return *o.Releasever
}

// GetReleaseverOk returns a tuple with the Releasever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchesRequest) GetReleaseverOk() (*string, bool) {
	if o == nil || o.Releasever == nil {
		return nil, false
	}
	return o.Releasever, true
}

// HasReleasever returns a boolean if a field has been set.
func (o *PatchesRequest) HasReleasever() bool {
	if o != nil && o.Releasever != nil {
		return true
	}

	return false
}

// SetReleasever gets a reference to the given string and assigns it to the Releasever field.
func (o *PatchesRequest) SetReleasever(v string) {
	o.Releasever = &v
}

// GetBasearch returns the Basearch field value if set, zero value otherwise.
func (o *PatchesRequest) GetBasearch() string {
	if o == nil || o.Basearch == nil {
		var ret string
		return ret
	}
	return *o.Basearch
}

// GetBasearchOk returns a tuple with the Basearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchesRequest) GetBasearchOk() (*string, bool) {
	if o == nil || o.Basearch == nil {
		return nil, false
	}
	return o.Basearch, true
}

// HasBasearch returns a boolean if a field has been set.
func (o *PatchesRequest) HasBasearch() bool {
	if o != nil && o.Basearch != nil {
		return true
	}

	return false
}

// SetBasearch gets a reference to the given string and assigns it to the Basearch field.
func (o *PatchesRequest) SetBasearch(v string) {
	o.Basearch = &v
}

func (o PatchesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["package_list"] = o.PackageList
	}
	if o.RepositoryList != nil {
		toSerialize["repository_list"] = o.RepositoryList
	}
	if o.ModulesList != nil {
		toSerialize["modules_list"] = o.ModulesList
	}
	if o.Releasever != nil {
		toSerialize["releasever"] = o.Releasever
	}
	if o.Basearch != nil {
		toSerialize["basearch"] = o.Basearch
	}
	return json.Marshal(toSerialize)
}

type NullablePatchesRequest struct {
	value *PatchesRequest
	isSet bool
}

func (v NullablePatchesRequest) Get() *PatchesRequest {
	return v.value
}

func (v *NullablePatchesRequest) Set(val *PatchesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchesRequest(val *PatchesRequest) *NullablePatchesRequest {
	return &NullablePatchesRequest{value: val, isSet: true}
}

func (v NullablePatchesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


