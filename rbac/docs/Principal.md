# Principal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Email** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsOrgAdmin** | Pointer to **bool** |  | [optional] 

## Methods

### NewPrincipal

`func NewPrincipal(username string, email string, ) *Principal`

NewPrincipal instantiates a new Principal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalWithDefaults

`func NewPrincipalWithDefaults() *Principal`

NewPrincipalWithDefaults instantiates a new Principal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *Principal) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Principal) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Principal) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *Principal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Principal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Principal) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *Principal) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Principal) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Principal) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Principal) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Principal) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Principal) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Principal) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Principal) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetIsActive

`func (o *Principal) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Principal) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Principal) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Principal) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsOrgAdmin

`func (o *Principal) GetIsOrgAdmin() bool`

GetIsOrgAdmin returns the IsOrgAdmin field if non-nil, zero value otherwise.

### GetIsOrgAdminOk

`func (o *Principal) GetIsOrgAdminOk() (*bool, bool)`

GetIsOrgAdminOk returns a tuple with the IsOrgAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgAdmin

`func (o *Principal) SetIsOrgAdmin(v bool)`

SetIsOrgAdmin sets IsOrgAdmin field to given value.

### HasIsOrgAdmin

`func (o *Principal) HasIsOrgAdmin() bool`

HasIsOrgAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


