# \SapSystemApi

All URIs are relative to *http://localhost/api/inventory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSystemProfileGetSapSids**](SapSystemApi.md#ApiSystemProfileGetSapSids) | **Get** /system_profile/sap_sids | get sap system values
[**ApiSystemProfileGetSapSystem**](SapSystemApi.md#ApiSystemProfileGetSapSystem) | **Get** /system_profile/sap_system | get sap system values



## ApiSystemProfileGetSapSids

> SystemProfileSapSystemOut ApiSystemProfileGetSapSids(ctx).Tags(tags).PerPage(perPage).Page(page).Staleness(staleness).RegisteredWith(registeredWith).Filter(filter).Execute()

get sap system values



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
    tags := []string{"Inner_example"} // []string | filters out hosts not tagged by the given tags (optional)
    perPage := 987 // int32 | A number of items to return per page. (optional) (default to 50)
    page := 987 // int32 | A page number of the items to return. (optional) (default to 1)
    staleness := []string{"Staleness_example"} // []string | Culling states of the hosts. Default: fresh,stale,unknown (optional) (default to ["fresh","stale","unknown"])
    registeredWith := "registeredWith_example" // string | Filters out any host not registered with the specified service (optional)
    filter := TODO // map[string]interface{} | Filters hosts based on system_profile fields (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SapSystemApi.ApiSystemProfileGetSapSids(context.Background()).Tags(tags).PerPage(perPage).Page(page).Staleness(staleness).RegisteredWith(registeredWith).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SapSystemApi.ApiSystemProfileGetSapSids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSystemProfileGetSapSids`: SystemProfileSapSystemOut
    fmt.Fprintf(os.Stdout, "Response from `SapSystemApi.ApiSystemProfileGetSapSids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSystemProfileGetSapSidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**[]string**](string.md) | filters out hosts not tagged by the given tags | 
 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **staleness** | [**[]string**](string.md) | Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **registeredWith** | **string** | Filters out any host not registered with the specified service | 
 **filter** | [**map[string]interface{}**](.md) | Filters hosts based on system_profile fields | 

### Return type

[**SystemProfileSapSystemOut**](SystemProfileSapSystemOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSystemProfileGetSapSystem

> SystemProfileSapSystemOut ApiSystemProfileGetSapSystem(ctx).Tags(tags).PerPage(perPage).Page(page).Staleness(staleness).RegisteredWith(registeredWith).Filter(filter).Execute()

get sap system values



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
    tags := []string{"Inner_example"} // []string | filters out hosts not tagged by the given tags (optional)
    perPage := 987 // int32 | A number of items to return per page. (optional) (default to 50)
    page := 987 // int32 | A page number of the items to return. (optional) (default to 1)
    staleness := []string{"Staleness_example"} // []string | Culling states of the hosts. Default: fresh,stale,unknown (optional) (default to ["fresh","stale","unknown"])
    registeredWith := "registeredWith_example" // string | Filters out any host not registered with the specified service (optional)
    filter := TODO // map[string]interface{} | Filters hosts based on system_profile fields (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SapSystemApi.ApiSystemProfileGetSapSystem(context.Background()).Tags(tags).PerPage(perPage).Page(page).Staleness(staleness).RegisteredWith(registeredWith).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SapSystemApi.ApiSystemProfileGetSapSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSystemProfileGetSapSystem`: SystemProfileSapSystemOut
    fmt.Fprintf(os.Stdout, "Response from `SapSystemApi.ApiSystemProfileGetSapSystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSystemProfileGetSapSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**[]string**](string.md) | filters out hosts not tagged by the given tags | 
 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **staleness** | [**[]string**](string.md) | Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **registeredWith** | **string** | Filters out any host not registered with the specified service | 
 **filter** | [**map[string]interface{}**](.md) | Filters hosts based on system_profile fields | 

### Return type

[**SystemProfileSapSystemOut**](SystemProfileSapSystemOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

