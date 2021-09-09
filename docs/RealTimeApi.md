# \RealTimeApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRealtimeBreakdown**](RealTimeApi.md#GetRealtimeBreakdown) | **Get** /data/v1/realtime/metrics/{REALTIME_METRIC_ID}/breakdown | Get Real-Time Breakdown
[**GetRealtimeHistogramTimeseries**](RealTimeApi.md#GetRealtimeHistogramTimeseries) | **Get** /data/v1/realtime/metrics/{REALTIME_HISTOGRAM_METRIC_ID}/histogram-timeseries | Get Real-Time Histogram Timeseries
[**GetRealtimeTimeseries**](RealTimeApi.md#GetRealtimeTimeseries) | **Get** /data/v1/realtime/metrics/{REALTIME_METRIC_ID}/timeseries | Get Real-Time Timeseries
[**ListRealtimeDimensions**](RealTimeApi.md#ListRealtimeDimensions) | **Get** /data/v1/realtime/dimensions | List Real-Time Dimensions
[**ListRealtimeMetrics**](RealTimeApi.md#ListRealtimeMetrics) | **Get** /data/v1/realtime/metrics | List Real-Time Metrics


# **GetRealtimeBreakdown**
> GetRealTimeBreakdownResponse GetRealtimeBreakdown(ctx, rEALTIMEMETRICID, optional)
Get Real-Time Breakdown

Gets breakdown information for a specific dimension and metric along with the number of concurrent viewers and negative impact score.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rEALTIMEMETRICID** | **string**| ID of the Realtime Metric | 
 **optional** | ***GetRealtimeBreakdownOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRealtimeBreakdownOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dimension** | **optional.String**| Dimension the specified value belongs to | 
 **timestamp** | **optional.Float64**| Timestamp to limit results by. This value must be provided as a unix timestamp. Defaults to the current unix timestamp. | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 
 **orderBy** | **optional.String**| Value to order the results by | 
 **orderDirection** | **optional.String**| Sort order. | 

### Return type

[**GetRealTimeBreakdownResponse**](GetRealTimeBreakdownResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRealtimeHistogramTimeseries**
> GetRealTimeHistogramTimeseriesResponse GetRealtimeHistogramTimeseries(ctx, rEALTIMEHISTOGRAMMETRICID, optional)
Get Real-Time Histogram Timeseries

Gets histogram timeseries information for a specific metric.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rEALTIMEHISTOGRAMMETRICID** | **string**| ID of the Realtime Histogram Metric | 
 **optional** | ***GetRealtimeHistogramTimeseriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRealtimeHistogramTimeseriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | [**optional.Interface of []string**](string.md)| Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 

### Return type

[**GetRealTimeHistogramTimeseriesResponse**](GetRealTimeHistogramTimeseriesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRealtimeTimeseries**
> GetRealTimeTimeseriesResponse GetRealtimeTimeseries(ctx, rEALTIMEMETRICID, optional)
Get Real-Time Timeseries

Gets Time series information for a specific metric along with the number of concurrent viewers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rEALTIMEMETRICID** | **string**| ID of the Realtime Metric | 
 **optional** | ***GetRealtimeTimeseriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRealtimeTimeseriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | [**optional.Interface of []string**](string.md)| Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 

### Return type

[**GetRealTimeTimeseriesResponse**](GetRealTimeTimeseriesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRealtimeDimensions**
> ListRealTimeDimensionsResponse ListRealtimeDimensions(ctx, )
List Real-Time Dimensions

Lists available real-time dimensions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListRealTimeDimensionsResponse**](ListRealTimeDimensionsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRealtimeMetrics**
> ListRealTimeMetricsResponse ListRealtimeMetrics(ctx, )
List Real-Time Metrics

Lists available real-time metrics.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListRealTimeMetricsResponse**](ListRealTimeMetricsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

