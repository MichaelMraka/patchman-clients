# RoleOut

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

## Methods

### NewRoleOut

`func NewRoleOut(name string, uuid string, created string, modified string, ) *RoleOut`

NewRoleOut instantiates a new RoleOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleOutWithDefaults

`func NewRoleOutWithDefaults() *RoleOut`

NewRoleOutWithDefaults instantiates a new RoleOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleOut) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *RoleOut) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RoleOut) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RoleOut) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RoleOut) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *RoleOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *RoleOut) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoleOut) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoleOut) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCreated

`func (o *RoleOut) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RoleOut) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RoleOut) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetModified

`func (o *RoleOut) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RoleOut) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RoleOut) SetModified(v string)`

SetModified sets Modified field to given value.


### GetPolicyCount

`func (o *RoleOut) GetPolicyCount() int32`

GetPolicyCount returns the PolicyCount field if non-nil, zero value otherwise.

### GetPolicyCountOk

`func (o *RoleOut) GetPolicyCountOk() (*int32, bool)`

GetPolicyCountOk returns a tuple with the PolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCount

`func (o *RoleOut) SetPolicyCount(v int32)`

SetPolicyCount sets PolicyCount field to given value.

### HasPolicyCount

`func (o *RoleOut) HasPolicyCount() bool`

HasPolicyCount returns a boolean if a field has been set.

### GetAccessCount

`func (o *RoleOut) GetAccessCount() int32`

GetAccessCount returns the AccessCount field if non-nil, zero value otherwise.

### GetAccessCountOk

`func (o *RoleOut) GetAccessCountOk() (*int32, bool)`

GetAccessCountOk returns a tuple with the AccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCount

`func (o *RoleOut) SetAccessCount(v int32)`

SetAccessCount sets AccessCount field to given value.

### HasAccessCount

`func (o *RoleOut) HasAccessCount() bool`

HasAccessCount returns a boolean if a field has been set.

### GetApplications

`func (o *RoleOut) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RoleOut) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *RoleOut) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *RoleOut) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetSystem

`func (o *RoleOut) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *RoleOut) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *RoleOut) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *RoleOut) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetPlatformDefault

`func (o *RoleOut) GetPlatformDefault() bool`

GetPlatformDefault returns the PlatformDefault field if non-nil, zero value otherwise.

### GetPlatformDefaultOk

`func (o *RoleOut) GetPlatformDefaultOk() (*bool, bool)`

GetPlatformDefaultOk returns a tuple with the PlatformDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformDefault

`func (o *RoleOut) SetPlatformDefault(v bool)`

SetPlatformDefault sets PlatformDefault field to given value.

### HasPlatformDefault

`func (o *RoleOut) HasPlatformDefault() bool`

HasPlatformDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


