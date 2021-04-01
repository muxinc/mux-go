# \VideoViewsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVideoView**](VideoViewsApi.md#GetVideoView) | **Get** /data/v1/video-views/{VIDEO_VIEW_ID} | Get a Video View
[**ListVideoViews**](VideoViewsApi.md#ListVideoViews) | **Get** /data/v1/video-views | List Video Views



## GetVideoView

> VideoViewResponse GetVideoView(ctx, vIDEOVIEWID).Execute()

Get a Video View



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
    vIDEOVIEWID := "abcd1234" // string | ID of the Video View

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideoViewsApi.GetVideoView(context.Background(), vIDEOVIEWID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoViewsApi.GetVideoView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVideoView`: VideoViewResponse
    fmt.Fprintf(os.Stdout, "Response from `VideoViewsApi.GetVideoView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vIDEOVIEWID** | **string** | ID of the Video View | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VideoViewResponse**](VideoViewResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVideoViews

> ListVideoViewsResponse ListVideoViews(ctx).Limit(limit).Page(page).ViewerId(viewerId).ErrorId(errorId).OrderDirection(orderDirection).Filters(filters).Timeframe(timeframe).Execute()

List Video Views



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
    viewerId := "viewerId_example" // string | Viewer ID to filter results by. This value may be provided by the integration, or may be created by Mux. (optional)
    errorId := int32(56) // int32 | Filter video views by the provided error ID (as returned in the error_type_id field in the list video views endpoint). If you provide any as the error ID, this will filter the results to those with any error. (optional)
    orderDirection := "orderDirection_example" // string | Sort order. (optional)
    filters := []string{"Inner_example"} // []string | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US).  Possible filter names are the same as returned by the List Filters endpoint.  (optional)
    timeframe := []string{"Inner_example"} // []string | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]=). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]=1498867200&timeframe[]=1498953600    * duration string e.g. timeframe[]=24:hours or timeframe[]=7:days.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideoViewsApi.ListVideoViews(context.Background()).Limit(limit).Page(page).ViewerId(viewerId).ErrorId(errorId).OrderDirection(orderDirection).Filters(filters).Timeframe(timeframe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoViewsApi.ListVideoViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVideoViews`: ListVideoViewsResponse
    fmt.Fprintf(os.Stdout, "Response from `VideoViewsApi.ListVideoViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVideoViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of items to include in the response | [default to 25]
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **viewerId** | **string** | Viewer ID to filter results by. This value may be provided by the integration, or may be created by Mux. | 
 **errorId** | **int32** | Filter video views by the provided error ID (as returned in the error_type_id field in the list video views endpoint). If you provide any as the error ID, this will filter the results to those with any error. | 
 **orderDirection** | **string** | Sort order. | 
 **filters** | **[]string** | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US).  Possible filter names are the same as returned by the List Filters endpoint.  | 
 **timeframe** | **[]string** | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600    * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 

### Return type

[**ListVideoViewsResponse**](ListVideoViewsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

