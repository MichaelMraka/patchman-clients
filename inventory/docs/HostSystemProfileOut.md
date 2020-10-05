# HostSystemProfileOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SystemProfile** | Pointer to [**SystemProfile**](SystemProfile.md) |  | [optional] 

## Methods

### NewHostSystemProfileOut

`func NewHostSystemProfileOut() *HostSystemProfileOut`

NewHostSystemProfileOut instantiates a new HostSystemProfileOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostSystemProfileOutWithDefaults

`func NewHostSystemProfileOutWithDefaults() *HostSystemProfileOut`

NewHostSystemProfileOutWithDefaults instantiates a new HostSystemProfileOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostSystemProfileOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostSystemProfileOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostSystemProfileOut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostSystemProfileOut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSystemProfile

`func (o *HostSystemProfileOut) GetSystemProfile() SystemProfile`

GetSystemProfile returns the SystemProfile field if non-nil, zero value otherwise.

### GetSystemProfileOk

`func (o *HostSystemProfileOut) GetSystemProfileOk() (*SystemProfile, bool)`

GetSystemProfileOk returns a tuple with the SystemProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemProfile

`func (o *HostSystemProfileOut) SetSystemProfile(v SystemProfile)`

SetSystemProfile sets SystemProfile field to given value.

### HasSystemProfile

`func (o *HostSystemProfileOut) HasSystemProfile() bool`

HasSystemProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


