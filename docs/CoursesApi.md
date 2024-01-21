# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3CoursesTermCodeCourseIdGet**](CoursesApi.md#V3CoursesTermCodeCourseIdGet) | **Get** /v3/Courses/{termCode}/{courseId} | Get Course catalog data filtered by Term and Course ID, multiple Courses are usually cross listed
[**V3CoursesTermCodeCourseIdOfferNumberGet**](CoursesApi.md#V3CoursesTermCodeCourseIdOfferNumberGet) | **Get** /v3/Courses/{termCode}/{courseId}/{offerNumber} | Get Course catalog data filtered by Term, Course ID, and offer number
[**V3CoursesTermCodeGet**](CoursesApi.md#V3CoursesTermCodeGet) | **Get** /v3/Courses/{termCode} | Get all Course data for a Term
[**V3CoursesTermCodeSubjectCatalogNumberGet**](CoursesApi.md#V3CoursesTermCodeSubjectCatalogNumberGet) | **Get** /v3/Courses/{termCode}/{subject}/{catalogNumber} | Get Course catalog data filtered by Term, Subject, and catalog number
[**V3CoursesTermCodeSubjectGet**](CoursesApi.md#V3CoursesTermCodeSubjectGet) | **Get** /v3/Courses/{termCode}/{subject} | Get Course catalog data filtered by Term and Subject code

# **V3CoursesTermCodeCourseIdGet**
> []Course V3CoursesTermCodeCourseIdGet(ctx, termCode, courseId)
Get Course catalog data filtered by Term and Course ID, multiple Courses are usually cross listed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termCode** | **string**| Term code to filter on | 
  **courseId** | **int32**| Course ID to filter on | 

### Return type

[**[]Course**](Course.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3CoursesTermCodeCourseIdOfferNumberGet**
> Course V3CoursesTermCodeCourseIdOfferNumberGet(ctx, termCode, courseId, offerNumber)
Get Course catalog data filtered by Term, Course ID, and offer number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termCode** | **string**| Term code to filter on | 
  **courseId** | **int32**| Course ID to filter on | 
  **offerNumber** | **int32**| Offer number to filter on | 

### Return type

[**Course**](Course.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3CoursesTermCodeGet**
> []Course V3CoursesTermCodeGet(ctx, termCode)
Get all Course data for a Term

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termCode** | **string**|  | 

### Return type

[**[]Course**](Course.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3CoursesTermCodeSubjectCatalogNumberGet**
> Course V3CoursesTermCodeSubjectCatalogNumberGet(ctx, termCode, subject, catalogNumber)
Get Course catalog data filtered by Term, Subject, and catalog number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termCode** | **string**| Term code to filter on | 
  **subject** | **string**| Academic Subject code to filter on | 
  **catalogNumber** | **string**| Course catalog number to filter on, ie: 101, 111W, 239 | 

### Return type

[**Course**](Course.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3CoursesTermCodeSubjectGet**
> []Course V3CoursesTermCodeSubjectGet(ctx, termCode, subject)
Get Course catalog data filtered by Term and Subject code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termCode** | **string**| Term code to filter on | 
  **subject** | **string**| Academic Subject code to filter on | 

### Return type

[**[]Course**](Course.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

