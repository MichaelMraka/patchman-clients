# ReposResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**Pages** | Pointer to **float32** |  | [optional] 
**RepositoryList** | Pointer to [**map[string][]map[string]interface{}**](array.md) |  | [optional] 

## Methods

### NewReposResponse

`func NewReposResponse() *ReposResponse`

NewReposResponse instantiates a new ReposResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposResponseWithDefaults

`func NewReposResponseWithDefaults() *ReposResponse`

NewReposResponseWithDefaults instantiates a new ReposResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ReposResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ReposResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ReposResponse) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ReposResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ReposResponse) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ReposResponse) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ReposResponse) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ReposResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPages

`func (o *ReposResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ReposResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ReposResponse) SetPages(v float32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ReposResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetRepositoryList

`func (o *ReposResponse) GetRepositoryList() map[string][]map[string]interface{}`

GetRepositoryList returns the RepositoryList field if non-nil, zero value otherwise.

### GetRepositoryListOk

`func (o *ReposResponse) GetRepositoryListOk() (*map[string][]map[string]interface{}, bool)`

GetRepositoryListOk returns a tuple with the RepositoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryList

`func (o *ReposResponse) SetRepositoryList(v map[string][]map[string]interface{})`

SetRepositoryList sets RepositoryList field to given value.

### HasRepositoryList

`func (o *ReposResponse) HasRepositoryList() bool`

HasRepositoryList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


