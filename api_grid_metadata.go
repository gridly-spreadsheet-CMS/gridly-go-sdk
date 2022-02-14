/*
Gridly API

Gridly API documentation

API version: 3.21.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

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

// GridMetadataApiService GridMetadataApi service
type GridMetadataApiService service

type GridMetadataApiApiCreateRequest struct {
	ctx _context.Context
	ApiService *GridMetadataApiService
	gridId string
	createMetadataDTO *CreateMetadata
}

// createMetadataDTO
func (r GridMetadataApiApiCreateRequest) CreateMetadataDTO(createMetadataDTO CreateMetadata) GridMetadataApiApiCreateRequest {
	r.createMetadataDTO = &createMetadataDTO
	return r
}

func (r GridMetadataApiApiCreateRequest) Execute() (Metadata, *_nethttp.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create create

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gridId gridId
 @return GridMetadataApiApiCreateRequest
*/
func (a *GridMetadataApiService) Create(ctx _context.Context, gridId string) GridMetadataApiApiCreateRequest {
	return GridMetadataApiApiCreateRequest{
		ApiService: a,
		ctx: ctx,
		gridId: gridId,
	}
}

// Execute executes the request
//  @return Metadata
func (a *GridMetadataApiService) CreateExecute(r GridMetadataApiApiCreateRequest) (Metadata, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  Metadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GridMetadataApiService.Create")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grids/{gridId}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"gridId"+"}", _neturl.PathEscape(parameterToString(r.gridId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.createMetadataDTO == nil {
		return localVarReturnValue, nil, reportError("createMetadataDTO is required and must be specified")
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
	localVarPostBody = r.createMetadataDTO
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type GridMetadataApiApiDeleteRequest struct {
	ctx _context.Context
	ApiService *GridMetadataApiService
	gridId string
	metadataId string
}


func (r GridMetadataApiApiDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete delete

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gridId gridId
 @param metadataId metadataId
 @return GridMetadataApiApiDeleteRequest
*/
func (a *GridMetadataApiService) Delete(ctx _context.Context, gridId string, metadataId string) GridMetadataApiApiDeleteRequest {
	return GridMetadataApiApiDeleteRequest{
		ApiService: a,
		ctx: ctx,
		gridId: gridId,
		metadataId: metadataId,
	}
}

// Execute executes the request
func (a *GridMetadataApiService) DeleteExecute(r GridMetadataApiApiDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GridMetadataApiService.Delete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grids/{gridId}/metadata/{metadataId}"
	localVarPath = strings.Replace(localVarPath, "{"+"gridId"+"}", _neturl.PathEscape(parameterToString(r.gridId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"metadataId"+"}", _neturl.PathEscape(parameterToString(r.metadataId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type GridMetadataApiApiGetRequest struct {
	ctx _context.Context
	ApiService *GridMetadataApiService
	gridId string
	metadataId string
}


func (r GridMetadataApiApiGetRequest) Execute() (Metadata, *_nethttp.Response, error) {
	return r.ApiService.GetExecute(r)
}

/*
Get get

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gridId gridId
 @param metadataId metadataId
 @return GridMetadataApiApiGetRequest
*/
func (a *GridMetadataApiService) Get(ctx _context.Context, gridId string, metadataId string) GridMetadataApiApiGetRequest {
	return GridMetadataApiApiGetRequest{
		ApiService: a,
		ctx: ctx,
		gridId: gridId,
		metadataId: metadataId,
	}
}

// Execute executes the request
//  @return Metadata
func (a *GridMetadataApiService) GetExecute(r GridMetadataApiApiGetRequest) (Metadata, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  Metadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GridMetadataApiService.Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grids/{gridId}/metadata/{metadataId}"
	localVarPath = strings.Replace(localVarPath, "{"+"gridId"+"}", _neturl.PathEscape(parameterToString(r.gridId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"metadataId"+"}", _neturl.PathEscape(parameterToString(r.metadataId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type GridMetadataApiApiListRequest struct {
	ctx _context.Context
	ApiService *GridMetadataApiService
	gridId string
}


func (r GridMetadataApiApiListRequest) Execute() ([]Metadata, *_nethttp.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List list

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gridId gridId
 @return GridMetadataApiApiListRequest
*/
func (a *GridMetadataApiService) List(ctx _context.Context, gridId string) GridMetadataApiApiListRequest {
	return GridMetadataApiApiListRequest{
		ApiService: a,
		ctx: ctx,
		gridId: gridId,
	}
}

// Execute executes the request
//  @return []Metadata
func (a *GridMetadataApiService) ListExecute(r GridMetadataApiApiListRequest) ([]Metadata, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Metadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GridMetadataApiService.List")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grids/{gridId}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"gridId"+"}", _neturl.PathEscape(parameterToString(r.gridId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type GridMetadataApiApiUpdateRequest struct {
	ctx _context.Context
	ApiService *GridMetadataApiService
	gridId string
	metadataId string
	updateMetadataDTO *UpdateMetadata
}

// updateMetadataDTO
func (r GridMetadataApiApiUpdateRequest) UpdateMetadataDTO(updateMetadataDTO UpdateMetadata) GridMetadataApiApiUpdateRequest {
	r.updateMetadataDTO = &updateMetadataDTO
	return r
}

func (r GridMetadataApiApiUpdateRequest) Execute() (Metadata, *_nethttp.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update update

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gridId gridId
 @param metadataId metadataId
 @return GridMetadataApiApiUpdateRequest
*/
func (a *GridMetadataApiService) Update(ctx _context.Context, gridId string, metadataId string) GridMetadataApiApiUpdateRequest {
	return GridMetadataApiApiUpdateRequest{
		ApiService: a,
		ctx: ctx,
		gridId: gridId,
		metadataId: metadataId,
	}
}

// Execute executes the request
//  @return Metadata
func (a *GridMetadataApiService) UpdateExecute(r GridMetadataApiApiUpdateRequest) (Metadata, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  Metadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GridMetadataApiService.Update")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/grids/{gridId}/metadata/{metadataId}"
	localVarPath = strings.Replace(localVarPath, "{"+"gridId"+"}", _neturl.PathEscape(parameterToString(r.gridId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"metadataId"+"}", _neturl.PathEscape(parameterToString(r.metadataId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.updateMetadataDTO == nil {
		return localVarReturnValue, nil, reportError("updateMetadataDTO is required and must be specified")
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
	localVarPostBody = r.updateMetadataDTO
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
