# PrincipalOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Email** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsOrgAdmin** | Pointer to **bool** |  | [optional] 
**Uuid** | **string** |  | 

## Methods

### NewPrincipalOut

`func NewPrincipalOut(username string, email string, uuid string, ) *PrincipalOut`

NewPrincipalOut instantiates a new PrincipalOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalOutWithDefaults

`func NewPrincipalOutWithDefaults() *PrincipalOut`

NewPrincipalOutWithDefaults instantiates a new PrincipalOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *PrincipalOut) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PrincipalOut) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PrincipalOut) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *PrincipalOut) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PrincipalOut) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PrincipalOut) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *PrincipalOut) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PrincipalOut) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PrincipalOut) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PrincipalOut) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PrincipalOut) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PrincipalOut) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PrincipalOut) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PrincipalOut) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetIsActive

`func (o *PrincipalOut) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PrincipalOut) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PrincipalOut) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PrincipalOut) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsOrgAdmin

`func (o *PrincipalOut) GetIsOrgAdmin() bool`

GetIsOrgAdmin returns the IsOrgAdmin field if non-nil, zero value otherwise.

### GetIsOrgAdminOk

`func (o *PrincipalOut) GetIsOrgAdminOk() (*bool, bool)`

GetIsOrgAdminOk returns a tuple with the IsOrgAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgAdmin

`func (o *PrincipalOut) SetIsOrgAdmin(v bool)`

SetIsOrgAdmin sets IsOrgAdmin field to given value.

### HasIsOrgAdmin

`func (o *PrincipalOut) HasIsOrgAdmin() bool`

HasIsOrgAdmin returns a boolean if a field has been set.

### GetUuid

`func (o *PrincipalOut) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PrincipalOut) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PrincipalOut) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


