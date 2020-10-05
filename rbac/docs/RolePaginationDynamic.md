# RolePaginationDynamic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Data** | [**[]RoleOutDynamic**](RoleOutDynamic.md) |  | 

## Methods

### NewRolePaginationDynamic

`func NewRolePaginationDynamic(data []RoleOutDynamic, ) *RolePaginationDynamic`

NewRolePaginationDynamic instantiates a new RolePaginationDynamic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePaginationDynamicWithDefaults

`func NewRolePaginationDynamicWithDefaults() *RolePaginationDynamic`

NewRolePaginationDynamicWithDefaults instantiates a new RolePaginationDynamic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *RolePaginationDynamic) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RolePaginationDynamic) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RolePaginationDynamic) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RolePaginationDynamic) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *RolePaginationDynamic) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RolePaginationDynamic) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RolePaginationDynamic) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RolePaginationDynamic) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *RolePaginationDynamic) GetData() []RoleOutDynamic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RolePaginationDynamic) GetDataOk() (*[]RoleOutDynamic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RolePaginationDynamic) SetData(v []RoleOutDynamic)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


