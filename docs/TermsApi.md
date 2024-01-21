# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3TermsCodeGet**](TermsApi.md#V3TermsCodeGet) | **Get** /v3/Terms/{code} | Gets Term data for a specific Term filtered by code
[**V3TermsCurrentGet**](TermsApi.md#V3TermsCurrentGet) | **Get** /v3/Terms/current | Gets the current Term data
[**V3TermsForacademicyearYearGet**](TermsApi.md#V3TermsForacademicyearYearGet) | **Get** /v3/Terms/foracademicyear/{year} | Gets Term data for terms that are part of a specific academic year
[**V3TermsGet**](TermsApi.md#V3TermsGet) | **Get** /v3/Terms | Gets all Term data that is effective at the current time

# **V3TermsCodeGet**
> Term V3TermsCodeGet(ctx, code)
Gets Term data for a specific Term filtered by code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| Term 4 digit Code | 

### Return type

[**Term**](Term.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3TermsCurrentGet**
> Term V3TermsCurrentGet(ctx, )
Gets the current Term data

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Term**](Term.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3TermsForacademicyearYearGet**
> []Term V3TermsForacademicyearYearGet(ctx, year)
Gets Term data for terms that are part of a specific academic year

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Academic year associated to Terms | 

### Return type

[**[]Term**](Term.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3TermsGet**
> []Term V3TermsGet(ctx, )
Gets all Term data that is effective at the current time

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Term**](Term.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

