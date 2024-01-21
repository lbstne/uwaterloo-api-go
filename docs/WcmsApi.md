# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3WcmsGet**](WcmsApi.md#V3WcmsGet) | **Get** /v3/Wcms | Retrieves information about all active and published WCMS sites
[**V3WcmsIdEventsGet**](WcmsApi.md#V3WcmsIdEventsGet) | **Get** /v3/Wcms/{id}/events | Retrieves all event items for a specific WCMS site by Site Id
[**V3WcmsIdGet**](WcmsApi.md#V3WcmsIdGet) | **Get** /v3/Wcms/{id} | Retrieves information about a specific WCMS site by Site Id
[**V3WcmsIdNewsGet**](WcmsApi.md#V3WcmsIdNewsGet) | **Get** /v3/Wcms/{id}/news | Retrieves all news items for a specific WCMS site by Site Id
[**V3WcmsIdOpportunitiesGet**](WcmsApi.md#V3WcmsIdOpportunitiesGet) | **Get** /v3/Wcms/{id}/opportunities | Retrieves all opportunity items for a specific WCMS site by Site Id
[**V3WcmsIdPostsGet**](WcmsApi.md#V3WcmsIdPostsGet) | **Get** /v3/Wcms/{id}/posts | Retrieves all blog post items for a specific WCMS site by Site Id
[**V3WcmsLatesteventsMaxItemsGet**](WcmsApi.md#V3WcmsLatesteventsMaxItemsGet) | **Get** /v3/Wcms/latestevents/{maxItems} | Retrieves the latest events items across all WCMS sites, ordered by event start date
[**V3WcmsLatestnewsMaxItemsGet**](WcmsApi.md#V3WcmsLatestnewsMaxItemsGet) | **Get** /v3/Wcms/latestnews/{maxItems} | Retrieves the latest news items across all WCMS sites, ordered by posted date
[**V3WcmsLatestopportunitiesMaxItemsGet**](WcmsApi.md#V3WcmsLatestopportunitiesMaxItemsGet) | **Get** /v3/Wcms/latestopportunities/{maxItems} | Retrieves the latest opportunity items across all WCMS sites, ordered by posted/open date
[**V3WcmsLatestpostsMaxItemsGet**](WcmsApi.md#V3WcmsLatestpostsMaxItemsGet) | **Get** /v3/Wcms/latestposts/{maxItems} | Retrieves the latest blog post items across all WCMS sites, ordered by posted date

# **V3WcmsGet**
> []Site V3WcmsGet(ctx, )
Retrieves information about all active and published WCMS sites

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Site**](Site.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3WcmsIdEventsGet**
> []SiteEvent V3WcmsIdEventsGet(ctx, id, optional)
Retrieves all event items for a specific WCMS site by Site Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Unique Id for the Site | 
 **optional** | ***WcmsApiV3WcmsIdEventsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WcmsApiV3WcmsIdEventsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxItems** | **optional.Int32**| (Optional) Maximum amount of items to retrieve | 
 **newestFirst** | **optional.Bool**| (Optional) Requires maxItems parameter, sorts items by created date, newest first | 

### Return type

[**[]SiteEvent**](SiteEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3WcmsIdGet**
> Site V3WcmsIdGet(ctx, id)
Retrieves information about a specific WCMS site by Site Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Unique site Id | 

### Return type

[**Site**](Site.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3WcmsIdNewsGet**
> []SiteNews V3WcmsIdNewsGet(ctx, id, optional)
Retrieves all news items for a specific WCMS site by Site Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Unique Id for the Site | 
 **optional** | ***WcmsApiV3WcmsIdNewsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WcmsApiV3WcmsIdNewsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxItems** | **optional.Int32**| (Optional) Maximum amount of items to retrieve | 
 **newestFirst** | **optional.Bool**| (Optional) Requires maxItems parameter, sorts items by created date, newest first | 

### Return type

[**[]SiteNews**](SiteNews.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3WcmsIdOpportunitiesGet**
> []SiteOpportunity V3WcmsIdOpportunitiesGet(ctx, id, optional)
Retrieves all opportunity items for a specific WCMS site by Site Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Unique Id for the Site | 
 **optional** | ***WcmsApiV3WcmsIdOpportunitiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WcmsApiV3WcmsIdOpportunitiesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxItems** | **optional.Int32**| (Optional) Maximum amount of items to retrieve | 
 **newestFirst** | **optional.Bool**| (Optional) Requires maxItems parameter, sorts items by created date, newest first | 

### Return type

[**[]SiteOpportunity**](SiteOpportunity.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3WcmsIdPostsGet**
> []SiteBlogPost V3WcmsIdPostsGet(ctx, id, optional)
Retrieves all blog post items for a specific WCMS site by Site Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Unique Id for the Site | 
 **optional** | ***WcmsApiV3WcmsIdPostsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WcmsApiV3WcmsIdPostsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxItems** | **optional.Int32**| (Optional) Maximum amount of items to retrieve | 
 **newestFirst** | **optional.Bool**| (Optional) Requires maxItems parameter, sorts items by created date, newest first | 

### Return type

[**[]SiteBlogPost**](SiteBlogPost.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3WcmsLatesteventsMaxItemsGet**
> []SiteEvent V3WcmsLatesteventsMaxItemsGet(ctx, maxItems)
Retrieves the latest events items across all WCMS sites, ordered by event start date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maxItems** | **int32**| Number of items to retrieve, default 15, maximum 25 | 

### Return type

[**[]SiteEvent**](SiteEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3WcmsLatestnewsMaxItemsGet**
> []SiteNews V3WcmsLatestnewsMaxItemsGet(ctx, maxItems)
Retrieves the latest news items across all WCMS sites, ordered by posted date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maxItems** | **int32**| Number of items to retrieve, default 15, maximum 25 | 

### Return type

[**[]SiteNews**](SiteNews.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3WcmsLatestopportunitiesMaxItemsGet**
> []SiteOpportunity V3WcmsLatestopportunitiesMaxItemsGet(ctx, maxItems)
Retrieves the latest opportunity items across all WCMS sites, ordered by posted/open date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maxItems** | **int32**| Number of items to retrieve, default 15, maximum 25 | 

### Return type

[**[]SiteOpportunity**](SiteOpportunity.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3WcmsLatestpostsMaxItemsGet**
> []SiteBlogPost V3WcmsLatestpostsMaxItemsGet(ctx, maxItems)
Retrieves the latest blog post items across all WCMS sites, ordered by posted date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maxItems** | **int32**| Number of items to retrieve, default 15, maximum 25 | 

### Return type

[**[]SiteBlogPost**](SiteBlogPost.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

