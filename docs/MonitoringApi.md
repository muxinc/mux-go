# \MonitoringApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMonitoringBreakdown**](MonitoringApi.md#GetMonitoringBreakdown) | **Get** /data/v1/monitoring/metrics/{MONITORING_METRIC_ID}/breakdown | Get Monitoring Breakdown
[**GetMonitoringBreakdownTimeseries**](MonitoringApi.md#GetMonitoringBreakdownTimeseries) | **Get** /data/v1/monitoring/metrics/{MONITORING_METRIC_ID}/breakdown-timeseries | Get Monitoring Breakdown Timeseries
[**GetMonitoringHistogramTimeseries**](MonitoringApi.md#GetMonitoringHistogramTimeseries) | **Get** /data/v1/monitoring/metrics/{MONITORING_HISTOGRAM_METRIC_ID}/histogram-timeseries | Get Monitoring Histogram Timeseries
[**GetMonitoringTimeseries**](MonitoringApi.md#GetMonitoringTimeseries) | **Get** /data/v1/monitoring/metrics/{MONITORING_METRIC_ID}/timeseries | Get Monitoring Timeseries
[**ListMonitoringDimensions**](MonitoringApi.md#ListMonitoringDimensions) | **Get** /data/v1/monitoring/dimensions | List Monitoring Dimensions
[**ListMonitoringMetrics**](MonitoringApi.md#ListMonitoringMetrics) | **Get** /data/v1/monitoring/metrics | List Monitoring Metrics


# **GetMonitoringBreakdown**
> GetMonitoringBreakdownResponse GetMonitoringBreakdown(ctx, mONITORINGMETRICID, optional)
Get Monitoring Breakdown

Gets breakdown information for a specific dimension and metric along with the number of concurrent viewers and negative impact score.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mONITORINGMETRICID** | **string**| ID of the Monitoring Metric | 
 **optional** | ***GetMonitoringBreakdownOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMonitoringBreakdownOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dimension** | **optional.String**| Dimension the specified value belongs to | 
 **timestamp** | **optional.Int32**| Timestamp to limit results by. This value must be provided as a unix timestamp. Defaults to the current unix timestamp. | 
 **filters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match conditions from provided key:value pairs. Must be provided as an array query string parameter.  To exclude rows that match a certain condition, prepend a &#x60;!&#x60; character to the dimension.  Possible filter names are the same as returned by the List Monitoring Dimensions endpoint.  Example:    * &#x60;filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;!country:US&#x60;  | 
 **orderBy** | **optional.String**| Value to order the results by | 
 **orderDirection** | **optional.String**| Sort order. | 

### Return type

[**GetMonitoringBreakdownResponse**](GetMonitoringBreakdownResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitoringBreakdownTimeseries**
> GetMonitoringBreakdownTimeseriesResponse GetMonitoringBreakdownTimeseries(ctx, mONITORINGMETRICID, optional)
Get Monitoring Breakdown Timeseries

Gets timeseries of breakdown information for a specific dimension and metric. Each datapoint in the response represents 5 seconds worth of data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mONITORINGMETRICID** | **string**| ID of the Monitoring Metric | 
 **optional** | ***GetMonitoringBreakdownTimeseriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMonitoringBreakdownTimeseriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dimension** | **optional.String**| Dimension the specified value belongs to | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;).  The default for this is the last 60 seconds of available data. Timeframes larger than 10 minutes are not allowed, and must be within the last 24 hours.  | 
 **filters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match conditions from provided key:value pairs. Must be provided as an array query string parameter.  To exclude rows that match a certain condition, prepend a &#x60;!&#x60; character to the dimension.  Possible filter names are the same as returned by the List Monitoring Dimensions endpoint.  Example:    * &#x60;filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;!country:US&#x60;  | 
 **limit** | **optional.Int32**| Number of items to include in each timestamp&#39;s &#x60;value&#x60; list.  The default is 10, and the maximum is 100.  | [default to 10]
 **orderBy** | **optional.String**| Value to order the results by | 
 **orderDirection** | **optional.String**| Sort order. | 

### Return type

[**GetMonitoringBreakdownTimeseriesResponse**](GetMonitoringBreakdownTimeseriesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitoringHistogramTimeseries**
> GetMonitoringHistogramTimeseriesResponse GetMonitoringHistogramTimeseries(ctx, mONITORINGHISTOGRAMMETRICID, optional)
Get Monitoring Histogram Timeseries

Gets histogram timeseries information for a specific metric.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mONITORINGHISTOGRAMMETRICID** | **string**| ID of the Monitoring Histogram Metric | 
 **optional** | ***GetMonitoringHistogramTimeseriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMonitoringHistogramTimeseriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match conditions from provided key:value pairs. Must be provided as an array query string parameter.  To exclude rows that match a certain condition, prepend a &#x60;!&#x60; character to the dimension.  Possible filter names are the same as returned by the List Monitoring Dimensions endpoint.  Example:    * &#x60;filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;!country:US&#x60;  | 

### Return type

[**GetMonitoringHistogramTimeseriesResponse**](GetMonitoringHistogramTimeseriesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMonitoringTimeseries**
> GetMonitoringTimeseriesResponse GetMonitoringTimeseries(ctx, mONITORINGMETRICID, optional)
Get Monitoring Timeseries

Gets Time series information for a specific metric along with the number of concurrent viewers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mONITORINGMETRICID** | **string**| ID of the Monitoring Metric | 
 **optional** | ***GetMonitoringTimeseriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMonitoringTimeseriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match conditions from provided key:value pairs. Must be provided as an array query string parameter.  To exclude rows that match a certain condition, prepend a &#x60;!&#x60; character to the dimension.  Possible filter names are the same as returned by the List Monitoring Dimensions endpoint.  Example:    * &#x60;filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;!country:US&#x60;  | 
 **timestamp** | **optional.Int32**| Timestamp to use as the start of the timeseries data. This value must be provided as a unix timestamp. Defaults to 30 minutes ago. | 

### Return type

[**GetMonitoringTimeseriesResponse**](GetMonitoringTimeseriesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMonitoringDimensions**
> ListMonitoringDimensionsResponse ListMonitoringDimensions(ctx, )
List Monitoring Dimensions

Lists available monitoring dimensions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListMonitoringDimensionsResponse**](ListMonitoringDimensionsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMonitoringMetrics**
> ListMonitoringMetricsResponse ListMonitoringMetrics(ctx, )
List Monitoring Metrics

Lists available monitoring metrics.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListMonitoringMetricsResponse**](ListMonitoringMetricsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

