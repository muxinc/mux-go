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

// IncidentsApiService IncidentsApi service
type IncidentsApiService service

type ApiGetIncidentRequest struct {
	ctx _context.Context
	ApiService *IncidentsApiService
	iNCIDENTID string
}


func (r ApiGetIncidentRequest) Execute() (IncidentResponse, *_nethttp.Response, error) {
	return r.ApiService.GetIncidentExecute(r)
}

/*
 * GetIncident Get an Incident
 * Returns the details of an incident

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param iNCIDENTID ID of the Incident
 * @return ApiGetIncidentRequest
 */
func (a *IncidentsApiService) GetIncident(ctx _context.Context, iNCIDENTID string) ApiGetIncidentRequest {
	return ApiGetIncidentRequest{
		ApiService: a,
		ctx: ctx,
		iNCIDENTID: iNCIDENTID,
	}
}

/*
 * Execute executes the request
 * @return IncidentResponse
 */
func (a *IncidentsApiService) GetIncidentExecute(r ApiGetIncidentRequest) (IncidentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IncidentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IncidentsApiService.GetIncident")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/incidents/{INCIDENT_ID}"
	localVarPath = strings.Replace(localVarPath, "{"+"INCIDENT_ID"+"}", _neturl.PathEscape(parameterToString(r.iNCIDENTID, "")), -1)

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

type ApiListIncidentsRequest struct {
	ctx _context.Context
	ApiService *IncidentsApiService
	limit *int32
	page *int32
	orderBy *string
	orderDirection *string
	status *string
	severity *string
}

func (r ApiListIncidentsRequest) Limit(limit int32) ApiListIncidentsRequest {
	r.limit = &limit
	return r
}
func (r ApiListIncidentsRequest) Page(page int32) ApiListIncidentsRequest {
	r.page = &page
	return r
}
func (r ApiListIncidentsRequest) OrderBy(orderBy string) ApiListIncidentsRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiListIncidentsRequest) OrderDirection(orderDirection string) ApiListIncidentsRequest {
	r.orderDirection = &orderDirection
	return r
}
func (r ApiListIncidentsRequest) Status(status string) ApiListIncidentsRequest {
	r.status = &status
	return r
}
func (r ApiListIncidentsRequest) Severity(severity string) ApiListIncidentsRequest {
	r.severity = &severity
	return r
}

func (r ApiListIncidentsRequest) Execute() (ListIncidentsResponse, *_nethttp.Response, error) {
	return r.ApiService.ListIncidentsExecute(r)
}

/*
 * ListIncidents List Incidents
 * Returns a list of incidents

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListIncidentsRequest
 */
func (a *IncidentsApiService) ListIncidents(ctx _context.Context) ApiListIncidentsRequest {
	return ApiListIncidentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ListIncidentsResponse
 */
func (a *IncidentsApiService) ListIncidentsExecute(r ApiListIncidentsRequest) (ListIncidentsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListIncidentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IncidentsApiService.ListIncidents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/incidents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("order_by", parameterToString(*r.orderBy, ""))
	}
	if r.orderDirection != nil {
		localVarQueryParams.Add("order_direction", parameterToString(*r.orderDirection, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.severity != nil {
		localVarQueryParams.Add("severity", parameterToString(*r.severity, ""))
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

type ApiListRelatedIncidentsRequest struct {
	ctx _context.Context
	ApiService *IncidentsApiService
	iNCIDENTID string
	limit *int32
	page *int32
	orderBy *string
	orderDirection *string
}

func (r ApiListRelatedIncidentsRequest) Limit(limit int32) ApiListRelatedIncidentsRequest {
	r.limit = &limit
	return r
}
func (r ApiListRelatedIncidentsRequest) Page(page int32) ApiListRelatedIncidentsRequest {
	r.page = &page
	return r
}
func (r ApiListRelatedIncidentsRequest) OrderBy(orderBy string) ApiListRelatedIncidentsRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiListRelatedIncidentsRequest) OrderDirection(orderDirection string) ApiListRelatedIncidentsRequest {
	r.orderDirection = &orderDirection
	return r
}

func (r ApiListRelatedIncidentsRequest) Execute() (ListRelatedIncidentsResponse, *_nethttp.Response, error) {
	return r.ApiService.ListRelatedIncidentsExecute(r)
}

/*
 * ListRelatedIncidents List Related Incidents
 * Returns all the incidents that seem related to a specific incident

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param iNCIDENTID ID of the Incident
 * @return ApiListRelatedIncidentsRequest
 */
func (a *IncidentsApiService) ListRelatedIncidents(ctx _context.Context, iNCIDENTID string) ApiListRelatedIncidentsRequest {
	return ApiListRelatedIncidentsRequest{
		ApiService: a,
		ctx: ctx,
		iNCIDENTID: iNCIDENTID,
	}
}

/*
 * Execute executes the request
 * @return ListRelatedIncidentsResponse
 */
func (a *IncidentsApiService) ListRelatedIncidentsExecute(r ApiListRelatedIncidentsRequest) (ListRelatedIncidentsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListRelatedIncidentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IncidentsApiService.ListRelatedIncidents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/incidents/{INCIDENT_ID}/related"
	localVarPath = strings.Replace(localVarPath, "{"+"INCIDENT_ID"+"}", _neturl.PathEscape(parameterToString(r.iNCIDENTID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
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
