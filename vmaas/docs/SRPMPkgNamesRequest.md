# SRPMPkgNamesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SrpmNameList** | **[]string** |  | 
**ContentSetList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSRPMPkgNamesRequest

`func NewSRPMPkgNamesRequest(srpmNameList []string, ) *SRPMPkgNamesRequest`

NewSRPMPkgNamesRequest instantiates a new SRPMPkgNamesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSRPMPkgNamesRequestWithDefaults

`func NewSRPMPkgNamesRequestWithDefaults() *SRPMPkgNamesRequest`

NewSRPMPkgNamesRequestWithDefaults instantiates a new SRPMPkgNamesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrpmNameList

`func (o *SRPMPkgNamesRequest) GetSrpmNameList() []string`

GetSrpmNameList returns the SrpmNameList field if non-nil, zero value otherwise.

### GetSrpmNameListOk

`func (o *SRPMPkgNamesRequest) GetSrpmNameListOk() (*[]string, bool)`

GetSrpmNameListOk returns a tuple with the SrpmNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrpmNameList

`func (o *SRPMPkgNamesRequest) SetSrpmNameList(v []string)`

SetSrpmNameList sets SrpmNameList field to given value.


### GetContentSetList

`func (o *SRPMPkgNamesRequest) GetContentSetList() []string`

GetContentSetList returns the ContentSetList field if non-nil, zero value otherwise.

### GetContentSetListOk

`func (o *SRPMPkgNamesRequest) GetContentSetListOk() (*[]string, bool)`

GetContentSetListOk returns a tuple with the ContentSetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSetList

`func (o *SRPMPkgNamesRequest) SetContentSetList(v []string)`

SetContentSetList sets ContentSetList field to given value.

### HasContentSetList

`func (o *SRPMPkgNamesRequest) HasContentSetList() bool`

HasContentSetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


