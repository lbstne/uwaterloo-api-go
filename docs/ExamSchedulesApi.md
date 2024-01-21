# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3ExamSchedulesCodeGet**](ExamSchedulesApi.md#V3ExamSchedulesCodeGet) | **Get** /v3/ExamSchedules/{code} | Returns Exam Schedule data for the requested Term
[**V3ExamSchedulesGet**](ExamSchedulesApi.md#V3ExamSchedulesGet) | **Get** /v3/ExamSchedules | Returns Exam Schedule data for the current Term

# **V3ExamSchedulesCodeGet**
> []Exam V3ExamSchedulesCodeGet(ctx, code)
Returns Exam Schedule data for the requested Term

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| Waterloo Term code | 

### Return type

[**[]Exam**](Exam.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3ExamSchedulesGet**
> []Exam V3ExamSchedulesGet(ctx, )
Returns Exam Schedule data for the current Term

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Exam**](Exam.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

