# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3FoodServicesFranchisesGet**](FoodServicesApi.md#V3FoodServicesFranchisesGet) | **Get** /v3/FoodServices/franchises | Get all food service Franchise data
[**V3FoodServicesFranchisesIdGet**](FoodServicesApi.md#V3FoodServicesFranchisesIdGet) | **Get** /v3/FoodServices/franchises/{id} | Get specific food sercices Franchise data by Id
[**V3FoodServicesFranchisesNameGet**](FoodServicesApi.md#V3FoodServicesFranchisesNameGet) | **Get** /v3/FoodServices/franchises/{name} | Get specific food services Franchise data by Franchise name
[**V3FoodServicesOutletsGet**](FoodServicesApi.md#V3FoodServicesOutletsGet) | **Get** /v3/FoodServices/outlets | Get all food service Outlet data
[**V3FoodServicesOutletsIdGet**](FoodServicesApi.md#V3FoodServicesOutletsIdGet) | **Get** /v3/FoodServices/outlets/{id} | Get specific food services Outlet data by Id
[**V3FoodServicesOutletsNameGet**](FoodServicesApi.md#V3FoodServicesOutletsNameGet) | **Get** /v3/FoodServices/outlets/{name} | Get specific food services Outlet data by Outlet name

# **V3FoodServicesFranchisesGet**
> FoodServicesFranchises V3FoodServicesFranchisesGet(ctx, )
Get all food service Franchise data

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FoodServicesFranchises**](FoodServicesFranchises.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3FoodServicesFranchisesIdGet**
> FoodServicesFranchise V3FoodServicesFranchisesIdGet(ctx, id)
Get specific food sercices Franchise data by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| WCMS object ID representing the Franchise | 

### Return type

[**FoodServicesFranchise**](FoodServicesFranchise.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3FoodServicesFranchisesNameGet**
> FoodServicesFranchise V3FoodServicesFranchisesNameGet(ctx, name)
Get specific food services Franchise data by Franchise name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name to filter Franchise by | 

### Return type

[**FoodServicesFranchise**](FoodServicesFranchise.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3FoodServicesOutletsGet**
> FoodServicesOutlets V3FoodServicesOutletsGet(ctx, )
Get all food service Outlet data

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FoodServicesOutlets**](FoodServicesOutlets.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3FoodServicesOutletsIdGet**
> FoodServicesOutlet V3FoodServicesOutletsIdGet(ctx, id)
Get specific food services Outlet data by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| WCMS object ID representing the Outlet | 

### Return type

[**FoodServicesOutlet**](FoodServicesOutlet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3FoodServicesOutletsNameGet**
> FoodServicesOutlet V3FoodServicesOutletsNameGet(ctx, name)
Get specific food services Outlet data by Outlet name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name to filter Outlet by | 

### Return type

[**FoodServicesOutlet**](FoodServicesOutlet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

