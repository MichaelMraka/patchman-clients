# BulkHostOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BulkHostOutDetails**](BulkHostOutDetails.md) | List of hosts that were created, updated or failed to be processed. | [optional] 
**Errors** | Pointer to **int32** | Number of items in the \&quot;data\&quot; list that contain errors. | [optional] 
**Total** | Pointer to **int32** | Total number of items in the \&quot;data\&quot; list. | [optional] 

## Methods

### NewBulkHostOut

`func NewBulkHostOut() *BulkHostOut`

NewBulkHostOut instantiates a new BulkHostOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkHostOutWithDefaults

`func NewBulkHostOutWithDefaults() *BulkHostOut`

NewBulkHostOutWithDefaults instantiates a new BulkHostOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BulkHostOut) GetData() []BulkHostOutDetails`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkHostOut) GetDataOk() (*[]BulkHostOutDetails, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkHostOut) SetData(v []BulkHostOutDetails)`

SetData sets Data field to given value.

### HasData

`func (o *BulkHostOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *BulkHostOut) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkHostOut) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkHostOut) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkHostOut) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTotal

`func (o *BulkHostOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BulkHostOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BulkHostOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BulkHostOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


