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

// UpdatesResponse struct for UpdatesResponse
type UpdatesResponse struct {
	UpdateList *map[string]UpdatesResponseUpdateList `json:"update_list,omitempty"`
	RepositoryList *[]string `json:"repository_list,omitempty"`
	ModulesList *[]UpdatesRequestModulesList `json:"modules_list,omitempty"`
	Releasever *string `json:"releasever,omitempty"`
	Basearch *string `json:"basearch,omitempty"`
}

// NewUpdatesResponse instantiates a new UpdatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatesResponse() *UpdatesResponse {
	this := UpdatesResponse{}
	return &this
}

// NewUpdatesResponseWithDefaults instantiates a new UpdatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatesResponseWithDefaults() *UpdatesResponse {
	this := UpdatesResponse{}
	return &this
}

// GetUpdateList returns the UpdateList field value if set, zero value otherwise.
func (o *UpdatesResponse) GetUpdateList() map[string]UpdatesResponseUpdateList {
	if o == nil || o.UpdateList == nil {
		var ret map[string]UpdatesResponseUpdateList
		return ret
	}
	return *o.UpdateList
}

// GetUpdateListOk returns a tuple with the UpdateList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponse) GetUpdateListOk() (*map[string]UpdatesResponseUpdateList, bool) {
	if o == nil || o.UpdateList == nil {
		return nil, false
	}
	return o.UpdateList, true
}

// HasUpdateList returns a boolean if a field has been set.
func (o *UpdatesResponse) HasUpdateList() bool {
	if o != nil && o.UpdateList != nil {
		return true
	}

	return false
}

// SetUpdateList gets a reference to the given map[string]UpdatesResponseUpdateList and assigns it to the UpdateList field.
func (o *UpdatesResponse) SetUpdateList(v map[string]UpdatesResponseUpdateList) {
	o.UpdateList = &v
}

// GetRepositoryList returns the RepositoryList field value if set, zero value otherwise.
func (o *UpdatesResponse) GetRepositoryList() []string {
	if o == nil || o.RepositoryList == nil {
		var ret []string
		return ret
	}
	return *o.RepositoryList
}

// GetRepositoryListOk returns a tuple with the RepositoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponse) GetRepositoryListOk() (*[]string, bool) {
	if o == nil || o.RepositoryList == nil {
		return nil, false
	}
	return o.RepositoryList, true
}

// HasRepositoryList returns a boolean if a field has been set.
func (o *UpdatesResponse) HasRepositoryList() bool {
	if o != nil && o.RepositoryList != nil {
		return true
	}

	return false
}

// SetRepositoryList gets a reference to the given []string and assigns it to the RepositoryList field.
func (o *UpdatesResponse) SetRepositoryList(v []string) {
	o.RepositoryList = &v
}

// GetModulesList returns the ModulesList field value if set, zero value otherwise.
func (o *UpdatesResponse) GetModulesList() []UpdatesRequestModulesList {
	if o == nil || o.ModulesList == nil {
		var ret []UpdatesRequestModulesList
		return ret
	}
	return *o.ModulesList
}

// GetModulesListOk returns a tuple with the ModulesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponse) GetModulesListOk() (*[]UpdatesRequestModulesList, bool) {
	if o == nil || o.ModulesList == nil {
		return nil, false
	}
	return o.ModulesList, true
}

// HasModulesList returns a boolean if a field has been set.
func (o *UpdatesResponse) HasModulesList() bool {
	if o != nil && o.ModulesList != nil {
		return true
	}

	return false
}

// SetModulesList gets a reference to the given []UpdatesRequestModulesList and assigns it to the ModulesList field.
func (o *UpdatesResponse) SetModulesList(v []UpdatesRequestModulesList) {
	o.ModulesList = &v
}

// GetReleasever returns the Releasever field value if set, zero value otherwise.
func (o *UpdatesResponse) GetReleasever() string {
	if o == nil || o.Releasever == nil {
		var ret string
		return ret
	}
	return *o.Releasever
}

// GetReleaseverOk returns a tuple with the Releasever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponse) GetReleaseverOk() (*string, bool) {
	if o == nil || o.Releasever == nil {
		return nil, false
	}
	return o.Releasever, true
}

// HasReleasever returns a boolean if a field has been set.
func (o *UpdatesResponse) HasReleasever() bool {
	if o != nil && o.Releasever != nil {
		return true
	}

	return false
}

// SetReleasever gets a reference to the given string and assigns it to the Releasever field.
func (o *UpdatesResponse) SetReleasever(v string) {
	o.Releasever = &v
}

// GetBasearch returns the Basearch field value if set, zero value otherwise.
func (o *UpdatesResponse) GetBasearch() string {
	if o == nil || o.Basearch == nil {
		var ret string
		return ret
	}
	return *o.Basearch
}

// GetBasearchOk returns a tuple with the Basearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesResponse) GetBasearchOk() (*string, bool) {
	if o == nil || o.Basearch == nil {
		return nil, false
	}
	return o.Basearch, true
}

// HasBasearch returns a boolean if a field has been set.
func (o *UpdatesResponse) HasBasearch() bool {
	if o != nil && o.Basearch != nil {
		return true
	}

	return false
}

// SetBasearch gets a reference to the given string and assigns it to the Basearch field.
func (o *UpdatesResponse) SetBasearch(v string) {
	o.Basearch = &v
}

func (o UpdatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UpdateList != nil {
		toSerialize["update_list"] = o.UpdateList
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

type NullableUpdatesResponse struct {
	value *UpdatesResponse
	isSet bool
}

func (v NullableUpdatesResponse) Get() *UpdatesResponse {
	return v.value
}

func (v *NullableUpdatesResponse) Set(val *UpdatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatesResponse(val *UpdatesResponse) *NullableUpdatesResponse {
	return &NullableUpdatesResponse{value: val, isSet: true}
}

func (v NullableUpdatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


