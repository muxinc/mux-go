# \TranscriptionVocabulariesApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranscriptionVocabulary**](TranscriptionVocabulariesApi.md#CreateTranscriptionVocabulary) | **Post** /video/v1/transcription-vocabularies | Create a Transcription Vocabulary
[**DeleteTranscriptionVocabulary**](TranscriptionVocabulariesApi.md#DeleteTranscriptionVocabulary) | **Delete** /video/v1/transcription-vocabularies/{TRANSCRIPTION_VOCABULARY_ID} | Delete a Transcription Vocabulary
[**GetTranscriptionVocabulary**](TranscriptionVocabulariesApi.md#GetTranscriptionVocabulary) | **Get** /video/v1/transcription-vocabularies/{TRANSCRIPTION_VOCABULARY_ID} | Retrieve a Transcription Vocabulary
[**ListTranscriptionVocabularies**](TranscriptionVocabulariesApi.md#ListTranscriptionVocabularies) | **Get** /video/v1/transcription-vocabularies | List Transcription Vocabularies
[**UpdateTranscriptionVocabulary**](TranscriptionVocabulariesApi.md#UpdateTranscriptionVocabulary) | **Put** /video/v1/transcription-vocabularies/{TRANSCRIPTION_VOCABULARY_ID} | Update a Transcription Vocabulary


# **CreateTranscriptionVocabulary**
> TranscriptionVocabularyResponse CreateTranscriptionVocabulary(ctx, createTranscriptionVocabularyRequest)
Create a Transcription Vocabulary

Create a new Transcription Vocabulary.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createTranscriptionVocabularyRequest** | [**CreateTranscriptionVocabularyRequest**](CreateTranscriptionVocabularyRequest.md)|  | 

### Return type

[**TranscriptionVocabularyResponse**](TranscriptionVocabularyResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTranscriptionVocabulary**
> DeleteTranscriptionVocabulary(ctx, tRANSCRIPTIONVOCABULARYID)
Delete a Transcription Vocabulary

Deletes a Transcription Vocabulary. The Transcription Vocabulary's ID will be disassociated from any live streams using it. Transcription Vocabularies can be deleted while associated live streams are active. However, the words and phrases in the deleted Transcription Vocabulary will remain attached to those streams while they are active.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tRANSCRIPTIONVOCABULARYID** | **string**| The ID of the Transcription Vocabulary. | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTranscriptionVocabulary**
> TranscriptionVocabularyResponse GetTranscriptionVocabulary(ctx, tRANSCRIPTIONVOCABULARYID)
Retrieve a Transcription Vocabulary

Retrieves the details of a Transcription Vocabulary that has previously been created. Supply the unique Transcription Vocabulary ID and Mux will return the corresponding Transcription Vocabulary information. The same information is returned when creating a Transcription Vocabulary.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tRANSCRIPTIONVOCABULARYID** | **string**| The ID of the Transcription Vocabulary. | 

### Return type

[**TranscriptionVocabularyResponse**](TranscriptionVocabularyResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTranscriptionVocabularies**
> ListTranscriptionVocabulariesResponse ListTranscriptionVocabularies(ctx, optional)
List Transcription Vocabularies

List all Transcription Vocabularies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListTranscriptionVocabulariesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListTranscriptionVocabulariesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of items to include in the response | [default to 10]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]

### Return type

[**ListTranscriptionVocabulariesResponse**](ListTranscriptionVocabulariesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTranscriptionVocabulary**
> TranscriptionVocabularyResponse UpdateTranscriptionVocabulary(ctx, tRANSCRIPTIONVOCABULARYID, updateTranscriptionVocabularyRequest)
Update a Transcription Vocabulary

Updates the details of a previously-created Transcription Vocabulary. Updates to Transcription Vocabularies are allowed while associated live streams are active. However, updates will not be applied to those streams while they are active.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tRANSCRIPTIONVOCABULARYID** | **string**| The ID of the Transcription Vocabulary. | 
  **updateTranscriptionVocabularyRequest** | [**UpdateTranscriptionVocabularyRequest**](UpdateTranscriptionVocabularyRequest.md)|  | 

### Return type

[**TranscriptionVocabularyResponse**](TranscriptionVocabularyResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

