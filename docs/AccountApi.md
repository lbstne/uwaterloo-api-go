# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AccountConfirmPost**](AccountApi.md#V3AccountConfirmPost) | **Post** /v3/Account/Confirm | Use this method to confirm your email and activate your account
[**V3AccountEmailApiKeyResetGet**](AccountApi.md#V3AccountEmailApiKeyResetGet) | **Get** /v3/Account/{email}/{apiKey}/reset | User this method to put your account in pending confirmation status and generate a new API key. Your old key will no longer grant access. The account will need to be confirmed again before the new key grants access.
[**V3AccountEmailGet**](AccountApi.md#V3AccountEmailGet) | **Get** /v3/Account/{email} | Use this method to see if an email has already been registered for an API key
[**V3AccountEmailNotifyGet**](AccountApi.md#V3AccountEmailNotifyGet) | **Get** /v3/Account/{email}/notify | Use this method to have us re-send the confirmation email to an account pending confirmation, if it exists. The activation code will be reset in the process.
[**V3AccountRegisterPost**](AccountApi.md#V3AccountRegisterPost) | **Post** /v3/Account/Register | Use this method to request an API key and begin your registration process

# **V3AccountConfirmPost**
> V3AccountConfirmPost(ctx, optional)
Use this method to confirm your email and activate your account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiV3AccountConfirmPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiV3AccountConfirmPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| Email address you used to register | 
 **code** | [**optional.Interface of string**](.md)| Activation code we sent you in the confirmation email | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3AccountEmailApiKeyResetGet**
> V3AccountEmailApiKeyResetGet(ctx, email, apiKey)
User this method to put your account in pending confirmation status and generate a new API key. Your old key will no longer grant access. The account will need to be confirmed again before the new key grants access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| Email address of associated account to reset | 
  **apiKey** | **string**| Current API account key | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3AccountEmailGet**
> V3AccountEmailGet(ctx, email)
Use this method to see if an email has already been registered for an API key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| Email address to check | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3AccountEmailNotifyGet**
> V3AccountEmailNotifyGet(ctx, email)
Use this method to have us re-send the confirmation email to an account pending confirmation, if it exists. The activation code will be reset in the process.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| Email address to re-send confirmation email to | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V3AccountRegisterPost**
> V3AccountRegisterPost(ctx, optional)
Use this method to request an API key and begin your registration process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiV3AccountRegisterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiV3AccountRegisterPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| A unique email that we can use to identify you and contact you. We will ask to confirm it. | 
 **project** | **optional.String**| A name and description of your project | 
 **uri** | **optional.String**| The web address of your project: it&#x27;s live location, app store entry, or similar | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

