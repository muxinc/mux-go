# \ErrorsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListErrors**](ErrorsApi.md#ListErrors) | **Get** /data/v1/errors | List Errors


# **ListErrors**
> ListErrorsResponse ListErrors(ctx, optional)
List Errors

Returns a list of errors.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListErrorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListErrorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**optional.Interface of []string**](string.md)| Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600   * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 

### Return type

[**ListErrorsResponse**](ListErrorsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

