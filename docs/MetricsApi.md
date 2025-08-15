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

Returns timeseries data for a specific metric.  Each interval represented in the data array contains an array with the following values:   * the first element is the interval time   * the second element is the calculated metric value   * the third element is the number of views in the interval that have a valid metric value 

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

 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;).  Accepted formats are...    * array of epoch timestamps e.g. &#x60;timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600&#x60;   * duration string e.g. &#x60;timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days&#x60;  | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter results using key:value pairs. Must be provided as an array query string parameter.  **Basic filtering:** * &#x60;filters[]&#x3D;dimension:value&#x60; - Include rows where dimension equals value * &#x60;filters[]&#x3D;!dimension:value&#x60; - Exclude rows where dimension equals value  **For trace dimensions (like video_cdn_trace):** * &#x60;filters[]&#x3D;+dimension:value&#x60; - Include rows where trace contains value * &#x60;filters[]&#x3D;-dimension:value&#x60; - Exclude rows where trace contains value * &#x60;filters[]&#x3D;dimension:[value1,value2]&#x60; - Exact trace match  **Examples:** * &#x60;filters[]&#x3D;country:US&#x60; - US views only * &#x60;filters[]&#x3D;+video_cdn_trace:fastly&#x60; - Views using Fastly CDN  | 
 **metricFilters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match inequality conditions from provided metric comparison clauses. Must be provided as an array query string parameter.  Possible filterable metrics are the same as the set of metric ids, with the exceptions of &#x60;exits_before_video_start&#x60;, &#x60;unique_viewers&#x60;, &#x60;video_startup_failure_percentage&#x60;, &#x60;view_dropped_percentage&#x60;, and &#x60;views&#x60;.  Example:    * &#x60;metric_filters[]&#x3D;aggregate_startup_time&gt;&#x3D;1000&#x60;  | 
 **measurement** | **optional.String**| Measurement for the provided metric. If omitted, the default for the metric will be used. The default measurement for each metric is: \&quot;sum\&quot; : &#x60;ad_attempt_count&#x60;, &#x60;ad_break_count&#x60;, &#x60;ad_break_error_count&#x60;, &#x60;ad_error_count&#x60;, &#x60;ad_impression_count&#x60;, &#x60;playing_time&#x60; \&quot;median\&quot; : &#x60;ad_preroll_startup_time&#x60;, &#x60;aggregate_startup_time&#x60;, &#x60;content_startup_time&#x60;, &#x60;max_downscale_percentage&#x60;, &#x60;max_upscale_percentage&#x60;, &#x60;page_load_time&#x60;, &#x60;player_average_live_latency&#x60;, &#x60;player_startup_time&#x60;, &#x60;rebuffer_count&#x60;, &#x60;rebuffer_duration&#x60;, &#x60;requests_for_first_preroll&#x60;, &#x60;video_startup_preroll_load_time&#x60;, &#x60;video_startup_preroll_request_time&#x60;, &#x60;video_startup_time&#x60;, &#x60;view_average_request_latency&#x60;, &#x60;view_average_request_throughput&#x60;, &#x60;view_max_request_latency&#x60;, &#x60;weighted_average_bitrate&#x60; \&quot;avg\&quot; : &#x60;ad_break_error_percentage&#x60;, &#x60;ad_error_percentage&#x60;, &#x60;ad_exit_before_start_count&#x60;, &#x60;ad_exit_before_start_percentage&#x60;, &#x60;ad_playback_failure_percentage&#x60;, &#x60;ad_startup_error_count&#x60;, &#x60;ad_startup_error_percentage&#x60;, &#x60;content_playback_failure_percentage&#x60;, &#x60;downscale_percentage&#x60;, &#x60;exits_before_video_start&#x60;, &#x60;playback_business_exception_percentage&#x60;, &#x60;playback_failure_percentage&#x60;, &#x60;playback_success_score&#x60;, &#x60;rebuffer_frequency&#x60;, &#x60;rebuffer_percentage&#x60;, &#x60;seek_latency&#x60;, &#x60;smoothness_score&#x60;, &#x60;startup_time_score&#x60;, &#x60;upscale_percentage&#x60;, &#x60;video_quality_score&#x60;, &#x60;video_startup_business_exception_percentage&#x60;, &#x60;video_startup_failure_percentage&#x60;, &#x60;view_dropped_percentage&#x60;, &#x60;viewer_experience_score&#x60; \&quot;count\&quot; : &#x60;started_views&#x60;, &#x60;unique_viewers&#x60;  | 
 **orderDirection** | **optional.String**| Sort order. | 
 **groupBy** | **optional.String**| Time granularity to group results by. If this value is omitted, a default granularity is chosen based on the timeframe.  For timeframes of less than 90 minutes, the default granularity is &#x60;minute&#x60;. Between 90 minutes and 6 hours, the default granularity is &#x60;ten_minutes&#x60;. Between 6 hours and 15 days inclusive, the default granularity is &#x60;hour&#x60;. The granularity of timeframes that exceed 15 days is &#x60;day&#x60;. This default behavior is subject to change; it is strongly suggested that you explicitly specify the granularity.  | 

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

 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;).  Accepted formats are...    * array of epoch timestamps e.g. &#x60;timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600&#x60;   * duration string e.g. &#x60;timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days&#x60;  | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter results using key:value pairs. Must be provided as an array query string parameter.  **Basic filtering:** * &#x60;filters[]&#x3D;dimension:value&#x60; - Include rows where dimension equals value * &#x60;filters[]&#x3D;!dimension:value&#x60; - Exclude rows where dimension equals value  **For trace dimensions (like video_cdn_trace):** * &#x60;filters[]&#x3D;+dimension:value&#x60; - Include rows where trace contains value * &#x60;filters[]&#x3D;-dimension:value&#x60; - Exclude rows where trace contains value * &#x60;filters[]&#x3D;dimension:[value1,value2]&#x60; - Exact trace match  **Examples:** * &#x60;filters[]&#x3D;country:US&#x60; - US views only * &#x60;filters[]&#x3D;+video_cdn_trace:fastly&#x60; - Views using Fastly CDN  | 
 **metricFilters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match inequality conditions from provided metric comparison clauses. Must be provided as an array query string parameter.  Possible filterable metrics are the same as the set of metric ids, with the exceptions of &#x60;exits_before_video_start&#x60;, &#x60;unique_viewers&#x60;, &#x60;video_startup_failure_percentage&#x60;, &#x60;view_dropped_percentage&#x60;, and &#x60;views&#x60;.  Example:    * &#x60;metric_filters[]&#x3D;aggregate_startup_time&gt;&#x3D;1000&#x60;  | 
 **measurement** | **optional.String**| Measurement for the provided metric. If omitted, the default for the metric will be used. The default measurement for each metric is: \&quot;sum\&quot; : &#x60;ad_attempt_count&#x60;, &#x60;ad_break_count&#x60;, &#x60;ad_break_error_count&#x60;, &#x60;ad_error_count&#x60;, &#x60;ad_impression_count&#x60;, &#x60;playing_time&#x60; \&quot;median\&quot; : &#x60;ad_preroll_startup_time&#x60;, &#x60;aggregate_startup_time&#x60;, &#x60;content_startup_time&#x60;, &#x60;max_downscale_percentage&#x60;, &#x60;max_upscale_percentage&#x60;, &#x60;page_load_time&#x60;, &#x60;player_average_live_latency&#x60;, &#x60;player_startup_time&#x60;, &#x60;rebuffer_count&#x60;, &#x60;rebuffer_duration&#x60;, &#x60;requests_for_first_preroll&#x60;, &#x60;video_startup_preroll_load_time&#x60;, &#x60;video_startup_preroll_request_time&#x60;, &#x60;video_startup_time&#x60;, &#x60;view_average_request_latency&#x60;, &#x60;view_average_request_throughput&#x60;, &#x60;view_max_request_latency&#x60;, &#x60;weighted_average_bitrate&#x60; \&quot;avg\&quot; : &#x60;ad_break_error_percentage&#x60;, &#x60;ad_error_percentage&#x60;, &#x60;ad_exit_before_start_count&#x60;, &#x60;ad_exit_before_start_percentage&#x60;, &#x60;ad_playback_failure_percentage&#x60;, &#x60;ad_startup_error_count&#x60;, &#x60;ad_startup_error_percentage&#x60;, &#x60;content_playback_failure_percentage&#x60;, &#x60;downscale_percentage&#x60;, &#x60;exits_before_video_start&#x60;, &#x60;playback_business_exception_percentage&#x60;, &#x60;playback_failure_percentage&#x60;, &#x60;playback_success_score&#x60;, &#x60;rebuffer_frequency&#x60;, &#x60;rebuffer_percentage&#x60;, &#x60;seek_latency&#x60;, &#x60;smoothness_score&#x60;, &#x60;startup_time_score&#x60;, &#x60;upscale_percentage&#x60;, &#x60;video_quality_score&#x60;, &#x60;video_startup_business_exception_percentage&#x60;, &#x60;video_startup_failure_percentage&#x60;, &#x60;view_dropped_percentage&#x60;, &#x60;viewer_experience_score&#x60; \&quot;count\&quot; : &#x60;started_views&#x60;, &#x60;unique_viewers&#x60;  | 

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

