# UpdatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateList** | Pointer to [**map[string]UpdatesResponseUpdateList**](UpdatesResponse_update_list.md) |  | [optional] 
**RepositoryList** | Pointer to **[]string** |  | [optional] 
**ModulesList** | Pointer to [**[]UpdatesRequestModulesList**](UpdatesRequest_modules_list.md) |  | [optional] 
**Releasever** | Pointer to **string** |  | [optional] 
**Basearch** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatesResponse

`func NewUpdatesResponse() *UpdatesResponse`

NewUpdatesResponse instantiates a new UpdatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatesResponseWithDefaults

`func NewUpdatesResponseWithDefaults() *UpdatesResponse`

NewUpdatesResponseWithDefaults instantiates a new UpdatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateList

`func (o *UpdatesResponse) GetUpdateList() map[string]UpdatesResponseUpdateList`

GetUpdateList returns the UpdateList field if non-nil, zero value otherwise.

### GetUpdateListOk

`func (o *UpdatesResponse) GetUpdateListOk() (*map[string]UpdatesResponseUpdateList, bool)`

GetUpdateListOk returns a tuple with the UpdateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateList

`func (o *UpdatesResponse) SetUpdateList(v map[string]UpdatesResponseUpdateList)`

SetUpdateList sets UpdateList field to given value.

### HasUpdateList

`func (o *UpdatesResponse) HasUpdateList() bool`

HasUpdateList returns a boolean if a field has been set.

### GetRepositoryList

`func (o *UpdatesResponse) GetRepositoryList() []string`

GetRepositoryList returns the RepositoryList field if non-nil, zero value otherwise.

### GetRepositoryListOk

`func (o *UpdatesResponse) GetRepositoryListOk() (*[]string, bool)`

GetRepositoryListOk returns a tuple with the RepositoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryList

`func (o *UpdatesResponse) SetRepositoryList(v []string)`

SetRepositoryList sets RepositoryList field to given value.

### HasRepositoryList

`func (o *UpdatesResponse) HasRepositoryList() bool`

HasRepositoryList returns a boolean if a field has been set.

### GetModulesList

`func (o *UpdatesResponse) GetModulesList() []UpdatesRequestModulesList`

GetModulesList returns the ModulesList field if non-nil, zero value otherwise.

### GetModulesListOk

`func (o *UpdatesResponse) GetModulesListOk() (*[]UpdatesRequestModulesList, bool)`

GetModulesListOk returns a tuple with the ModulesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesList

`func (o *UpdatesResponse) SetModulesList(v []UpdatesRequestModulesList)`

SetModulesList sets ModulesList field to given value.

### HasModulesList

`func (o *UpdatesResponse) HasModulesList() bool`

HasModulesList returns a boolean if a field has been set.

### GetReleasever

`func (o *UpdatesResponse) GetReleasever() string`

GetReleasever returns the Releasever field if non-nil, zero value otherwise.

### GetReleaseverOk

`func (o *UpdatesResponse) GetReleaseverOk() (*string, bool)`

GetReleaseverOk returns a tuple with the Releasever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasever

`func (o *UpdatesResponse) SetReleasever(v string)`

SetReleasever sets Releasever field to given value.

### HasReleasever

`func (o *UpdatesResponse) HasReleasever() bool`

HasReleasever returns a boolean if a field has been set.

### GetBasearch

`func (o *UpdatesResponse) GetBasearch() string`

GetBasearch returns the Basearch field if non-nil, zero value otherwise.

### GetBasearchOk

`func (o *UpdatesResponse) GetBasearchOk() (*string, bool)`

GetBasearchOk returns a tuple with the Basearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasearch

`func (o *UpdatesResponse) SetBasearch(v string)`

SetBasearch sets Basearch field to given value.

### HasBasearch

`func (o *UpdatesResponse) HasBasearch() bool`

HasBasearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


