# \LiveStreamsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLiveStream**](LiveStreamsApi.md#CreateLiveStream) | **Post** /video/v1/live-streams | Create a live stream
[**CreateLiveStreamPlaybackId**](LiveStreamsApi.md#CreateLiveStreamPlaybackId) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids | Create a live stream playback ID
[**CreateLiveStreamSimulcastTarget**](LiveStreamsApi.md#CreateLiveStreamSimulcastTarget) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets | Create a live stream simulcast target
[**DeleteLiveStream**](LiveStreamsApi.md#DeleteLiveStream) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID} | Delete a live stream
[**DeleteLiveStreamPlaybackId**](LiveStreamsApi.md#DeleteLiveStreamPlaybackId) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids/{PLAYBACK_ID} | Delete a live stream playback ID
[**DeleteLiveStreamSimulcastTarget**](LiveStreamsApi.md#DeleteLiveStreamSimulcastTarget) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets/{SIMULCAST_TARGET_ID} | Delete a Live Stream Simulcast Target
[**DisableLiveStream**](LiveStreamsApi.md#DisableLiveStream) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/disable | Disable a live stream
[**EnableLiveStream**](LiveStreamsApi.md#EnableLiveStream) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/enable | Enable a live stream
[**GetLiveStream**](LiveStreamsApi.md#GetLiveStream) | **Get** /video/v1/live-streams/{LIVE_STREAM_ID} | Retrieve a live stream
[**GetLiveStreamPlaybackId**](LiveStreamsApi.md#GetLiveStreamPlaybackId) | **Get** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids/{PLAYBACK_ID} | Retrieve a live stream playback ID
[**GetLiveStreamSimulcastTarget**](LiveStreamsApi.md#GetLiveStreamSimulcastTarget) | **Get** /video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets/{SIMULCAST_TARGET_ID} | Retrieve a Live Stream Simulcast Target
[**ListLiveStreams**](LiveStreamsApi.md#ListLiveStreams) | **Get** /video/v1/live-streams | List live streams
[**ResetStreamKey**](LiveStreamsApi.md#ResetStreamKey) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/reset-stream-key | Reset a live stream’s stream key
[**SignalLiveStreamComplete**](LiveStreamsApi.md#SignalLiveStreamComplete) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/complete | Signal a live stream is finished
[**UpdateLiveStreamEmbeddedSubtitles**](LiveStreamsApi.md#UpdateLiveStreamEmbeddedSubtitles) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/embedded-subtitles | Update a live stream&#39;s embedded subtitles


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

# **CreateLiveStreamSimulcastTarget**
> SimulcastTargetResponse CreateLiveStreamSimulcastTarget(ctx, lIVESTREAMID, createSimulcastTargetRequest)
Create a live stream simulcast target

Create a simulcast target for the parent live stream. Simulcast target can only be created when the parent live stream is in idle state. Only one simulcast target can be created at a time with this API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 
  **createSimulcastTargetRequest** | [**CreateSimulcastTargetRequest**](CreateSimulcastTargetRequest.md)|  | 

### Return type

[**SimulcastTargetResponse**](SimulcastTargetResponse.md)

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

# **DeleteLiveStreamSimulcastTarget**
> DeleteLiveStreamSimulcastTarget(ctx, lIVESTREAMID, sIMULCASTTARGETID)
Delete a Live Stream Simulcast Target

Delete the simulcast target using the simulcast target ID returned when creating the simulcast target. Simulcast Target can only be deleted when the parent live stream is in idle state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 
  **sIMULCASTTARGETID** | **string**| The ID of the simulcast target. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableLiveStream**
> DisableLiveStreamResponse DisableLiveStream(ctx, lIVESTREAMID)
Disable a live stream

Disables a live stream, making it reject incoming RTMP streams until re-enabled. The API also ends the live stream recording immediately when active. Ending the live stream recording adds the `EXT-X-ENDLIST` tag to the HLS manifest which notifies the player that this live stream is over.  Mux also closes the encoder connection immediately. Any attempt from the encoder to re-establish connection will fail till the live stream is re-enabled. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 

### Return type

[**DisableLiveStreamResponse**](DisableLiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableLiveStream**
> EnableLiveStreamResponse EnableLiveStream(ctx, lIVESTREAMID)
Enable a live stream

Enables a live stream, allowing it to accept an incoming RTMP stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 

### Return type

[**EnableLiveStreamResponse**](EnableLiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

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

# **GetLiveStreamPlaybackId**
> GetLiveStreamPlaybackIdResponse GetLiveStreamPlaybackId(ctx, lIVESTREAMID, pLAYBACKID)
Retrieve a live stream playback ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 
  **pLAYBACKID** | **string**| The live stream&#39;s playback ID. | 

### Return type

[**GetLiveStreamPlaybackIdResponse**](GetLiveStreamPlaybackIDResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLiveStreamSimulcastTarget**
> SimulcastTargetResponse GetLiveStreamSimulcastTarget(ctx, lIVESTREAMID, sIMULCASTTARGETID)
Retrieve a Live Stream Simulcast Target

Retrieves the details of the simulcast target created for the parent live stream. Supply the unique live stream ID and simulcast target ID that was returned in the response of create simulcast target request, and Mux will return the corresponding information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 
  **sIMULCASTTARGETID** | **string**| The ID of the simulcast target. | 

### Return type

[**SimulcastTargetResponse**](SimulcastTargetResponse.md)

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
 **streamKey** | **optional.String**| Filter response to return live stream for this stream key only | 

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

(Optional) End the live stream recording immediately instead of waiting for the reconnect_window. `EXT-X-ENDLIST` tag is added to the HLS manifest which notifies the player that this live stream is over.  Mux does not close the encoder connection immediately. Encoders are often configured to re-establish connections immediately which would result in a new recorded asset. For this reason, Mux waits for 60s before closing the connection with the encoder. This 60s timeframe is meant to give encoder operators a chance to disconnect from their end. 

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

# **UpdateLiveStreamEmbeddedSubtitles**
> LiveStreamResponse UpdateLiveStreamEmbeddedSubtitles(ctx, lIVESTREAMID, updateLiveStreamEmbeddedSubtitlesRequest)
Update a live stream's embedded subtitles

Configures a live stream to receive embedded subtitles including captions and translations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lIVESTREAMID** | **string**| The live stream ID | 
  **updateLiveStreamEmbeddedSubtitlesRequest** | [**UpdateLiveStreamEmbeddedSubtitlesRequest**](UpdateLiveStreamEmbeddedSubtitlesRequest.md)|  | 

### Return type

[**LiveStreamResponse**](LiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

