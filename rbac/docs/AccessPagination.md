# AccessPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Data** | [**[]Access**](Access.md) |  | 

## Methods

### NewAccessPagination

`func NewAccessPagination(data []Access, ) *AccessPagination`

NewAccessPagination instantiates a new AccessPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPaginationWithDefaults

`func NewAccessPaginationWithDefaults() *AccessPagination`

NewAccessPaginationWithDefaults instantiates a new AccessPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AccessPagination) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AccessPagination) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AccessPagination) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AccessPagination) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *AccessPagination) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccessPagination) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccessPagination) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AccessPagination) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *AccessPagination) GetData() []Access`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccessPagination) GetDataOk() (*[]Access, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccessPagination) SetData(v []Access)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


