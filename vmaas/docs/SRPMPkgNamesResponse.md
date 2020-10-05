# SRPMPkgNamesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **interface{}** |  | [optional] 
**SrpmNameList** | Pointer to [**map[string]map[string][]string**](map.md) |  | [optional] 

## Methods

### NewSRPMPkgNamesResponse

`func NewSRPMPkgNamesResponse() *SRPMPkgNamesResponse`

NewSRPMPkgNamesResponse instantiates a new SRPMPkgNamesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSRPMPkgNamesResponseWithDefaults

`func NewSRPMPkgNamesResponseWithDefaults() *SRPMPkgNamesResponse`

NewSRPMPkgNamesResponseWithDefaults instantiates a new SRPMPkgNamesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *SRPMPkgNamesResponse) GetLastChange() interface{}`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *SRPMPkgNamesResponse) GetLastChangeOk() (*interface{}, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *SRPMPkgNamesResponse) SetLastChange(v interface{})`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *SRPMPkgNamesResponse) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *SRPMPkgNamesResponse) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *SRPMPkgNamesResponse) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetSrpmNameList

`func (o *SRPMPkgNamesResponse) GetSrpmNameList() map[string]map[string][]string`

GetSrpmNameList returns the SrpmNameList field if non-nil, zero value otherwise.

### GetSrpmNameListOk

`func (o *SRPMPkgNamesResponse) GetSrpmNameListOk() (*map[string]map[string][]string, bool)`

GetSrpmNameListOk returns a tuple with the SrpmNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrpmNameList

`func (o *SRPMPkgNamesResponse) SetSrpmNameList(v map[string]map[string][]string)`

SetSrpmNameList sets SrpmNameList field to given value.

### HasSrpmNameList

`func (o *SRPMPkgNamesResponse) HasSrpmNameList() bool`

HasSrpmNameList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


