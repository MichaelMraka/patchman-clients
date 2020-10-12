# \AccessApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrincipalAccess**](AccessApi.md#GetPrincipalAccess) | **Get** /access/ | Get the permitted access for a principal in the tenant (defaults to principal from the identity header)



## GetPrincipalAccess

> AccessPagination GetPrincipalAccess(ctx).Application(application).Username(username).Limit(limit).Offset(offset).Execute()

Get the permitted access for a principal in the tenant (defaults to principal from the identity header)



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
    application := "application_example" // string | The application name(s) to obtain access for the principal. This is an exact match. When no application is supplied, all permissions for the principal are returned. You may also use a comma-separated list to match on multiple applications.
    username := "username_example" // string | Unique username of the principal to obtain access for (only available for admins, and if supplied, takes precedence over the identity header). (optional)
    limit := 987 // int32 | Parameter for selecting the amount of data returned. (optional) (default to 10)
    offset := 987 // int32 | Parameter for selecting the offset of data. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessApi.GetPrincipalAccess(context.Background()).Application(application).Username(username).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetPrincipalAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrincipalAccess`: AccessPagination
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetPrincipalAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string** | The application name(s) to obtain access for the principal. This is an exact match. When no application is supplied, all permissions for the principal are returned. You may also use a comma-separated list to match on multiple applications. | 
 **username** | **string** | Unique username of the principal to obtain access for (only available for admins, and if supplied, takes precedence over the identity header). | 
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

