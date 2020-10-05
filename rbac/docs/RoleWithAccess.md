# RoleWithAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**Created** | **string** |  | 
**Modified** | **string** |  | 
**PolicyCount** | Pointer to **int32** |  | [optional] 
**AccessCount** | Pointer to **int32** |  | [optional] 
**Applications** | Pointer to **[]string** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] [default to false]
**PlatformDefault** | Pointer to **bool** |  | [optional] [default to false]
**Access** | [**[]Access**](Access.md) |  | 

## Methods

### NewRoleWithAccess

`func NewRoleWithAccess(name string, uuid string, created string, modified string, access []Access, ) *RoleWithAccess`

NewRoleWithAccess instantiates a new RoleWithAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithAccessWithDefaults

`func NewRoleWithAccessWithDefaults() *RoleWithAccess`

NewRoleWithAccessWithDefaults instantiates a new RoleWithAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleWithAccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleWithAccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleWithAccess) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *RoleWithAccess) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RoleWithAccess) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RoleWithAccess) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RoleWithAccess) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *RoleWithAccess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleWithAccess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleWithAccess) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleWithAccess) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *RoleWithAccess) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoleWithAccess) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoleWithAccess) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCreated

`func (o *RoleWithAccess) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RoleWithAccess) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RoleWithAccess) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetModified

`func (o *RoleWithAccess) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RoleWithAccess) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RoleWithAccess) SetModified(v string)`

SetModified sets Modified field to given value.


### GetPolicyCount

`func (o *RoleWithAccess) GetPolicyCount() int32`

GetPolicyCount returns the PolicyCount field if non-nil, zero value otherwise.

### GetPolicyCountOk

`func (o *RoleWithAccess) GetPolicyCountOk() (*int32, bool)`

GetPolicyCountOk returns a tuple with the PolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCount

`func (o *RoleWithAccess) SetPolicyCount(v int32)`

SetPolicyCount sets PolicyCount field to given value.

### HasPolicyCount

`func (o *RoleWithAccess) HasPolicyCount() bool`

HasPolicyCount returns a boolean if a field has been set.

### GetAccessCount

`func (o *RoleWithAccess) GetAccessCount() int32`

GetAccessCount returns the AccessCount field if non-nil, zero value otherwise.

### GetAccessCountOk

`func (o *RoleWithAccess) GetAccessCountOk() (*int32, bool)`

GetAccessCountOk returns a tuple with the AccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCount

`func (o *RoleWithAccess) SetAccessCount(v int32)`

SetAccessCount sets AccessCount field to given value.

### HasAccessCount

`func (o *RoleWithAccess) HasAccessCount() bool`

HasAccessCount returns a boolean if a field has been set.

### GetApplications

`func (o *RoleWithAccess) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RoleWithAccess) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *RoleWithAccess) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *RoleWithAccess) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetSystem

`func (o *RoleWithAccess) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *RoleWithAccess) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *RoleWithAccess) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *RoleWithAccess) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetPlatformDefault

`func (o *RoleWithAccess) GetPlatformDefault() bool`

GetPlatformDefault returns the PlatformDefault field if non-nil, zero value otherwise.

### GetPlatformDefaultOk

`func (o *RoleWithAccess) GetPlatformDefaultOk() (*bool, bool)`

GetPlatformDefaultOk returns a tuple with the PlatformDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformDefault

`func (o *RoleWithAccess) SetPlatformDefault(v bool)`

SetPlatformDefault sets PlatformDefault field to given value.

### HasPlatformDefault

`func (o *RoleWithAccess) HasPlatformDefault() bool`

HasPlatformDefault returns a boolean if a field has been set.

### GetAccess

`func (o *RoleWithAccess) GetAccess() []Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RoleWithAccess) GetAccessOk() (*[]Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RoleWithAccess) SetAccess(v []Access)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


