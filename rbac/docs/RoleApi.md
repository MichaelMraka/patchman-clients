# \RoleApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoles**](RoleApi.md#CreateRoles) | **Post** /roles/ | Create a roles for a tenant
[**DeleteRole**](RoleApi.md#DeleteRole) | **Delete** /roles/{uuid}/ | Delete a role in the tenant
[**GetRole**](RoleApi.md#GetRole) | **Get** /roles/{uuid}/ | Get a role in the tenant
[**GetRoleAccess**](RoleApi.md#GetRoleAccess) | **Get** /roles/{uuid}/access/ | Get access for a role in the tenant
[**ListRoles**](RoleApi.md#ListRoles) | **Get** /roles/ | List the roles for a tenant
[**UpdateRole**](RoleApi.md#UpdateRole) | **Put** /roles/{uuid}/ | Update a Role in the tenant



## CreateRoles

> RoleWithAccess CreateRoles(ctx).RoleIn(roleIn).Execute()

Create a roles for a tenant

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
    roleIn := openapiclient.RoleIn{Name: "Name_example", DisplayName: "DisplayName_example", Description: "Description_example", Access: []Access{openapiclient.Access{Permission: "Permission_example", ResourceDefinitions: []ResourceDefinition{openapiclient.ResourceDefinition{AttributeFilter: openapiclient.ResourceDefinitionFilter{Key: "Key_example", Operation: "Operation_example", Value: "Value_example"}})})} // RoleIn | Role to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleApi.CreateRoles(context.Background()).RoleIn(roleIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.CreateRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoles`: RoleWithAccess
    fmt.Fprintf(os.Stdout, "Response from `RoleApi.CreateRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleIn** | [**RoleIn**](RoleIn.md) | Role to create | 

### Return type

[**RoleWithAccess**](RoleWithAccess.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, uuid).Execute()

Delete a role in the tenant

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
    uuid := TODO // string | ID of role to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleApi.DeleteRole(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of role to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> RoleWithAccess GetRole(ctx, uuid).Scope(scope).Execute()

Get a role in the tenant

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
    uuid := TODO // string | ID of role to get
    scope := "scope_example" // string | Parameter for filtering resource by scope. (optional) (default to "account")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleApi.GetRole(context.Background(), uuid).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: RoleWithAccess
    fmt.Fprintf(os.Stdout, "Response from `RoleApi.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of role to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **string** | Parameter for filtering resource by scope. | [default to &quot;account&quot;]

### Return type

[**RoleWithAccess**](RoleWithAccess.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleAccess

> AccessPagination GetRoleAccess(ctx, uuid).Limit(limit).Offset(offset).Execute()

Get access for a role in the tenant

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
    uuid := TODO // string | ID of the role
    limit := 987 // int32 | Parameter for selecting the amount of data returned. (optional) (default to 10)
    offset := 987 // int32 | Parameter for selecting the offset of data. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleApi.GetRoleAccess(context.Background(), uuid).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.GetRoleAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleAccess`: AccessPagination
    fmt.Fprintf(os.Stdout, "Response from `RoleApi.GetRoleAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]

### Return type

[**AccessPagination**](AccessPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> RolePaginationDynamic ListRoles(ctx).Limit(limit).Offset(offset).Name(name).NameMatch(nameMatch).Scope(scope).OrderBy(orderBy).AddFields(addFields).Username(username).Application(application).Permission(permission).Execute()

List the roles for a tenant



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
    name := "name_example" // string | Parameter for filtering resource by name using string contains search. (optional)
    nameMatch := "nameMatch_example" // string | Parameter for specifying the matching criteria for an object's name. (optional)
    scope := "scope_example" // string | Parameter for filtering resource by scope. (optional) (default to "account")
    orderBy := "orderBy_example" // string | Parameter for ordering resource by value. For inverse ordering, supply '-' before the param value, such as: ?order_by=-name (optional)
    addFields := []string{"AddFields_example"} // []string | Parameter for add list of fields to display for roles. (optional)
    username := "username_example" // string | Unique username of the principal to obtain roles for (only available for admins, and if supplied, takes precedence over the identity header). (optional)
    application := "application_example" // string | The application name(s) to filter roles by, from permissions. This is an exact match. You may also use a comma-separated list to match on multiple applications. (optional)
    permission := "permission_example" // string | The permission(s) to filter roles by. This is an exact match. You may also use a comma-separated list to match on multiple permissions. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleApi.ListRoles(context.Background()).Limit(limit).Offset(offset).Name(name).NameMatch(nameMatch).Scope(scope).OrderBy(orderBy).AddFields(addFields).Username(username).Application(application).Permission(permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: RolePaginationDynamic
    fmt.Fprintf(os.Stdout, "Response from `RoleApi.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **name** | **string** | Parameter for filtering resource by name using string contains search. | 
 **nameMatch** | **string** | Parameter for specifying the matching criteria for an object&#39;s name. | 
 **scope** | **string** | Parameter for filtering resource by scope. | [default to &quot;account&quot;]
 **orderBy** | **string** | Parameter for ordering resource by value. For inverse ordering, supply &#39;-&#39; before the param value, such as: ?order_by&#x3D;-name | 
 **addFields** | [**[]string**](string.md) | Parameter for add list of fields to display for roles. | 
 **username** | **string** | Unique username of the principal to obtain roles for (only available for admins, and if supplied, takes precedence over the identity header). | 
 **application** | **string** | The application name(s) to filter roles by, from permissions. This is an exact match. You may also use a comma-separated list to match on multiple applications. | 
 **permission** | **string** | The permission(s) to filter roles by. This is an exact match. You may also use a comma-separated list to match on multiple permissions. | 

### Return type

[**RolePaginationDynamic**](RolePaginationDynamic.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> UpdateRole(ctx, uuid).RoleWithAccess(roleWithAccess).Execute()

Update a Role in the tenant

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
    uuid := TODO // string | ID of role to update
    roleWithAccess := openapiclient.RoleWithAccess{Name: "Name_example", DisplayName: "DisplayName_example", Description: "Description_example", Uuid: "Uuid_example", Created: time.Now(), Modified: time.Now(), PolicyCount: 123, AccessCount: 123, Applications: []string{"Applications_example"), System: false, PlatformDefault: false, Access: []Access{openapiclient.Access{Permission: "Permission_example", ResourceDefinitions: []ResourceDefinition{openapiclient.ResourceDefinition{AttributeFilter: openapiclient.ResourceDefinitionFilter{Key: "Key_example", Operation: "Operation_example", Value: "Value_example"}})})} // RoleWithAccess | Update to a Role

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleApi.UpdateRole(context.Background(), uuid).RoleWithAccess(roleWithAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApi.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of role to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleWithAccess** | [**RoleWithAccess**](RoleWithAccess.md) | Update to a Role | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

