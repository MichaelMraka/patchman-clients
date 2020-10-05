# RPMPkgNamesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RpmNameList** | **[]string** |  | 
**ContentSetList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRPMPkgNamesRequest

`func NewRPMPkgNamesRequest(rpmNameList []string, ) *RPMPkgNamesRequest`

NewRPMPkgNamesRequest instantiates a new RPMPkgNamesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRPMPkgNamesRequestWithDefaults

`func NewRPMPkgNamesRequestWithDefaults() *RPMPkgNamesRequest`

NewRPMPkgNamesRequestWithDefaults instantiates a new RPMPkgNamesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRpmNameList

`func (o *RPMPkgNamesRequest) GetRpmNameList() []string`

GetRpmNameList returns the RpmNameList field if non-nil, zero value otherwise.

### GetRpmNameListOk

`func (o *RPMPkgNamesRequest) GetRpmNameListOk() (*[]string, bool)`

GetRpmNameListOk returns a tuple with the RpmNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmNameList

`func (o *RPMPkgNamesRequest) SetRpmNameList(v []string)`

SetRpmNameList sets RpmNameList field to given value.


### GetContentSetList

`func (o *RPMPkgNamesRequest) GetContentSetList() []string`

GetContentSetList returns the ContentSetList field if non-nil, zero value otherwise.

### GetContentSetListOk

`func (o *RPMPkgNamesRequest) GetContentSetListOk() (*[]string, bool)`

GetContentSetListOk returns a tuple with the ContentSetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSetList

`func (o *RPMPkgNamesRequest) SetContentSetList(v []string)`

SetContentSetList sets ContentSetList field to given value.

### HasContentSetList

`func (o *RPMPkgNamesRequest) HasContentSetList() bool`

HasContentSetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


