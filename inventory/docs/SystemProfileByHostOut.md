# SystemProfileByHostOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | A number of entries on the current page. | 
**Page** | **int32** | A current page number. | 
**PerPage** | **int32** | A page size – a number of entries per single page. | 
**Results** | [**[]HostSystemProfileOut**](HostSystemProfileOut.md) | Actual host search query result entries. | 
**Total** | **int32** | A total count of the found entries. | 

## Methods

### NewSystemProfileByHostOut

`func NewSystemProfileByHostOut(count int32, page int32, perPage int32, results []HostSystemProfileOut, total int32, ) *SystemProfileByHostOut`

NewSystemProfileByHostOut instantiates a new SystemProfileByHostOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileByHostOutWithDefaults

`func NewSystemProfileByHostOutWithDefaults() *SystemProfileByHostOut`

NewSystemProfileByHostOutWithDefaults instantiates a new SystemProfileByHostOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SystemProfileByHostOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SystemProfileByHostOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SystemProfileByHostOut) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *SystemProfileByHostOut) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SystemProfileByHostOut) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SystemProfileByHostOut) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *SystemProfileByHostOut) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *SystemProfileByHostOut) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *SystemProfileByHostOut) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetResults

`func (o *SystemProfileByHostOut) GetResults() []HostSystemProfileOut`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SystemProfileByHostOut) GetResultsOk() (*[]HostSystemProfileOut, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SystemProfileByHostOut) SetResults(v []HostSystemProfileOut)`

SetResults sets Results field to given value.


### GetTotal

`func (o *SystemProfileByHostOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SystemProfileByHostOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SystemProfileByHostOut) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


