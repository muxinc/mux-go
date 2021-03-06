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
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// RealTimeApiService RealTimeApi service
type RealTimeApiService service

type ApiGetRealtimeBreakdownRequest struct {
	ctx _context.Context
	ApiService *RealTimeApiService
	rEALTIMEMETRICID string
	dimension *string
	timestamp *float32
	filters *[]string
	orderBy *string
	orderDirection *string
}

func (r ApiGetRealtimeBreakdownRequest) Dimension(dimension string) ApiGetRealtimeBreakdownRequest {
	r.dimension = &dimension
	return r
}
func (r ApiGetRealtimeBreakdownRequest) Timestamp(timestamp float32) ApiGetRealtimeBreakdownRequest {
	r.timestamp = &timestamp
	return r
}
func (r ApiGetRealtimeBreakdownRequest) Filters(filters []string) ApiGetRealtimeBreakdownRequest {
	r.filters = &filters
	return r
}
func (r ApiGetRealtimeBreakdownRequest) OrderBy(orderBy string) ApiGetRealtimeBreakdownRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiGetRealtimeBreakdownRequest) OrderDirection(orderDirection string) ApiGetRealtimeBreakdownRequest {
	r.orderDirection = &orderDirection
	return r
}

func (r ApiGetRealtimeBreakdownRequest) Execute() (GetRealTimeBreakdownResponse, *_nethttp.Response, error) {
	return r.ApiService.GetRealtimeBreakdownExecute(r)
}

/*
 * GetRealtimeBreakdown Get Real-Time Breakdown
 * Gets breakdown information for a specific dimension and metric along with the number of concurrent viewers and negative impact score.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param rEALTIMEMETRICID ID of the Realtime Metric
 * @return ApiGetRealtimeBreakdownRequest
 */
func (a *RealTimeApiService) GetRealtimeBreakdown(ctx _context.Context, rEALTIMEMETRICID string) ApiGetRealtimeBreakdownRequest {
	return ApiGetRealtimeBreakdownRequest{
		ApiService: a,
		ctx: ctx,
		rEALTIMEMETRICID: rEALTIMEMETRICID,
	}
}

/*
 * Execute executes the request
 * @return GetRealTimeBreakdownResponse
 */
