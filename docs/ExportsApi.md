# \ExportsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExports**](ExportsApi.md#ListExports) | **Get** /data/v1/exports | List property video view export links
[**ListExportsViews**](ExportsApi.md#ListExportsViews) | **Get** /data/v1/exports/views | List available property view exports



## ListExports

> ListExportsResponse ListExports(ctx).Execute()

List property video view export links



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ListExports(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ListExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExports`: ListExportsResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ListExports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListExportsRequest struct via the builder pattern


### Return type

[**ListExportsResponse**](ListExportsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExportsViews

> ListVideoViewExportsResponse ListExportsViews(ctx).Execute()

List available property view exports



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ListExportsViews(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ListExportsViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExportsViews`: ListVideoViewExportsResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ListExportsViews`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListExportsViewsRequest struct via the builder pattern


### Return type

[**ListVideoViewExportsResponse**](ListVideoViewExportsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

