# \URLSigningKeysApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUrlSigningKey**](URLSigningKeysApi.md#CreateUrlSigningKey) | **Post** /video/v1/signing-keys | Create a URL signing key
[**DeleteUrlSigningKey**](URLSigningKeysApi.md#DeleteUrlSigningKey) | **Delete** /video/v1/signing-keys/{SIGNING_KEY_ID} | Delete a URL signing key
[**GetUrlSigningKey**](URLSigningKeysApi.md#GetUrlSigningKey) | **Get** /video/v1/signing-keys/{SIGNING_KEY_ID} | Retrieve a URL signing key
[**ListUrlSigningKeys**](URLSigningKeysApi.md#ListUrlSigningKeys) | **Get** /video/v1/signing-keys | List URL signing keys


# **CreateUrlSigningKey**
> SigningKeyResponse CreateUrlSigningKey(ctx, )
Create a URL signing key

Creates a new signing key pair. When creating a new signing key, the API will generate a 2048-bit RSA key-pair and return the private key and a generated key-id; the public key will be stored at Mux to validate signed tokens.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SigningKeyResponse**](SigningKeyResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUrlSigningKey**
> DeleteUrlSigningKey(ctx, sIGNINGKEYID)
Delete a URL signing key

Deletes an existing signing key. Use with caution, as this will invalidate any existing signatures and no URLs can be signed using the key again.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sIGNINGKEYID** | **string**| The ID of the signing key. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUrlSigningKey**
> SigningKeyResponse GetUrlSigningKey(ctx, sIGNINGKEYID)
Retrieve a URL signing key

Retrieves the details of a URL signing key that has previously been created. Supply the unique signing key ID that was returned from your previous request, and Mux will return the corresponding signing key information. **The private key is not returned in this response.** 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sIGNINGKEYID** | **string**| The ID of the signing key. | 

### Return type

[**SigningKeyResponse**](SigningKeyResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUrlSigningKeys**
> ListSigningKeysResponse ListUrlSigningKeys(ctx, optional)
List URL signing keys

Returns a list of URL signing keys.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListUrlSigningKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListUrlSigningKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]

### Return type

[**ListSigningKeysResponse**](ListSigningKeysResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

