# ReposRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**RepositoryList** | **[]string** |  | 
**ModifiedSince** | Pointer to **string** | Return only repositories changed after the given date | [optional] 

## Methods

### NewReposRequest

`func NewReposRequest(repositoryList []string, ) *ReposRequest`

NewReposRequest instantiates a new ReposRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposRequestWithDefaults

`func NewReposRequestWithDefaults() *ReposRequest`

NewReposRequestWithDefaults instantiates a new ReposRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ReposRequest) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ReposRequest) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ReposRequest) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ReposRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ReposRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ReposRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ReposRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ReposRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetRepositoryList

`func (o *ReposRequest) GetRepositoryList() []string`

GetRepositoryList returns the RepositoryList field if non-nil, zero value otherwise.

### GetRepositoryListOk

`func (o *ReposRequest) GetRepositoryListOk() (*[]string, bool)`

GetRepositoryListOk returns a tuple with the RepositoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryList

`func (o *ReposRequest) SetRepositoryList(v []string)`

SetRepositoryList sets RepositoryList field to given value.


### GetModifiedSince

`func (o *ReposRequest) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *ReposRequest) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *ReposRequest) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *ReposRequest) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


