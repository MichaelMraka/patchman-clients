# PackagesResponsePackageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SourcePackage** | Pointer to **string** |  | [optional] 
**PackageList** | Pointer to **[]string** |  | [optional] 
**Repositories** | Pointer to [**[]PackagesResponseRepositories**](PackagesResponse_repositories.md) |  | [optional] 

## Methods

### NewPackagesResponsePackageList

`func NewPackagesResponsePackageList() *PackagesResponsePackageList`

NewPackagesResponsePackageList instantiates a new PackagesResponsePackageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesResponsePackageListWithDefaults

`func NewPackagesResponsePackageListWithDefaults() *PackagesResponsePackageList`

NewPackagesResponsePackageListWithDefaults instantiates a new PackagesResponsePackageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *PackagesResponsePackageList) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PackagesResponsePackageList) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PackagesResponsePackageList) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PackagesResponsePackageList) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *PackagesResponsePackageList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackagesResponsePackageList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackagesResponsePackageList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackagesResponsePackageList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourcePackage

`func (o *PackagesResponsePackageList) GetSourcePackage() string`

GetSourcePackage returns the SourcePackage field if non-nil, zero value otherwise.

### GetSourcePackageOk

`func (o *PackagesResponsePackageList) GetSourcePackageOk() (*string, bool)`

GetSourcePackageOk returns a tuple with the SourcePackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePackage

`func (o *PackagesResponsePackageList) SetSourcePackage(v string)`

SetSourcePackage sets SourcePackage field to given value.

### HasSourcePackage

`func (o *PackagesResponsePackageList) HasSourcePackage() bool`

HasSourcePackage returns a boolean if a field has been set.

### GetPackageList

`func (o *PackagesResponsePackageList) GetPackageList() []string`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *PackagesResponsePackageList) GetPackageListOk() (*[]string, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *PackagesResponsePackageList) SetPackageList(v []string)`

SetPackageList sets PackageList field to given value.

### HasPackageList

`func (o *PackagesResponsePackageList) HasPackageList() bool`

HasPackageList returns a boolean if a field has been set.

### GetRepositories

`func (o *PackagesResponsePackageList) GetRepositories() []PackagesResponseRepositories`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *PackagesResponsePackageList) GetRepositoriesOk() (*[]PackagesResponseRepositories, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *PackagesResponsePackageList) SetRepositories(v []PackagesResponseRepositories)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *PackagesResponsePackageList) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


