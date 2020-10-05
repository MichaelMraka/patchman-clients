# RoleOutDynamicAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyCount** | **int32** |  | 
**AccessCount** | **int32** |  | 
**Applications** | **[]string** |  | 
**System** | **bool** |  | [default to false]
**PlatformDefault** | **bool** |  | [default to false]
**GroupsInCount** | Pointer to **int32** |  | [optional] 
**GroupsIn** | Pointer to [**[]AdditionalGroup**](AdditionalGroup.md) |  | [optional] 

## Methods

### NewRoleOutDynamicAllOf

`func NewRoleOutDynamicAllOf(policyCount int32, accessCount int32, applications []string, system bool, platformDefault bool, ) *RoleOutDynamicAllOf`

NewRoleOutDynamicAllOf instantiates a new RoleOutDynamicAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleOutDynamicAllOfWithDefaults

`func NewRoleOutDynamicAllOfWithDefaults() *RoleOutDynamicAllOf`

NewRoleOutDynamicAllOfWithDefaults instantiates a new RoleOutDynamicAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyCount

`func (o *RoleOutDynamicAllOf) GetPolicyCount() int32`

GetPolicyCount returns the PolicyCount field if non-nil, zero value otherwise.

### GetPolicyCountOk

`func (o *RoleOutDynamicAllOf) GetPolicyCountOk() (*int32, bool)`

GetPolicyCountOk returns a tuple with the PolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCount

`func (o *RoleOutDynamicAllOf) SetPolicyCount(v int32)`

SetPolicyCount sets PolicyCount field to given value.


### GetAccessCount

`func (o *RoleOutDynamicAllOf) GetAccessCount() int32`

GetAccessCount returns the AccessCount field if non-nil, zero value otherwise.

### GetAccessCountOk

`func (o *RoleOutDynamicAllOf) GetAccessCountOk() (*int32, bool)`

GetAccessCountOk returns a tuple with the AccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCount

`func (o *RoleOutDynamicAllOf) SetAccessCount(v int32)`

SetAccessCount sets AccessCount field to given value.


### GetApplications

`func (o *RoleOutDynamicAllOf) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RoleOutDynamicAllOf) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *RoleOutDynamicAllOf) SetApplications(v []string)`

SetApplications sets Applications field to given value.


### GetSystem

`func (o *RoleOutDynamicAllOf) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *RoleOutDynamicAllOf) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *RoleOutDynamicAllOf) SetSystem(v bool)`

SetSystem sets System field to given value.


### GetPlatformDefault

`func (o *RoleOutDynamicAllOf) GetPlatformDefault() bool`

GetPlatformDefault returns the PlatformDefault field if non-nil, zero value otherwise.

### GetPlatformDefaultOk

`func (o *RoleOutDynamicAllOf) GetPlatformDefaultOk() (*bool, bool)`

GetPlatformDefaultOk returns a tuple with the PlatformDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformDefault

`func (o *RoleOutDynamicAllOf) SetPlatformDefault(v bool)`

SetPlatformDefault sets PlatformDefault field to given value.


### GetGroupsInCount

`func (o *RoleOutDynamicAllOf) GetGroupsInCount() int32`

GetGroupsInCount returns the GroupsInCount field if non-nil, zero value otherwise.

### GetGroupsInCountOk

`func (o *RoleOutDynamicAllOf) GetGroupsInCountOk() (*int32, bool)`

GetGroupsInCountOk returns a tuple with the GroupsInCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsInCount

`func (o *RoleOutDynamicAllOf) SetGroupsInCount(v int32)`

SetGroupsInCount sets GroupsInCount field to given value.

### HasGroupsInCount

`func (o *RoleOutDynamicAllOf) HasGroupsInCount() bool`

HasGroupsInCount returns a boolean if a field has been set.

### GetGroupsIn

`func (o *RoleOutDynamicAllOf) GetGroupsIn() []AdditionalGroup`

GetGroupsIn returns the GroupsIn field if non-nil, zero value otherwise.

### GetGroupsInOk

`func (o *RoleOutDynamicAllOf) GetGroupsInOk() (*[]AdditionalGroup, bool)`

GetGroupsInOk returns a tuple with the GroupsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsIn

`func (o *RoleOutDynamicAllOf) SetGroupsIn(v []AdditionalGroup)`

SetGroupsIn sets GroupsIn field to given value.

### HasGroupsIn

`func (o *RoleOutDynamicAllOf) HasGroupsIn() bool`

HasGroupsIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


