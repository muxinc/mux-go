# \URLSigningKeysApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUrlSigningKey**](URLSigningKeysApi.md#CreateUrlSigningKey) | **Post** /video/v1/signing-keys | Create a URL signing key
[**DeleteUrlSigningKey**](URLSigningKeysApi.md#DeleteUrlSigningKey) | **Delete** /video/v1/signing-keys/{SIGNING_KEY_ID} | Delete a URL signing key
[**GetUrlSigningKey**](URLSigningKeysApi.md#GetUrlSigningKey) | **Get** /video/v1/signing-keys/{SIGNING_KEY_ID} | Retrieve a URL signing key
[**ListUrlSigningKeys**](URLSigningKeysApi.md#ListUrlSigningKeys) | **Get** /video/v1/signing-keys | List URL signing keys



## CreateUrlSigningKey

> SigningKeyResponse CreateUrlSigningKey(ctx).Execute()

Create a URL signing key



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
    resp, r, err := api_client.URLSigningKeysApi.CreateUrlSigningKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `URLSigningKeysApi.CreateUrlSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUrlSigningKey`: SigningKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `URLSigningKeysApi.CreateUrlSigningKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUrlSigningKeyRequest struct via the builder pattern


### Return type

[**SigningKeyResponse**](SigningKeyResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUrlSigningKey

> DeleteUrlSigningKey(ctx, sIGNINGKEYID).Execute()

Delete a URL signing key



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
    sIGNINGKEYID := "sIGNINGKEYID_example" // string | The ID of the signing key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.URLSigningKeysApi.DeleteUrlSigningKey(context.Background(), sIGNINGKEYID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `URLSigningKeysApi.DeleteUrlSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sIGNINGKEYID** | **string** | The ID of the signing key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUrlSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUrlSigningKey

> SigningKeyResponse GetUrlSigningKey(ctx, sIGNINGKEYID).Execute()

Retrieve a URL signing key



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
    sIGNINGKEYID := "sIGNINGKEYID_example" // string | The ID of the signing key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.URLSigningKeysApi.GetUrlSigningKey(context.Background(), sIGNINGKEYID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `URLSigningKeysApi.GetUrlSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUrlSigningKey`: SigningKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `URLSigningKeysApi.GetUrlSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sIGNINGKEYID** | **string** | The ID of the signing key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUrlSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigningKeyResponse**](SigningKeyResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUrlSigningKeys

> ListSigningKeysResponse ListUrlSigningKeys(ctx).Limit(limit).Page(page).Execute()

List URL signing keys



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
    resp, r, err := api_client.URLSigningKeysApi.ListUrlSigningKeys(context.Background()).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `URLSigningKeysApi.ListUrlSigningKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUrlSigningKeys`: ListSigningKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `URLSigningKeysApi.ListUrlSigningKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUrlSigningKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of items to include in the response | [default to 25]
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]

### Return type

[**ListSigningKeysResponse**](ListSigningKeysResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

