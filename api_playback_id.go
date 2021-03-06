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

// PlaybackIDApiService PlaybackIDApi service
type PlaybackIDApiService service

type ApiGetAssetOrLivestreamIdRequest struct {
	ctx _context.Context
	ApiService *PlaybackIDApiService
	pLAYBACKID string
}


func (r ApiGetAssetOrLivestreamIdRequest) Execute() (GetAssetOrLiveStreamIdResponse, *_nethttp.Response, error) {
	return r.ApiService.GetAssetOrLivestreamIdExecute(r)
}

/*
 * GetAssetOrLivestreamId Retrieve an Asset or Live Stream ID
 * Retrieves the Identifier of the Asset or Live Stream associated with the Playback ID.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pLAYBACKID The live stream's playback ID.
 * @return ApiGetAssetOrLivestreamIdRequest
 */
func (a *PlaybackIDApiService) GetAssetOrLivestreamId(ctx _context.Context, pLAYBACKID string) ApiGetAssetOrLivestreamIdRequest {
	return ApiGetAssetOrLivestreamIdRequest{
		ApiService: a,
		ctx: ctx,
		pLAYBACKID: pLAYBACKID,
	}
}

/*
 * Execute executes the request
 * @return GetAssetOrLiveStreamIdResponse
 */
func (a *PlaybackIDApiService) GetAssetOrLivestreamIdExecute(r ApiGetAssetOrLivestreamIdRequest) (GetAssetOrLiveStreamIdResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetAssetOrLiveStreamIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlaybackIDApiService.GetAssetOrLivestreamId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video/v1/playback-ids/{PLAYBACK_ID}"
	localVarPath = strings.Replace(localVarPath, "{"+"PLAYBACK_ID"+"}", _neturl.PathEscape(parameterToString(r.pLAYBACKID, "")), -1)

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
