/*
Gridly API

Gridly API documentation

API version: 3.27.2
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
	"os"
)

// Linger please
var (
	_ _context.Context
)

// ViewFileApiService ViewFileApi service
type ViewFileApiService service

type ViewFileApiApiDeleteRequest struct {
	ctx _context.Context
	ApiService *ViewFileApiService
	columnId *string
	recordId *string
	viewId string
	deleteFileDTO *DeleteFile
}

// columnId
func (r ViewFileApiApiDeleteRequest) ColumnId(columnId string) ViewFileApiApiDeleteRequest {
	r.columnId = &columnId
	return r
}
// recordId
func (r ViewFileApiApiDeleteRequest) RecordId(recordId string) ViewFileApiApiDeleteRequest {
	r.recordId = &recordId
	return r
}
// deleteFileDTO
func (r ViewFileApiApiDeleteRequest) DeleteFileDTO(deleteFileDTO DeleteFile) ViewFileApiApiDeleteRequest {
	r.deleteFileDTO = &deleteFileDTO
	return r
}

func (r ViewFileApiApiDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete delete

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return ViewFileApiApiDeleteRequest
*/
func (a *ViewFileApiService) Delete(ctx _context.Context, viewId string) ViewFileApiApiDeleteRequest {
	return ViewFileApiApiDeleteRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
func (a *ViewFileApiService) DeleteExecute(r ViewFileApiApiDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewFileApiService.Delete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/files"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", _neturl.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.columnId == nil {
		return nil, reportError("columnId is required and must be specified")
	}
	if r.recordId == nil {
		return nil, reportError("recordId is required and must be specified")
	}
	if r.deleteFileDTO == nil {
		return nil, reportError("deleteFileDTO is required and must be specified")
	}

	localVarQueryParams.Add("columnId", parameterToString(*r.columnId, ""))
	localVarQueryParams.Add("recordId", parameterToString(*r.recordId, ""))
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
	localVarPostBody = r.deleteFileDTO
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

type ViewFileApiApiDownloadRequest struct {
	ctx _context.Context
	ApiService *ViewFileApiService
	fileId string
	viewId string
}


func (r ViewFileApiApiDownloadRequest) Execute() (*os.File, *_nethttp.Response, error) {
	return r.ApiService.DownloadExecute(r)
}

/*
Download download

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId fileId
 @param viewId viewId
 @return ViewFileApiApiDownloadRequest
*/
func (a *ViewFileApiService) Download(ctx _context.Context, fileId string, viewId string) ViewFileApiApiDownloadRequest {
	return ViewFileApiApiDownloadRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *ViewFileApiService) DownloadExecute(r ViewFileApiApiDownloadRequest) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewFileApiService.Download")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/files/{fileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", _neturl.PathEscape(parameterToString(r.fileId, "")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

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

type ViewFileApiApiUploadRequest struct {
	ctx _context.Context
	ApiService *ViewFileApiService
	viewId string
	columnId *string
	file **os.File
	recordId *string
}

// columnId
func (r ViewFileApiApiUploadRequest) ColumnId(columnId string) ViewFileApiApiUploadRequest {
	r.columnId = &columnId
	return r
}
// file
func (r ViewFileApiApiUploadRequest) File(file *os.File) ViewFileApiApiUploadRequest {
	r.file = &file
	return r
}
// recordId
func (r ViewFileApiApiUploadRequest) RecordId(recordId string) ViewFileApiApiUploadRequest {
	r.recordId = &recordId
	return r
}

func (r ViewFileApiApiUploadRequest) Execute() (UploadedFile, *_nethttp.Response, error) {
	return r.ApiService.UploadExecute(r)
}

/*
Upload upload

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return ViewFileApiApiUploadRequest
*/
func (a *ViewFileApiService) Upload(ctx _context.Context, viewId string) ViewFileApiApiUploadRequest {
	return ViewFileApiApiUploadRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return UploadedFile
func (a *ViewFileApiService) UploadExecute(r ViewFileApiApiUploadRequest) (UploadedFile, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  UploadedFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewFileApiService.Upload")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/files"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", _neturl.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.columnId == nil {
		return localVarReturnValue, nil, reportError("columnId is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}
	if r.recordId == nil {
		return localVarReturnValue, nil, reportError("recordId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormParams.Add("columnId", parameterToString(*r.columnId, ""))
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"

	fileLocalVarFile := *r.file
	if fileLocalVarFile != nil {
		fbs, _ := _ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	localVarFormParams.Add("recordId", parameterToString(*r.recordId, ""))
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

type ViewFileApiApiUploadZipRequest struct {
	ctx _context.Context
	ApiService *ViewFileApiService
	viewId string
	columnId *string
	file **os.File
	fileMappings *string
}

// columnId
func (r ViewFileApiApiUploadZipRequest) ColumnId(columnId string) ViewFileApiApiUploadZipRequest {
	r.columnId = &columnId
	return r
}
// file
func (r ViewFileApiApiUploadZipRequest) File(file *os.File) ViewFileApiApiUploadZipRequest {
	r.file = &file
	return r
}
// fileMappings
func (r ViewFileApiApiUploadZipRequest) FileMappings(fileMappings string) ViewFileApiApiUploadZipRequest {
	r.fileMappings = &fileMappings
	return r
}

func (r ViewFileApiApiUploadZipRequest) Execute() ([]Record, *_nethttp.Response, error) {
	return r.ApiService.UploadZipExecute(r)
}

/*
UploadZip uploadZip

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param viewId viewId
 @return ViewFileApiApiUploadZipRequest
*/
func (a *ViewFileApiService) UploadZip(ctx _context.Context, viewId string) ViewFileApiApiUploadZipRequest {
	return ViewFileApiApiUploadZipRequest{
		ApiService: a,
		ctx: ctx,
		viewId: viewId,
	}
}

// Execute executes the request
//  @return []Record
func (a *ViewFileApiService) UploadZipExecute(r ViewFileApiApiUploadZipRequest) ([]Record, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Record
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ViewFileApiService.UploadZip")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/views/{viewId}/files/zip"
	localVarPath = strings.Replace(localVarPath, "{"+"viewId"+"}", _neturl.PathEscape(parameterToString(r.viewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.columnId == nil {
		return localVarReturnValue, nil, reportError("columnId is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}
	if r.fileMappings == nil {
		return localVarReturnValue, nil, reportError("fileMappings is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormParams.Add("columnId", parameterToString(*r.columnId, ""))
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"

	fileLocalVarFile := *r.file
	if fileLocalVarFile != nil {
		fbs, _ := _ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	localVarFormParams.Add("fileMappings", parameterToString(*r.fileMappings, ""))
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
