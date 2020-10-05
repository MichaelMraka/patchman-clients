# GroupWithPrincipals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**Created** | **string** |  | 
**Modified** | **string** |  | 
**Principals** | [**[]Principal**](Principal.md) |  | 

## Methods

### NewGroupWithPrincipals

`func NewGroupWithPrincipals(name string, uuid string, created string, modified string, principals []Principal, ) *GroupWithPrincipals`

NewGroupWithPrincipals instantiates a new GroupWithPrincipals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithPrincipalsWithDefaults

`func NewGroupWithPrincipalsWithDefaults() *GroupWithPrincipals`

NewGroupWithPrincipalsWithDefaults instantiates a new GroupWithPrincipals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupWithPrincipals) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupWithPrincipals) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupWithPrincipals) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GroupWithPrincipals) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupWithPrincipals) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupWithPrincipals) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupWithPrincipals) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *GroupWithPrincipals) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GroupWithPrincipals) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GroupWithPrincipals) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCreated

`func (o *GroupWithPrincipals) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupWithPrincipals) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupWithPrincipals) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetModified

`func (o *GroupWithPrincipals) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GroupWithPrincipals) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GroupWithPrincipals) SetModified(v string)`

SetModified sets Modified field to given value.


### GetPrincipals

`func (o *GroupWithPrincipals) GetPrincipals() []Principal`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *GroupWithPrincipals) GetPrincipalsOk() (*[]Principal, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *GroupWithPrincipals) SetPrincipals(v []Principal)`

SetPrincipals sets Principals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


