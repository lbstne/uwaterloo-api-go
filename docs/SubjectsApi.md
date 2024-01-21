# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3SubjectsAssociatedtoOrganizationCodeGet**](SubjectsApi.md#V3SubjectsAssociatedtoOrganizationCodeGet) | **Get** /v3/Subjects/associatedto/{organizationCode} | Gets Subject data for Subjects associated to an Academic Organization by Organization code
[**V3SubjectsCodeGet**](SubjectsApi.md#V3SubjectsCodeGet) | **Get** /v3/Subjects/{code} | Gets Subject data filtered by Subject code
[**V3SubjectsGet**](SubjectsApi.md#V3SubjectsGet) | **Get** /v3/Subjects | Gets all Subject data

# **V3SubjectsAssociatedtoOrganizationCodeGet**
> []Subject V3SubjectsAssociatedtoOrganizationCodeGet(ctx, organizationCode)
Gets Subject data for Subjects associated to an Academic Organization by Organization code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationCode** | **string**| Academic Organization Code that associates to the Subjects | 

### Return type

[**[]Subject**](Subject.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3SubjectsCodeGet**
> Subject V3SubjectsCodeGet(ctx, code)
Gets Subject data filtered by Subject code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| Specific Subject code | 

### Return type

[**Subject**](Subject.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3SubjectsGet**
> []Subject V3SubjectsGet(ctx, )
Gets all Subject data

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Subject**](Subject.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

