# \AnnotationsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnnotation**](AnnotationsApi.md#CreateAnnotation) | **Post** /data/v1/annotations | Create Annotation
[**DeleteAnnotation**](AnnotationsApi.md#DeleteAnnotation) | **Delete** /data/v1/annotations/{ANNOTATION_ID} | Delete Annotation
[**GetAnnotation**](AnnotationsApi.md#GetAnnotation) | **Get** /data/v1/annotations/{ANNOTATION_ID} | Get Annotation
[**ListAnnotations**](AnnotationsApi.md#ListAnnotations) | **Get** /data/v1/annotations | List Annotations
[**UpdateAnnotation**](AnnotationsApi.md#UpdateAnnotation) | **Patch** /data/v1/annotations/{ANNOTATION_ID} | Update Annotation


# **CreateAnnotation**
> AnnotationResponse CreateAnnotation(ctx, annotationInput)
Create Annotation

Creates a new annotation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **annotationInput** | [**AnnotationInput**](AnnotationInput.md)|  | 

### Return type

[**AnnotationResponse**](AnnotationResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAnnotation**
> DeleteAnnotation(ctx, aNNOTATIONID)
Delete Annotation

Deletes an annotation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aNNOTATIONID** | [**string**](.md)| The annotation ID | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnnotation**
> AnnotationResponse GetAnnotation(ctx, aNNOTATIONID)
Get Annotation

Returns the details of a specific annotation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aNNOTATIONID** | [**string**](.md)| The annotation ID | 

### Return type

[**AnnotationResponse**](AnnotationResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAnnotations**
> ListAnnotationsResponse ListAnnotations(ctx, optional)
List Annotations

Returns a list of annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAnnotationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAnnotationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **orderDirection** | **optional.String**| Sort order. | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;).  Accepted formats are...    * array of epoch timestamps e.g. &#x60;timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600&#x60;   * duration string e.g. &#x60;timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days&#x60;  | 

### Return type

[**ListAnnotationsResponse**](ListAnnotationsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAnnotation**
> AnnotationResponse UpdateAnnotation(ctx, aNNOTATIONID, annotationInput)
Update Annotation

Updates an existing annotation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aNNOTATIONID** | [**string**](.md)| The annotation ID | 
  **annotationInput** | [**AnnotationInput**](AnnotationInput.md)|  | 

### Return type

[**AnnotationResponse**](AnnotationResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

