# {{classname}}

All URIs are relative to *https://deep-index.moralis.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlock**](NativeApi.md#GetBlock) | **Get** /block/{block_number_or_hash} | Gets block contents by block hash
[**GetContractEvents**](NativeApi.md#GetContractEvents) | **Post** /{address}/events | Gets events by topic
[**GetDateToBlock**](NativeApi.md#GetDateToBlock) | **Get** /dateToBlock | Gets the closest block of the provided date
[**GetLogsByAddress**](NativeApi.md#GetLogsByAddress) | **Get** /{address}/logs | Gets address logs
[**GetNFTTransfersByBlock**](NativeApi.md#GetNFTTransfersByBlock) | **Get** /block/{block_number_or_hash}/nft/transfers | Gets NFT transfers by block number or block hash
[**GetTransaction**](NativeApi.md#GetTransaction) | **Get** /transaction/{transaction_hash} | Get transaction details by transaction hash
[**RunContractFunction**](NativeApi.md#RunContractFunction) | **Post** /{address}/function | Runs a function of a contract abi

# **GetBlock**
> Block GetBlock(ctx, blockNumberOrHash, optional)
Gets block contents by block hash

Gets the contents of a block by block hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockNumberOrHash** | **string**| The block hash or block number | 
 **optional** | ***NativeApiGetBlockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativeApiGetBlockOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 

### Return type

[**Block**](block.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContractEvents**
> []LogEvent GetContractEvents(ctx, address, topic, optional)
Gets events by topic

Gets events in descending order based on block number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| address | 
  **topic** | **string**| The topic of the event | 
 **optional** | ***NativeApiGetContractEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativeApiGetContractEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)| ABI of the specific event | 
 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 
 **providerUrl** | **optional.**| web3 provider url to user when using local dev chain | 
 **fromBlock** | **optional.**| The minimum block number from where to get the logs * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toBlock** | **optional.**| The maximum block number from where to get the logs. * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **fromDate** | **optional.**| The date from where to get the logs (any format that is accepted by momentjs) * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toDate** | **optional.**| Get the logs to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **offset** | **optional.**| offset | 
 **limit** | **optional.**| limit | 

### Return type

[**[]LogEvent**](logEvent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDateToBlock**
> BlockDate GetDateToBlock(ctx, date, optional)
Gets the closest block of the provided date

Gets the closest block of the provided date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **date** | **string**| Unix date in miliseconds or a datestring (any format that is accepted by momentjs) | 
 **optional** | ***NativeApiGetDateToBlockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativeApiGetDateToBlockOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **providerUrl** | **optional.String**| web3 provider url to user when using local dev chain | 

### Return type

[**BlockDate**](blockDate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogsByAddress**
> LogEventByAddress GetLogsByAddress(ctx, address, optional)
Gets address logs

Gets the logs from an address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| address | 
 **optional** | ***NativeApiGetLogsByAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativeApiGetLogsByAddressOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 
 **blockNumber** | **optional.String**| The block number * Provide the param &#x27;block_numer&#x27; or (&#x27;from_block&#x27; and / or &#x27;to_block&#x27;) * If &#x27;block_numer&#x27; is provided in conbinaison with &#x27;from_block&#x27; and / or &#x27;to_block&#x27;, &#x27;block_number&#x27; will will be used  | 
 **fromBlock** | **optional.String**| The minimum block number from where to get the logs * Provide the param &#x27;block_numer&#x27; or (&#x27;from_block&#x27; and / or &#x27;to_block&#x27;) * If &#x27;block_numer&#x27; is provided in conbinaison with &#x27;from_block&#x27; and / or &#x27;to_block&#x27;, &#x27;block_number&#x27; will will be used  | 
 **toBlock** | **optional.String**| The maximum block number from where to get the logs * Provide the param &#x27;block_numer&#x27; or (&#x27;from_block&#x27; and / or &#x27;to_block&#x27;) * If &#x27;block_numer&#x27; is provided in conbinaison with &#x27;from_block&#x27; and / or &#x27;to_block&#x27;, &#x27;block_number&#x27; will will be used  | 
 **fromDate** | **optional.String**| The date from where to get the logs (any format that is accepted by momentjs) * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used. * If &#x27;from_date&#x27; and the block params are provided, the block params will be used. Please refer to the blocks params sections (block_number,from_block and to_block) on how to use them  | 
 **toDate** | **optional.String**| Get the logs to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used. * If &#x27;to_date&#x27; and the block params are provided, the block params will be used. Please refer to the blocks params sections (block_number,from_block and to_block) on how to use them  | 
 **topic0** | **optional.String**| topic0 | 
 **topic1** | **optional.String**| topic1 | 
 **topic2** | **optional.String**| topic2 | 
 **topic3** | **optional.String**| topic3 | 

### Return type

[**LogEventByAddress**](logEventByAddress.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFTTransfersByBlock**
> NftTransferCollection GetNFTTransfersByBlock(ctx, blockNumberOrHash, optional)
Gets NFT transfers by block number or block hash

Gets NFT transfers by block number or block hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockNumberOrHash** | **string**| The block hash or block number | 
 **optional** | ***NativeApiGetNFTTransfersByBlockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativeApiGetNFTTransfersByBlockOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 
 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| limit | [default to 500]
 **cursor** | **optional.String**| The cursor returned in the last response (for getting the next page)  | 

### Return type

[**NftTransferCollection**](nftTransferCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransaction**
> BlockTransaction GetTransaction(ctx, transactionHash, optional)
Get transaction details by transaction hash

Gets the contents of a block transaction by hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionHash** | **string**| The transaction hash | 
 **optional** | ***NativeApiGetTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativeApiGetTransactionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 

### Return type

[**BlockTransaction**](blockTransaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunContractFunction**
> string RunContractFunction(ctx, body, address, functionName, optional)
Runs a function of a contract abi

Runs a given function of a contract abi and returns readonly data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RunContractDto**](RunContractDto.md)| Body | 
  **address** | **string**| address | 
  **functionName** | **string**| function_name | 
 **optional** | ***NativeApiRunContractFunctionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativeApiRunContractFunctionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 
 **providerUrl** | **optional.**| web3 provider url to user when using local dev chain | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

