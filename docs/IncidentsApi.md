# \IncidentsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIncident**](IncidentsApi.md#GetIncident) | **Get** /data/v1/incidents/{INCIDENT_ID} | Get an Incident
[**ListIncidents**](IncidentsApi.md#ListIncidents) | **Get** /data/v1/incidents | List Incidents
[**ListRelatedIncidents**](IncidentsApi.md#ListRelatedIncidents) | **Get** /data/v1/incidents/{INCIDENT_ID}/related | List Related Incidents


# **GetIncident**
> IncidentResponse GetIncident(ctx, iNCIDENTID)
Get an Incident

Returns the details of an incident.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iNCIDENTID** | **string**| ID of the Incident | 

### Return type

[**IncidentResponse**](IncidentResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIncidents**
> ListIncidentsResponse ListIncidents(ctx, optional)
List Incidents

Returns a list of incidents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListIncidentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListIncidentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **orderBy** | **optional.String**| Value to order the results by | 
 **orderDirection** | **optional.String**| Sort order. | 
 **status** | **optional.String**| Status to filter incidents by | 
 **severity** | **optional.String**| Severity to filter incidents by | 

### Return type

[**ListIncidentsResponse**](ListIncidentsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRelatedIncidents**
> ListRelatedIncidentsResponse ListRelatedIncidents(ctx, iNCIDENTID, optional)
List Related Incidents

Returns all the incidents that seem related to a specific incident.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iNCIDENTID** | **string**| ID of the Incident | 
 **optional** | ***ListRelatedIncidentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListRelatedIncidentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **orderBy** | **optional.String**| Value to order the results by | 
 **orderDirection** | **optional.String**| Sort order. | 

### Return type

[**ListRelatedIncidentsResponse**](ListRelatedIncidentsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

