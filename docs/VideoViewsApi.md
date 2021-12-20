# \VideoViewsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVideoView**](VideoViewsApi.md#GetVideoView) | **Get** /data/v1/video-views/{VIDEO_VIEW_ID} | Get a Video View
[**ListVideoViews**](VideoViewsApi.md#ListVideoViews) | **Get** /data/v1/video-views | List Video Views


# **GetVideoView**
> VideoViewResponse GetVideoView(ctx, vIDEOVIEWID)
Get a Video View

Returns the details of a video view.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vIDEOVIEWID** | **string**| ID of the Video View | 

### Return type

[**VideoViewResponse**](VideoViewResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVideoViews**
> ListVideoViewsResponse ListVideoViews(ctx, optional)
List Video Views

Returns a list of video views.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVideoViewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListVideoViewsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **viewerId** | **optional.String**| Viewer ID to filter results by. This value may be provided by the integration, or may be created by Mux. | 
 **errorId** | **optional.Int32**| Filter video views by the provided error ID (as returned in the error_type_id field in the list video views endpoint). If you provide any as the error ID, this will filter the results to those with any error. | 
 **orderDirection** | **optional.String**| Sort order. | 
 **filters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match conditions from provided key:value pairs. Must be provided as an array query string parameter.  To exclude rows that match a certain condition, prepend a &#x60;!&#x60; character to the dimension.  Possible filter names are the same as returned by the List Filters endpoint.  Example:    * &#x60;filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;!country:US&#x60;  | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;).  Accepted formats are...    * array of epoch timestamps e.g. &#x60;timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600&#x60;   * duration string e.g. &#x60;timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days&#x60;  | 

### Return type

[**ListVideoViewsResponse**](ListVideoViewsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

