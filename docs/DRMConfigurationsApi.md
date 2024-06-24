# \DRMConfigurationsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDrmConfiguration**](DRMConfigurationsApi.md#GetDrmConfiguration) | **Get** /video/v1/drm-configurations/{DRM_CONFIGURATION_ID} | Retrieve a DRM Configuration
[**ListDrmConfigurations**](DRMConfigurationsApi.md#ListDrmConfigurations) | **Get** /video/v1/drm-configurations | List DRM Configurations


# **GetDrmConfiguration**
> DrmConfigurationResponse GetDrmConfiguration(ctx, dRMCONFIGURATIONID)
Retrieve a DRM Configuration

Retrieves a single DRM Configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dRMCONFIGURATIONID** | **string**| The DRM Configuration ID. | 

### Return type

[**DrmConfigurationResponse**](DRMConfigurationResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDrmConfigurations**
> ListDrmConfigurationsResponse ListDrmConfigurations(ctx, optional)
List DRM Configurations

Returns a list of DRM Configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDrmConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDrmConfigurationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]

### Return type

[**ListDrmConfigurationsResponse**](ListDRMConfigurationsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

