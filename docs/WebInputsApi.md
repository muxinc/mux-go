# \WebInputsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebInput**](WebInputsApi.md#CreateWebInput) | **Post** /video/v1/web-inputs | Create a new Web Input
[**DeleteWebInput**](WebInputsApi.md#DeleteWebInput) | **Delete** /video/v1/web-inputs/{WEB_INPUT_ID} | Delete a Web Input
[**GetWebInput**](WebInputsApi.md#GetWebInput) | **Get** /video/v1/web-inputs/{WEB_INPUT_ID} | Retrieve a Web Input
[**LaunchWebInput**](WebInputsApi.md#LaunchWebInput) | **Put** /video/v1/web-inputs/{WEB_INPUT_ID}/launch | Launch a Web Input
[**ListWebInputs**](WebInputsApi.md#ListWebInputs) | **Get** /video/v1/web-inputs | List Web Inputs
[**ReloadWebInput**](WebInputsApi.md#ReloadWebInput) | **Put** /video/v1/web-inputs/{WEB_INPUT_ID}/reload | Reload a Web Input
[**ShutdownWebInput**](WebInputsApi.md#ShutdownWebInput) | **Put** /video/v1/web-inputs/{WEB_INPUT_ID}/shutdown | Shut down a Web Input
[**UpdateWebInputUrl**](WebInputsApi.md#UpdateWebInputUrl) | **Put** /video/v1/web-inputs/{WEB_INPUT_ID}/url | Update Web Input URL


# **CreateWebInput**
> WebInputResponse CreateWebInput(ctx, createWebInputRequest)
Create a new Web Input

Create a new Web Input

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createWebInputRequest** | [**CreateWebInputRequest**](CreateWebInputRequest.md)|  | 

### Return type

[**WebInputResponse**](WebInputResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebInput**
> DeleteWebInput(ctx, wEBINPUTID)
Delete a Web Input

Deletes a Web Input and all its data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **wEBINPUTID** | **string**| The Web Input ID | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebInput**
> WebInputResponse GetWebInput(ctx, wEBINPUTID)
Retrieve a Web Input

Retrieve a single Web Input's info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **wEBINPUTID** | **string**| The Web Input ID | 

### Return type

[**WebInputResponse**](WebInputResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LaunchWebInput**
> LaunchWebInputResponse LaunchWebInput(ctx, wEBINPUTID)
Launch a Web Input

Launches the browsers instance, loads the URL specified, and then starts streaming to the specified Live Stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **wEBINPUTID** | **string**| The Web Input ID | 

### Return type

[**LaunchWebInputResponse**](LaunchWebInputResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebInputs**
> ListWebInputsResponse ListWebInputs(ctx, optional)
List Web Inputs

List Web Inputs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListWebInputsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListWebInputsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]

### Return type

[**ListWebInputsResponse**](ListWebInputsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReloadWebInput**
> ReloadWebInputResponse ReloadWebInput(ctx, wEBINPUTID)
Reload a Web Input

Reloads the page that a Web Input is displaying.  Note: Using this when the Web Input is streaming will display the page reloading. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **wEBINPUTID** | **string**| The Web Input ID | 

### Return type

[**ReloadWebInputResponse**](ReloadWebInputResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShutdownWebInput**
> ShutdownWebInputResponse ShutdownWebInput(ctx, wEBINPUTID)
Shut down a Web Input

Ends streaming to the specified Live Stream, and then shuts down the Web Input browser instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **wEBINPUTID** | **string**| The Web Input ID | 

### Return type

[**ShutdownWebInputResponse**](ShutdownWebInputResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebInputUrl**
> WebInputResponse UpdateWebInputUrl(ctx, wEBINPUTID, updateWebInputUrlRequest)
Update Web Input URL

Changes the URL that a Web Input loads when it launches.  Note: This can only be called when the Web Input is idle. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **wEBINPUTID** | **string**| The Web Input ID | 
  **updateWebInputUrlRequest** | [**UpdateWebInputUrlRequest**](UpdateWebInputUrlRequest.md)|  | 

### Return type

[**WebInputResponse**](WebInputResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

