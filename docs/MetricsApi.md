# \MetricsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetricTimeseriesData**](MetricsApi.md#GetMetricTimeseriesData) | **Get** /data/v1/metrics/{METRIC_ID}/timeseries | Get metric timeseries data
[**GetOverallValues**](MetricsApi.md#GetOverallValues) | **Get** /data/v1/metrics/{METRIC_ID}/overall | Get Overall values
[**ListAllMetricValues**](MetricsApi.md#ListAllMetricValues) | **Get** /data/v1/metrics/comparison | List all metric values
[**ListBreakdownValues**](MetricsApi.md#ListBreakdownValues) | **Get** /data/v1/metrics/{METRIC_ID}/breakdown | List breakdown values
[**ListInsights**](MetricsApi.md#ListInsights) | **Get** /data/v1/metrics/{METRIC_ID}/insights | List Insights



## GetMetricTimeseriesData

> GetMetricTimeseriesDataResponse GetMetricTimeseriesData(ctx, mETRICID).Timeframe(timeframe).Filters(filters).Measurement(measurement).OrderDirection(orderDirection).GroupBy(groupBy).Execute()

Get metric timeseries data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mETRICID := "video_startup_time" // string | ID of the Metric
    timeframe := []string{"Inner_example"} // []string | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]=). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]=1498867200&timeframe[]=1498953600   * duration string e.g. timeframe[]=24:hours or timeframe[]=7:days.  (optional)
    filters := []string{"Inner_example"} // []string | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US). Possible filter names are the same as returned by the List Filters endpoint.  (optional)
    measurement := "measurement_example" // string | Measurement for the provided metric. If omitted, the deafult for the metric will be used. (optional)
    orderDirection := "orderDirection_example" // string | Sort order. (optional)
    groupBy := "groupBy_example" // string | Time granularity to group results by. If this value is omitted, a default granularity is chosen based on the supplied timeframe. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetMetricTimeseriesData(context.Background(), mETRICID).Timeframe(timeframe).Filters(filters).Measurement(measurement).OrderDirection(orderDirection).GroupBy(groupBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetricTimeseriesData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricTimeseriesData`: GetMetricTimeseriesDataResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetMetricTimeseriesData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mETRICID** | **string** | ID of the Metric | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricTimeseriesDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeframe** | **[]string** | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600   * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 
 **filters** | **[]string** | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 
 **measurement** | **string** | Measurement for the provided metric. If omitted, the deafult for the metric will be used. | 
 **orderDirection** | **string** | Sort order. | 
 **groupBy** | **string** | Time granularity to group results by. If this value is omitted, a default granularity is chosen based on the supplied timeframe. | 

### Return type

[**GetMetricTimeseriesDataResponse**](GetMetricTimeseriesDataResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverallValues

> GetOverallValuesResponse GetOverallValues(ctx, mETRICID).Timeframe(timeframe).Filters(filters).Measurement(measurement).Execute()

Get Overall values



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mETRICID := "video_startup_time" // string | ID of the Metric
    timeframe := []string{"Inner_example"} // []string | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]=). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]=1498867200&timeframe[]=1498953600   * duration string e.g. timeframe[]=24:hours or timeframe[]=7:days.  (optional)
    filters := []string{"Inner_example"} // []string | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US). Possible filter names are the same as returned by the List Filters endpoint.  (optional)
    measurement := "measurement_example" // string | Measurement for the provided metric. If omitted, the deafult for the metric will be used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetOverallValues(context.Background(), mETRICID).Timeframe(timeframe).Filters(filters).Measurement(measurement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetOverallValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOverallValues`: GetOverallValuesResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetOverallValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mETRICID** | **string** | ID of the Metric | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOverallValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeframe** | **[]string** | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600   * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 
 **filters** | **[]string** | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 
 **measurement** | **string** | Measurement for the provided metric. If omitted, the deafult for the metric will be used. | 

### Return type

[**GetOverallValuesResponse**](GetOverallValuesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllMetricValues

> ListAllMetricValuesResponse ListAllMetricValues(ctx).Timeframe(timeframe).Filters(filters).Dimension(dimension).Value(value).Execute()

List all metric values



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    timeframe := []string{"Inner_example"} // []string | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]=). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]=1498867200&timeframe[]=1498953600   * duration string e.g. timeframe[]=24:hours or timeframe[]=7:days.  (optional)
    filters := []string{"Inner_example"} // []string | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US). Possible filter names are the same as returned by the List Filters endpoint.  (optional)
    dimension := "dimension_example" // string | Dimension the specified value belongs to (optional)
    value := "value_example" // string | Value to show all available metrics for (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.ListAllMetricValues(context.Background()).Timeframe(timeframe).Filters(filters).Dimension(dimension).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListAllMetricValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllMetricValues`: ListAllMetricValuesResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.ListAllMetricValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllMetricValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeframe** | **[]string** | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600   * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 
 **filters** | **[]string** | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 
 **dimension** | **string** | Dimension the specified value belongs to | 
 **value** | **string** | Value to show all available metrics for | 

### Return type

[**ListAllMetricValuesResponse**](ListAllMetricValuesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBreakdownValues

> ListBreakdownValuesResponse ListBreakdownValues(ctx, mETRICID).GroupBy(groupBy).Measurement(measurement).Filters(filters).Limit(limit).Page(page).OrderBy(orderBy).OrderDirection(orderDirection).Timeframe(timeframe).Execute()

List breakdown values



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mETRICID := "video_startup_time" // string | ID of the Metric
    groupBy := "groupBy_example" // string | Breakdown value to group the results by (optional)
    measurement := "measurement_example" // string | Measurement for the provided metric. If omitted, the deafult for the metric will be used. (optional)
    filters := []string{"Inner_example"} // []string | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US). Possible filter names are the same as returned by the List Filters endpoint.  (optional)
    limit := int32(56) // int32 | Number of items to include in the response (optional) (default to 25)
    page := int32(56) // int32 | Offset by this many pages, of the size of `limit` (optional) (default to 1)
    orderBy := "orderBy_example" // string | Value to order the results by (optional)
    orderDirection := "orderDirection_example" // string | Sort order. (optional)
    timeframe := []string{"Inner_example"} // []string | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]=). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]=1498867200&timeframe[]=1498953600   * duration string e.g. timeframe[]=24:hours or timeframe[]=7:days.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.ListBreakdownValues(context.Background(), mETRICID).GroupBy(groupBy).Measurement(measurement).Filters(filters).Limit(limit).Page(page).OrderBy(orderBy).OrderDirection(orderDirection).Timeframe(timeframe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListBreakdownValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBreakdownValues`: ListBreakdownValuesResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.ListBreakdownValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mETRICID** | **string** | ID of the Metric | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBreakdownValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBy** | **string** | Breakdown value to group the results by | 
 **measurement** | **string** | Measurement for the provided metric. If omitted, the deafult for the metric will be used. | 
 **filters** | **[]string** | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 
 **limit** | **int32** | Number of items to include in the response | [default to 25]
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **orderBy** | **string** | Value to order the results by | 
 **orderDirection** | **string** | Sort order. | 
 **timeframe** | **[]string** | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600   * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 

### Return type

[**ListBreakdownValuesResponse**](ListBreakdownValuesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInsights

> ListInsightsResponse ListInsights(ctx, mETRICID).Measurement(measurement).OrderDirection(orderDirection).Timeframe(timeframe).Execute()

List Insights



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mETRICID := "video_startup_time" // string | ID of the Metric
    measurement := "measurement_example" // string | Measurement for the provided metric. If omitted, the deafult for the metric will be used. (optional)
    orderDirection := "orderDirection_example" // string | Sort order. (optional)
    timeframe := []string{"Inner_example"} // []string | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]=). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]=1498867200&timeframe[]=1498953600   * duration string e.g. timeframe[]=24:hours or timeframe[]=7:days.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.ListInsights(context.Background(), mETRICID).Measurement(measurement).OrderDirection(orderDirection).Timeframe(timeframe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInsights`: ListInsightsResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.ListInsights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mETRICID** | **string** | ID of the Metric | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **measurement** | **string** | Measurement for the provided metric. If omitted, the deafult for the metric will be used. | 
 **orderDirection** | **string** | Sort order. | 
 **timeframe** | **[]string** | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600   * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 

### Return type

[**ListInsightsResponse**](ListInsightsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

