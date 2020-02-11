# \AssetsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAsset**](AssetsApi.md#CreateAsset) | **Post** /video/v1/assets | Create an asset
[**CreateAssetPlaybackId**](AssetsApi.md#CreateAssetPlaybackId) | **Post** /video/v1/assets/{ASSET_ID}/playback-ids | Create a playback ID
[**CreateAssetTrack**](AssetsApi.md#CreateAssetTrack) | **Post** /video/v1/assets/{ASSET_ID}/tracks | Create an asset track
[**DeleteAsset**](AssetsApi.md#DeleteAsset) | **Delete** /video/v1/assets/{ASSET_ID} | Delete an asset
[**DeleteAssetPlaybackId**](AssetsApi.md#DeleteAssetPlaybackId) | **Delete** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Delete a playback ID
[**DeleteAssetTrack**](AssetsApi.md#DeleteAssetTrack) | **Delete** /video/v1/assets/{ASSET_ID}/tracks/{TRACK_ID} | Delete an asset track
[**GetAsset**](AssetsApi.md#GetAsset) | **Get** /video/v1/assets/{ASSET_ID} | Retrieve an asset
[**GetAssetInputInfo**](AssetsApi.md#GetAssetInputInfo) | **Get** /video/v1/assets/{ASSET_ID}/input-info | Retrieve asset input info
[**GetAssetPlaybackId**](AssetsApi.md#GetAssetPlaybackId) | **Get** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Retrieve a playback ID
[**ListAssets**](AssetsApi.md#ListAssets) | **Get** /video/v1/assets | List assets
[**UpdateAssetMp4Support**](AssetsApi.md#UpdateAssetMp4Support) | **Put** /video/v1/assets/{ASSET_ID}/mp4-support | Update MP4 support


# **CreateAsset**
> AssetResponse CreateAsset(ctx, createAssetRequest)
Create an asset

Create a new Mux Video asset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createAssetRequest** | [**CreateAssetRequest**](CreateAssetRequest.md)|  | 

### Return type

[**AssetResponse**](AssetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAssetPlaybackId**
> CreatePlaybackIdResponse CreateAssetPlaybackId(ctx, aSSETID, createPlaybackIdRequest)
Create a playback ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **createPlaybackIdRequest** | [**CreatePlaybackIdRequest**](CreatePlaybackIdRequest.md)|  | 

### Return type

[**CreatePlaybackIdResponse**](CreatePlaybackIDResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAssetTrack**
> CreateTrackResponse CreateAssetTrack(ctx, aSSETID, createTrackRequest)
Create an asset track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **createTrackRequest** | [**CreateTrackRequest**](CreateTrackRequest.md)|  | 

### Return type

[**CreateTrackResponse**](CreateTrackResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAsset**
> DeleteAsset(ctx, aSSETID)
Delete an asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAssetPlaybackId**
> DeleteAssetPlaybackId(ctx, aSSETID, pLAYBACKID)
Delete a playback ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **pLAYBACKID** | **string**| The live stream&#39;s playback ID. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAssetTrack**
> DeleteAssetTrack(ctx, aSSETID, tRACKID)
Delete an asset track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **tRACKID** | **string**| The track ID. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsset**
> AssetResponse GetAsset(ctx, aSSETID)
Retrieve an asset

Retrieves the details of an asset that has previously been created. Supply the unique asset ID that was returned from your previous request, and Mux will return the corresponding asset information. The same information is returned when creating an asset.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 

### Return type

[**AssetResponse**](AssetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssetInputInfo**
> GetAssetInputInfoResponse GetAssetInputInfo(ctx, aSSETID)
Retrieve asset input info

Returns a list of the input objects that were used to create the asset along with any settings that were applied to each input.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 

### Return type

[**GetAssetInputInfoResponse**](GetAssetInputInfoResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssetPlaybackId**
> GetAssetPlaybackIdResponse GetAssetPlaybackId(ctx, aSSETID, pLAYBACKID)
Retrieve a playback ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **pLAYBACKID** | **string**| The live stream&#39;s playback ID. | 

### Return type

[**GetAssetPlaybackIdResponse**](GetAssetPlaybackIDResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAssets**
> ListAssetsResponse ListAssets(ctx, optional)
List assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAssetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]

### Return type

[**ListAssetsResponse**](ListAssetsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAssetMp4Support**
> AssetResponse UpdateAssetMp4Support(ctx, aSSETID, updateAssetMp4SupportRequest)
Update MP4 support

Allows you add or remove mp4 support for assets that were created without it. Currently there are two values supported in this request, `standard` and `none`. `none` means that an asset *does not* have mp4 support, so submitting a request with `mp4_support` set to `none` will delete the mp4 assets from the asset in question.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **updateAssetMp4SupportRequest** | [**UpdateAssetMp4SupportRequest**](UpdateAssetMp4SupportRequest.md)|  | 

### Return type

[**AssetResponse**](AssetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

