# \ExperimentalApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppRPMPkgNamesHandlerGetGet**](ExperimentalApi.md#AppRPMPkgNamesHandlerGetGet) | **Get** /v1/package_names/rpms/{rpm} | 
[**AppRPMPkgNamesHandlerPostPost**](ExperimentalApi.md#AppRPMPkgNamesHandlerPostPost) | **Post** /v1/package_names/rpms | 
[**AppSRPMPkgNamesHandlerGetGet**](ExperimentalApi.md#AppSRPMPkgNamesHandlerGetGet) | **Get** /v1/package_names/srpms/{srpm} | 
[**AppSRPMPkgNamesHandlerPostPost**](ExperimentalApi.md#AppSRPMPkgNamesHandlerPostPost) | **Post** /v1/package_names/srpms | 



## AppRPMPkgNamesHandlerGetGet

> RPMPkgNamesResponse AppRPMPkgNamesHandlerGetGet(ctx, rpm).Execute()





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
    rpm := "rpm_example" // string | Package name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentalApi.AppRPMPkgNamesHandlerGetGet(context.Background(), rpm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalApi.AppRPMPkgNamesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRPMPkgNamesHandlerGetGet`: RPMPkgNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `ExperimentalApi.AppRPMPkgNamesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpm** | **string** | Package name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRPMPkgNamesHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RPMPkgNamesResponse**](RPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRPMPkgNamesHandlerPostPost

> RPMPkgNamesResponse AppRPMPkgNamesHandlerPostPost(ctx).RPMPkgNamesRequest(rPMPkgNamesRequest).Execute()





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
    rPMPkgNamesRequest := openapiclient.RPMPkgNamesRequest{RpmNameList: []string{"RpmNameList_example"), ContentSetList: []string{"ContentSetList_example")} // RPMPkgNamesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentalApi.AppRPMPkgNamesHandlerPostPost(context.Background()).RPMPkgNamesRequest(rPMPkgNamesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalApi.AppRPMPkgNamesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRPMPkgNamesHandlerPostPost`: RPMPkgNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `ExperimentalApi.AppRPMPkgNamesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppRPMPkgNamesHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rPMPkgNamesRequest** | [**RPMPkgNamesRequest**](RPMPkgNamesRequest.md) |  | 

### Return type

[**RPMPkgNamesResponse**](RPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSRPMPkgNamesHandlerGetGet

> SRPMPkgNamesResponse AppSRPMPkgNamesHandlerGetGet(ctx, srpm).Execute()





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
    srpm := "srpm_example" // string | Source package name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentalApi.AppSRPMPkgNamesHandlerGetGet(context.Background(), srpm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalApi.AppSRPMPkgNamesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSRPMPkgNamesHandlerGetGet`: SRPMPkgNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `ExperimentalApi.AppSRPMPkgNamesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srpm** | **string** | Source package name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSRPMPkgNamesHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SRPMPkgNamesResponse**](SRPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSRPMPkgNamesHandlerPostPost

> SRPMPkgNamesResponse AppSRPMPkgNamesHandlerPostPost(ctx).SRPMPkgNamesRequest(sRPMPkgNamesRequest).Execute()





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
    sRPMPkgNamesRequest := openapiclient.SRPMPkgNamesRequest{SrpmNameList: []string{"SrpmNameList_example"), ContentSetList: []string{"ContentSetList_example")} // SRPMPkgNamesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentalApi.AppSRPMPkgNamesHandlerPostPost(context.Background()).SRPMPkgNamesRequest(sRPMPkgNamesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalApi.AppSRPMPkgNamesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSRPMPkgNamesHandlerPostPost`: SRPMPkgNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `ExperimentalApi.AppSRPMPkgNamesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppSRPMPkgNamesHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sRPMPkgNamesRequest** | [**SRPMPkgNamesRequest**](SRPMPkgNamesRequest.md) |  | 

### Return type

[**SRPMPkgNamesResponse**](SRPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

