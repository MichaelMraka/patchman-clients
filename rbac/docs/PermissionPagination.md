# PermissionPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Data** | [**[]Permission**](Permission.md) |  | 

## Methods

### NewPermissionPagination

`func NewPermissionPagination(data []Permission, ) *PermissionPagination`

NewPermissionPagination instantiates a new PermissionPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionPaginationWithDefaults

`func NewPermissionPaginationWithDefaults() *PermissionPagination`

NewPermissionPaginationWithDefaults instantiates a new PermissionPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PermissionPagination) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PermissionPagination) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PermissionPagination) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PermissionPagination) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *PermissionPagination) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PermissionPagination) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PermissionPagination) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PermissionPagination) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *PermissionPagination) GetData() []Permission`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PermissionPagination) GetDataOk() (*[]Permission, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PermissionPagination) SetData(v []Permission)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


