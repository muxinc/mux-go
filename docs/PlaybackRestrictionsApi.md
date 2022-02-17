# \PlaybackRestrictionsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlaybackRestriction**](PlaybackRestrictionsApi.md#CreatePlaybackRestriction) | **Post** /video/v1/playback-restrictions | Create a Playback Restriction
[**DeletePlaybackRestriction**](PlaybackRestrictionsApi.md#DeletePlaybackRestriction) | **Delete** /video/v1/playback-restrictions/{PLAYBACK_RESTRICTION_ID} | Delete a Playback Restriction
[**GetPlaybackRestriction**](PlaybackRestrictionsApi.md#GetPlaybackRestriction) | **Get** /video/v1/playback-restrictions/{PLAYBACK_RESTRICTION_ID} | Retrieve a Playback Restriction
[**ListPlaybackRestrictions**](PlaybackRestrictionsApi.md#ListPlaybackRestrictions) | **Get** /video/v1/playback-restrictions | List Playback Restrictions
[**UpdateReferrerDomainRestriction**](PlaybackRestrictionsApi.md#UpdateReferrerDomainRestriction) | **Put** /video/v1/playback-restrictions/{PLAYBACK_RESTRICTION_ID}/referrer | Update the Referrer Playback Restriction


# **CreatePlaybackRestriction**
> PlaybackRestrictionResponse CreatePlaybackRestriction(ctx, createPlaybackRestrictionRequest)
Create a Playback Restriction

Create a new Playback Restriction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createPlaybackRestrictionRequest** | [**CreatePlaybackRestrictionRequest**](CreatePlaybackRestrictionRequest.md)|  | 

### Return type

[**PlaybackRestrictionResponse**](PlaybackRestrictionResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePlaybackRestriction**
> DeletePlaybackRestriction(ctx, pLAYBACKRESTRICTIONID)
Delete a Playback Restriction

Deletes a single Playback Restriction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pLAYBACKRESTRICTIONID** | **string**| ID of the Playback Restriction. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlaybackRestriction**
> PlaybackRestrictionResponse GetPlaybackRestriction(ctx, pLAYBACKRESTRICTIONID)
Retrieve a Playback Restriction

Retrieves a Playback Restriction associated with the unique identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pLAYBACKRESTRICTIONID** | **string**| ID of the Playback Restriction. | 

### Return type

[**PlaybackRestrictionResponse**](PlaybackRestrictionResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPlaybackRestrictions**
> ListPlaybackRestrictionsResponse ListPlaybackRestrictions(ctx, optional)
List Playback Restrictions

Returns a list of all Playback Restrictions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPlaybackRestrictionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPlaybackRestrictionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]

### Return type

[**ListPlaybackRestrictionsResponse**](ListPlaybackRestrictionsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReferrerDomainRestriction**
> PlaybackRestrictionResponse UpdateReferrerDomainRestriction(ctx, pLAYBACKRESTRICTIONID, body)
Update the Referrer Playback Restriction

Allows you to modify the list of domains or change how Mux validates playback requests without the `Referer` HTTP header. The Referrer restriction fully replaces the old list with this new list of domains.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pLAYBACKRESTRICTIONID** | **string**| ID of the Playback Restriction. | 
  **body** | **ReferrerDomainRestriction**|  | 

### Return type

[**PlaybackRestrictionResponse**](PlaybackRestrictionResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

