# GroupWithPrincipalsAndRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**Created** | **string** |  | 
**Modified** | **string** |  | 
**Principals** | [**[]Principal**](Principal.md) |  | 
**Roles** | [**[]RoleOut**](RoleOut.md) |  | 

## Methods

### NewGroupWithPrincipalsAndRoles

`func NewGroupWithPrincipalsAndRoles(name string, uuid string, created string, modified string, principals []Principal, roles []RoleOut, ) *GroupWithPrincipalsAndRoles`

NewGroupWithPrincipalsAndRoles instantiates a new GroupWithPrincipalsAndRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithPrincipalsAndRolesWithDefaults

`func NewGroupWithPrincipalsAndRolesWithDefaults() *GroupWithPrincipalsAndRoles`

NewGroupWithPrincipalsAndRolesWithDefaults instantiates a new GroupWithPrincipalsAndRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupWithPrincipalsAndRoles) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupWithPrincipalsAndRoles) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupWithPrincipalsAndRoles) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GroupWithPrincipalsAndRoles) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupWithPrincipalsAndRoles) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupWithPrincipalsAndRoles) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupWithPrincipalsAndRoles) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *GroupWithPrincipalsAndRoles) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GroupWithPrincipalsAndRoles) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GroupWithPrincipalsAndRoles) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCreated

`func (o *GroupWithPrincipalsAndRoles) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupWithPrincipalsAndRoles) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupWithPrincipalsAndRoles) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetModified

`func (o *GroupWithPrincipalsAndRoles) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GroupWithPrincipalsAndRoles) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GroupWithPrincipalsAndRoles) SetModified(v string)`

SetModified sets Modified field to given value.


### GetPrincipals

`func (o *GroupWithPrincipalsAndRoles) GetPrincipals() []Principal`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *GroupWithPrincipalsAndRoles) GetPrincipalsOk() (*[]Principal, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *GroupWithPrincipalsAndRoles) SetPrincipals(v []Principal)`

SetPrincipals sets Principals field to given value.


### GetRoles

`func (o *GroupWithPrincipalsAndRoles) GetRoles() []RoleOut`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupWithPrincipalsAndRoles) GetRolesOk() (*[]RoleOut, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupWithPrincipalsAndRoles) SetRoles(v []RoleOut)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


