# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3HolidayDatesPaidholidaysGet**](HolidayDatesApi.md#V3HolidayDatesPaidholidaysGet) | **Get** /v3/HolidayDates/paidholidays | Retrieves data for all paid holidays as published by Human Resources
[**V3HolidayDatesPaidholidaysIcsGet**](HolidayDatesApi.md#V3HolidayDatesPaidholidaysIcsGet) | **Get** /v3/HolidayDates/paidholidays/ics | Retrieves data for all paid holidays as published by Human Resources, as an ICS format feed. Allows anonymous access.
[**V3HolidayDatesPaidholidaysYearGet**](HolidayDatesApi.md#V3HolidayDatesPaidholidaysYearGet) | **Get** /v3/HolidayDates/paidholidays/{year} | Retrieves data for paid holidays associated to the given year as published by Human Resources

# **V3HolidayDatesPaidholidaysGet**
> []PaidHoliday V3HolidayDatesPaidholidaysGet(ctx, )
Retrieves data for all paid holidays as published by Human Resources

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PaidHoliday**](PaidHoliday.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3HolidayDatesPaidholidaysIcsGet**
> V3HolidayDatesPaidholidaysIcsGet(ctx, )
Retrieves data for all paid holidays as published by Human Resources, as an ICS format feed. Allows anonymous access.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3HolidayDatesPaidholidaysYearGet**
> []PaidHoliday V3HolidayDatesPaidholidaysYearGet(ctx, year)
Retrieves data for paid holidays associated to the given year as published by Human Resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**|  | 

### Return type

[**[]PaidHoliday**](PaidHoliday.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

