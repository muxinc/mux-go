# \DeliveryUsageApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDeliveryUsage**](DeliveryUsageApi.md#ListDeliveryUsage) | **Get** /video/v1/delivery-usage | List Usage


# **ListDeliveryUsage**
> ListDeliveryUsageResponse ListDeliveryUsage(ctx, optional)
List Usage

Returns a list of delivery usage records and their associated Asset IDs or Live Stream IDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDeliveryUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDeliveryUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 100]
 **assetId** | **optional.String**| Filter response to return delivery usage for this asset only. | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Time window to get delivery usage information. timeframe[0] indicates the start time, timeframe[1] indicates the end time in seconds since the Unix epoch. Default time window is 1 hour representing usage from 13th to 12th hour from when the request is made. | 

### Return type

[**ListDeliveryUsageResponse**](ListDeliveryUsageResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