List all of the values across every breakdown for a specific metric.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAllMetricValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAllMetricValuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;).  Accepted formats are...    * array of epoch timestamps e.g. &#x60;timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600&#x60;   * duration string e.g. &#x60;timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days&#x60;  | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter results using key:value pairs. Must be provided as an array query string parameter.  **Basic filtering:** * &#x60;filters[]&#x3D;dimension:value&#x60; - Include rows where dimension equals value * &#x60;filters[]&#x3D;!dimension:value&#x60; - Exclude rows where dimension equals value  **For trace dimensions (like video_cdn_trace):** * &#x60;filters[]&#x3D;+dimension:value&#x60; - Include rows where trace contains value * &#x60;filters[]&#x3D;-dimension:value&#x60; - Exclude rows where trace contains value * &#x60;filters[]&#x3D;dimension:[value1,value2]&#x60; - Exact trace match  **Examples:** * &#x60;filters[]&#x3D;country:US&#x60; - US views only * &#x60;filters[]&#x3D;+video_cdn_trace:fastly&#x60; - Views using Fastly CDN  | 
 **metricFilters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match inequality conditions from provided metric comparison clauses. Must be provided as an array query string parameter.  Possible filterable metrics are the same as the set of metric ids, with the exceptions of &#x60;exits_before_video_start&#x60;, &#x60;unique_viewers&#x60;, &#x60;video_startup_failure_percentage&#x60;, &#x60;view_dropped_percentage&#x60;, and &#x60;views&#x60;.  Example:    * &#x60;metric_filters[]&#x3D;aggregate_startup_time&gt;&#x3D;1000&#x60;  | 
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

