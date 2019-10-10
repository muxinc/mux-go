// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type VideoViewsApiService service

func (a *VideoViewsApiService) GetVideoView(vIDEOVIEWID string, opts ...APIOption) (VideoViewResponse, error) {
	var (
		localVarAPIOptions   = new(APIOptions)
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VideoViewResponse
	)

	for _, opt := range opts {
		opt(localVarAPIOptions)
	}

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/data/v1/video-views/{VIDEO_VIEW_ID}"
	localVarPath = strings.Replace(localVarPath, "{"+"VIDEO_VIEW_ID"+"}", fmt.Sprintf("%v", vIDEOVIEWID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	// Check for common HTTP error status codes
	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return localVarReturnValue, err
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

type ListVideoViewsParams struct {
	Limit          int32
	Page           int32
	ViewerId       string
	ErrorId        int32
	OrderDirection string
	Filters        []string
	Timeframe      []string
}

// ListVideoViews optionally accepts the APIOption of WithParams(*ListVideoViewsParams).
func (a *VideoViewsApiService) ListVideoViews(opts ...APIOption) (ListVideoViewsResponse, error) {
	var (
		localVarAPIOptions   = new(APIOptions)
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListVideoViewsResponse
	)

	for _, opt := range opts {
		opt(localVarAPIOptions)
	}

	localVarOptionals, ok := localVarAPIOptions.params.(*ListVideoViewsParams)
	if localVarAPIOptions.params != nil && !ok {
		return localVarReturnValue, reportError("provided params were not of type *ListVideoViewsParams")
	}

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/data/v1/video-views"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && isSet(localVarOptionals.Limit) {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit, ""))
	}
	if localVarOptionals != nil && isSet(localVarOptionals.Page) {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page, ""))
	}
	if localVarOptionals != nil && isSet(localVarOptionals.ViewerId) {
		localVarQueryParams.Add("viewer_id", parameterToString(localVarOptionals.ViewerId, ""))
	}
	if localVarOptionals != nil && isSet(localVarOptionals.ErrorId) {
		localVarQueryParams.Add("error_id", parameterToString(localVarOptionals.ErrorId, ""))
	}
	if localVarOptionals != nil && isSet(localVarOptionals.OrderDirection) {
		localVarQueryParams.Add("order_direction", parameterToString(localVarOptionals.OrderDirection, ""))
	}
	if localVarOptionals != nil && isSet(localVarOptionals.Filters) {
		// This will "always work" for Mux's use case, since we always treat collections in query params as "multi" types.
		// The first version of this code checked the collectionFormat, but that's just wasted CPU cycles right now.
		for _, v := range localVarOptionals.Filters {
			localVarQueryParams.Add("filters[]", v)
		}
	}
	if localVarOptionals != nil && isSet(localVarOptionals.Timeframe) {
		// This will "always work" for Mux's use case, since we always treat collections in query params as "multi" types.
		// The first version of this code checked the collectionFormat, but that's just wasted CPU cycles right now.
		for _, v := range localVarOptionals.Timeframe {
			localVarQueryParams.Add("timeframe[]", v)
		}
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	// Check for common HTTP error status codes
	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return localVarReturnValue, err
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}
