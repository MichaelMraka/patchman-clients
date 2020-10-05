# Access

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | **string** |  | 
**ResourceDefinitions** | [**[]ResourceDefinition**](ResourceDefinition.md) |  | 

## Methods

### NewAccess

`func NewAccess(permission string, resourceDefinitions []ResourceDefinition, ) *Access`

NewAccess instantiates a new Access object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessWithDefaults

`func NewAccessWithDefaults() *Access`

NewAccessWithDefaults instantiates a new Access object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *Access) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Access) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Access) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetResourceDefinitions

`func (o *Access) GetResourceDefinitions() []ResourceDefinition`

GetResourceDefinitions returns the ResourceDefinitions field if non-nil, zero value otherwise.

### GetResourceDefinitionsOk

`func (o *Access) GetResourceDefinitionsOk() (*[]ResourceDefinition, bool)`

GetResourceDefinitionsOk returns a tuple with the ResourceDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDefinitions

`func (o *Access) SetResourceDefinitions(v []ResourceDefinition)`

SetResourceDefinitions sets ResourceDefinitions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


