# PackagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageList** | Pointer to [**map[string]PackagesResponsePackageList**](PackagesResponse_package_list.md) |  | [optional] 

## Methods

### NewPackagesResponse

`func NewPackagesResponse() *PackagesResponse`

NewPackagesResponse instantiates a new PackagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesResponseWithDefaults

`func NewPackagesResponseWithDefaults() *PackagesResponse`

NewPackagesResponseWithDefaults instantiates a new PackagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageList

`func (o *PackagesResponse) GetPackageList() map[string]PackagesResponsePackageList`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *PackagesResponse) GetPackageListOk() (*map[string]PackagesResponsePackageList, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *PackagesResponse) SetPackageList(v map[string]PackagesResponsePackageList)`

SetPackageList sets PackageList field to given value.

### HasPackageList

`func (o *PackagesResponse) HasPackageList() bool`

HasPackageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


