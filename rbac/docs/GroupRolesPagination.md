# GroupRolesPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Data** | [**[]RoleOut**](RoleOut.md) |  | 

## Methods

### NewGroupRolesPagination

`func NewGroupRolesPagination(data []RoleOut, ) *GroupRolesPagination`

NewGroupRolesPagination instantiates a new GroupRolesPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRolesPaginationWithDefaults

`func NewGroupRolesPaginationWithDefaults() *GroupRolesPagination`

NewGroupRolesPaginationWithDefaults instantiates a new GroupRolesPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GroupRolesPagination) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GroupRolesPagination) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GroupRolesPagination) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GroupRolesPagination) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *GroupRolesPagination) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GroupRolesPagination) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GroupRolesPagination) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GroupRolesPagination) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *GroupRolesPagination) GetData() []RoleOut`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupRolesPagination) GetDataOk() (*[]RoleOut, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupRolesPagination) SetData(v []RoleOut)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


