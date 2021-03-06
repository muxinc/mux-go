/*
 * Mux API
 *
 * Mux is how developers build online video. This API encompasses both Mux Video and Mux Data functionality to help you build your video-related projects better and faster than ever before. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package muxgo

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DirectUploadsApiService DirectUploadsApi service
type DirectUploadsApiService service

type ApiCancelDirectUploadRequest struct {
	ctx _context.Context
	ApiService *DirectUploadsApiService
	uPLOADID string
}


func (r ApiCancelDirectUploadRequest) Execute() (UploadResponse, *_nethttp.Response, error) {
	return r.ApiService.CancelDirectUploadExecute(r)
}

/*
 * CancelDirectUpload Cancel a direct upload
 * Cancels a direct upload and marks it as cancelled. If a pending upload finishes after this
request, no asset will be created. This request will only succeed if the upload is still in
the `waiting` state.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uPLOADID ID of the Upload
 * @return ApiCancelDirectUploadRequest
 */
func (a *DirectUploadsApiService) CancelDirectUpload(ctx _context.Context, uPLOADID string) ApiCancelDirectUploadRequest {
	return ApiCancelDirectUploadRequest{
		ApiService: a,
		ctx: ctx,
		uPLOADID: uPLOADID,
	}
}

/*
 * Execute executes the request
 * @return UploadResponse
 */
func (a *DirectUploadsApiService) CancelDirectUploadExecute(r ApiCancelDirectUploadRequest) (UploadResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UploadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectUploadsApiService.CancelDirectUpload")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video/v1/uploads/{UPLOAD_ID}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"UPLOAD_ID"+"}", _neturl.PathEscape(parameterToString(r.uPLOADID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateDirectUploadRequest struct {
	ctx _context.Context
	ApiService *DirectUploadsApiService
	createUploadRequest *CreateUploadRequest
}

func (r ApiCreateDirectUploadRequest) CreateUploadRequest(createUploadRequest CreateUploadRequest) ApiCreateDirectUploadRequest {
	r.createUploadRequest = &createUploadRequest
	return r
}

func (r ApiCreateDirectUploadRequest) Execute() (UploadResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateDirectUploadExecute(r)
}

/*
 * CreateDirectUpload Create a new direct upload URL
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateDirectUploadRequest
 */
func (a *DirectUploadsApiService) CreateDirectUpload(ctx _context.Context) ApiCreateDirectUploadRequest {
	return ApiCreateDirectUploadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return UploadResponse
 */
func (a *DirectUploadsApiService) CreateDirectUploadExecute(r ApiCreateDirectUploadRequest) (UploadResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UploadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectUploadsApiService.CreateDirectUpload")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video/v1/uploads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.createUploadRequest == nil {
		return localVarReturnValue, nil, reportError("createUploadRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createUploadRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetDirectUploadRequest struct {
	ctx _context.Context
	ApiService *DirectUploadsApiService
	uPLOADID string
}


func (r ApiGetDirectUploadRequest) Execute() (UploadResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDirectUploadExecute(r)
}

/*
 * GetDirectUpload Retrieve a single direct upload's info
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uPLOADID ID of the Upload
 * @return ApiGetDirectUploadRequest
 */
func (a *DirectUploadsApiService) GetDirectUpload(ctx _context.Context, uPLOADID string) ApiGetDirectUploadRequest {
	return ApiGetDirectUploadRequest{
		ApiService: a,
		ctx: ctx,
		uPLOADID: uPLOADID,
	}
}

/*
 * Execute executes the request
 * @return UploadResponse
 */
func (a *DirectUploadsApiService) GetDirectUploadExecute(r ApiGetDirectUploadRequest) (UploadResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UploadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectUploadsApiService.GetDirectUpload")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video/v1/uploads/{UPLOAD_ID}"
	localVarPath = strings.Replace(localVarPath, "{"+"UPLOAD_ID"+"}", _neturl.PathEscape(parameterToString(r.uPLOADID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListDirectUploadsRequest struct {
	ctx _context.Context
	ApiService *DirectUploadsApiService
	limit *int32
	page *int32
}

func (r ApiListDirectUploadsRequest) Limit(limit int32) ApiListDirectUploadsRequest {
	r.limit = &limit
	return r
}
func (r ApiListDirectUploadsRequest) Page(page int32) ApiListDirectUploadsRequest {
	r.page = &page
	return r
}

func (r ApiListDirectUploadsRequest) Execute() (ListUploadsResponse, *_nethttp.Response, error) {
	return r.ApiService.ListDirectUploadsExecute(r)
}

/*
 * ListDirectUploads List direct uploads
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListDirectUploadsRequest
 */
func (a *DirectUploadsApiService) ListDirectUploads(ctx _context.Context) ApiListDirectUploadsRequest {
	return ApiListDirectUploadsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ListUploadsResponse
 */
func (a *DirectUploadsApiService) ListDirectUploadsExecute(r ApiListDirectUploadsRequest) (ListUploadsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListUploadsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectUploadsApiService.ListDirectUploads")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video/v1/uploads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
