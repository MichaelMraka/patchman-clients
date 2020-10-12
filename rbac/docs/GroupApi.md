# \GroupApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrincipalToGroup**](GroupApi.md#AddPrincipalToGroup) | **Post** /groups/{uuid}/principals/ | Add a principal to a group in the tenant
[**AddRoleToGroup**](GroupApi.md#AddRoleToGroup) | **Post** /groups/{uuid}/roles/ | Add a role to a group in the tenant
[**CreateGroup**](GroupApi.md#CreateGroup) | **Post** /groups/ | Create a group in a tenant
[**DeleteGroup**](GroupApi.md#DeleteGroup) | **Delete** /groups/{uuid}/ | Delete a group in the tenant
[**DeletePrincipalFromGroup**](GroupApi.md#DeletePrincipalFromGroup) | **Delete** /groups/{uuid}/principals/ | Remove a principal from a group in the tenant
[**DeleteRoleFromGroup**](GroupApi.md#DeleteRoleFromGroup) | **Delete** /groups/{uuid}/roles/ | Remove a role from a group in the tenant
[**GetGroup**](GroupApi.md#GetGroup) | **Get** /groups/{uuid}/ | Get a group in the tenant
[**GetPrincipalsFromGroup**](GroupApi.md#GetPrincipalsFromGroup) | **Get** /groups/{uuid}/principals/ | Get a list of principals from a group in the tenant
[**ListGroups**](GroupApi.md#ListGroups) | **Get** /groups/ | List the groups for a tenant
[**ListRolesForGroup**](GroupApi.md#ListRolesForGroup) | **Get** /groups/{uuid}/roles/ | List the roles for a group in the tenant
[**UpdateGroup**](GroupApi.md#UpdateGroup) | **Put** /groups/{uuid}/ | Udate a group in the tenant



## AddPrincipalToGroup

> GroupWithPrincipalsAndRoles AddPrincipalToGroup(ctx, uuid).GroupPrincipalIn(groupPrincipalIn).Execute()

Add a principal to a group in the tenant

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
    uuid := TODO // string | ID of group to update
    groupPrincipalIn := openapiclient.GroupPrincipalIn{Principals: []PrincipalIn{openapiclient.PrincipalIn{Username: "Username_example"})} // GroupPrincipalIn | Principal to add to a group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.AddPrincipalToGroup(context.Background(), uuid).GroupPrincipalIn(groupPrincipalIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.AddPrincipalToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPrincipalToGroup`: GroupWithPrincipalsAndRoles
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.AddPrincipalToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of group to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPrincipalToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupPrincipalIn** | [**GroupPrincipalIn**](GroupPrincipalIn.md) | Principal to add to a group | 

### Return type

[**GroupWithPrincipalsAndRoles**](GroupWithPrincipalsAndRoles.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleToGroup

> InlineResponse200 AddRoleToGroup(ctx, uuid).GroupRoleIn(groupRoleIn).Execute()

Add a role to a group in the tenant

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
    uuid := TODO // string | ID of group to update
    groupRoleIn := openapiclient.GroupRoleIn{Roles: []string{"Roles_example")} // GroupRoleIn | Role to add to a group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.AddRoleToGroup(context.Background(), uuid).GroupRoleIn(groupRoleIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.AddRoleToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleToGroup`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.AddRoleToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of group to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRoleIn** | [**GroupRoleIn**](GroupRoleIn.md) | Role to add to a group | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> GroupOut CreateGroup(ctx).Group(group).Execute()

Create a group in a tenant

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
    group := openapiclient.Group{Name: "Name_example", Description: "Description_example"} // Group | Group to create in tenant

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.CreateGroup(context.Background()).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: GroupOut
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**Group**](Group.md) | Group to create in tenant | 

### Return type

[**GroupOut**](GroupOut.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, uuid).Execute()

Delete a group in the tenant

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
    uuid := TODO // string | ID of group to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.DeleteGroup(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of group to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrincipalFromGroup

> DeletePrincipalFromGroup(ctx, uuid).Usernames(usernames).Execute()

Remove a principal from a group in the tenant

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
    uuid := TODO // string | ID of group to update
    usernames := "usernames_example" // string | A comma separated list of usernames for principals to remove from the group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.DeletePrincipalFromGroup(context.Background(), uuid).Usernames(usernames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeletePrincipalFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of group to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrincipalFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usernames** | **string** | A comma separated list of usernames for principals to remove from the group | 

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


## DeleteRoleFromGroup

> DeleteRoleFromGroup(ctx, uuid).Roles(roles).Execute()

Remove a role from a group in the tenant

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
    uuid := TODO // string | ID of group to update
    roles := "roles_example" // string | A comma separated list of role UUIDs for roles to remove from the group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.DeleteRoleFromGroup(context.Background(), uuid).Roles(roles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteRoleFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of group to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roles** | **string** | A comma separated list of role UUIDs for roles to remove from the group | 

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


## GetGroup

> GroupWithPrincipalsAndRoles GetGroup(ctx, uuid).Execute()

Get a group in the tenant

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
    uuid := TODO // string | ID of group to get

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.GetGroup(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: GroupWithPrincipalsAndRoles
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of group to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupWithPrincipalsAndRoles**](GroupWithPrincipalsAndRoles.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalsFromGroup

> PrincipalPagination GetPrincipalsFromGroup(ctx, uuid).PrincipalUsername(principalUsername).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()

Get a list of principals from a group in the tenant



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
    uuid := TODO // string | ID of group from which to get principals
    principalUsername := "principalUsername_example" // string | Parameter for filtering group principals by principal `username` using string contains search. (optional)
    limit := 987 // int32 | Parameter for selecting the amount of data returned. (optional) (default to 10)
    offset := 987 // int32 | Parameter for selecting the offset of data. (optional) (default to 0)
    orderBy := "orderBy_example" // string | Parameter for ordering principals by value. For inverse ordering, supply '-' before the param value, such as: ?order_by=-username (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.GetPrincipalsFromGroup(context.Background(), uuid).PrincipalUsername(principalUsername).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetPrincipalsFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrincipalsFromGroup`: PrincipalPagination
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetPrincipalsFromGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of group from which to get principals | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalsFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **principalUsername** | **string** | Parameter for filtering group principals by principal &#x60;username&#x60; using string contains search. | 
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **orderBy** | **string** | Parameter for ordering principals by value. For inverse ordering, supply &#39;-&#39; before the param value, such as: ?order_by&#x3D;-username | 

### Return type

[**PrincipalPagination**](PrincipalPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> GroupPagination ListGroups(ctx).Limit(limit).Offset(offset).Name(name).NameMatch(nameMatch).Scope(scope).Username(username).Uuid(uuid).RoleNames(roleNames).RoleDiscriminator(roleDiscriminator).OrderBy(orderBy).Execute()

List the groups for a tenant



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
    username := "username_example" // string | A username for a principal to filter for groups (optional)
    uuid := []string{"Inner_example"} // []string | A list of UUIDs to filter listed groups. (optional)
    roleNames := []string{"Inner_example"} // []string | List of role name to filter for groups. It is exact match but case-insensitive (optional)
    roleDiscriminator := "roleDiscriminator_example" // string | Discriminator that works with role_names to indicate matching all/any of the role names (optional)
    orderBy := "orderBy_example" // string | Parameter for ordering resource by value. For inverse ordering, supply '-' before the param value, such as: ?order_by=-name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.ListGroups(context.Background()).Limit(limit).Offset(offset).Name(name).NameMatch(nameMatch).Scope(scope).Username(username).Uuid(uuid).RoleNames(roleNames).RoleDiscriminator(roleDiscriminator).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.ListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroups`: GroupPagination
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.ListGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **name** | **string** | Parameter for filtering resource by name using string contains search. | 
 **nameMatch** | **string** | Parameter for specifying the matching criteria for an object&#39;s name. | 
 **scope** | **string** | Parameter for filtering resource by scope. | [default to &quot;account&quot;]
 **username** | **string** | A username for a principal to filter for groups | 
 **uuid** | [**[]string**](string.md) | A list of UUIDs to filter listed groups. | 
 **roleNames** | [**[]string**](string.md) | List of role name to filter for groups. It is exact match but case-insensitive | 
 **roleDiscriminator** | **string** | Discriminator that works with role_names to indicate matching all/any of the role names | 
 **orderBy** | **string** | Parameter for ordering resource by value. For inverse ordering, supply &#39;-&#39; before the param value, such as: ?order_by&#x3D;-name | 

### Return type

[**GroupPagination**](GroupPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRolesForGroup

> GroupRolesPagination ListRolesForGroup(ctx, uuid).Exclude(exclude).RoleName(roleName).RoleDescription(roleDescription).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()

List the roles for a group in the tenant



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
    uuid := TODO // string | ID of group
    exclude := true // bool | If this is set to true, the result would be roles excluding the ones in the group (optional) (default to false)
    roleName := "roleName_example" // string | Parameter for filtering group roles by role `name` using string contains search. (optional)
    roleDescription := "roleDescription_example" // string | Parameter for filtering group roles by role `description` using string contains search. (optional)
    limit := 987 // int32 | Parameter for selecting the amount of data returned. (optional) (default to 10)
    offset := 987 // int32 | Parameter for selecting the offset of data. (optional) (default to 0)
    orderBy := "orderBy_example" // string | Parameter for ordering resource by value. For inverse ordering, supply '-' before the param value, such as: ?order_by=-name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.ListRolesForGroup(context.Background(), uuid).Exclude(exclude).RoleName(roleName).RoleDescription(roleDescription).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.ListRolesForGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRolesForGroup`: GroupRolesPagination
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.ListRolesForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRolesForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exclude** | **bool** | If this is set to true, the result would be roles excluding the ones in the group | [default to false]
 **roleName** | **string** | Parameter for filtering group roles by role &#x60;name&#x60; using string contains search. | 
 **roleDescription** | **string** | Parameter for filtering group roles by role &#x60;description&#x60; using string contains search. | 
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **orderBy** | **string** | Parameter for ordering resource by value. For inverse ordering, supply &#39;-&#39; before the param value, such as: ?order_by&#x3D;-name | 

### Return type

[**GroupRolesPagination**](GroupRolesPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> GroupOut UpdateGroup(ctx, uuid).Group(group).Execute()

Udate a group in the tenant

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
    uuid := TODO // string | ID of group to update
    group := openapiclient.Group{Name: "Name_example", Description: "Description_example"} // Group | Group to update in tenant

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.UpdateGroup(context.Background(), uuid).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: GroupOut
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of group to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**Group**](Group.md) | Group to update in tenant | 

### Return type

[**GroupOut**](GroupOut.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

