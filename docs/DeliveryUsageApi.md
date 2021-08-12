# \DeliveryUsageApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDeliveryUsage**](DeliveryUsageApi.md#ListDeliveryUsage) | **Get** /video/v1/delivery-usage | List Usage



## ListDeliveryUsage

> ListDeliveryUsageResponse ListDeliveryUsage(ctx).Page(page).Limit(limit).AssetId(assetId).Timeframe(timeframe).Execute()

List Usage



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
    page := int32(56) // int32 | Offset by this many pages, of the size of `limit` (optional) (default to 1)
    limit := int32(56) // int32 | Number of items to include in the response (optional) (default to 100)
    assetId := "assetId_example" // string | Filter response to return delivery usage for this asset only. (optional)
    timeframe := []string{"Inner_example"} // []string | Time window to get delivery usage information. timeframe[0] indicates the start time, timeframe[1] indicates the end time in seconds since the Unix epoch. Default time window is 1 hour representing usage from 13th to 12th hour from when the request is made. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeliveryUsageApi.ListDeliveryUsage(context.Background()).Page(page).Limit(limit).AssetId(assetId).Timeframe(timeframe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryUsageApi.ListDeliveryUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeliveryUsage`: ListDeliveryUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `DeliveryUsageApi.ListDeliveryUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeliveryUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **limit** | **int32** | Number of items to include in the response | [default to 100]
 **assetId** | **string** | Filter response to return delivery usage for this asset only. | 
 **timeframe** | **[]string** | Time window to get delivery usage information. timeframe[0] indicates the start time, timeframe[1] indicates the end time in seconds since the Unix epoch. Default time window is 1 hour representing usage from 13th to 12th hour from when the request is made. | 

### Return type

[**ListDeliveryUsageResponse**](ListDeliveryUsageResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

