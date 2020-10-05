# VulnerabilitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageList** | **[]string** |  | 
**RepositoryList** | Pointer to **[]string** |  | [optional] 
**ModulesList** | Pointer to [**[]UpdatesRequestModulesList**](UpdatesRequest_modules_list.md) |  | [optional] 
**Releasever** | Pointer to **string** |  | [optional] 
**Basearch** | Pointer to **string** |  | [optional] 

## Methods

### NewVulnerabilitiesRequest

`func NewVulnerabilitiesRequest(packageList []string, ) *VulnerabilitiesRequest`

NewVulnerabilitiesRequest instantiates a new VulnerabilitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnerabilitiesRequestWithDefaults

`func NewVulnerabilitiesRequestWithDefaults() *VulnerabilitiesRequest`

NewVulnerabilitiesRequestWithDefaults instantiates a new VulnerabilitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageList

`func (o *VulnerabilitiesRequest) GetPackageList() []string`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *VulnerabilitiesRequest) GetPackageListOk() (*[]string, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *VulnerabilitiesRequest) SetPackageList(v []string)`

SetPackageList sets PackageList field to given value.


### GetRepositoryList

`func (o *VulnerabilitiesRequest) GetRepositoryList() []string`

GetRepositoryList returns the RepositoryList field if non-nil, zero value otherwise.

### GetRepositoryListOk

`func (o *VulnerabilitiesRequest) GetRepositoryListOk() (*[]string, bool)`

GetRepositoryListOk returns a tuple with the RepositoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryList

`func (o *VulnerabilitiesRequest) SetRepositoryList(v []string)`

SetRepositoryList sets RepositoryList field to given value.

### HasRepositoryList

`func (o *VulnerabilitiesRequest) HasRepositoryList() bool`

HasRepositoryList returns a boolean if a field has been set.

### GetModulesList

`func (o *VulnerabilitiesRequest) GetModulesList() []UpdatesRequestModulesList`

GetModulesList returns the ModulesList field if non-nil, zero value otherwise.

### GetModulesListOk

`func (o *VulnerabilitiesRequest) GetModulesListOk() (*[]UpdatesRequestModulesList, bool)`

GetModulesListOk returns a tuple with the ModulesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesList

`func (o *VulnerabilitiesRequest) SetModulesList(v []UpdatesRequestModulesList)`

SetModulesList sets ModulesList field to given value.

### HasModulesList

`func (o *VulnerabilitiesRequest) HasModulesList() bool`

HasModulesList returns a boolean if a field has been set.

### GetReleasever

`func (o *VulnerabilitiesRequest) GetReleasever() string`

GetReleasever returns the Releasever field if non-nil, zero value otherwise.

### GetReleaseverOk

`func (o *VulnerabilitiesRequest) GetReleaseverOk() (*string, bool)`

GetReleaseverOk returns a tuple with the Releasever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasever

`func (o *VulnerabilitiesRequest) SetReleasever(v string)`

SetReleasever sets Releasever field to given value.

### HasReleasever

`func (o *VulnerabilitiesRequest) HasReleasever() bool`

HasReleasever returns a boolean if a field has been set.

### GetBasearch

`func (o *VulnerabilitiesRequest) GetBasearch() string`

GetBasearch returns the Basearch field if non-nil, zero value otherwise.

### GetBasearchOk

`func (o *VulnerabilitiesRequest) GetBasearchOk() (*string, bool)`

GetBasearchOk returns a tuple with the Basearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasearch

`func (o *VulnerabilitiesRequest) SetBasearch(v string)`

SetBasearch sets Basearch field to given value.

### HasBasearch

`func (o *VulnerabilitiesRequest) HasBasearch() bool`

HasBasearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


