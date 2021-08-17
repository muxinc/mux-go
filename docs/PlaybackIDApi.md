# \PlaybackIDApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssetOrLivestreamId**](PlaybackIDApi.md#GetAssetOrLivestreamId) | **Get** /video/v1/playback-ids/{PLAYBACK_ID} | Retrieve an Asset or Live Stream ID


# **GetAssetOrLivestreamId**
> GetAssetOrLiveStreamIdResponse GetAssetOrLivestreamId(ctx, pLAYBACKID)
Retrieve an Asset or Live Stream ID

Retrieves the Identifier of the Asset or Live Stream associated with the Playback ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pLAYBACKID** | **string**| The live stream&#39;s playback ID. | 

### Return type

[**GetAssetOrLiveStreamIdResponse**](GetAssetOrLiveStreamIdResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

