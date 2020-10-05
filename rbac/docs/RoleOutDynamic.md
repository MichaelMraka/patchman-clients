# RoleOutDynamic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**Created** | **string** |  | 
**Modified** | **string** |  | 
**PolicyCount** | **int32** |  | 
**AccessCount** | **int32** |  | 
**Applications** | **[]string** |  | 
**System** | **bool** |  | [default to false]
**PlatformDefault** | **bool** |  | [default to false]
**GroupsInCount** | Pointer to **int32** |  | [optional] 
**GroupsIn** | Pointer to [**[]AdditionalGroup**](AdditionalGroup.md) |  | [optional] 

## Methods

### NewRoleOutDynamic

`func NewRoleOutDynamic(name string, uuid string, created string, modified string, policyCount int32, accessCount int32, applications []string, system bool, platformDefault bool, ) *RoleOutDynamic`

NewRoleOutDynamic instantiates a new RoleOutDynamic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleOutDynamicWithDefaults

`func NewRoleOutDynamicWithDefaults() *RoleOutDynamic`

NewRoleOutDynamicWithDefaults instantiates a new RoleOutDynamic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleOutDynamic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleOutDynamic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleOutDynamic) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *RoleOutDynamic) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RoleOutDynamic) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RoleOutDynamic) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RoleOutDynamic) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *RoleOutDynamic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleOutDynamic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleOutDynamic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleOutDynamic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *RoleOutDynamic) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoleOutDynamic) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoleOutDynamic) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCreated

`func (o *RoleOutDynamic) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RoleOutDynamic) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RoleOutDynamic) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetModified

`func (o *RoleOutDynamic) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RoleOutDynamic) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RoleOutDynamic) SetModified(v string)`

SetModified sets Modified field to given value.


### GetPolicyCount

`func (o *RoleOutDynamic) GetPolicyCount() int32`

GetPolicyCount returns the PolicyCount field if non-nil, zero value otherwise.

### GetPolicyCountOk

`func (o *RoleOutDynamic) GetPolicyCountOk() (*int32, bool)`

GetPolicyCountOk returns a tuple with the PolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCount

`func (o *RoleOutDynamic) SetPolicyCount(v int32)`

SetPolicyCount sets PolicyCount field to given value.


### GetAccessCount

`func (o *RoleOutDynamic) GetAccessCount() int32`

GetAccessCount returns the AccessCount field if non-nil, zero value otherwise.

### GetAccessCountOk

`func (o *RoleOutDynamic) GetAccessCountOk() (*int32, bool)`

GetAccessCountOk returns a tuple with the AccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCount

`func (o *RoleOutDynamic) SetAccessCount(v int32)`

SetAccessCount sets AccessCount field to given value.


### GetApplications

`func (o *RoleOutDynamic) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RoleOutDynamic) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *RoleOutDynamic) SetApplications(v []string)`

SetApplications sets Applications field to given value.


### GetSystem

`func (o *RoleOutDynamic) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *RoleOutDynamic) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *RoleOutDynamic) SetSystem(v bool)`

SetSystem sets System field to given value.


### GetPlatformDefault

`func (o *RoleOutDynamic) GetPlatformDefault() bool`

GetPlatformDefault returns the PlatformDefault field if non-nil, zero value otherwise.

### GetPlatformDefaultOk

`func (o *RoleOutDynamic) GetPlatformDefaultOk() (*bool, bool)`

GetPlatformDefaultOk returns a tuple with the PlatformDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformDefault

`func (o *RoleOutDynamic) SetPlatformDefault(v bool)`

SetPlatformDefault sets PlatformDefault field to given value.


### GetGroupsInCount

`func (o *RoleOutDynamic) GetGroupsInCount() int32`

GetGroupsInCount returns the GroupsInCount field if non-nil, zero value otherwise.

### GetGroupsInCountOk

`func (o *RoleOutDynamic) GetGroupsInCountOk() (*int32, bool)`

GetGroupsInCountOk returns a tuple with the GroupsInCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsInCount

`func (o *RoleOutDynamic) SetGroupsInCount(v int32)`

SetGroupsInCount sets GroupsInCount field to given value.

### HasGroupsInCount

`func (o *RoleOutDynamic) HasGroupsInCount() bool`

HasGroupsInCount returns a boolean if a field has been set.

### GetGroupsIn

`func (o *RoleOutDynamic) GetGroupsIn() []AdditionalGroup`

GetGroupsIn returns the GroupsIn field if non-nil, zero value otherwise.

### GetGroupsInOk

`func (o *RoleOutDynamic) GetGroupsInOk() (*[]AdditionalGroup, bool)`

GetGroupsInOk returns a tuple with the GroupsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsIn

`func (o *RoleOutDynamic) SetGroupsIn(v []AdditionalGroup)`

SetGroupsIn sets GroupsIn field to given value.

### HasGroupsIn

`func (o *RoleOutDynamic) HasGroupsIn() bool`

HasGroupsIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


