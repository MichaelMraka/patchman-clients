# ErrataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**ErrataList** | **[]string** |  | 
**ModifiedSince** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**OneOfstringarray**](oneOf&lt;string,array&gt;.md) |  | [optional] 
**Severity** | Pointer to [**OneOfstringarray**](oneOf&lt;string,array&gt;.md) |  | [optional] 

## Methods

### NewErrataRequest

`func NewErrataRequest(errataList []string, ) *ErrataRequest`

NewErrataRequest instantiates a new ErrataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrataRequestWithDefaults

`func NewErrataRequestWithDefaults() *ErrataRequest`

NewErrataRequestWithDefaults instantiates a new ErrataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ErrataRequest) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ErrataRequest) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ErrataRequest) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ErrataRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ErrataRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ErrataRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ErrataRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ErrataRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetErrataList

`func (o *ErrataRequest) GetErrataList() []string`

GetErrataList returns the ErrataList field if non-nil, zero value otherwise.

### GetErrataListOk

`func (o *ErrataRequest) GetErrataListOk() (*[]string, bool)`

GetErrataListOk returns a tuple with the ErrataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrataList

`func (o *ErrataRequest) SetErrataList(v []string)`

SetErrataList sets ErrataList field to given value.


### GetModifiedSince

`func (o *ErrataRequest) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *ErrataRequest) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *ErrataRequest) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *ErrataRequest) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.

### GetType

`func (o *ErrataRequest) GetType() OneOfstringarray`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrataRequest) GetTypeOk() (*OneOfstringarray, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrataRequest) SetType(v OneOfstringarray)`

SetType sets Type field to given value.

### HasType

`func (o *ErrataRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeverity

`func (o *ErrataRequest) GetSeverity() OneOfstringarray`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ErrataRequest) GetSeverityOk() (*OneOfstringarray, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ErrataRequest) SetSeverity(v OneOfstringarray)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ErrataRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


