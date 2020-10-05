# PolicyExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**Created** | **string** |  | 
**Modified** | **string** |  | 
**Group** | [**GroupOut**](GroupOut.md) |  | 
**Roles** | [**[]RoleOut**](RoleOut.md) |  | 

## Methods

### NewPolicyExtended

`func NewPolicyExtended(name string, uuid string, created string, modified string, group GroupOut, roles []RoleOut, ) *PolicyExtended`

NewPolicyExtended instantiates a new PolicyExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyExtendedWithDefaults

`func NewPolicyExtendedWithDefaults() *PolicyExtended`

NewPolicyExtendedWithDefaults instantiates a new PolicyExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyExtended) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyExtended) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyExtended) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicyExtended) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyExtended) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyExtended) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyExtended) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *PolicyExtended) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PolicyExtended) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PolicyExtended) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCreated

`func (o *PolicyExtended) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PolicyExtended) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PolicyExtended) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetModified

`func (o *PolicyExtended) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PolicyExtended) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PolicyExtended) SetModified(v string)`

SetModified sets Modified field to given value.


### GetGroup

`func (o *PolicyExtended) GetGroup() GroupOut`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PolicyExtended) GetGroupOk() (*GroupOut, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PolicyExtended) SetGroup(v GroupOut)`

SetGroup sets Group field to given value.


### GetRoles

`func (o *PolicyExtended) GetRoles() []RoleOut`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PolicyExtended) GetRolesOk() (*[]RoleOut, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PolicyExtended) SetRoles(v []RoleOut)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


