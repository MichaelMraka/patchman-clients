# CvesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**Pages** | Pointer to **float32** |  | [optional] 
**CveList** | Pointer to [**map[string]CvesResponseCveList**](CvesResponse_cve_list.md) |  | [optional] 
**ModifiedSince** | Pointer to **string** |  | [optional] 

## Methods

### NewCvesResponse

`func NewCvesResponse() *CvesResponse`

NewCvesResponse instantiates a new CvesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCvesResponseWithDefaults

`func NewCvesResponseWithDefaults() *CvesResponse`

NewCvesResponseWithDefaults instantiates a new CvesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CvesResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CvesResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CvesResponse) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *CvesResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *CvesResponse) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CvesResponse) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CvesResponse) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CvesResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPages

`func (o *CvesResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *CvesResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *CvesResponse) SetPages(v float32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *CvesResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetCveList

`func (o *CvesResponse) GetCveList() map[string]CvesResponseCveList`

GetCveList returns the CveList field if non-nil, zero value otherwise.

### GetCveListOk

`func (o *CvesResponse) GetCveListOk() (*map[string]CvesResponseCveList, bool)`

GetCveListOk returns a tuple with the CveList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveList

`func (o *CvesResponse) SetCveList(v map[string]CvesResponseCveList)`

SetCveList sets CveList field to given value.

### HasCveList

`func (o *CvesResponse) HasCveList() bool`

HasCveList returns a boolean if a field has been set.

### GetModifiedSince

`func (o *CvesResponse) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *CvesResponse) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *CvesResponse) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *CvesResponse) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


