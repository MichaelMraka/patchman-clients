# UpdatesResponseUpdateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableUpdates** | Pointer to [**[]UpdatesResponseAvailableUpdates**](UpdatesResponse_available_updates.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatesResponseUpdateList

`func NewUpdatesResponseUpdateList() *UpdatesResponseUpdateList`

NewUpdatesResponseUpdateList instantiates a new UpdatesResponseUpdateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatesResponseUpdateListWithDefaults

`func NewUpdatesResponseUpdateListWithDefaults() *UpdatesResponseUpdateList`

NewUpdatesResponseUpdateListWithDefaults instantiates a new UpdatesResponseUpdateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableUpdates

`func (o *UpdatesResponseUpdateList) GetAvailableUpdates() []UpdatesResponseAvailableUpdates`

GetAvailableUpdates returns the AvailableUpdates field if non-nil, zero value otherwise.

### GetAvailableUpdatesOk

`func (o *UpdatesResponseUpdateList) GetAvailableUpdatesOk() (*[]UpdatesResponseAvailableUpdates, bool)`

GetAvailableUpdatesOk returns a tuple with the AvailableUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUpdates

`func (o *UpdatesResponseUpdateList) SetAvailableUpdates(v []UpdatesResponseAvailableUpdates)`

SetAvailableUpdates sets AvailableUpdates field to given value.

### HasAvailableUpdates

`func (o *UpdatesResponseUpdateList) HasAvailableUpdates() bool`

HasAvailableUpdates returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatesResponseUpdateList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatesResponseUpdateList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatesResponseUpdateList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatesResponseUpdateList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSummary

`func (o *UpdatesResponseUpdateList) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *UpdatesResponseUpdateList) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *UpdatesResponseUpdateList) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *UpdatesResponseUpdateList) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


