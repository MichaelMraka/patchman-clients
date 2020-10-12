# \PrincipalApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPrincipals**](PrincipalApi.md#ListPrincipals) | **Get** /principals/ | List the principals for a tenant



## ListPrincipals

> PrincipalPagination ListPrincipals(ctx).Limit(limit).Offset(offset).Usernames(usernames).SortOrder(sortOrder).Email(email).Status(status).AdminOnly(adminOnly).Execute()

List the principals for a tenant



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
    usernames := "usernames_example" // string | Usernames of principals to get (optional)
    sortOrder := "sortOrder_example" // string | The sort order of the query, either ascending or descending (optional)
    email := "email_example" // string | Exact e-mail address of principal to search for (optional)
    status := "status_example" // string | Set the status of users to get back. Could not be used with: usernames, email, admin_only (optional) (default to "enabled")
    adminOnly := "adminOnly_example" // string | Get only admin users within an account. Setting this would ignore the parameters: usernames, email (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrincipalApi.ListPrincipals(context.Background()).Limit(limit).Offset(offset).Usernames(usernames).SortOrder(sortOrder).Email(email).Status(status).AdminOnly(adminOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalApi.ListPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrincipals`: PrincipalPagination
    fmt.Fprintf(os.Stdout, "Response from `PrincipalApi.ListPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **usernames** | **string** | Usernames of principals to get | 
 **sortOrder** | **string** | The sort order of the query, either ascending or descending | 
 **email** | **string** | Exact e-mail address of principal to search for | 
 **status** | **string** | Set the status of users to get back. Could not be used with: usernames, email, admin_only | [default to &quot;enabled&quot;]
 **adminOnly** | **string** | Get only admin users within an account. Setting this would ignore the parameters: usernames, email | [default to &quot;false&quot;]

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

