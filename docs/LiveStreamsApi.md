# \LiveStreamsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLiveStream**](LiveStreamsApi.md#CreateLiveStream) | **Post** /video/v1/live-streams | Create a live stream
[**CreateLiveStreamPlaybackId**](LiveStreamsApi.md#CreateLiveStreamPlaybackId) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids | Create a live stream playback ID
[**DeleteLiveStream**](LiveStreamsApi.md#DeleteLiveStream) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID} | Delete a live stream
[**DeleteLiveStreamPlaybackId**](LiveStreamsApi.md#DeleteLiveStreamPlaybackId) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids/{PLAYBACK_ID} | Delete a live stream playback ID
[**GetLiveStream**](LiveStreamsApi.md#GetLiveStream) | **Get** /video/v1/live-streams/{LIVE_STREAM_ID} | Retrieve a live stream
[**ListLiveStreams**](LiveStreamsApi.md#ListLiveStreams) | **Get** /video/v1/live-streams | List live streams
[**ResetStreamKey**](LiveStreamsApi.md#ResetStreamKey) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/reset-stream-key | Reset a live stream’s stream key
[**SignalLiveStreamComplete**](LiveStreamsApi.md#SignalLiveStreamComplete) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/complete | Signal a live stream is finished


# **CreateLiveStream**
> LiveStreamResponse CreateLiveStream(ctx, createLiveStreamRequest)
Create a live stream

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createLiveStreamRequest** | [**CreateLiveStreamRequest**](CreateLiveStreamRequest.md)|  | 

### Return type

[**LiveStreamResponse**](LiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLiveStreamPlaybackId**
> CreatePlaybackIdResponse CreateLiveStreamPlaybackId(ctx, lIVESTREAMID, createPlaybackIdRequest)
Create a live stream playback ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 
  **createPlaybackIdRequest** | [**CreatePlaybackIdRequest**](CreatePlaybackIdRequest.md)|  | 

### Return type

[**CreatePlaybackIdResponse**](CreatePlaybackIDResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLiveStream**
> DeleteLiveStream(ctx, lIVESTREAMID)
Delete a live stream

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLiveStreamPlaybackId**
> DeleteLiveStreamPlaybackId(ctx, lIVESTREAMID, pLAYBACKID)
Delete a live stream playback ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 
  **pLAYBACKID** | **string**| The live stream&#39;s playback ID. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLiveStream**
> LiveStreamResponse GetLiveStream(ctx, lIVESTREAMID)
Retrieve a live stream

Retrieves the details of a live stream that has previously been created. Supply the unique live stream ID that was returned from your previous request, and Mux will return the corresponding live stream information. The same information is returned when creating a live stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 

### Return type

[**LiveStreamResponse**](LiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLiveStreams**
> ListLiveStreamsResponse ListLiveStreams(ctx, optional)
List live streams

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListLiveStreamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListLiveStreamsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]

### Return type

[**ListLiveStreamsResponse**](ListLiveStreamsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetStreamKey**
> LiveStreamResponse ResetStreamKey(ctx, lIVESTREAMID)
Reset a live stream’s stream key

Reset a live stream key if you want to immediately stop the current stream key from working and create a new stream key that can be used for future broadcasts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 

### Return type

[**LiveStreamResponse**](LiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignalLiveStreamComplete**
> SignalLiveStreamCompleteResponse SignalLiveStreamComplete(ctx, lIVESTREAMID)
Signal a live stream is finished

(Optional) Make the recorded asset available immediately instead of waiting for the reconnect_window.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 

### Return type

[**SignalLiveStreamCompleteResponse**](SignalLiveStreamCompleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

