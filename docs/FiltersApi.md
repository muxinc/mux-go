# \FiltersApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFilterValues**](FiltersApi.md#ListFilterValues) | **Get** /data/v1/filters/{FILTER_ID} | Lists values for a specific filter
[**ListFilters**](FiltersApi.md#ListFilters) | **Get** /data/v1/filters | List Filters



## ListFilterValues

> ListFilterValuesResponse ListFilterValues(ctx, fILTERID).Limit(limit).Page(page).Filters(filters).Timeframe(timeframe).Execute()

Lists values for a specific filter



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
    fILTERID := "abcd1234" // string | ID of the Filter
    limit := int32(56) // int32 | Number of items to include in the response (optional) (default to 25)
    page := int32(56) // int32 | Offset by this many pages, of the size of `limit` (optional) (default to 1)
    filters := []string{"Inner_example"} // []string | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US).  Possible filter names are the same as returned by the List Filters endpoint.  (optional)
    timeframe := []string{"Inner_example"} // []string | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]=). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]=1498867200&timeframe[]=1498953600    * duration string e.g. timeframe[]=24:hours or timeframe[]=7:days.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FiltersApi.ListFilterValues(context.Background(), fILTERID).Limit(limit).Page(page).Filters(filters).Timeframe(timeframe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersApi.ListFilterValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFilterValues`: ListFilterValuesResponse
    fmt.Fprintf(os.Stdout, "Response from `FiltersApi.ListFilterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fILTERID** | **string** | ID of the Filter | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFilterValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of items to include in the response | [default to 25]
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **filters** | **[]string** | Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US).  Possible filter names are the same as returned by the List Filters endpoint.  | 
 **timeframe** | **[]string** | Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600    * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 

### Return type

[**ListFilterValuesResponse**](ListFilterValuesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFilters

> ListFiltersResponse ListFilters(ctx).Execute()

List Filters



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
    resp, r, err := api_client.FiltersApi.ListFilters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersApi.ListFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFilters`: ListFiltersResponse
    fmt.Fprintf(os.Stdout, "Response from `FiltersApi.ListFilters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFiltersRequest struct via the builder pattern


### Return type

[**ListFiltersResponse**](ListFiltersResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

