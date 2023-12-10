/*
Apono

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apono

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AccessRequestTemplatesAPIService AccessRequestTemplatesAPI service
type AccessRequestTemplatesAPIService service

type ApiCreateAccessRequestTemplateRequest struct {
	ctx                                             context.Context
	ApiService                                      *AccessRequestTemplatesAPIService
	createAndUpdateAccessRequestTemplateClientModel *CreateAndUpdateAccessRequestTemplateClientModel
}

func (r ApiCreateAccessRequestTemplateRequest) CreateAndUpdateAccessRequestTemplateClientModel(createAndUpdateAccessRequestTemplateClientModel CreateAndUpdateAccessRequestTemplateClientModel) ApiCreateAccessRequestTemplateRequest {
	r.createAndUpdateAccessRequestTemplateClientModel = &createAndUpdateAccessRequestTemplateClientModel
	return r
}

func (r ApiCreateAccessRequestTemplateRequest) Execute() (*AccessRequestTemplateClientModel, *http.Response, error) {
	return r.ApiService.CreateAccessRequestTemplateExecute(r)
}

/*
CreateAccessRequestTemplate Create access request template

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateAccessRequestTemplateRequest
*/
func (a *AccessRequestTemplatesAPIService) CreateAccessRequestTemplate(ctx context.Context) ApiCreateAccessRequestTemplateRequest {
	return ApiCreateAccessRequestTemplateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccessRequestTemplateClientModel
func (a *AccessRequestTemplatesAPIService) CreateAccessRequestTemplateExecute(r ApiCreateAccessRequestTemplateRequest) (*AccessRequestTemplateClientModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessRequestTemplateClientModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessRequestTemplatesAPIService.CreateAccessRequestTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/client/v1/access-request-templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createAndUpdateAccessRequestTemplateClientModel == nil {
		return localVarReturnValue, nil, reportError("createAndUpdateAccessRequestTemplateClientModel is required and must be specified")
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
	localVarPostBody = r.createAndUpdateAccessRequestTemplateClientModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteAccessRequestTemplateRequest struct {
	ctx        context.Context
	ApiService *AccessRequestTemplatesAPIService
	id         string
}

func (r ApiDeleteAccessRequestTemplateRequest) Execute() (*MessageResponse, *http.Response, error) {
	return r.ApiService.DeleteAccessRequestTemplateExecute(r)
}

/*
DeleteAccessRequestTemplate Delete access request template

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiDeleteAccessRequestTemplateRequest
*/
func (a *AccessRequestTemplatesAPIService) DeleteAccessRequestTemplate(ctx context.Context, id string) ApiDeleteAccessRequestTemplateRequest {
	return ApiDeleteAccessRequestTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return MessageResponse
func (a *AccessRequestTemplatesAPIService) DeleteAccessRequestTemplateExecute(r ApiDeleteAccessRequestTemplateRequest) (*MessageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MessageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessRequestTemplatesAPIService.DeleteAccessRequestTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/client/v1/access-request-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAccessRequestTemplateRequest struct {
	ctx             context.Context
	ApiService      *AccessRequestTemplatesAPIService
	id              string
	filterAvailable *bool
}

func (r ApiGetAccessRequestTemplateRequest) FilterAvailable(filterAvailable bool) ApiGetAccessRequestTemplateRequest {
	r.filterAvailable = &filterAvailable
	return r
}

func (r ApiGetAccessRequestTemplateRequest) Execute() (*AccessRequestTemplateClientModel, *http.Response, error) {
	return r.ApiService.GetAccessRequestTemplateExecute(r)
}

/*
GetAccessRequestTemplate Get access request template by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiGetAccessRequestTemplateRequest
*/
func (a *AccessRequestTemplatesAPIService) GetAccessRequestTemplate(ctx context.Context, id string) ApiGetAccessRequestTemplateRequest {
	return ApiGetAccessRequestTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AccessRequestTemplateClientModel
func (a *AccessRequestTemplatesAPIService) GetAccessRequestTemplateExecute(r ApiGetAccessRequestTemplateRequest) (*AccessRequestTemplateClientModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessRequestTemplateClientModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessRequestTemplatesAPIService.GetAccessRequestTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/client/v1/access-request-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterAvailable != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter-available", r.filterAvailable, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAccessRequestTemplatesRequest struct {
	ctx        context.Context
	ApiService *AccessRequestTemplatesAPIService
	limit      *int32
	search     *string
	skip       *int32
}

func (r ApiListAccessRequestTemplatesRequest) Limit(limit int32) ApiListAccessRequestTemplatesRequest {
	r.limit = &limit
	return r
}

func (r ApiListAccessRequestTemplatesRequest) Search(search string) ApiListAccessRequestTemplatesRequest {
	r.search = &search
	return r
}

func (r ApiListAccessRequestTemplatesRequest) Skip(skip int32) ApiListAccessRequestTemplatesRequest {
	r.skip = &skip
	return r
}

func (r ApiListAccessRequestTemplatesRequest) Execute() (*PaginatedClientResponseModelAccessRequestTemplateClientModel, *http.Response, error) {
	return r.ApiService.ListAccessRequestTemplatesExecute(r)
}

/*
ListAccessRequestTemplates List access request templates

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListAccessRequestTemplatesRequest
*/
func (a *AccessRequestTemplatesAPIService) ListAccessRequestTemplates(ctx context.Context) ApiListAccessRequestTemplatesRequest {
	return ApiListAccessRequestTemplatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PaginatedClientResponseModelAccessRequestTemplateClientModel
func (a *AccessRequestTemplatesAPIService) ListAccessRequestTemplatesExecute(r ApiListAccessRequestTemplatesRequest) (*PaginatedClientResponseModelAccessRequestTemplateClientModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedClientResponseModelAccessRequestTemplateClientModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessRequestTemplatesAPIService.ListAccessRequestTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/client/v1/access-request-templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	} else {
		var defaultValue int32 = 0
		r.skip = &defaultValue
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateAccessRequestTemplateRequest struct {
	ctx                                             context.Context
	ApiService                                      *AccessRequestTemplatesAPIService
	id                                              string
	createAndUpdateAccessRequestTemplateClientModel *CreateAndUpdateAccessRequestTemplateClientModel
}

func (r ApiUpdateAccessRequestTemplateRequest) CreateAndUpdateAccessRequestTemplateClientModel(createAndUpdateAccessRequestTemplateClientModel CreateAndUpdateAccessRequestTemplateClientModel) ApiUpdateAccessRequestTemplateRequest {
	r.createAndUpdateAccessRequestTemplateClientModel = &createAndUpdateAccessRequestTemplateClientModel
	return r
}

func (r ApiUpdateAccessRequestTemplateRequest) Execute() (*AccessRequestTemplateClientModel, *http.Response, error) {
	return r.ApiService.UpdateAccessRequestTemplateExecute(r)
}

/*
UpdateAccessRequestTemplate Update access request template

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiUpdateAccessRequestTemplateRequest
*/
func (a *AccessRequestTemplatesAPIService) UpdateAccessRequestTemplate(ctx context.Context, id string) ApiUpdateAccessRequestTemplateRequest {
	return ApiUpdateAccessRequestTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AccessRequestTemplateClientModel
func (a *AccessRequestTemplatesAPIService) UpdateAccessRequestTemplateExecute(r ApiUpdateAccessRequestTemplateRequest) (*AccessRequestTemplateClientModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessRequestTemplateClientModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessRequestTemplatesAPIService.UpdateAccessRequestTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/client/v1/access-request-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createAndUpdateAccessRequestTemplateClientModel == nil {
		return localVarReturnValue, nil, reportError("createAndUpdateAccessRequestTemplateClientModel is required and must be specified")
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
	localVarPostBody = r.createAndUpdateAccessRequestTemplateClientModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
