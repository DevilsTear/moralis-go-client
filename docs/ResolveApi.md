# {{classname}}

All URIs are relative to *https://deep-index.moralis.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResolveAddress**](ResolveApi.md#ResolveAddress) | **Get** /resolve/{address}/reverse | Return the ENS domain when available (Only for ETH)
[**ResolveDomain**](ResolveApi.md#ResolveDomain) | **Get** /resolve/{domain} | Resolves an Unstoppable domain and returns the address

# **ResolveAddress**
> Ens ResolveAddress(ctx, address)
Return the ENS domain when available (Only for ETH)

Resolves an ETH address and find the ENS name 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address to be resolved | 

### Return type

[**Ens**](ens.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResolveDomain**
> Resolve ResolveDomain(ctx, domain, optional)
Resolves an Unstoppable domain and returns the address

Resolves an Unstoppable domain and returns the address 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain to be resolved | 
 **optional** | ***ResolveApiResolveDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResolveApiResolveDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currency** | **optional.String**| The currency to query | [default to eth]

### Return type

[**Resolve**](resolve.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

