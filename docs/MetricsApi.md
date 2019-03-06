# \MetricsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetricTimeseriesData**](MetricsApi.md#GetMetricTimeseriesData) | **Get** /data/v1/metrics/{METRIC_ID}/timeseries | Get metric timeseries data
[**GetOverallValues**](MetricsApi.md#GetOverallValues) | **Get** /data/v1/metrics/{METRIC_ID}/overall | Get Overall values
[**ListAllMetricValues**](MetricsApi.md#ListAllMetricValues) | **Get** /data/v1/metrics/comparison | List all metric values
[**ListBreakdownValues**](MetricsApi.md#ListBreakdownValues) | **Get** /data/v1/metrics/{METRIC_ID}/breakdown | List breakdown values
[**ListInsights**](MetricsApi.md#ListInsights) | **Get** /data/v1/metrics/{METRIC_ID}/insights | List Insights


# **GetMetricTimeseriesData**
> GetMetricTimeseriesDataResponse GetMetricTimeseriesData(ctx, mETRICID, optional)
Get metric timeseries data

Returns timeseries data for a specific metric 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mETRICID** | **string**| ID of the Metric | 
 **optional** | ***GetMetricTimeseriesDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMetricTimeseriesDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600    * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US).  Possible filter names are the same as returned by the List Filters endpoint.  | 
 **measurement** | **optional.String**| Measurement for the provided metric. If omitted, the deafult for the metric will be used. | 
 **orderDirection** | **optional.String**| Sort order. | 
 **groupBy** | **optional.String**| Time granularity to group results by. If this value is omitted, a default granularity is chosen based on the supplied timeframe. | 

### Return type

[**GetMetricTimeseriesDataResponse**](GetMetricTimeseriesDataResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOverallValues**
> GetOverallValuesResponse GetOverallValues(ctx, mETRICID, optional)
Get Overall values

Returns the overall value for a specific metric, as well as the total view count, watch time, and the Mux Global metric value for the metric. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mETRICID** | **string**| ID of the Metric | 
 **optional** | ***GetOverallValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOverallValuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600    * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US).  Possible filter names are the same as returned by the List Filters endpoint.  | 
 **measurement** | **optional.String**| Measurement for the provided metric. If omitted, the deafult for the metric will be used. | 

### Return type

[**GetOverallValuesResponse**](GetOverallValuesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllMetricValues**
> ListAllMetricValuesResponse ListAllMetricValues(ctx, optional)
List all metric values

List all of the values across every breakdown for a specific metric 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAllMetricValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAllMetricValuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600    * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US).  Possible filter names are the same as returned by the List Filters endpoint.  | 
 **dimension** | **optional.String**| Dimension the specified value belongs to | 
 **value** | **optional.String**| Value to show all available metrics for | 

### Return type

[**ListAllMetricValuesResponse**](ListAllMetricValuesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBreakdownValues**
> ListBreakdownValuesResponse ListBreakdownValues(ctx, mETRICID, optional)
List breakdown values

List the breakdown values for a specific metric 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mETRICID** | **string**| ID of the Metric | 
 **optional** | ***ListBreakdownValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListBreakdownValuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBy** | **optional.String**| Breakdown value to group the results by | 
 **measurement** | **optional.String**| Measurement for the provided metric. If omitted, the deafult for the metric will be used. | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US).  Possible filter names are the same as returned by the List Filters endpoint.  | 
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **orderBy** | **optional.String**| Value to order the results by | 
 **orderDirection** | **optional.String**| Sort order. | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600    * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 

### Return type

[**ListBreakdownValuesResponse**](ListBreakdownValuesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInsights**
> ListInsightsResponse ListInsights(ctx, mETRICID, optional)
List Insights

Returns a list of insights for a metric. These are the worst performing values across all breakdowns sorted by how much they negatively impact a specific metric. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mETRICID** | **string**| ID of the Metric | 
 **optional** | ***ListInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListInsightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **measurement** | **optional.String**| Measurement for the provided metric. If omitted, the deafult for the metric will be used. | 
 **orderDirection** | **optional.String**| Sort order. | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600    * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 

### Return type

[**ListInsightsResponse**](ListInsightsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

