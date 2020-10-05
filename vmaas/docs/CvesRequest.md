# CvesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**CveList** | **[]string** |  | 
**ModifiedSince** | Pointer to **string** |  | [optional] 
**PublishedSince** | Pointer to **string** |  | [optional] 
**RhOnly** | Pointer to **bool** |  | [optional] 
**ErrataAssociated** | Pointer to **bool** | Return only those CVEs which are associated with at least one errata. Defaults to false. | [optional] 

## Methods

### NewCvesRequest

`func NewCvesRequest(cveList []string, ) *CvesRequest`

NewCvesRequest instantiates a new CvesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCvesRequestWithDefaults

`func NewCvesRequestWithDefaults() *CvesRequest`

NewCvesRequestWithDefaults instantiates a new CvesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CvesRequest) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CvesRequest) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CvesRequest) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *CvesRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *CvesRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CvesRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CvesRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CvesRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetCveList

`func (o *CvesRequest) GetCveList() []string`

GetCveList returns the CveList field if non-nil, zero value otherwise.

### GetCveListOk

`func (o *CvesRequest) GetCveListOk() (*[]string, bool)`

GetCveListOk returns a tuple with the CveList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveList

`func (o *CvesRequest) SetCveList(v []string)`

SetCveList sets CveList field to given value.


### GetModifiedSince

`func (o *CvesRequest) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *CvesRequest) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *CvesRequest) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *CvesRequest) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.

### GetPublishedSince

`func (o *CvesRequest) GetPublishedSince() string`

GetPublishedSince returns the PublishedSince field if non-nil, zero value otherwise.

### GetPublishedSinceOk

`func (o *CvesRequest) GetPublishedSinceOk() (*string, bool)`

GetPublishedSinceOk returns a tuple with the PublishedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedSince

`func (o *CvesRequest) SetPublishedSince(v string)`

SetPublishedSince sets PublishedSince field to given value.

### HasPublishedSince

`func (o *CvesRequest) HasPublishedSince() bool`

HasPublishedSince returns a boolean if a field has been set.

### GetRhOnly

`func (o *CvesRequest) GetRhOnly() bool`

GetRhOnly returns the RhOnly field if non-nil, zero value otherwise.

### GetRhOnlyOk

`func (o *CvesRequest) GetRhOnlyOk() (*bool, bool)`

GetRhOnlyOk returns a tuple with the RhOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhOnly

`func (o *CvesRequest) SetRhOnly(v bool)`

SetRhOnly sets RhOnly field to given value.

### HasRhOnly

`func (o *CvesRequest) HasRhOnly() bool`

HasRhOnly returns a boolean if a field has been set.

### GetErrataAssociated

`func (o *CvesRequest) GetErrataAssociated() bool`

GetErrataAssociated returns the ErrataAssociated field if non-nil, zero value otherwise.

### GetErrataAssociatedOk

`func (o *CvesRequest) GetErrataAssociatedOk() (*bool, bool)`

GetErrataAssociatedOk returns a tuple with the ErrataAssociated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrataAssociated

`func (o *CvesRequest) SetErrataAssociated(v bool)`

SetErrataAssociated sets ErrataAssociated field to given value.

### HasErrataAssociated

`func (o *CvesRequest) HasErrataAssociated() bool`

HasErrataAssociated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


