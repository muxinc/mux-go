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

// URLSigningKeysApiService URLSigningKeysApi service
type URLSigningKeysApiService service

type ApiCreateUrlSigningKeyRequest struct {
	ctx _context.Context
	ApiService *URLSigningKeysApiService
}


func (r ApiCreateUrlSigningKeyRequest) Execute() (SigningKeyResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateUrlSigningKeyExecute(r)
}

/*
 * CreateUrlSigningKey Create a URL signing key
 * Creates a new signing key pair. When creating a new signing key, the API will generate a 2048-bit RSA key-pair and return the private key and a generated key-id; the public key will be stored at Mux to validate signed tokens.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateUrlSigningKeyRequest
 */
func (a *URLSigningKeysApiService) CreateUrlSigningKey(ctx _context.Context) ApiCreateUrlSigningKeyRequest {
	return ApiCreateUrlSigningKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SigningKeyResponse
 */
func (a *URLSigningKeysApiService) CreateUrlSigningKeyExecute(r ApiCreateUrlSigningKeyRequest) (SigningKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SigningKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "URLSigningKeysApiService.CreateUrlSigningKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video/v1/signing-keys"

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

type ApiDeleteUrlSigningKeyRequest struct {
	ctx _context.Context
	ApiService *URLSigningKeysApiService
	sIGNINGKEYID string
}


func (r ApiDeleteUrlSigningKeyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteUrlSigningKeyExecute(r)
}

/*
 * DeleteUrlSigningKey Delete a URL signing key
 * Deletes an existing signing key. Use with caution, as this will invalidate any existing signatures and no URLs can be signed using the key again.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sIGNINGKEYID The ID of the signing key.
 * @return ApiDeleteUrlSigningKeyRequest
 */
func (a *URLSigningKeysApiService) DeleteUrlSigningKey(ctx _context.Context, sIGNINGKEYID string) ApiDeleteUrlSigningKeyRequest {
	return ApiDeleteUrlSigningKeyRequest{
		ApiService: a,
		ctx: ctx,
		sIGNINGKEYID: sIGNINGKEYID,
	}
}

/*
 * Execute executes the request
 */
func (a *URLSigningKeysApiService) DeleteUrlSigningKeyExecute(r ApiDeleteUrlSigningKeyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "URLSigningKeysApiService.DeleteUrlSigningKey")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video/v1/signing-keys/{SIGNING_KEY_ID}"
	localVarPath = strings.Replace(localVarPath, "{"+"SIGNING_KEY_ID"+"}", _neturl.PathEscape(parameterToString(r.sIGNINGKEYID, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetUrlSigningKeyRequest struct {
	ctx _context.Context
	ApiService *URLSigningKeysApiService
	sIGNINGKEYID string
}


func (r ApiGetUrlSigningKeyRequest) Execute() (SigningKeyResponse, *_nethttp.Response, error) {
	return r.ApiService.GetUrlSigningKeyExecute(r)
}

/*
 * GetUrlSigningKey Retrieve a URL signing key
 * Retrieves the details of a URL signing key that has previously
been created. Supply the unique signing key ID that was returned from your
previous request, and Mux will return the corresponding signing key information.
**The private key is not returned in this response.**

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sIGNINGKEYID The ID of the signing key.
 * @return ApiGetUrlSigningKeyRequest
 */
func (a *URLSigningKeysApiService) GetUrlSigningKey(ctx _context.Context, sIGNINGKEYID string) ApiGetUrlSigningKeyRequest {
	return ApiGetUrlSigningKeyRequest{
		ApiService: a,
		ctx: ctx,
		sIGNINGKEYID: sIGNINGKEYID,
	}
}

/*
 * Execute executes the request
 * @return SigningKeyResponse
 */
func (a *URLSigningKeysApiService) GetUrlSigningKeyExecute(r ApiGetUrlSigningKeyRequest) (SigningKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SigningKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "URLSigningKeysApiService.GetUrlSigningKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video/v1/signing-keys/{SIGNING_KEY_ID}"
	localVarPath = strings.Replace(localVarPath, "{"+"SIGNING_KEY_ID"+"}", _neturl.PathEscape(parameterToString(r.sIGNINGKEYID, "")), -1)

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

type ApiListUrlSigningKeysRequest struct {
	ctx _context.Context
	ApiService *URLSigningKeysApiService
	limit *int32
	page *int32
}

func (r ApiListUrlSigningKeysRequest) Limit(limit int32) ApiListUrlSigningKeysRequest {
	r.limit = &limit
	return r
}
func (r ApiListUrlSigningKeysRequest) Page(page int32) ApiListUrlSigningKeysRequest {
	r.page = &page
	return r
}

func (r ApiListUrlSigningKeysRequest) Execute() (ListSigningKeysResponse, *_nethttp.Response, error) {
	return r.ApiService.ListUrlSigningKeysExecute(r)
}

/*
 * ListUrlSigningKeys List URL signing keys
 * Returns a list of URL signing keys.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListUrlSigningKeysRequest
 */
func (a *URLSigningKeysApiService) ListUrlSigningKeys(ctx _context.Context) ApiListUrlSigningKeysRequest {
	return ApiListUrlSigningKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ListSigningKeysResponse
 */
func (a *URLSigningKeysApiService) ListUrlSigningKeysExecute(r ApiListUrlSigningKeysRequest) (ListSigningKeysResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListSigningKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "URLSigningKeysApiService.ListUrlSigningKeys")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video/v1/signing-keys"

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
