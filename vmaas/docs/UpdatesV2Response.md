# UpdatesV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateList** | Pointer to [**map[string]UpdatesV2ResponseUpdateList**](UpdatesV2Response_update_list.md) |  | [optional] 
**RepositoryList** | Pointer to **[]string** |  | [optional] 
**ModulesList** | Pointer to [**[]UpdatesRequestModulesList**](UpdatesRequest_modules_list.md) |  | [optional] 
**Releasever** | Pointer to **string** |  | [optional] 
**Basearch** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatesV2Response

`func NewUpdatesV2Response() *UpdatesV2Response`

NewUpdatesV2Response instantiates a new UpdatesV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatesV2ResponseWithDefaults

`func NewUpdatesV2ResponseWithDefaults() *UpdatesV2Response`

NewUpdatesV2ResponseWithDefaults instantiates a new UpdatesV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateList

`func (o *UpdatesV2Response) GetUpdateList() map[string]UpdatesV2ResponseUpdateList`

GetUpdateList returns the UpdateList field if non-nil, zero value otherwise.

### GetUpdateListOk

`func (o *UpdatesV2Response) GetUpdateListOk() (*map[string]UpdatesV2ResponseUpdateList, bool)`

GetUpdateListOk returns a tuple with the UpdateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateList

`func (o *UpdatesV2Response) SetUpdateList(v map[string]UpdatesV2ResponseUpdateList)`

SetUpdateList sets UpdateList field to given value.

### HasUpdateList

`func (o *UpdatesV2Response) HasUpdateList() bool`

HasUpdateList returns a boolean if a field has been set.

### GetRepositoryList

`func (o *UpdatesV2Response) GetRepositoryList() []string`

GetRepositoryList returns the RepositoryList field if non-nil, zero value otherwise.

### GetRepositoryListOk

`func (o *UpdatesV2Response) GetRepositoryListOk() (*[]string, bool)`

GetRepositoryListOk returns a tuple with the RepositoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryList

`func (o *UpdatesV2Response) SetRepositoryList(v []string)`

SetRepositoryList sets RepositoryList field to given value.

### HasRepositoryList

`func (o *UpdatesV2Response) HasRepositoryList() bool`

HasRepositoryList returns a boolean if a field has been set.

### GetModulesList

`func (o *UpdatesV2Response) GetModulesList() []UpdatesRequestModulesList`

GetModulesList returns the ModulesList field if non-nil, zero value otherwise.

### GetModulesListOk

`func (o *UpdatesV2Response) GetModulesListOk() (*[]UpdatesRequestModulesList, bool)`

GetModulesListOk returns a tuple with the ModulesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesList

`func (o *UpdatesV2Response) SetModulesList(v []UpdatesRequestModulesList)`

SetModulesList sets ModulesList field to given value.

### HasModulesList

`func (o *UpdatesV2Response) HasModulesList() bool`

HasModulesList returns a boolean if a field has been set.

### GetReleasever

`func (o *UpdatesV2Response) GetReleasever() string`

GetReleasever returns the Releasever field if non-nil, zero value otherwise.

### GetReleaseverOk

`func (o *UpdatesV2Response) GetReleaseverOk() (*string, bool)`

GetReleaseverOk returns a tuple with the Releasever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasever

`func (o *UpdatesV2Response) SetReleasever(v string)`

SetReleasever sets Releasever field to given value.

### HasReleasever

`func (o *UpdatesV2Response) HasReleasever() bool`

HasReleasever returns a boolean if a field has been set.

### GetBasearch

`func (o *UpdatesV2Response) GetBasearch() string`

GetBasearch returns the Basearch field if non-nil, zero value otherwise.

### GetBasearchOk

`func (o *UpdatesV2Response) GetBasearchOk() (*string, bool)`

GetBasearchOk returns a tuple with the Basearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasearch

`func (o *UpdatesV2Response) SetBasearch(v string)`

SetBasearch sets Basearch field to given value.

### HasBasearch

`func (o *UpdatesV2Response) HasBasearch() bool`

HasBasearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


