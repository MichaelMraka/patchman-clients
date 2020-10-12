# \PolicyApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicies**](PolicyApi.md#CreatePolicies) | **Post** /policies/ | Create a policy in a tenant
[**DeletePolicy**](PolicyApi.md#DeletePolicy) | **Delete** /policies/{uuid}/ | Delete a policy in the tenant
[**GetPolicy**](PolicyApi.md#GetPolicy) | **Get** /policies/{uuid}/ | Get a policy in the tenant
[**ListPolicies**](PolicyApi.md#ListPolicies) | **Get** /policies/ | List the policies in the tenant
[**UpdatePolicy**](PolicyApi.md#UpdatePolicy) | **Put** /policies/{uuid}/ | Update a policy in the tenant



## CreatePolicies

> PolicyExtended CreatePolicies(ctx).PolicyIn(policyIn).Execute()

Create a policy in a tenant

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
    policyIn := openapiclient.PolicyIn{Name: "Name_example", Description: "Description_example", Group: "Group_example", Roles: []string{"Roles_example")} // PolicyIn | Policy to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.CreatePolicies(context.Background()).PolicyIn(policyIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.CreatePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicies`: PolicyExtended
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.CreatePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyIn** | [**PolicyIn**](PolicyIn.md) | Policy to create | 

### Return type

[**PolicyExtended**](PolicyExtended.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy(ctx, uuid).Execute()

Delete a policy in the tenant

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
    uuid := TODO // string | ID of policy to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.DeletePolicy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of policy to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


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


## GetPolicy

> PolicyExtended GetPolicy(ctx, uuid).Execute()

Get a policy in the tenant

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
    uuid := TODO // string | ID of policy to get

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.GetPolicy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: PolicyExtended
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of policy to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyExtended**](PolicyExtended.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> PolicyPagination ListPolicies(ctx).Limit(limit).Offset(offset).Name(name).Scope(scope).GroupName(groupName).GroupUuid(groupUuid).OrderBy(orderBy).Execute()

List the policies in the tenant



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
    scope := "scope_example" // string | Parameter for filtering resource by scope. (optional) (default to "account")
    groupName := "groupName_example" // string | Parameter for filtering resource by group name using string contains search. (optional)
    groupUuid := TODO // string | Parameter for filtering resource by group uuid using UUID exact match. (optional)
    orderBy := "orderBy_example" // string | Parameter for ordering resource by value. For inverse ordering, supply '-' before the param value, such as: ?order_by=-name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.ListPolicies(context.Background()).Limit(limit).Offset(offset).Name(name).Scope(scope).GroupName(groupName).GroupUuid(groupUuid).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ListPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicies`: PolicyPagination
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **name** | **string** | Parameter for filtering resource by name using string contains search. | 
 **scope** | **string** | Parameter for filtering resource by scope. | [default to &quot;account&quot;]
 **groupName** | **string** | Parameter for filtering resource by group name using string contains search. | 
 **groupUuid** | [**string**](.md) | Parameter for filtering resource by group uuid using UUID exact match. | 
 **orderBy** | **string** | Parameter for ordering resource by value. For inverse ordering, supply &#39;-&#39; before the param value, such as: ?order_by&#x3D;-name | 

### Return type

[**PolicyPagination**](PolicyPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> PolicyExtended UpdatePolicy(ctx, uuid).PolicyIn(policyIn).Execute()

Update a policy in the tenant

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
    uuid := TODO // string | ID of policy to update
    policyIn := openapiclient.PolicyIn{Name: "Name_example", Description: "Description_example", Group: "Group_example", Roles: []string{"Roles_example")} // PolicyIn | Policy to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PolicyApi.UpdatePolicy(context.Background(), uuid).PolicyIn(policyIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.UpdatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicy`: PolicyExtended
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of policy to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyIn** | [**PolicyIn**](PolicyIn.md) | Policy to update | 

### Return type

[**PolicyExtended**](PolicyExtended.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

