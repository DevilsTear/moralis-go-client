# {{classname}}

All URIs are relative to *https://deep-index.moralis.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllTokenIds**](TokenApi.md#GetAllTokenIds) | **Get** /nft/{address} | Retrieves the unique NFTs inside a given contract
[**GetContractNFTTransfers**](TokenApi.md#GetContractNFTTransfers) | **Get** /nft/{address}/transfers | Gets NFT transfers of a given contract
[**GetNFTLowestPrice**](TokenApi.md#GetNFTLowestPrice) | **Get** /nft/{address}/lowestprice | Get the lowest price found for a nft token contract
[**GetNFTMetadata**](TokenApi.md#GetNFTMetadata) | **Get** /nft/{address}/metadata | Gets the global metadata for a given contract
[**GetNFTOwners**](TokenApi.md#GetNFTOwners) | **Get** /nft/{address}/owners | Gets the owners of the NFTs of a given contract
[**GetNFTTrades**](TokenApi.md#GetNFTTrades) | **Get** /nft/{address}/trades | Get nft trades by marketplaces
[**GetNftTransfersFromToBlock**](TokenApi.md#GetNftTransfersFromToBlock) | **Get** /nft/transfers | Gets NFT transfers from a block number to a block number
[**GetTokenAddressTransfers**](TokenApi.md#GetTokenAddressTransfers) | **Get** /erc20/{address}/transfers | Gets erc20 transactions of a token contract
[**GetTokenAllowance**](TokenApi.md#GetTokenAllowance) | **Get** /erc20/{address}/allowance | Gets the amount which the spender is allowed to withdraw from the owner.
[**GetTokenIdMetadata**](TokenApi.md#GetTokenIdMetadata) | **Get** /nft/{address}/{token_id} | Gets the NFT with the given id of a given contract
[**GetTokenIdOwners**](TokenApi.md#GetTokenIdOwners) | **Get** /nft/{address}/{token_id}/owners | Gets the owners of NFTs for a given contract
[**GetTokenMetadata**](TokenApi.md#GetTokenMetadata) | **Get** /erc20/metadata | Gets token metadata
[**GetTokenMetadataBySymbol**](TokenApi.md#GetTokenMetadataBySymbol) | **Get** /erc20/metadata/symbols | Gets token metadata
[**GetTokenPrice**](TokenApi.md#GetTokenPrice) | **Get** /erc20/{address}/price | Gets token price
[**GetWalletTokenIdTransfers**](TokenApi.md#GetWalletTokenIdTransfers) | **Get** /nft/{address}/{token_id}/transfers | Gets NFT transfers of a given contract
[**ReSyncMetadata**](TokenApi.md#ReSyncMetadata) | **Get** /nft/{address}/{token_id}/metadata/resync | resync the metadata for a given token_id
[**SearchNFTs**](TokenApi.md#SearchNFTs) | **Get** /nft/search | Retrieves the NFT data based on a metadata search
[**SyncNFTContract**](TokenApi.md#SyncNFTContract) | **Put** /nft/{address}/sync | Sync a Contract for NFT Index

# **GetAllTokenIds**
> NftCollection GetAllTokenIds(ctx, address, optional)
Retrieves the unique NFTs inside a given contract

Gets data, including metadata (where available), for all token ids for the given contract address. * Results are limited to 500 per page by default * Requests for contract addresses not yet indexed will automatically start the indexing process for that NFT collection 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
 **optional** | ***TokenApiGetAllTokenIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetAllTokenIdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
 **limit** | **optional.Int32**| limit | 
 **cursor** | **optional.String**| The cursor returned in the last response (for getting the next page)  | 

### Return type

[**NftCollection**](nftCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContractNFTTransfers**
> NftTransferCollection GetContractNFTTransfers(ctx, address, optional)
Gets NFT transfers of a given contract

Gets the transfers of the tokens matching the given parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
 **optional** | ***TokenApiGetContractNFTTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetContractNFTTransfersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
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

# **GetNFTLowestPrice**
> Trade GetNFTLowestPrice(ctx, address, optional)
Get the lowest price found for a nft token contract

Get the lowest price found for a nft token contract for the last x days (only trades paid in ETH)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
 **optional** | ***TokenApiGetNFTLowestPriceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetNFTLowestPriceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **days** | **optional.Int32**| The number of days to look back to find the lowest price If not provided 7 days will be the default  | 
 **providerUrl** | **optional.String**| web3 provider url to user when using local dev chain | 
 **marketplace** | **optional.String**| marketplace from where to get the trades (only opensea is supported at the moment) | [default to opensea]

### Return type

[**Trade**](trade.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFTMetadata**
> NftContractMetadata GetNFTMetadata(ctx, address, optional)
Gets the global metadata for a given contract

Gets the contract level metadata (name, symbol, base token uri) for the given contract * Requests for contract addresses not yet indexed will automatically start the indexing process for that NFT collection 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
 **optional** | ***TokenApiGetNFTMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetNFTMetadataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 

### Return type

[**NftContractMetadata**](nftContractMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFTOwners**
> NftOwnerCollection GetNFTOwners(ctx, address, optional)
Gets the owners of the NFTs of a given contract

Gets all owners of NFT items within a given contract collection * Use after /nft/contract/{token_address} to find out who owns each token id in a collection * Make sure to include a sort parm on a column like block_number_minted for consistent pagination results * Requests for contract addresses not yet indexed will automatically start the indexing process for that NFT collection 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
 **optional** | ***TokenApiGetNFTOwnersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetNFTOwnersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
 **limit** | **optional.Int32**| limit | 
 **cursor** | **optional.String**| The cursor returned in the last response (for getting the next page)  | 

### Return type

[**NftOwnerCollection**](nftOwnerCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFTTrades**
> TradeCollection GetNFTTrades(ctx, address, optional)
Get nft trades by marketplaces

Get the nft trades for a given contracts and marketplace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
 **optional** | ***TokenApiGetNFTTradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetNFTTradesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **fromBlock** | **optional.Int32**| The minimum block number from where to get the transfers * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toBlock** | **optional.String**| To get the reserves at this block number | 
 **fromDate** | **optional.String**| The date from where to get the transfers (any format that is accepted by momentjs) * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toDate** | **optional.String**| Get the reserves to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **providerUrl** | **optional.String**| web3 provider url to user when using local dev chain | 
 **marketplace** | **optional.String**| marketplace from where to get the trades (only opensea is supported at the moment) | [default to opensea]
 **cursor** | **optional.String**| cursor | 
 **limit** | **optional.Int32**| limit | 

### Return type

[**TradeCollection**](tradeCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNftTransfersFromToBlock**
> NftTransferCollection GetNftTransfersFromToBlock(ctx, optional)
Gets NFT transfers from a block number to a block number

Gets the transfers of the tokens from a block number to a block number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenApiGetNftTransfersFromToBlockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetNftTransfersFromToBlockOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **fromBlock** | **optional.Int32**| The minimum block number from where to get the transfers * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toBlock** | **optional.Int32**| The maximum block number from where to get the transfers. * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **fromDate** | **optional.String**| The date from where to get the transfers (any format that is accepted by momentjs) * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toDate** | **optional.String**| Get transfers up until this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
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

# **GetTokenAddressTransfers**
> Erc20TransactionCollection GetTokenAddressTransfers(ctx, address, optional)
Gets erc20 transactions of a token contract

Gets ERC20 token contract transactions in descending order based on block number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address of the token contract | 
 **optional** | ***TokenApiGetTokenAddressTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetTokenAddressTransfersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 
 **fromBlock** | **optional.Int32**| The minimum block number from where to get the transfers * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toBlock** | **optional.Int32**| The maximum block number from where to get the transfers. * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **fromDate** | **optional.String**| The date from where to get the transfers (any format that is accepted by momentjs) * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toDate** | **optional.String**| Get the transfers to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
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

# **GetTokenAllowance**
> Erc20Allowance GetTokenAllowance(ctx, address, ownerAddress, spenderAddress, optional)
Gets the amount which the spender is allowed to withdraw from the owner.

Gets the amount which the spender is allowed to withdraw from the spender

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address of the token contract | 
  **ownerAddress** | **string**| The address of the token owner | 
  **spenderAddress** | **string**| The address of the token spender | 
 **optional** | ***TokenApiGetTokenAllowanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetTokenAllowanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **providerUrl** | **optional.String**| web3 provider url to user when using local dev chain | 

### Return type

[**Erc20Allowance**](erc20Allowance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenIdMetadata**
> Nft GetTokenIdMetadata(ctx, address, tokenId, optional)
Gets the NFT with the given id of a given contract

Gets data, including metadata (where available), for the given token id of the given contract address. * Requests for contract addresses not yet indexed will automatically start the indexing process for that NFT collection 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
  **tokenId** | **string**| The id of the token | 
 **optional** | ***TokenApiGetTokenIdMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetTokenIdMetadataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]

### Return type

[**Nft**](nft.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenIdOwners**
> NftOwnerCollection GetTokenIdOwners(ctx, address, tokenId, optional)
Gets the owners of NFTs for a given contract

Gets all owners of NFT items within a given contract collection * Use after /nft/contract/{token_address} to find out who owns each token id in a collection * Make sure to include a sort parm on a column like block_number_minted for consistent pagination results * Requests for contract addresses not yet indexed will automatically start the indexing process for that NFT collection 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
  **tokenId** | **string**| The id of the token | 
 **optional** | ***TokenApiGetTokenIdOwnersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetTokenIdOwnersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
 **limit** | **optional.Int32**| limit | 
 **cursor** | **optional.String**| The cursor returned in the last response (for getting the next page)  | 

### Return type

[**NftOwnerCollection**](nftOwnerCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenMetadata**
> []Erc20Metadata GetTokenMetadata(ctx, addresses, optional)
Gets token metadata

Returns metadata (name, symbol, decimals, logo) for a given token contract address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addresses** | [**[]string**](string.md)| The addresses to get metadata for | 
 **optional** | ***TokenApiGetTokenMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetTokenMetadataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 
 **providerUrl** | **optional.String**| web3 provider url to user when using local dev chain | 

### Return type

[**[]Erc20Metadata**](erc20Metadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenMetadataBySymbol**
> []Erc20Metadata GetTokenMetadataBySymbol(ctx, symbols, optional)
Gets token metadata

Returns metadata (name, symbol, decimals, logo) for a given token contract address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbols** | [**[]string**](string.md)| The symbols to get metadata for | 
 **optional** | ***TokenApiGetTokenMetadataBySymbolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetTokenMetadataBySymbolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **subdomain** | **optional.String**| The subdomain of the moralis server to use (Only use when selecting local devchain as chain) | 

### Return type

[**[]Erc20Metadata**](erc20Metadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenPrice**
> Erc20Price GetTokenPrice(ctx, address, optional)
Gets token price

Returns the price nominated in the native token and usd for a given token contract address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address of the token contract | 
 **optional** | ***TokenApiGetTokenPriceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetTokenPriceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **providerUrl** | **optional.String**| web3 provider url to user when using local dev chain | 
 **exchange** | **optional.String**| The factory name or address of the token exchange | 
 **toBlock** | **optional.Int32**| to_block | 

### Return type

[**Erc20Price**](erc20Price.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWalletTokenIdTransfers**
> NftTransferCollection GetWalletTokenIdTransfers(ctx, address, tokenId, optional)
Gets NFT transfers of a given contract

Gets the transfers of the tokens matching the given parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
  **tokenId** | **string**| The id of the token | 
 **optional** | ***TokenApiGetWalletTokenIdTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetWalletTokenIdTransfersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
 **limit** | **optional.Int32**| limit | 
 **order** | **optional.String**| The field(s) to order on and if it should be ordered in ascending or descending order. Specified by: fieldName1.order,fieldName2.order. Example 1: \&quot;block_number\&quot;, \&quot;block_number.ASC\&quot;, \&quot;block_number.DESC\&quot;, Example 2: \&quot;block_number and contract_type\&quot;, \&quot;block_number.ASC,contract_type.DESC\&quot; | 
 **cursor** | **optional.String**| The cursor returned in the last response (for getting the next page)  | 

### Return type

[**NftTransferCollection**](nftTransferCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReSyncMetadata**
> MetadataResync ReSyncMetadata(ctx, address, tokenId, optional)
resync the metadata for a given token_id

ReSync the metadata for an NFT * The metadata(default) flag will request a the NFT's metadata from the already existing token_uri * The uri flag will fetch the latest token_uri from the given NFT address. In sync mode the metadata will also be fetched * The sync mode will make the endpoint synchronous so it will wait for the task to be completed before responding * The async mode(default) will make the endpoint asynchronous so we will wait for the task to be completed before responding 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
  **tokenId** | **string**| The id of the token | 
 **optional** | ***TokenApiReSyncMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiReSyncMetadataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **flag** | **optional.String**| The type of resync to operate | [default to metadata]
 **mode** | **optional.String**| To define the behaviour of the endpoint | [default to async]

### Return type

[**MetadataResync**](metadataResync.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchNFTs**
> NftMetadataCollection SearchNFTs(ctx, q, optional)
Retrieves the NFT data based on a metadata search

Gets NFTs that match a given metadata search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| The search string | 
 **optional** | ***TokenApiSearchNFTsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiSearchNFTsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 
 **format** | **optional.String**| The format of the token id | [default to decimal]
 **filter** | **optional.String**| What fields the search should match on. To look into the entire metadata set the value to &#x27;global&#x27;. To have a better response time you can look into a specific field like name | [default to global]
 **fromBlock** | **optional.Int32**| The minimum block number from where to start the search * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toBlock** | **optional.Int32**| The maximum block number from where to end the search * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **fromDate** | **optional.String**| The date from where to start the search (any format that is accepted by momentjs) * Provide the param &#x27;from_block&#x27; or &#x27;from_date&#x27; * If &#x27;from_date&#x27; and &#x27;from_block&#x27; are provided, &#x27;from_block&#x27; will be used.  | 
 **toDate** | **optional.String**| Get search results up until this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.  | 
 **addresses** | [**optional.Interface of []string**](string.md)| The addresses to get metadata for | 
 **cursor** | **optional.String**| cursor | 
 **limit** | **optional.Int32**| limit | 

### Return type

[**NftMetadataCollection**](nftMetadataCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncNFTContract**
> SyncNFTContract(ctx, address, optional)
Sync a Contract for NFT Index

Sync a Contract for NFT Index 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address of the contract | 
 **optional** | ***TokenApiSyncNFTContractOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiSyncNFTContractOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chain** | [**optional.Interface of ChainList**](.md)| The chain to query | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

