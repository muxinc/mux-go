# \PlaybackIDApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssetOrLivestreamId**](PlaybackIDApi.md#GetAssetOrLivestreamId) | **Get** /video/v1/playback-ids/{PLAYBACK_ID} | Retrieve an Asset or Live Stream ID



## GetAssetOrLivestreamId

> GetAssetOrLiveStreamIdResponse GetAssetOrLivestreamId(ctx, pLAYBACKID).Execute()

Retrieve an Asset or Live Stream ID



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
    pLAYBACKID := "pLAYBACKID_example" // string | The live stream's playback ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaybackIDApi.GetAssetOrLivestreamId(context.Background(), pLAYBACKID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaybackIDApi.GetAssetOrLivestreamId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetOrLivestreamId`: GetAssetOrLiveStreamIdResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaybackIDApi.GetAssetOrLivestreamId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pLAYBACKID** | **string** | The live stream&#39;s playback ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetOrLivestreamIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAssetOrLiveStreamIdResponse**](GetAssetOrLiveStreamIdResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

