# BulkHostOutDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Details about why a host failed to be created or updated. | [optional] 
**Host** | Pointer to [**CreateHostOut**](CreateHostOut.md) |  | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the results of the host create/update process | [optional] 
**Title** | Pointer to **string** | Short description of why a host failed to be created or updated. | [optional] 

## Methods

### NewBulkHostOutDetails

`func NewBulkHostOutDetails() *BulkHostOutDetails`

NewBulkHostOutDetails instantiates a new BulkHostOutDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkHostOutDetailsWithDefaults

`func NewBulkHostOutDetailsWithDefaults() *BulkHostOutDetails`

NewBulkHostOutDetailsWithDefaults instantiates a new BulkHostOutDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *BulkHostOutDetails) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BulkHostOutDetails) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BulkHostOutDetails) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *BulkHostOutDetails) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetHost

`func (o *BulkHostOutDetails) GetHost() CreateHostOut`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *BulkHostOutDetails) GetHostOk() (*CreateHostOut, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *BulkHostOutDetails) SetHost(v CreateHostOut)`

SetHost sets Host field to given value.

### HasHost

`func (o *BulkHostOutDetails) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetStatus

`func (o *BulkHostOutDetails) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkHostOutDetails) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkHostOutDetails) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkHostOutDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *BulkHostOutDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BulkHostOutDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BulkHostOutDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BulkHostOutDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


