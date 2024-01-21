# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AcademicOrganizationsGet**](AcademicOrganizationsApi.md#V3AcademicOrganizationsGet) | **Get** /v3/AcademicOrganizations | Gets all Academic Organization data
[**V3AcademicOrganizationsOrganizationCodeGet**](AcademicOrganizationsApi.md#V3AcademicOrganizationsOrganizationCodeGet) | **Get** /v3/AcademicOrganizations/{organizationCode} | Gets Academic Organization data for a specific entry by the Organization code

# **V3AcademicOrganizationsGet**
> []AcademicOrganization V3AcademicOrganizationsGet(ctx, )
Gets all Academic Organization data

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AcademicOrganization**](AcademicOrganization.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3AcademicOrganizationsOrganizationCodeGet**
> AcademicOrganization V3AcademicOrganizationsOrganizationCodeGet(ctx, organizationCode)
Gets Academic Organization data for a specific entry by the Organization code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationCode** | **string**| Unique Academic Organization code | 

### Return type

[**AcademicOrganization**](AcademicOrganization.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

