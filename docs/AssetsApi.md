# \AssetsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAsset**](AssetsApi.md#CreateAsset) | **Post** /video/v1/assets | Create an asset
[**CreateAssetPlaybackId**](AssetsApi.md#CreateAssetPlaybackId) | **Post** /video/v1/assets/{ASSET_ID}/playback-ids | Create a playback ID
[**CreateAssetStaticRendition**](AssetsApi.md#CreateAssetStaticRendition) | **Post** /video/v1/assets/{ASSET_ID}/static-renditions | Create a static rendition for an asset
[**CreateAssetTrack**](AssetsApi.md#CreateAssetTrack) | **Post** /video/v1/assets/{ASSET_ID}/tracks | Create an asset track
[**DeleteAsset**](AssetsApi.md#DeleteAsset) | **Delete** /video/v1/assets/{ASSET_ID} | Delete an asset
[**DeleteAssetPlaybackId**](AssetsApi.md#DeleteAssetPlaybackId) | **Delete** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Delete a playback ID
[**DeleteAssetStaticRendition**](AssetsApi.md#DeleteAssetStaticRendition) | **Delete** /video/v1/assets/{ASSET_ID}/static-renditions/{STATIC_RENDITION_ID} | Delete a single static rendition for an asset
[**DeleteAssetTrack**](AssetsApi.md#DeleteAssetTrack) | **Delete** /video/v1/assets/{ASSET_ID}/tracks/{TRACK_ID} | Delete an asset track
[**GenerateAssetTrackSubtitles**](AssetsApi.md#GenerateAssetTrackSubtitles) | **Post** /video/v1/assets/{ASSET_ID}/tracks/{TRACK_ID}/generate-subtitles | Generate track subtitles
[**GetAsset**](AssetsApi.md#GetAsset) | **Get** /video/v1/assets/{ASSET_ID} | Retrieve an asset
[**GetAssetInputInfo**](AssetsApi.md#GetAssetInputInfo) | **Get** /video/v1/assets/{ASSET_ID}/input-info | Retrieve asset input info
[**GetAssetPlaybackId**](AssetsApi.md#GetAssetPlaybackId) | **Get** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Retrieve a playback ID
[**ListAssets**](AssetsApi.md#ListAssets) | **Get** /video/v1/assets | List assets
[**UpdateAsset**](AssetsApi.md#UpdateAsset) | **Patch** /video/v1/assets/{ASSET_ID} | Update an asset
[**UpdateAssetMasterAccess**](AssetsApi.md#UpdateAssetMasterAccess) | **Put** /video/v1/assets/{ASSET_ID}/master-access | Update master access
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

Creates a playback ID that can be used to stream the asset to a viewer.

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

# **CreateAssetStaticRendition**
> CreateStaticRenditionResponse CreateAssetStaticRendition(ctx, aSSETID, createStaticRenditionRequest)
Create a static rendition for an asset

Creates a static rendition (i.e. MP4) for an asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **createStaticRenditionRequest** | [**CreateStaticRenditionRequest**](CreateStaticRenditionRequest.md)|  | 

### Return type

[**CreateStaticRenditionResponse**](CreateStaticRenditionResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAssetTrack**
> CreateTrackResponse CreateAssetTrack(ctx, aSSETID, createTrackRequest)
Create an asset track

Adds an asset track (for example, subtitles, or an alternate audio track) to an asset. Assets must be in the `ready` state before tracks can be added.

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

Deletes a video asset and all its data.

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

Deletes a playback ID, rendering it nonfunctional for viewing an asset's video content. Please note that deleting the playback ID removes access to the underlying asset; a viewer who started playback before the playback ID was deleted may be able to watch the entire video for a limited duration.

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

# **DeleteAssetStaticRendition**
> DeleteAssetStaticRendition(ctx, aSSETID, sTATICRENDITIONID)
Delete a single static rendition for an asset

Deletes a single static rendition for an asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **sTATICRENDITIONID** | **string**| The static rendition ID. | 

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

Removes a text or additional audio track from an asset. Neither video nor the primary audio track can be removed.

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

# **GenerateAssetTrackSubtitles**
> GenerateTrackSubtitlesResponse GenerateAssetTrackSubtitles(ctx, aSSETID, tRACKID, generateTrackSubtitlesRequest)
Generate track subtitles

Generates subtitles (captions) for a given audio track. This API can be used for up to 7 days after an asset is created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **tRACKID** | **string**| The track ID. | 
  **generateTrackSubtitlesRequest** | [**GenerateTrackSubtitlesRequest**](GenerateTrackSubtitlesRequest.md)|  | 

### Return type

[**GenerateTrackSubtitlesResponse**](GenerateTrackSubtitlesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

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

Retrieves information about the specified playback ID.

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

List all Mux assets.

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
 **cursor** | **optional.String**| This parameter is used to request pages beyond the first. You can find the cursor value in the &#x60;next_cursor&#x60; field of paginated responses. | 
 **liveStreamId** | **optional.String**| Filter response to return all the assets for this live stream only | 
 **uploadId** | **optional.String**| Filter response to return an asset created from this direct upload only | 

### Return type

[**ListAssetsResponse**](ListAssetsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAsset**
> AssetResponse UpdateAsset(ctx, aSSETID, updateAssetRequest)
Update an asset

Updates the details of an already-created Asset with the provided Asset ID. This currently supports only the `passthrough` field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **updateAssetRequest** | [**UpdateAssetRequest**](UpdateAssetRequest.md)|  | 

### Return type

[**AssetResponse**](AssetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAssetMasterAccess**
> AssetResponse UpdateAssetMasterAccess(ctx, aSSETID, updateAssetMasterAccessRequest)
Update master access

Allows you to add temporary access to the master (highest-quality) version of the asset in MP4 format. A URL will be created that can be used to download the master version for 24 hours. After 24 hours Master Access will revert to \"none\". This master version is not optimized for web and not meant to be streamed, only downloaded for purposes like archiving or editing the video offline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aSSETID** | **string**| The asset ID. | 
  **updateAssetMasterAccessRequest** | [**UpdateAssetMasterAccessRequest**](UpdateAssetMasterAccessRequest.md)|  | 

### Return type

[**AssetResponse**](AssetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAssetMp4Support**
> AssetResponse UpdateAssetMp4Support(ctx, aSSETID, updateAssetMp4SupportRequest)
Update MP4 support

This method has been deprecated. Please see the [Static Rendition API](https://www.mux.com/docs/guides/enable-static-mp4-renditions#after-asset-creation). Allows you to add or remove mp4 support for assets that were created without it. The values supported are `capped-1080p`, `audio-only`, `audio-only,capped-1080p`, `standard`(deprecated),  and `none`. `none` means that an asset *does not* have mp4 support, so submitting a request with `mp4_support` set to `none` will delete the mp4 assets from the asset in question. 

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

