# SystemProfileSapSystemOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of items on the current page | [optional] 
**Results** | Pointer to [**[]OneOfobjectobject**](oneOf&lt;object,object&gt;.md) | The list of sap_system values on the account | [optional] 
**Total** | Pointer to **int32** | Total number of items | [optional] 

## Methods

### NewSystemProfileSapSystemOut

`func NewSystemProfileSapSystemOut() *SystemProfileSapSystemOut`

NewSystemProfileSapSystemOut instantiates a new SystemProfileSapSystemOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileSapSystemOutWithDefaults

`func NewSystemProfileSapSystemOutWithDefaults() *SystemProfileSapSystemOut`

NewSystemProfileSapSystemOutWithDefaults instantiates a new SystemProfileSapSystemOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SystemProfileSapSystemOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SystemProfileSapSystemOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SystemProfileSapSystemOut) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SystemProfileSapSystemOut) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *SystemProfileSapSystemOut) GetResults() []OneOfobjectobject`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SystemProfileSapSystemOut) GetResultsOk() (*[]OneOfobjectobject, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SystemProfileSapSystemOut) SetResults(v []OneOfobjectobject)`

SetResults sets Results field to given value.

### HasResults

`func (o *SystemProfileSapSystemOut) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *SystemProfileSapSystemOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SystemProfileSapSystemOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SystemProfileSapSystemOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SystemProfileSapSystemOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


