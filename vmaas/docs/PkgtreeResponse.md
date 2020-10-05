# PkgtreeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **interface{}** |  | [optional] 
**PackageNameList** | Pointer to [**map[string][]map[string]interface{}**](array.md) |  | [optional] 

## Methods

### NewPkgtreeResponse

`func NewPkgtreeResponse() *PkgtreeResponse`

NewPkgtreeResponse instantiates a new PkgtreeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgtreeResponseWithDefaults

`func NewPkgtreeResponseWithDefaults() *PkgtreeResponse`

NewPkgtreeResponseWithDefaults instantiates a new PkgtreeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *PkgtreeResponse) GetLastChange() interface{}`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *PkgtreeResponse) GetLastChangeOk() (*interface{}, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *PkgtreeResponse) SetLastChange(v interface{})`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *PkgtreeResponse) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *PkgtreeResponse) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *PkgtreeResponse) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetPackageNameList

`func (o *PkgtreeResponse) GetPackageNameList() map[string][]map[string]interface{}`

GetPackageNameList returns the PackageNameList field if non-nil, zero value otherwise.

### GetPackageNameListOk

`func (o *PkgtreeResponse) GetPackageNameListOk() (*map[string][]map[string]interface{}, bool)`

GetPackageNameListOk returns a tuple with the PackageNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageNameList

`func (o *PkgtreeResponse) SetPackageNameList(v map[string][]map[string]interface{})`

SetPackageNameList sets PackageNameList field to given value.

### HasPackageNameList

`func (o *PkgtreeResponse) HasPackageNameList() bool`

HasPackageNameList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


