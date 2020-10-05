# UpdatesV3Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageList** | **[]string** |  | 
**RepositoryList** | Pointer to **[]string** |  | [optional] 
**ModulesList** | Pointer to [**[]UpdatesRequestModulesList**](UpdatesRequest_modules_list.md) |  | [optional] 
**Releasever** | Pointer to **string** |  | [optional] 
**Basearch** | Pointer to **string** |  | [optional] 
**SecurityOnly** | Pointer to **bool** |  | [optional] 
**LatestOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdatesV3Request

`func NewUpdatesV3Request(packageList []string, ) *UpdatesV3Request`

NewUpdatesV3Request instantiates a new UpdatesV3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatesV3RequestWithDefaults

`func NewUpdatesV3RequestWithDefaults() *UpdatesV3Request`

NewUpdatesV3RequestWithDefaults instantiates a new UpdatesV3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageList

`func (o *UpdatesV3Request) GetPackageList() []string`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *UpdatesV3Request) GetPackageListOk() (*[]string, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *UpdatesV3Request) SetPackageList(v []string)`

SetPackageList sets PackageList field to given value.


### GetRepositoryList

`func (o *UpdatesV3Request) GetRepositoryList() []string`

GetRepositoryList returns the RepositoryList field if non-nil, zero value otherwise.

### GetRepositoryListOk

`func (o *UpdatesV3Request) GetRepositoryListOk() (*[]string, bool)`

GetRepositoryListOk returns a tuple with the RepositoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryList

`func (o *UpdatesV3Request) SetRepositoryList(v []string)`

SetRepositoryList sets RepositoryList field to given value.

### HasRepositoryList

`func (o *UpdatesV3Request) HasRepositoryList() bool`

HasRepositoryList returns a boolean if a field has been set.

### GetModulesList

`func (o *UpdatesV3Request) GetModulesList() []UpdatesRequestModulesList`

GetModulesList returns the ModulesList field if non-nil, zero value otherwise.

### GetModulesListOk

`func (o *UpdatesV3Request) GetModulesListOk() (*[]UpdatesRequestModulesList, bool)`

GetModulesListOk returns a tuple with the ModulesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesList

`func (o *UpdatesV3Request) SetModulesList(v []UpdatesRequestModulesList)`

SetModulesList sets ModulesList field to given value.

### HasModulesList

`func (o *UpdatesV3Request) HasModulesList() bool`

HasModulesList returns a boolean if a field has been set.

### GetReleasever

`func (o *UpdatesV3Request) GetReleasever() string`

GetReleasever returns the Releasever field if non-nil, zero value otherwise.

### GetReleaseverOk

`func (o *UpdatesV3Request) GetReleaseverOk() (*string, bool)`

GetReleaseverOk returns a tuple with the Releasever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasever

`func (o *UpdatesV3Request) SetReleasever(v string)`

SetReleasever sets Releasever field to given value.

### HasReleasever

`func (o *UpdatesV3Request) HasReleasever() bool`

HasReleasever returns a boolean if a field has been set.

### GetBasearch

`func (o *UpdatesV3Request) GetBasearch() string`

GetBasearch returns the Basearch field if non-nil, zero value otherwise.

### GetBasearchOk

`func (o *UpdatesV3Request) GetBasearchOk() (*string, bool)`

GetBasearchOk returns a tuple with the Basearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasearch

`func (o *UpdatesV3Request) SetBasearch(v string)`

SetBasearch sets Basearch field to given value.

### HasBasearch

`func (o *UpdatesV3Request) HasBasearch() bool`

HasBasearch returns a boolean if a field has been set.

### GetSecurityOnly

`func (o *UpdatesV3Request) GetSecurityOnly() bool`

GetSecurityOnly returns the SecurityOnly field if non-nil, zero value otherwise.

### GetSecurityOnlyOk

`func (o *UpdatesV3Request) GetSecurityOnlyOk() (*bool, bool)`

GetSecurityOnlyOk returns a tuple with the SecurityOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityOnly

`func (o *UpdatesV3Request) SetSecurityOnly(v bool)`

SetSecurityOnly sets SecurityOnly field to given value.

### HasSecurityOnly

`func (o *UpdatesV3Request) HasSecurityOnly() bool`

HasSecurityOnly returns a boolean if a field has been set.

### GetLatestOnly

`func (o *UpdatesV3Request) GetLatestOnly() bool`

GetLatestOnly returns the LatestOnly field if non-nil, zero value otherwise.

### GetLatestOnlyOk

`func (o *UpdatesV3Request) GetLatestOnlyOk() (*bool, bool)`

GetLatestOnlyOk returns a tuple with the LatestOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestOnly

`func (o *UpdatesV3Request) SetLatestOnly(v bool)`

SetLatestOnly sets LatestOnly field to given value.

### HasLatestOnly

`func (o *UpdatesV3Request) HasLatestOnly() bool`

HasLatestOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


