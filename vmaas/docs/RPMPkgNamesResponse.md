# RPMPkgNamesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **interface{}** |  | [optional] 
**RpmNameList** | Pointer to [**map[string][]string**](array.md) |  | [optional] 

## Methods

### NewRPMPkgNamesResponse

`func NewRPMPkgNamesResponse() *RPMPkgNamesResponse`

NewRPMPkgNamesResponse instantiates a new RPMPkgNamesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRPMPkgNamesResponseWithDefaults

`func NewRPMPkgNamesResponseWithDefaults() *RPMPkgNamesResponse`

NewRPMPkgNamesResponseWithDefaults instantiates a new RPMPkgNamesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *RPMPkgNamesResponse) GetLastChange() interface{}`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *RPMPkgNamesResponse) GetLastChangeOk() (*interface{}, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *RPMPkgNamesResponse) SetLastChange(v interface{})`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *RPMPkgNamesResponse) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *RPMPkgNamesResponse) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *RPMPkgNamesResponse) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetRpmNameList

`func (o *RPMPkgNamesResponse) GetRpmNameList() map[string][]string`

GetRpmNameList returns the RpmNameList field if non-nil, zero value otherwise.

### GetRpmNameListOk

`func (o *RPMPkgNamesResponse) GetRpmNameListOk() (*map[string][]string, bool)`

GetRpmNameListOk returns a tuple with the RpmNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmNameList

`func (o *RPMPkgNamesResponse) SetRpmNameList(v map[string][]string)`

SetRpmNameList sets RpmNameList field to given value.

### HasRpmNameList

`func (o *RPMPkgNamesResponse) HasRpmNameList() bool`

HasRpmNameList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


