# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3LocationsGeojsonGet**](LocationsApi.md#V3LocationsGeojsonGet) | **Get** /v3/Locations/geojson | Get all building location data as GEO JSON
[**V3LocationsGet**](LocationsApi.md#V3LocationsGet) | **Get** /v3/Locations | Get all building location data
[**V3LocationsLocationCodeGeojsonGet**](LocationsApi.md#V3LocationsLocationCodeGeojsonGet) | **Get** /v3/Locations/{locationCode}/geojson | Gets building by matched building code, case insensitive, as GEO JSON
[**V3LocationsLocationCodeGet**](LocationsApi.md#V3LocationsLocationCodeGet) | **Get** /v3/Locations/{locationCode} | Gets building by matched building code, case insensitive
[**V3LocationsSearchLocationNameGeojsonGet**](LocationsApi.md#V3LocationsSearchLocationNameGeojsonGet) | **Get** /v3/Locations/search/{locationName}/geojson | Gets buildings by matched building name, contains, case insensitive, as GEO JSON
[**V3LocationsSearchLocationNameGet**](LocationsApi.md#V3LocationsSearchLocationNameGet) | **Get** /v3/Locations/search/{locationName} | Gets buildings by matched building name, contains, case insensitive

# **V3LocationsGeojsonGet**
> LocationGeo V3LocationsGeojsonGet(ctx, )
Get all building location data as GEO JSON

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LocationGeo**](LocationGeo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3LocationsGet**
> []Location V3LocationsGet(ctx, )
Get all building location data

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Location**](Location.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3LocationsLocationCodeGeojsonGet**
> LocationGeo V3LocationsLocationCodeGeojsonGet(ctx, locationCode)
Gets building by matched building code, case insensitive, as GEO JSON

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationCode** | **string**| Building code to match | 

### Return type

[**LocationGeo**](LocationGeo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3LocationsLocationCodeGet**
> Location V3LocationsLocationCodeGet(ctx, locationCode)
Gets building by matched building code, case insensitive

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationCode** | **string**| Building code to match | 

### Return type

[**Location**](Location.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3LocationsSearchLocationNameGeojsonGet**
> LocationGeo V3LocationsSearchLocationNameGeojsonGet(ctx, locationName)
Gets buildings by matched building name, contains, case insensitive, as GEO JSON

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationName** | **string**| Text to match in building name | 

### Return type

[**LocationGeo**](LocationGeo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3LocationsSearchLocationNameGet**
> []Location V3LocationsSearchLocationNameGet(ctx, locationName)
Gets buildings by matched building name, contains, case insensitive

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationName** | **string**| Text to match in building name | 

### Return type

[**[]Location**](Location.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

