# \RealTimeApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRealtimeBreakdown**](RealTimeApi.md#GetRealtimeBreakdown) | **Get** /data/v1/realtime/metrics/{REALTIME_METRIC_ID}/breakdown | Get Real-Time Breakdown
[**GetRealtimeHistogramTimeseries**](RealTimeApi.md#GetRealtimeHistogramTimeseries) | **Get** /data/v1/realtime/metrics/{REALTIME_HISTOGRAM_METRIC_ID}/histogram-timeseries | Get Real-Time Histogram Timeseries
[**GetRealtimeTimeseries**](RealTimeApi.md#GetRealtimeTimeseries) | **Get** /data/v1/realtime/metrics/{REALTIME_METRIC_ID}/timeseries | Get Real-Time Timeseries
[**ListRealtimeDimensions**](RealTimeApi.md#ListRealtimeDimensions) | **Get** /data/v1/realtime/dimensions | List Real-Time Dimensions
[**ListRealtimeMetrics**](RealTimeApi.md#ListRealtimeMetrics) | **Get** /data/v1/realtime/metrics | List Real-Time Metrics



## GetRealtimeBreakdown

> GetRealTimeBreakdownResponse GetRealtimeBreakdown(ctx, rEALTIMEMETRICID).Dimension(dimension).Timestamp(timestamp).Filters(filters).OrderBy(orderBy).OrderDirection(orderDirection).Execute()

Get Real-Time Breakdown



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
    rEALTIMEMETRICID := "current-concurrent-viewers" // string | ID of the Realtime Metric
    dimension := "dimension_example" // string | Dimension the specified value belongs to (optional)
    timestamp := float32(8.14) // float32 | Timestamp to limit results by. This value must be provided as a unix timestamp. Defaults to the current unix timestamp. (optional)
    filters := []string{"Inner_example"} // []string | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US). Possible filter names are the same as returned by the List Filters endpoint.  (optional)
    orderBy := "orderBy_example" // string | Value to order the results by (optional)
    orderDirection := "orderDirection_example" // string | Sort order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RealTimeApi.GetRealtimeBreakdown(context.Background(), rEALTIMEMETRICID).Dimension(dimension).Timestamp(timestamp).Filters(filters).OrderBy(orderBy).OrderDirection(orderDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeApi.GetRealtimeBreakdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRealtimeBreakdown`: GetRealTimeBreakdownResponse
    fmt.Fprintf(os.Stdout, "Response from `RealTimeApi.GetRealtimeBreakdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rEALTIMEMETRICID** | **string** | ID of the Realtime Metric | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealtimeBreakdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dimension** | **string** | Dimension the specified value belongs to | 
 **timestamp** | **float32** | Timestamp to limit results by. This value must be provided as a unix timestamp. Defaults to the current unix timestamp. | 
 **filters** | **[]string** | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 
 **orderBy** | **string** | Value to order the results by | 
 **orderDirection** | **string** | Sort order. | 

### Return type

[**GetRealTimeBreakdownResponse**](GetRealTimeBreakdownResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealtimeHistogramTimeseries

> GetRealTimeHistogramTimeseriesResponse GetRealtimeHistogramTimeseries(ctx, rEALTIMEHISTOGRAMMETRICID).Filters(filters).Execute()

Get Real-Time Histogram Timeseries



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
    rEALTIMEHISTOGRAMMETRICID := "video-startup-time" // string | ID of the Realtime Histogram Metric
    filters := []string{"Inner_example"} // []string | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US). Possible filter names are the same as returned by the List Filters endpoint.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RealTimeApi.GetRealtimeHistogramTimeseries(context.Background(), rEALTIMEHISTOGRAMMETRICID).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeApi.GetRealtimeHistogramTimeseries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRealtimeHistogramTimeseries`: GetRealTimeHistogramTimeseriesResponse
    fmt.Fprintf(os.Stdout, "Response from `RealTimeApi.GetRealtimeHistogramTimeseries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rEALTIMEHISTOGRAMMETRICID** | **string** | ID of the Realtime Histogram Metric | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealtimeHistogramTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **[]string** | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 

### Return type

[**GetRealTimeHistogramTimeseriesResponse**](GetRealTimeHistogramTimeseriesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealtimeTimeseries

> GetRealTimeTimeseriesResponse GetRealtimeTimeseries(ctx, rEALTIMEMETRICID).Filters(filters).Execute()

Get Real-Time Timeseries



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
    rEALTIMEMETRICID := "current-concurrent-viewers" // string | ID of the Realtime Metric
    filters := []string{"Inner_example"} // []string | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US). Possible filter names are the same as returned by the List Filters endpoint.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RealTimeApi.GetRealtimeTimeseries(context.Background(), rEALTIMEMETRICID).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeApi.GetRealtimeTimeseries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRealtimeTimeseries`: GetRealTimeTimeseriesResponse
    fmt.Fprintf(os.Stdout, "Response from `RealTimeApi.GetRealtimeTimeseries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rEALTIMEMETRICID** | **string** | ID of the Realtime Metric | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealtimeTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **[]string** | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 

### Return type

[**GetRealTimeTimeseriesResponse**](GetRealTimeTimeseriesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRealtimeDimensions

> ListRealTimeDimensionsResponse ListRealtimeDimensions(ctx).Execute()

List Real-Time Dimensions



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RealTimeApi.ListRealtimeDimensions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeApi.ListRealtimeDimensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRealtimeDimensions`: ListRealTimeDimensionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RealTimeApi.ListRealtimeDimensions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRealtimeDimensionsRequest struct via the builder pattern


### Return type

[**ListRealTimeDimensionsResponse**](ListRealTimeDimensionsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRealtimeMetrics

> ListRealTimeMetricsResponse ListRealtimeMetrics(ctx).Execute()

List Real-Time Metrics



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RealTimeApi.ListRealtimeMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeApi.ListRealtimeMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRealtimeMetrics`: ListRealTimeMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `RealTimeApi.ListRealtimeMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRealtimeMetricsRequest struct via the builder pattern


### Return type

[**ListRealTimeMetricsResponse**](ListRealTimeMetricsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

