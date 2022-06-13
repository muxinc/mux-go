# \SpacesApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpace**](SpacesApi.md#CreateSpace) | **Post** /video/v1/spaces | Create a space
[**CreateSpaceBroadcast**](SpacesApi.md#CreateSpaceBroadcast) | **Post** /video/v1/spaces/{SPACE_ID}/broadcasts | Create a space broadcast
[**DeleteSpace**](SpacesApi.md#DeleteSpace) | **Delete** /video/v1/spaces/{SPACE_ID} | Delete a space
[**DeleteSpaceBroadcast**](SpacesApi.md#DeleteSpaceBroadcast) | **Delete** /video/v1/spaces/{SPACE_ID}/broadcasts/{BROADCAST_ID} | Delete a space broadcast
[**GetSpace**](SpacesApi.md#GetSpace) | **Get** /video/v1/spaces/{SPACE_ID} | Retrieve a space
[**GetSpaceBroadcast**](SpacesApi.md#GetSpaceBroadcast) | **Get** /video/v1/spaces/{SPACE_ID}/broadcasts/{BROADCAST_ID} | Retrieve space broadcast
[**ListSpaces**](SpacesApi.md#ListSpaces) | **Get** /video/v1/spaces | List spaces
[**StartSpaceBroadcast**](SpacesApi.md#StartSpaceBroadcast) | **Post** /video/v1/spaces/{SPACE_ID}/broadcasts/{BROADCAST_ID}/start | Start a space broadcast
[**StopSpaceBroadcast**](SpacesApi.md#StopSpaceBroadcast) | **Post** /video/v1/spaces/{SPACE_ID}/broadcasts/{BROADCAST_ID}/stop | Stop a space broadcast


# **CreateSpace**
> SpaceResponse CreateSpace(ctx, createSpaceRequest)
Create a space

Create a new space. Spaces are used to build [real-time video applications.](https://mux.com/real-time-video)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSpaceRequest** | [**CreateSpaceRequest**](CreateSpaceRequest.md)|  | 

### Return type

[**SpaceResponse**](SpaceResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSpaceBroadcast**
> BroadcastResponse CreateSpaceBroadcast(ctx, sPACEID, createBroadcastRequest)
Create a space broadcast

Creates a new broadcast. Broadcasts are used to create composited versions of your space, which can be broadcast to live streams.  **Note:** By default only a single broadcast destination can be specified. Contact Mux support if you need more.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sPACEID** | **string**| The space ID. | 
  **createBroadcastRequest** | [**CreateBroadcastRequest**](CreateBroadcastRequest.md)|  | 

### Return type

[**BroadcastResponse**](BroadcastResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpace**
> DeleteSpace(ctx, sPACEID)
Delete a space

Deletes a space. Spaces can only be deleted when `idle`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sPACEID** | **string**| The space ID. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpaceBroadcast**
> DeleteSpaceBroadcast(ctx, sPACEID, bROADCASTID)
Delete a space broadcast

Deletes a single broadcast of a specific space. Broadcasts can only be deleted when `idle`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sPACEID** | **string**| The space ID. | 
  **bROADCASTID** | **string**| The broadcast ID. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpace**
> SpaceResponse GetSpace(ctx, sPACEID)
Retrieve a space

Retrieves the details of a space that has previously been created. Supply the unique space ID that was returned from your create space request, and Mux will return the information about the corresponding space. The same information is returned when creating a space.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sPACEID** | **string**| The space ID. | 

### Return type

[**SpaceResponse**](SpaceResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpaceBroadcast**
> BroadcastResponse GetSpaceBroadcast(ctx, sPACEID, bROADCASTID)
Retrieve space broadcast

Retrieves the details of a broadcast of a specific space.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sPACEID** | **string**| The space ID. | 
  **bROADCASTID** | **string**| The broadcast ID. | 

### Return type

[**BroadcastResponse**](BroadcastResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSpaces**
> ListSpacesResponse ListSpaces(ctx, optional)
List spaces

List all spaces in the current enviroment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSpacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSpacesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]

### Return type

[**ListSpacesResponse**](ListSpacesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSpaceBroadcast**
> StartSpaceBroadcastResponse StartSpaceBroadcast(ctx, sPACEID, bROADCASTID)
Start a space broadcast

Starts broadcasting a space to the associated destination. Broadcasts can only be started when the space is `active` (when there are participants connected).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sPACEID** | **string**| The space ID. | 
  **bROADCASTID** | **string**| The broadcast ID. | 

### Return type

[**StartSpaceBroadcastResponse**](StartSpaceBroadcastResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopSpaceBroadcast**
> StopSpaceBroadcastResponse StopSpaceBroadcast(ctx, sPACEID, bROADCASTID)
Stop a space broadcast

Stops broadcasting a space, causing the destination live stream to become idle. This API also automatically calls `complete` on the destination live stream. Broadcasts are also automatically stopped when a space becomes idle.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sPACEID** | **string**| The space ID. | 
  **bROADCASTID** | **string**| The broadcast ID. | 

### Return type

[**StopSpaceBroadcastResponse**](StopSpaceBroadcastResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

