# GroupOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**Created** | **string** |  | 
**Modified** | **string** |  | 
**PrincipalCount** | Pointer to **int32** |  | [optional] 
**RoleCount** | Pointer to **int32** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] [default to false]
**PlatformDefault** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGroupOut

`func NewGroupOut(name string, uuid string, created string, modified string, ) *GroupOut`

NewGroupOut instantiates a new GroupOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupOutWithDefaults

`func NewGroupOutWithDefaults() *GroupOut`

NewGroupOutWithDefaults instantiates a new GroupOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GroupOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *GroupOut) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GroupOut) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GroupOut) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCreated

`func (o *GroupOut) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupOut) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupOut) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetModified

`func (o *GroupOut) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GroupOut) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GroupOut) SetModified(v string)`

SetModified sets Modified field to given value.


### GetPrincipalCount

`func (o *GroupOut) GetPrincipalCount() int32`

GetPrincipalCount returns the PrincipalCount field if non-nil, zero value otherwise.

### GetPrincipalCountOk

`func (o *GroupOut) GetPrincipalCountOk() (*int32, bool)`

GetPrincipalCountOk returns a tuple with the PrincipalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalCount

`func (o *GroupOut) SetPrincipalCount(v int32)`

SetPrincipalCount sets PrincipalCount field to given value.

### HasPrincipalCount

`func (o *GroupOut) HasPrincipalCount() bool`

HasPrincipalCount returns a boolean if a field has been set.

### GetRoleCount

`func (o *GroupOut) GetRoleCount() int32`

GetRoleCount returns the RoleCount field if non-nil, zero value otherwise.

### GetRoleCountOk

`func (o *GroupOut) GetRoleCountOk() (*int32, bool)`

GetRoleCountOk returns a tuple with the RoleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCount

`func (o *GroupOut) SetRoleCount(v int32)`

SetRoleCount sets RoleCount field to given value.

### HasRoleCount

`func (o *GroupOut) HasRoleCount() bool`

HasRoleCount returns a boolean if a field has been set.

### GetSystem

`func (o *GroupOut) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *GroupOut) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *GroupOut) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *GroupOut) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetPlatformDefault

`func (o *GroupOut) GetPlatformDefault() bool`

GetPlatformDefault returns the PlatformDefault field if non-nil, zero value otherwise.

### GetPlatformDefaultOk

`func (o *GroupOut) GetPlatformDefaultOk() (*bool, bool)`

GetPlatformDefaultOk returns a tuple with the PlatformDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformDefault

`func (o *GroupOut) SetPlatformDefault(v bool)`

SetPlatformDefault sets PlatformDefault field to given value.

### HasPlatformDefault

`func (o *GroupOut) HasPlatformDefault() bool`

HasPlatformDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


