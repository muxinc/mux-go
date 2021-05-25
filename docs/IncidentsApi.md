# \IncidentsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIncident**](IncidentsApi.md#GetIncident) | **Get** /data/v1/incidents/{INCIDENT_ID} | Get an Incident
[**ListIncidents**](IncidentsApi.md#ListIncidents) | **Get** /data/v1/incidents | List Incidents
[**ListRelatedIncidents**](IncidentsApi.md#ListRelatedIncidents) | **Get** /data/v1/incidents/{INCIDENT_ID}/related | List Related Incidents



## GetIncident

> IncidentResponse GetIncident(ctx, iNCIDENTID).Execute()

Get an Incident



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
    iNCIDENTID := "abcd1234" // string | ID of the Incident

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.GetIncident(context.Background(), iNCIDENTID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncident`: IncidentResponse
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iNCIDENTID** | **string** | ID of the Incident | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IncidentResponse**](IncidentResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncidents

> ListIncidentsResponse ListIncidents(ctx).Limit(limit).Page(page).OrderBy(orderBy).OrderDirection(orderDirection).Status(status).Severity(severity).Execute()

List Incidents



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
    limit := int32(56) // int32 | Number of items to include in the response (optional) (default to 25)
    page := int32(56) // int32 | Offset by this many pages, of the size of `limit` (optional) (default to 1)
    orderBy := "orderBy_example" // string | Value to order the results by (optional)
    orderDirection := "orderDirection_example" // string | Sort order. (optional)
    status := "status_example" // string | Status to filter incidents by (optional)
    severity := "severity_example" // string | Severity to filter incidents by (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.ListIncidents(context.Background()).Limit(limit).Page(page).OrderBy(orderBy).OrderDirection(orderDirection).Status(status).Severity(severity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.ListIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncidents`: ListIncidentsResponse
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.ListIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of items to include in the response | [default to 25]
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **orderBy** | **string** | Value to order the results by | 
 **orderDirection** | **string** | Sort order. | 
 **status** | **string** | Status to filter incidents by | 
 **severity** | **string** | Severity to filter incidents by | 

### Return type

[**ListIncidentsResponse**](ListIncidentsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRelatedIncidents

> ListRelatedIncidentsResponse ListRelatedIncidents(ctx, iNCIDENTID).Limit(limit).Page(page).OrderBy(orderBy).OrderDirection(orderDirection).Execute()

List Related Incidents



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
    iNCIDENTID := "abcd1234" // string | ID of the Incident
    limit := int32(56) // int32 | Number of items to include in the response (optional) (default to 25)
    page := int32(56) // int32 | Offset by this many pages, of the size of `limit` (optional) (default to 1)
    orderBy := "orderBy_example" // string | Value to order the results by (optional)
    orderDirection := "orderDirection_example" // string | Sort order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.ListRelatedIncidents(context.Background(), iNCIDENTID).Limit(limit).Page(page).OrderBy(orderBy).OrderDirection(orderDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.ListRelatedIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRelatedIncidents`: ListRelatedIncidentsResponse
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.ListRelatedIncidents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iNCIDENTID** | **string** | ID of the Incident | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRelatedIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of items to include in the response | [default to 25]
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **orderBy** | **string** | Value to order the results by | 
 **orderDirection** | **string** | Sort order. | 

### Return type

[**ListRelatedIncidentsResponse**](ListRelatedIncidentsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

