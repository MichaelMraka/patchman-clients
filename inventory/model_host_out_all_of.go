/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inventory

import (
	"encoding/json"
)

// HostOutAllOf struct for HostOutAllOf
type HostOutAllOf struct {
	// A set of facts belonging to the host.
	Facts *[]FactSet `json:"facts,omitempty"`
}

// NewHostOutAllOf instantiates a new HostOutAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostOutAllOf() *HostOutAllOf {
	this := HostOutAllOf{}
	return &this
}

// NewHostOutAllOfWithDefaults instantiates a new HostOutAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostOutAllOfWithDefaults() *HostOutAllOf {
	this := HostOutAllOf{}
	return &this
}

// GetFacts returns the Facts field value if set, zero value otherwise.
func (o *HostOutAllOf) GetFacts() []FactSet {
	if o == nil || o.Facts == nil {
		var ret []FactSet
		return ret
	}
	return *o.Facts
}

// GetFactsOk returns a tuple with the Facts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOutAllOf) GetFactsOk() (*[]FactSet, bool) {
	if o == nil || o.Facts == nil {
		return nil, false
	}
	return o.Facts, true
}

// HasFacts returns a boolean if a field has been set.
func (o *HostOutAllOf) HasFacts() bool {
	if o != nil && o.Facts != nil {
		return true
	}

	return false
}

// SetFacts gets a reference to the given []FactSet and assigns it to the Facts field.
func (o *HostOutAllOf) SetFacts(v []FactSet) {
	o.Facts = &v
}

func (o HostOutAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Facts != nil {
		toSerialize["facts"] = o.Facts
	}
	return json.Marshal(toSerialize)
}

type NullableHostOutAllOf struct {
	value *HostOutAllOf
	isSet bool
}

func (v NullableHostOutAllOf) Get() *HostOutAllOf {
	return v.value
}

func (v *NullableHostOutAllOf) Set(val *HostOutAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHostOutAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHostOutAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostOutAllOf(val *HostOutAllOf) *NullableHostOutAllOf {
	return &NullableHostOutAllOf{value: val, isSet: true}
}

func (v NullableHostOutAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostOutAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