List the breakdown values for a specific metric.

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
 **measurement** | **optional.String**| Measurement for the provided metric. If omitted, the default for the metric will be used. The default measurement for each metric is: \&quot;sum\&quot; : &#x60;ad_attempt_count&#x60;, &#x60;ad_break_count&#x60;, &#x60;ad_break_error_count&#x60;, &#x60;ad_error_count&#x60;, &#x60;ad_impression_count&#x60;, &#x60;playing_time&#x60; \&quot;median\&quot; : &#x60;ad_preroll_startup_time&#x60;, &#x60;aggregate_startup_time&#x60;, &#x60;content_startup_time&#x60;, &#x60;max_downscale_percentage&#x60;, &#x60;max_upscale_percentage&#x60;, &#x60;page_load_time&#x60;, &#x60;player_average_live_latency&#x60;, &#x60;player_startup_time&#x60;, &#x60;rebuffer_count&#x60;, &#x60;rebuffer_duration&#x60;, &#x60;requests_for_first_preroll&#x60;, &#x60;video_startup_preroll_load_time&#x60;, &#x60;video_startup_preroll_request_time&#x60;, &#x60;video_startup_time&#x60;, &#x60;view_average_request_latency&#x60;, &#x60;view_average_request_throughput&#x60;, &#x60;view_max_request_latency&#x60;, &#x60;weighted_average_bitrate&#x60; \&quot;avg\&quot; : &#x60;ad_break_error_percentage&#x60;, &#x60;ad_error_percentage&#x60;, &#x60;ad_exit_before_start_count&#x60;, &#x60;ad_exit_before_start_percentage&#x60;, &#x60;ad_playback_failure_percentage&#x60;, &#x60;ad_startup_error_count&#x60;, &#x60;ad_startup_error_percentage&#x60;, &#x60;content_playback_failure_percentage&#x60;, &#x60;downscale_percentage&#x60;, &#x60;exits_before_video_start&#x60;, &#x60;playback_business_exception_percentage&#x60;, &#x60;playback_failure_percentage&#x60;, &#x60;playback_success_score&#x60;, &#x60;rebuffer_frequency&#x60;, &#x60;rebuffer_percentage&#x60;, &#x60;seek_latency&#x60;, &#x60;smoothness_score&#x60;, &#x60;startup_time_score&#x60;, &#x60;upscale_percentage&#x60;, &#x60;video_quality_score&#x60;, &#x60;video_startup_business_exception_percentage&#x60;, &#x60;video_startup_failure_percentage&#x60;, &#x60;view_dropped_percentage&#x60;, &#x60;viewer_experience_score&#x60; \&quot;count\&quot; : &#x60;started_views&#x60;, &#x60;unique_viewers&#x60;  | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter results using key:value pairs. Must be provided as an array query string parameter.  **Basic filtering:** * &#x60;filters[]&#x3D;dimension:value&#x60; - Include rows where dimension equals value * &#x60;filters[]&#x3D;!dimension:value&#x60; - Exclude rows where dimension equals value  **For trace dimensions (like video_cdn_trace):** * &#x60;filters[]&#x3D;+dimension:value&#x60; - Include rows where trace contains value * &#x60;filters[]&#x3D;-dimension:value&#x60; - Exclude rows where trace contains value * &#x60;filters[]&#x3D;dimension:[value1,value2]&#x60; - Exact trace match  **Examples:** * &#x60;filters[]&#x3D;country:US&#x60; - US views only * &#x60;filters[]&#x3D;+video_cdn_trace:fastly&#x60; - Views using Fastly CDN  | 
 **metricFilters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match inequality conditions from provided metric comparison clauses. Must be provided as an array query string parameter.  Possible filterable metrics are the same as the set of metric ids, with the exceptions of &#x60;exits_before_video_start&#x60;, &#x60;unique_viewers&#x60;, &#x60;video_startup_failure_percentage&#x60;, &#x60;view_dropped_percentage&#x60;, and &#x60;views&#x60;.  Example:    * &#x60;metric_filters[]&#x3D;aggregate_startup_time&gt;&#x3D;1000&#x60;  | 
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **orderBy** | **optional.String**| Value to order the results by | 
 **orderDirection** | **optional.String**| Sort order. | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;).  Accepted formats are...    * array of epoch timestamps e.g. &#x60;timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600&#x60;   * duration string e.g. &#x60;timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days&#x60;  | 

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

 **measurement** | **optional.String**| Measurement for the provided metric. If omitted, the default for the metric will be used. The default measurement for each metric is: \&quot;sum\&quot; : &#x60;ad_attempt_count&#x60;, &#x60;ad_break_count&#x60;, &#x60;ad_break_error_count&#x60;, &#x60;ad_error_count&#x60;, &#x60;ad_impression_count&#x60;, &#x60;playing_time&#x60; \&quot;median\&quot; : &#x60;ad_preroll_startup_time&#x60;, &#x60;aggregate_startup_time&#x60;, &#x60;content_startup_time&#x60;, &#x60;max_downscale_percentage&#x60;, &#x60;max_upscale_percentage&#x60;, &#x60;page_load_time&#x60;, &#x60;player_average_live_latency&#x60;, &#x60;player_startup_time&#x60;, &#x60;rebuffer_count&#x60;, &#x60;rebuffer_duration&#x60;, &#x60;requests_for_first_preroll&#x60;, &#x60;video_startup_preroll_load_time&#x60;, &#x60;video_startup_preroll_request_time&#x60;, &#x60;video_startup_time&#x60;, &#x60;view_average_request_latency&#x60;, &#x60;view_average_request_throughput&#x60;, &#x60;view_max_request_latency&#x60;, &#x60;weighted_average_bitrate&#x60; \&quot;avg\&quot; : &#x60;ad_break_error_percentage&#x60;, &#x60;ad_error_percentage&#x60;, &#x60;ad_exit_before_start_count&#x60;, &#x60;ad_exit_before_start_percentage&#x60;, &#x60;ad_playback_failure_percentage&#x60;, &#x60;ad_startup_error_count&#x60;, &#x60;ad_startup_error_percentage&#x60;, &#x60;content_playback_failure_percentage&#x60;, &#x60;downscale_percentage&#x60;, &#x60;exits_before_video_start&#x60;, &#x60;playback_business_exception_percentage&#x60;, &#x60;playback_failure_percentage&#x60;, &#x60;playback_success_score&#x60;, &#x60;rebuffer_frequency&#x60;, &#x60;rebuffer_percentage&#x60;, &#x60;seek_latency&#x60;, &#x60;smoothness_score&#x60;, &#x60;startup_time_score&#x60;, &#x60;upscale_percentage&#x60;, &#x60;video_quality_score&#x60;, &#x60;video_startup_business_exception_percentage&#x60;, &#x60;video_startup_failure_percentage&#x60;, &#x60;view_dropped_percentage&#x60;, &#x60;viewer_experience_score&#x60; \&quot;count\&quot; : &#x60;started_views&#x60;, &#x60;unique_viewers&#x60;  | 
 **orderDirection** | **optional.String**| Sort order. | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;).  Accepted formats are...    * array of epoch timestamps e.g. &#x60;timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600&#x60;   * duration string e.g. &#x60;timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days&#x60;  | 
 **filters** | [**optional.Interface of []string**](string.md)| Filter results using key:value pairs. Must be provided as an array query string parameter.  **Basic filtering:** * &#x60;filters[]&#x3D;dimension:value&#x60; - Include rows where dimension equals value * &#x60;filters[]&#x3D;!dimension:value&#x60; - Exclude rows where dimension equals value  **For trace dimensions (like video_cdn_trace):** * &#x60;filters[]&#x3D;+dimension:value&#x60; - Include rows where trace contains value * &#x60;filters[]&#x3D;-dimension:value&#x60; - Exclude rows where trace contains value * &#x60;filters[]&#x3D;dimension:[value1,value2]&#x60; - Exact trace match  **Examples:** * &#x60;filters[]&#x3D;country:US&#x60; - US views only * &#x60;filters[]&#x3D;+video_cdn_trace:fastly&#x60; - Views using Fastly CDN  | 
 **metricFilters** | [**optional.Interface of []string**](string.md)| Limit the results to rows that match inequality conditions from provided metric comparison clauses. Must be provided as an array query string parameter.  Possible filterable metrics are the same as the set of metric ids, with the exceptions of &#x60;exits_before_video_start&#x60;, &#x60;unique_viewers&#x60;, &#x60;video_startup_failure_percentage&#x60;, &#x60;view_dropped_percentage&#x60;, and &#x60;views&#x60;.  Example:    * &#x60;metric_filters[]&#x3D;aggregate_startup_time&gt;&#x3D;1000&#x60;  | 

### Return type

[**ListInsightsResponse**](ListInsightsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

