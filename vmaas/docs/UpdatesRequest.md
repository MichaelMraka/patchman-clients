# UpdatesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageList** | **[]string** |  | 
**RepositoryList** | Pointer to **[]string** |  | [optional] 
**ModulesList** | Pointer to [**[]UpdatesRequestModulesList**](UpdatesRequest_modules_list.md) |  | [optional] 
**Releasever** | Pointer to **string** |  | [optional] 
**Basearch** | Pointer to **string** |  | [optional] 
**LatestOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdatesRequest

`func NewUpdatesRequest(packageList []string, ) *UpdatesRequest`

NewUpdatesRequest instantiates a new UpdatesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatesRequestWithDefaults

`func NewUpdatesRequestWithDefaults() *UpdatesRequest`

NewUpdatesRequestWithDefaults instantiates a new UpdatesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageList

`func (o *UpdatesRequest) GetPackageList() []string`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *UpdatesRequest) GetPackageListOk() (*[]string, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *UpdatesRequest) SetPackageList(v []string)`

SetPackageList sets PackageList field to given value.


### GetRepositoryList

`func (o *UpdatesRequest) GetRepositoryList() []string`

GetRepositoryList returns the RepositoryList field if non-nil, zero value otherwise.

### GetRepositoryListOk

`func (o *UpdatesRequest) GetRepositoryListOk() (*[]string, bool)`

GetRepositoryListOk returns a tuple with the RepositoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryList

`func (o *UpdatesRequest) SetRepositoryList(v []string)`

SetRepositoryList sets RepositoryList field to given value.

### HasRepositoryList

`func (o *UpdatesRequest) HasRepositoryList() bool`

HasRepositoryList returns a boolean if a field has been set.

### GetModulesList

`func (o *UpdatesRequest) GetModulesList() []UpdatesRequestModulesList`

GetModulesList returns the ModulesList field if non-nil, zero value otherwise.

### GetModulesListOk

`func (o *UpdatesRequest) GetModulesListOk() (*[]UpdatesRequestModulesList, bool)`

GetModulesListOk returns a tuple with the ModulesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesList

`func (o *UpdatesRequest) SetModulesList(v []UpdatesRequestModulesList)`

SetModulesList sets ModulesList field to given value.

### HasModulesList

`func (o *UpdatesRequest) HasModulesList() bool`

HasModulesList returns a boolean if a field has been set.

### GetReleasever

`func (o *UpdatesRequest) GetReleasever() string`

GetReleasever returns the Releasever field if non-nil, zero value otherwise.

### GetReleaseverOk

`func (o *UpdatesRequest) GetReleaseverOk() (*string, bool)`

GetReleaseverOk returns a tuple with the Releasever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasever

`func (o *UpdatesRequest) SetReleasever(v string)`

SetReleasever sets Releasever field to given value.

### HasReleasever

`func (o *UpdatesRequest) HasReleasever() bool`

HasReleasever returns a boolean if a field has been set.

### GetBasearch

`func (o *UpdatesRequest) GetBasearch() string`

GetBasearch returns the Basearch field if non-nil, zero value otherwise.

### GetBasearchOk

`func (o *UpdatesRequest) GetBasearchOk() (*string, bool)`

GetBasearchOk returns a tuple with the Basearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasearch

`func (o *UpdatesRequest) SetBasearch(v string)`

SetBasearch sets Basearch field to given value.

### HasBasearch

`func (o *UpdatesRequest) HasBasearch() bool`

HasBasearch returns a boolean if a field has been set.

### GetLatestOnly

`func (o *UpdatesRequest) GetLatestOnly() bool`

GetLatestOnly returns the LatestOnly field if non-nil, zero value otherwise.

### GetLatestOnlyOk

`func (o *UpdatesRequest) GetLatestOnlyOk() (*bool, bool)`

GetLatestOnlyOk returns a tuple with the LatestOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestOnly

`func (o *UpdatesRequest) SetLatestOnly(v bool)`

SetLatestOnly sets LatestOnly field to given value.

### HasLatestOnly

`func (o *UpdatesRequest) HasLatestOnly() bool`

HasLatestOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


