# {{classname}}

All URIs are relative to *https://deep-index.moralis.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNFTTransfers**](AccountApi.md#GetNFTTransfers) | **Get** /{address}/nft/transfers | Gets NFT transfers to and from a given address
[**GetNFTs**](AccountApi.md#GetNFTs) | **Get** /{address}/nft | Gets the NFTs owned by a given address
[**GetNFTsForContract**](AccountApi.md#GetNFTsForContract) | **Get** /{address}/nft/{token_address} | Gets the NFTs owned by a given address
[**GetNativeBalance**](AccountApi.md#GetNativeBalance) | **Get** /{address}/balance | Gets native balance for a specific address.
[**GetTokenBalances**](AccountApi.md#GetTokenBalances) | **Get** /{address}/erc20 | Gets token balances for a specific address.
[**GetTokenTransfers**](AccountApi.md#GetTokenTransfers) | **Get** /{address}/erc20/transfers | Gets erc 20 token transactions
[**GetTransactions**](AccountApi.md#GetTransactions) | **Get** /{address} | Gets native transactions

# **GetNFTTransfers**
> NftTransferCollection GetNFTTransfers(ctx, address, optional)
Gets NFT transfers to and from a given address

Gets the transfers of the tokens matching the given parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The sender or recepient of the transfer | 
 **optional** | ***AccountApiGetNFTTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetNFTTransfersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
 **direction** | **optional.String**| The transfer direction | [default to both]
 **limit** | **optional.Int32**| limit | 
 **cursor** | **optional.String**| The cursor returned in the last response (for getting the next page)  | 

### Return type

[**NftTransferCollection**](nftTransferCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFTs**
> NftOwnerCollection GetNFTs(ctx, address, optional)
Gets the NFTs owned by a given address

Gets NFTs owned by the given address * The response will include status [SYNCED/SYNCING] based on the contracts being indexed. * Use the token_address param to get results for a specific contract only * Note results will include all indexed NFTs * Any request which includes the token_address param will start the indexing process for that NFT collection the very first time it is requested 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The owner of a given token | 
 **optional** | ***AccountApiGetNFTsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetNFTsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
 **limit** | **optional.Int32**| limit | 
 **tokenAddresses** | [**optional.Interface of []string**](string.md)| The addresses to get balances for (Optional) | 
 **cursor** | **optional.String**| The cursor returned in the last response (for getting the next page)  | 

### Return type

[**NftOwnerCollection**](nftOwnerCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFTsForContract**
> NftOwnerCollection GetNFTsForContract(ctx, address, tokenAddress, optional)
Gets the NFTs owned by a given address

Gets NFTs owned by the given address * Use the token_address param to get results for a specific contract only * Note results will include all indexed NFTs * Any request which includes the token_address param will start the indexing process for that NFT collection the very first time it is requested 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The owner of a given token | 
  **tokenAddress** | **string**| Address of the contract | 
 **optional** | ***AccountApiGetNFTsForContractOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetNFTsForContractOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
 **cursor** | **optional.String**| The cursor returned in the last response (for getting the next page)  | 
 **limit** | **optional.Int32**| limit | 

### Return type

[**NftOwnerCollection**](nftOwnerCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNativeBalance**
> NativeBalance GetNativeBalance(ctx, address, optional)
Gets native balance for a specific address.

Gets native balance for a specific address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address for which the native balance will be checked | 
 **optional** | ***AccountApiGetNativeBalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetNativeBalanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **providerUrl** | **optional.String**| web3 provider url to user when using local dev chain | 
 **toBlock** | **optional.Float64**| The block number on which the balances should be checked | 

### Return type

[**NativeBalance**](nativeBalance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenBalances**
> []Erc20TokenBalance GetTokenBalances(ctx, address, optional)
Gets token balances for a specific address.

Gets token balances for a specific address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address for which token balances will be checked | 
 **optional** | ***AccountApiGetTokenBalancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetTokenBalancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 
 **toBlock** | **optional.Float64**| The block number on which the balances should be checked | 
 **tokenAddresses** | [**optional.Interface of []string**](string.md)| The addresses to get balances for (Optional) | 

### Return type

[**[]Erc20TokenBalance**](erc20TokenBalance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenTransfers**
> Erc20TransactionCollection GetTokenTransfers(ctx, address, optional)
Gets erc 20 token transactions

Gets ERC20 token transactions in descending order based on block number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| address | 
 **optional** | ***AccountApiGetTokenTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetTokenTransfersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 
 **fromBlock** | **optional.Int32**| The minimum block number from where to get the transactions * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toBlock** | **optional.Int32**| The maximum block number from where to get the transactions. * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **fromDate** | **optional.String**| The date from where to get the transactions (any format that is accepted by momentjs) * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toDate** | **optional.String**| Get the transactions to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| limit | 

### Return type

[**Erc20TransactionCollection**](erc20TransactionCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactions**
> TransactionCollection GetTransactions(ctx, address, optional)
Gets native transactions

Gets native transactions in descending order based on block number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| address | 
 **optional** | ***AccountApiGetTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetTransactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 
 **fromBlock** | **optional.Int32**| The minimum block number from where to get the transactions * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toBlock** | **optional.Int32**| The maximum block number from where to get the transactions. * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **fromDate** | **optional.String**| The date from where to get the transactions (any format that is accepted by momentjs) * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toDate** | **optional.String**| Get the transactions to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| limit | 

### Return type

[**TransactionCollection**](transactionCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

