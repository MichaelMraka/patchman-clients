# \PermissionApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPermissionOptions**](PermissionApi.md#ListPermissionOptions) | **Get** /permissions/options/ | List the available options for fields of permissions for a tenant
[**ListPermissions**](PermissionApi.md#ListPermissions) | **Get** /permissions/ | List the permissions for a tenant



## ListPermissionOptions

> PermissionOptionsPagination ListPermissionOptions(ctx).Field(field).Limit(limit).Offset(offset).Application(application).ResourceType(resourceType).Verb(verb).Execute()

List the available options for fields of permissions for a tenant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    field := "field_example" // string | specify which fields of permission to display
    limit := 987 // int32 | Parameter for selecting the amount of data returned. (optional) (default to 10)
    offset := 987 // int32 | Parameter for selecting the offset of data. (optional) (default to 0)
    application := "application_example" // string | Filter returned options based on application. You may also use a comma-separated list to filter on multiple applications. (optional)
    resourceType := "resourceType_example" // string | Filter returned options based on resource_type. You may also use a comma-separated list to filter on multiple resource_types. (optional)
    verb := "verb_example" // string | Filter returned options based on verb. You may also use a comma-separated list to filter on multiple verbs. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionApi.ListPermissionOptions(context.Background()).Field(field).Limit(limit).Offset(offset).Application(application).ResourceType(resourceType).Verb(verb).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.ListPermissionOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermissionOptions`: PermissionOptionsPagination
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.ListPermissionOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | **string** | specify which fields of permission to display | 
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **application** | **string** | Filter returned options based on application. You may also use a comma-separated list to filter on multiple applications. | 
 **resourceType** | **string** | Filter returned options based on resource_type. You may also use a comma-separated list to filter on multiple resource_types. | 
 **verb** | **string** | Filter returned options based on verb. You may also use a comma-separated list to filter on multiple verbs. | 

### Return type

[**PermissionOptionsPagination**](PermissionOptionsPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissions

> PermissionPagination ListPermissions(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).Application(application).ResourceType(resourceType).Verb(verb).Permission(permission).ExcludeGlobals(excludeGlobals).Execute()

List the permissions for a tenant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := 987 // int32 | Parameter for selecting the amount of data returned. (optional) (default to 10)
    offset := 987 // int32 | Parameter for selecting the offset of data. (optional) (default to 0)
    orderBy := "orderBy_example" // string | Parameter for ordering resource by value. For inverse ordering, supply '-' before the param value, such as: ?order_by=-name (optional)
    application := "application_example" // string | Exact match for the application name of a permission. You may also use a comma-separated list to match on multiple applications. (optional)
    resourceType := "resourceType_example" // string | Exact match for the resource type name of a permission. You may also use a comma-separated list to match on multiple resource_types. (optional)
    verb := "verb_example" // string | Exact match for the operation verb name of a permission You may also use a comma-separated list to match on multiple verbs. (optional)
    permission := "permission_example" // string | Partial match for the aggregate permission value name of a permission object. (optional)
    excludeGlobals := "excludeGlobals_example" // string | If set to 'true', this will exclude any permission with a global allowance on either 'application', 'resource_type' or 'verb'. The default is 'false'. (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionApi.ListPermissions(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).Application(application).ResourceType(resourceType).Verb(verb).Permission(permission).ExcludeGlobals(excludeGlobals).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.ListPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermissions`: PermissionPagination
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.ListPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **orderBy** | **string** | Parameter for ordering resource by value. For inverse ordering, supply &#39;-&#39; before the param value, such as: ?order_by&#x3D;-name | 
 **application** | **string** | Exact match for the application name of a permission. You may also use a comma-separated list to match on multiple applications. | 
 **resourceType** | **string** | Exact match for the resource type name of a permission. You may also use a comma-separated list to match on multiple resource_types. | 
 **verb** | **string** | Exact match for the operation verb name of a permission You may also use a comma-separated list to match on multiple verbs. | 
 **permission** | **string** | Partial match for the aggregate permission value name of a permission object. | 
 **excludeGlobals** | **string** | If set to &#39;true&#39;, this will exclude any permission with a global allowance on either &#39;application&#39;, &#39;resource_type&#39; or &#39;verb&#39;. The default is &#39;false&#39;. | [default to &quot;false&quot;]

### Return type

[**PermissionPagination**](PermissionPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

