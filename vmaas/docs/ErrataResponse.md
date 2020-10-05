# ErrataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**Pages** | Pointer to **float32** |  | [optional] 
**ErrataList** | Pointer to [**map[string]ErrataResponseErrataList**](ErrataResponse_errata_list.md) |  | [optional] 
**ModifiedSince** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **[]string** |  | [optional] 

## Methods

### NewErrataResponse

`func NewErrataResponse() *ErrataResponse`

NewErrataResponse instantiates a new ErrataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrataResponseWithDefaults

`func NewErrataResponseWithDefaults() *ErrataResponse`

NewErrataResponseWithDefaults instantiates a new ErrataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ErrataResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ErrataResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ErrataResponse) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ErrataResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ErrataResponse) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ErrataResponse) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ErrataResponse) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ErrataResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPages

`func (o *ErrataResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ErrataResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ErrataResponse) SetPages(v float32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ErrataResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetErrataList

`func (o *ErrataResponse) GetErrataList() map[string]ErrataResponseErrataList`

GetErrataList returns the ErrataList field if non-nil, zero value otherwise.

### GetErrataListOk

`func (o *ErrataResponse) GetErrataListOk() (*map[string]ErrataResponseErrataList, bool)`

GetErrataListOk returns a tuple with the ErrataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrataList

`func (o *ErrataResponse) SetErrataList(v map[string]ErrataResponseErrataList)`

SetErrataList sets ErrataList field to given value.

### HasErrataList

`func (o *ErrataResponse) HasErrataList() bool`

HasErrataList returns a boolean if a field has been set.

### GetModifiedSince

`func (o *ErrataResponse) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *ErrataResponse) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *ErrataResponse) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *ErrataResponse) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.

### GetType

`func (o *ErrataResponse) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrataResponse) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrataResponse) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrataResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeverity

`func (o *ErrataResponse) GetSeverity() []string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ErrataResponse) GetSeverityOk() (*[]string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ErrataResponse) SetSeverity(v []string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ErrataResponse) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


