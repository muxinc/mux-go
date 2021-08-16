# \DirectUploadsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDirectUpload**](DirectUploadsApi.md#CancelDirectUpload) | **Put** /video/v1/uploads/{UPLOAD_ID}/cancel | Cancel a direct upload
[**CreateDirectUpload**](DirectUploadsApi.md#CreateDirectUpload) | **Post** /video/v1/uploads | Create a new direct upload URL
[**GetDirectUpload**](DirectUploadsApi.md#GetDirectUpload) | **Get** /video/v1/uploads/{UPLOAD_ID} | Retrieve a single direct upload&#39;s info
[**ListDirectUploads**](DirectUploadsApi.md#ListDirectUploads) | **Get** /video/v1/uploads | List direct uploads



## CancelDirectUpload

> UploadResponse CancelDirectUpload(ctx, uPLOADID).Execute()

Cancel a direct upload



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
    uPLOADID := "abcd1234" // string | ID of the Upload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectUploadsApi.CancelDirectUpload(context.Background(), uPLOADID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectUploadsApi.CancelDirectUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelDirectUpload`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectUploadsApi.CancelDirectUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uPLOADID** | **string** | ID of the Upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDirectUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UploadResponse**](UploadResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDirectUpload

> UploadResponse CreateDirectUpload(ctx).CreateUploadRequest(createUploadRequest).Execute()

Create a new direct upload URL

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
    createUploadRequest := *openapiclient.NewCreateUploadRequest(*openapiclient.NewCreateAssetRequest()) // CreateUploadRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectUploadsApi.CreateDirectUpload(context.Background()).CreateUploadRequest(createUploadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectUploadsApi.CreateDirectUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDirectUpload`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectUploadsApi.CreateDirectUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUploadRequest** | [**CreateUploadRequest**](CreateUploadRequest.md) |  | 

### Return type

[**UploadResponse**](UploadResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectUpload

> UploadResponse GetDirectUpload(ctx, uPLOADID).Execute()

Retrieve a single direct upload's info

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
    uPLOADID := "abcd1234" // string | ID of the Upload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectUploadsApi.GetDirectUpload(context.Background(), uPLOADID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectUploadsApi.GetDirectUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectUpload`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectUploadsApi.GetDirectUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uPLOADID** | **string** | ID of the Upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UploadResponse**](UploadResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDirectUploads

> ListUploadsResponse ListDirectUploads(ctx).Limit(limit).Page(page).Execute()

List direct uploads

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
    limit := int32(56) // int32 | Number of items to include in the response (optional) (default to 25)
    page := int32(56) // int32 | Offset by this many pages, of the size of `limit` (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectUploadsApi.ListDirectUploads(context.Background()).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectUploadsApi.ListDirectUploads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDirectUploads`: ListUploadsResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectUploadsApi.ListDirectUploads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDirectUploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of items to include in the response | [default to 25]
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]

### Return type

[**ListUploadsResponse**](ListUploadsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

