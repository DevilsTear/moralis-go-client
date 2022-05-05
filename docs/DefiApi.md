# {{classname}}

All URIs are relative to *https://deep-index.moralis.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPairAddress**](DefiApi.md#GetPairAddress) | **Get** /{token0_address}/{token1_address}/pairAddress | Get pair address based on token0 and token1 address
[**GetPairReserves**](DefiApi.md#GetPairReserves) | **Get** /{pair_address}/reserves | Get liquidity pair reserves for an Uniswap V2 based Exchange.

# **GetPairAddress**
> ReservesCollection GetPairAddress(ctx, exchange, token0Address, token1Address, optional)
Get pair address based on token0 and token1 address

Fetches and returns pair data of the provided token0+token1 combination. The token0 and token1 options are interchangable (ie. there is no different outcome in \"token0=WETH and token1=USDT\" or \"token0=USDT and token1=WETH\") 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exchange** | **string**| The factory name or address of the token exchange | 
  **token0Address** | **string**| Token0 address | 
  **token1Address** | **string**| Token1 address | 
 **optional** | ***DefiApiGetPairAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefiApiGetPairAddressOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **toBlock** | **optional.String**| To get the reserves at this block number | 
 **toDate** | **optional.String**| Get the reserves to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 

### Return type

[**ReservesCollection**](reservesCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPairReserves**
> ReservesCollection GetPairReserves(ctx, pairAddress, optional)
Get liquidity pair reserves for an Uniswap V2 based Exchange.

Get the liquidity reserves for a given pair address. Only Uniswap V2 based exchanges supported at the moment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairAddress** | **string**| Liquidity pair address | 
 **optional** | ***DefiApiGetPairReservesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefiApiGetPairReservesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **toBlock** | **optional.String**| To get the reserves at this block number | 
 **toDate** | **optional.String**| Get the reserves to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **providerUrl** | **optional.String**| web3 provider url to user when using local dev chain | 

### Return type

[**ReservesCollection**](reservesCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

