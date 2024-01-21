# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3ClassSchedulesTermCodeCourseIdGet**](ClassSchedulesApi.md#V3ClassSchedulesTermCodeCourseIdGet) | **Get** /v3/ClassSchedules/{termCode}/{courseId} | Get Class data for a scheduled Course by Course ID, in a specific Term
[**V3ClassSchedulesTermCodeGet**](ClassSchedulesApi.md#V3ClassSchedulesTermCodeGet) | **Get** /v3/ClassSchedules/{termCode} | Get the Course IDs that have one or more active and associated schedules in the given Term
[**V3ClassSchedulesTermCodeSubjectCatalogNumberGet**](ClassSchedulesApi.md#V3ClassSchedulesTermCodeSubjectCatalogNumberGet) | **Get** /v3/ClassSchedules/{termCode}/{subject}/{catalogNumber} | Get Class data for a scheduled Course by Subject and catalog number, in a specific Term

# **V3ClassSchedulesTermCodeCourseIdGet**
> []Class V3ClassSchedulesTermCodeCourseIdGet(ctx, termCode, courseId)
Get Class data for a scheduled Course by Course ID, in a specific Term

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termCode** | **string**| Waterloo Term code to filter on | 
  **courseId** | **int32**| Course ID to filter on | 

### Return type

[**[]Class**](Class.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3ClassSchedulesTermCodeGet**
> []string V3ClassSchedulesTermCodeGet(ctx, termCode)
Get the Course IDs that have one or more active and associated schedules in the given Term

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termCode** | **string**| Waterloo Term code to filter on | 

### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3ClassSchedulesTermCodeSubjectCatalogNumberGet**
> []Class V3ClassSchedulesTermCodeSubjectCatalogNumberGet(ctx, termCode, subject, catalogNumber)
Get Class data for a scheduled Course by Subject and catalog number, in a specific Term

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termCode** | **string**| Waterloo Term code to filter on | 
  **subject** | **string**| Academic Subject code to filter ron | 
  **catalogNumber** | **string**| Catalog number to filter on | 

### Return type

[**[]Class**](Class.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