func (a *RealTimeApiService) GetRealtimeBreakdownExecute(r ApiGetRealtimeBreakdownRequest) (GetRealTimeBreakdownResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetRealTimeBreakdownResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeApiService.GetRealtimeBreakdown")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/realtime/metrics/{REALTIME_METRIC_ID}/breakdown"
	localVarPath = strings.Replace(localVarPath, "{"+"REALTIME_METRIC_ID"+"}", _neturl.PathEscape(parameterToString(r.rEALTIMEMETRICID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.dimension != nil {
		localVarQueryParams.Add("dimension", parameterToString(*r.dimension, ""))
	}
	if r.timestamp != nil {
		localVarQueryParams.Add("timestamp", parameterToString(*r.timestamp, ""))
	}
	if r.filters != nil {
		t := *r.filters
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters[]", parameterToString(t, "multi"))
		}
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("order_by", parameterToString(*r.orderBy, ""))
	}
	if r.orderDirection != nil {
		localVarQueryParams.Add("order_direction", parameterToString(*r.orderDirection, ""))
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

type ApiGetRealtimeHistogramTimeseriesRequest struct {
	ctx _context.Context
	ApiService *RealTimeApiService
	rEALTIMEHISTOGRAMMETRICID string
	filters *[]string
}

func (r ApiGetRealtimeHistogramTimeseriesRequest) Filters(filters []string) ApiGetRealtimeHistogramTimeseriesRequest {
	r.filters = &filters
	return r
}

func (r ApiGetRealtimeHistogramTimeseriesRequest) Execute() (GetRealTimeHistogramTimeseriesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetRealtimeHistogramTimeseriesExecute(r)
}

/*
 * GetRealtimeHistogramTimeseries Get Real-Time Histogram Timeseries
 * Gets histogram timeseries information for a specific metric.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param rEALTIMEHISTOGRAMMETRICID ID of the Realtime Histogram Metric
 * @return ApiGetRealtimeHistogramTimeseriesRequest
 */
func (a *RealTimeApiService) GetRealtimeHistogramTimeseries(ctx _context.Context, rEALTIMEHISTOGRAMMETRICID string) ApiGetRealtimeHistogramTimeseriesRequest {
	return ApiGetRealtimeHistogramTimeseriesRequest{
		ApiService: a,
		ctx: ctx,
		rEALTIMEHISTOGRAMMETRICID: rEALTIMEHISTOGRAMMETRICID,
	}
}

/*
 * Execute executes the request
 * @return GetRealTimeHistogramTimeseriesResponse
 */
func (a *RealTimeApiService) GetRealtimeHistogramTimeseriesExecute(r ApiGetRealtimeHistogramTimeseriesRequest) (GetRealTimeHistogramTimeseriesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetRealTimeHistogramTimeseriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeApiService.GetRealtimeHistogramTimeseries")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/realtime/metrics/{REALTIME_HISTOGRAM_METRIC_ID}/histogram-timeseries"
	localVarPath = strings.Replace(localVarPath, "{"+"REALTIME_HISTOGRAM_METRIC_ID"+"}", _neturl.PathEscape(parameterToString(r.rEALTIMEHISTOGRAMMETRICID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filters != nil {
		t := *r.filters
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters[]", parameterToString(t, "multi"))
		}
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

type ApiGetRealtimeTimeseriesRequest struct {
	ctx _context.Context
	ApiService *RealTimeApiService
	rEALTIMEMETRICID string
	filters *[]string
}

func (r ApiGetRealtimeTimeseriesRequest) Filters(filters []string) ApiGetRealtimeTimeseriesRequest {
	r.filters = &filters
	return r
}

func (r ApiGetRealtimeTimeseriesRequest) Execute() (GetRealTimeTimeseriesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetRealtimeTimeseriesExecute(r)
}

/*
 * GetRealtimeTimeseries Get Real-Time Timeseries
 * Gets Time series information for a specific metric along with the number of concurrent viewers.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param rEALTIMEMETRICID ID of the Realtime Metric
 * @return ApiGetRealtimeTimeseriesRequest
 */
func (a *RealTimeApiService) GetRealtimeTimeseries(ctx _context.Context, rEALTIMEMETRICID string) ApiGetRealtimeTimeseriesRequest {
	return ApiGetRealtimeTimeseriesRequest{
		ApiService: a,
		ctx: ctx,
		rEALTIMEMETRICID: rEALTIMEMETRICID,
	}
}

/*
 * Execute executes the request
 * @return GetRealTimeTimeseriesResponse
 */
func (a *RealTimeApiService) GetRealtimeTimeseriesExecute(r ApiGetRealtimeTimeseriesRequest) (GetRealTimeTimeseriesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetRealTimeTimeseriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeApiService.GetRealtimeTimeseries")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/realtime/metrics/{REALTIME_METRIC_ID}/timeseries"
	localVarPath = strings.Replace(localVarPath, "{"+"REALTIME_METRIC_ID"+"}", _neturl.PathEscape(parameterToString(r.rEALTIMEMETRICID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filters != nil {
		t := *r.filters
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters[]", parameterToString(t, "multi"))
		}
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

type ApiListRealtimeDimensionsRequest struct {
	ctx _context.Context
	ApiService *RealTimeApiService
}


func (r ApiListRealtimeDimensionsRequest) Execute() (ListRealTimeDimensionsResponse, *_nethttp.Response, error) {
	return r.ApiService.ListRealtimeDimensionsExecute(r)
}

/*
 * ListRealtimeDimensions List Real-Time Dimensions
 * Lists availiable real-time dimensions

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListRealtimeDimensionsRequest
 */
func (a *RealTimeApiService) ListRealtimeDimensions(ctx _context.Context) ApiListRealtimeDimensionsRequest {
	return ApiListRealtimeDimensionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ListRealTimeDimensionsResponse
 */
func (a *RealTimeApiService) ListRealtimeDimensionsExecute(r ApiListRealtimeDimensionsRequest) (ListRealTimeDimensionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListRealTimeDimensionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeApiService.ListRealtimeDimensions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/realtime/dimensions"

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

type ApiListRealtimeMetricsRequest struct {
	ctx _context.Context
	ApiService *RealTimeApiService
}


func (r ApiListRealtimeMetricsRequest) Execute() (ListRealTimeMetricsResponse, *_nethttp.Response, error) {
	return r.ApiService.ListRealtimeMetricsExecute(r)
}

/*
 * ListRealtimeMetrics List Real-Time Metrics
 * Lists availiable real-time metrics.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListRealtimeMetricsRequest
 */
func (a *RealTimeApiService) ListRealtimeMetrics(ctx _context.Context) ApiListRealtimeMetricsRequest {
	return ApiListRealtimeMetricsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ListRealTimeMetricsResponse
 */
func (a *RealTimeApiService) ListRealtimeMetricsExecute(r ApiListRealtimeMetricsRequest) (ListRealTimeMetricsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListRealTimeMetricsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeApiService.ListRealtimeMetrics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/realtime/metrics"

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
