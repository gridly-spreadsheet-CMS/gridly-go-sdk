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

// ViewDependencyApiService ViewDependencyApi service
type ViewDependencyApiService service

type ViewDependencyApiApiCreateRequest struct {
	ctx _context.Context
	ApiService *ViewDependencyApiService
	viewId string
	createDependency *CreateDependency
}

// createDependency
func (r ViewDependencyApiApiCreateRequest) CreateDependency(createDependency CreateDependency) ViewDependencyApiApiCreateRequest {
	r.createDependency = &createDependency
	return r
}

func (r ViewDependencyApiApiCreateRequest) Execute() (Dependency, *_nethttp.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create create

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return ViewDependencyApiApiCreateRequest
*/
func (a *ViewDependencyApiService) Create(ctx _context.Context, viewId string) ViewDependencyApiApiCreateRequest {
	return ViewDependencyApiApiCreateRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return Dependency
func (a *ViewDependencyApiService) CreateExecute(r ViewDependencyApiApiCreateRequest) (Dependency, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  Dependency
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewDependencyApiService.Create")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/dependencies"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", _neturl.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.createDependency == nil {
		return localVarReturnValue, nil, reportError("createDependency is required and must be specified")
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
	localVarPostBody = r.createDependency
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

type ViewDependencyApiApiDeleteRequest struct {
	ctx _context.Context
	ApiService *ViewDependencyApiService
	viewId string
	deleteDependency *DeleteDependency
}

// deleteDependency
func (r ViewDependencyApiApiDeleteRequest) DeleteDependency(deleteDependency DeleteDependency) ViewDependencyApiApiDeleteRequest {
	r.deleteDependency = &deleteDependency
	return r
}

func (r ViewDependencyApiApiDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete delete

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return ViewDependencyApiApiDeleteRequest
*/
func (a *ViewDependencyApiService) Delete(ctx _context.Context, viewId string) ViewDependencyApiApiDeleteRequest {
	return ViewDependencyApiApiDeleteRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
func (a *ViewDependencyApiService) DeleteExecute(r ViewDependencyApiApiDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewDependencyApiService.Delete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/dependencies"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", _neturl.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.deleteDependency == nil {
		return nil, reportError("deleteDependency is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.deleteDependency
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

type ViewDependencyApiApiDelete_0Request struct {
	ctx _context.Context
	ApiService *ViewDependencyApiService
	dependencyId string
	viewId string
}


func (r ViewDependencyApiApiDelete_0Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.Delete_1Execute(r)
}

/*
Delete_0 delete

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dependencyId dependencyId
 @param viewId viewId
 @return ViewDependencyApiApiDelete_0Request
*/
func (a *ViewDependencyApiService) Delete_1(ctx _context.Context, dependencyId string, viewId string) ViewDependencyApiApiDelete_0Request {
	return ViewDependencyApiApiDelete_0Request{
		ApiService: a,
		ctx: ctx,
		dependencyId: dependencyId,
		viewId: viewId,
	}
}

// Execute executes the request
func (a *ViewDependencyApiService) Delete_1Execute(r ViewDependencyApiApiDelete_0Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewDependencyApiService.Delete_1")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/dependencies/{dependencyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"dependencyId"+"}", _neturl.PathEscape(parameterToString(r.dependencyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", _neturl.PathEscape(parameterToString(r.viewId, "")), -1)

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

type ViewDependencyApiApiGetRequest struct {
	ctx _context.Context
	ApiService *ViewDependencyApiService
	dependencyId string
	viewId string
}


func (r ViewDependencyApiApiGetRequest) Execute() (Dependency, *_nethttp.Response, error) {
	return r.ApiService.GetExecute(r)
}

/*
Get get

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dependencyId dependencyId
 @param viewId viewId
 @return ViewDependencyApiApiGetRequest
*/
func (a *ViewDependencyApiService) Get(ctx _context.Context, dependencyId string, viewId string) ViewDependencyApiApiGetRequest {
	return ViewDependencyApiApiGetRequest{
		ApiService: a,
		ctx: ctx,
		dependencyId: dependencyId,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return Dependency
func (a *ViewDependencyApiService) GetExecute(r ViewDependencyApiApiGetRequest) (Dependency, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  Dependency
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewDependencyApiService.Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/dependencies/{dependencyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"dependencyId"+"}", _neturl.PathEscape(parameterToString(r.dependencyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", _neturl.PathEscape(parameterToString(r.viewId, "")), -1)

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

type ViewDependencyApiApiListRequest struct {
	ctx _context.Context
	ApiService *ViewDependencyApiService
	viewId string
}


func (r ViewDependencyApiApiListRequest) Execute() ([]Dependency, *_nethttp.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List list

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return ViewDependencyApiApiListRequest
*/
func (a *ViewDependencyApiService) List(ctx _context.Context, viewId string) ViewDependencyApiApiListRequest {
	return ViewDependencyApiApiListRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return []Dependency
func (a *ViewDependencyApiService) ListExecute(r ViewDependencyApiApiListRequest) ([]Dependency, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Dependency
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewDependencyApiService.List")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/dependencies"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", _neturl.PathEscape(parameterToString(r.viewId, "")), -1)

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

type ViewDependencyApiApiUpdateRequest struct {
	ctx _context.Context
	ApiService *ViewDependencyApiService
	dependencyId string
	viewId string
	updateDependency *UpdateDependency
}

// updateDependency
func (r ViewDependencyApiApiUpdateRequest) UpdateDependency(updateDependency UpdateDependency) ViewDependencyApiApiUpdateRequest {
	r.updateDependency = &updateDependency
	return r
}

func (r ViewDependencyApiApiUpdateRequest) Execute() (Dependency, *_nethttp.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update update

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dependencyId dependencyId
 @param viewId viewId
 @return ViewDependencyApiApiUpdateRequest
*/
func (a *ViewDependencyApiService) Update(ctx _context.Context, dependencyId string, viewId string) ViewDependencyApiApiUpdateRequest {
	return ViewDependencyApiApiUpdateRequest{
		ApiService: a,
		ctx: ctx,
		dependencyId: dependencyId,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return Dependency
func (a *ViewDependencyApiService) UpdateExecute(r ViewDependencyApiApiUpdateRequest) (Dependency, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  Dependency
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewDependencyApiService.Update")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/dependencies/{dependencyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"dependencyId"+"}", _neturl.PathEscape(parameterToString(r.dependencyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", _neturl.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.updateDependency == nil {
		return localVarReturnValue, nil, reportError("updateDependency is required and must be specified")
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
	localVarPostBody = r.updateDependency
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
