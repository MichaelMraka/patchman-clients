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

// CrossAccountRequestOut struct for CrossAccountRequestOut
type CrossAccountRequestOut struct {
	RequestId *string `json:"request_id,omitempty"`
	TargetAccount *string `json:"target_account,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	Status *string `json:"status,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Roles *[]CrossAccountRequestWithRolesRoles `json:"roles,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// NewCrossAccountRequestOut instantiates a new CrossAccountRequestOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossAccountRequestOut() *CrossAccountRequestOut {
	this := CrossAccountRequestOut{}
	return &this
}

// NewCrossAccountRequestOutWithDefaults instantiates a new CrossAccountRequestOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossAccountRequestOutWithDefaults() *CrossAccountRequestOut {
	this := CrossAccountRequestOut{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CrossAccountRequestOut) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestOut) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CrossAccountRequestOut) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CrossAccountRequestOut) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTargetAccount returns the TargetAccount field value if set, zero value otherwise.
func (o *CrossAccountRequestOut) GetTargetAccount() string {
	if o == nil || o.TargetAccount == nil {
		var ret string
		return ret
	}
	return *o.TargetAccount
}

// GetTargetAccountOk returns a tuple with the TargetAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestOut) GetTargetAccountOk() (*string, bool) {
	if o == nil || o.TargetAccount == nil {
		return nil, false
	}
	return o.TargetAccount, true
}

// HasTargetAccount returns a boolean if a field has been set.
func (o *CrossAccountRequestOut) HasTargetAccount() bool {
	if o != nil && o.TargetAccount != nil {
		return true
	}

	return false
}

// SetTargetAccount gets a reference to the given string and assigns it to the TargetAccount field.
func (o *CrossAccountRequestOut) SetTargetAccount(v string) {
	o.TargetAccount = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CrossAccountRequestOut) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestOut) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CrossAccountRequestOut) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CrossAccountRequestOut) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CrossAccountRequestOut) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestOut) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CrossAccountRequestOut) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CrossAccountRequestOut) SetEndDate(v string) {
	o.EndDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CrossAccountRequestOut) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestOut) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CrossAccountRequestOut) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CrossAccountRequestOut) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CrossAccountRequestOut) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestOut) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CrossAccountRequestOut) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CrossAccountRequestOut) SetCreated(v time.Time) {
	o.Created = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CrossAccountRequestOut) GetRoles() []CrossAccountRequestWithRolesRoles {
	if o == nil || o.Roles == nil {
		var ret []CrossAccountRequestWithRolesRoles
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestOut) GetRolesOk() (*[]CrossAccountRequestWithRolesRoles, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CrossAccountRequestOut) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []CrossAccountRequestWithRolesRoles and assigns it to the Roles field.
func (o *CrossAccountRequestOut) SetRoles(v []CrossAccountRequestWithRolesRoles) {
	o.Roles = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CrossAccountRequestOut) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestOut) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CrossAccountRequestOut) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CrossAccountRequestOut) SetUserId(v string) {
	o.UserId = &v
}

func (o CrossAccountRequestOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.TargetAccount != nil {
		toSerialize["target_account"] = o.TargetAccount
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableCrossAccountRequestOut struct {
	value *CrossAccountRequestOut
	isSet bool
}

func (v NullableCrossAccountRequestOut) Get() *CrossAccountRequestOut {
	return v.value
}

func (v *NullableCrossAccountRequestOut) Set(val *CrossAccountRequestOut) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossAccountRequestOut) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossAccountRequestOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossAccountRequestOut(val *CrossAccountRequestOut) *NullableCrossAccountRequestOut {
	return &NullableCrossAccountRequestOut{value: val, isSet: true}
}

func (v NullableCrossAccountRequestOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossAccountRequestOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

