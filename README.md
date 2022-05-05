# Go API client for swagger

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://deep-index.moralis.io/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountApi* | [**GetNFTTransfers**](docs/AccountApi.md#getnfttransfers) | **Get** /{address}/nft/transfers | Gets NFT transfers to and from a given address
*AccountApi* | [**GetNFTs**](docs/AccountApi.md#getnfts) | **Get** /{address}/nft | Gets the NFTs owned by a given address
*AccountApi* | [**GetNFTsForContract**](docs/AccountApi.md#getnftsforcontract) | **Get** /{address}/nft/{token_address} | Gets the NFTs owned by a given address
*AccountApi* | [**GetNativeBalance**](docs/AccountApi.md#getnativebalance) | **Get** /{address}/balance | Gets native balance for a specific address.
*AccountApi* | [**GetTokenBalances**](docs/AccountApi.md#gettokenbalances) | **Get** /{address}/erc20 | Gets token balances for a specific address.
*AccountApi* | [**GetTokenTransfers**](docs/AccountApi.md#gettokentransfers) | **Get** /{address}/erc20/transfers | Gets erc 20 token transactions
*AccountApi* | [**GetTransactions**](docs/AccountApi.md#gettransactions) | **Get** /{address} | Gets native transactions
*DefiApi* | [**GetPairAddress**](docs/DefiApi.md#getpairaddress) | **Get** /{token0_address}/{token1_address}/pairAddress | Get pair address based on token0 and token1 address
*DefiApi* | [**GetPairReserves**](docs/DefiApi.md#getpairreserves) | **Get** /{pair_address}/reserves | Get liquidity pair reserves for an Uniswap V2 based Exchange.
*InfoApi* | [**EndpointWeights**](docs/InfoApi.md#endpointweights) | **Get** /info/endpointWeights | Returns the endpoint price list for rate limits and costs
*InfoApi* | [**Web3ApiVersion**](docs/InfoApi.md#web3apiversion) | **Get** /web3/version | Returns the web3 api version
*NativeApi* | [**GetBlock**](docs/NativeApi.md#getblock) | **Get** /block/{block_number_or_hash} | Gets block contents by block hash
*NativeApi* | [**GetContractEvents**](docs/NativeApi.md#getcontractevents) | **Post** /{address}/events | Gets events by topic
*NativeApi* | [**GetDateToBlock**](docs/NativeApi.md#getdatetoblock) | **Get** /dateToBlock | Gets the closest block of the provided date
*NativeApi* | [**GetLogsByAddress**](docs/NativeApi.md#getlogsbyaddress) | **Get** /{address}/logs | Gets address logs
*NativeApi* | [**GetNFTTransfersByBlock**](docs/NativeApi.md#getnfttransfersbyblock) | **Get** /block/{block_number_or_hash}/nft/transfers | Gets NFT transfers by block number or block hash
*NativeApi* | [**GetTransaction**](docs/NativeApi.md#gettransaction) | **Get** /transaction/{transaction_hash} | Get transaction details by transaction hash
*NativeApi* | [**RunContractFunction**](docs/NativeApi.md#runcontractfunction) | **Post** /{address}/function | Runs a function of a contract abi
*ResolveApi* | [**ResolveAddress**](docs/ResolveApi.md#resolveaddress) | **Get** /resolve/{address}/reverse | Return the ENS domain when available (Only for ETH)
*ResolveApi* | [**ResolveDomain**](docs/ResolveApi.md#resolvedomain) | **Get** /resolve/{domain} | Resolves an Unstoppable domain and returns the address
*StorageApi* | [**UploadFolder**](docs/StorageApi.md#uploadfolder) | **Post** /ipfs/uploadFolder | Uploads multiple files and place them in a folder directory
*TokenApi* | [**GetAllTokenIds**](docs/TokenApi.md#getalltokenids) | **Get** /nft/{address} | Retrieves the unique NFTs inside a given contract
*TokenApi* | [**GetContractNFTTransfers**](docs/TokenApi.md#getcontractnfttransfers) | **Get** /nft/{address}/transfers | Gets NFT transfers of a given contract
*TokenApi* | [**GetNFTLowestPrice**](docs/TokenApi.md#getnftlowestprice) | **Get** /nft/{address}/lowestprice | Get the lowest price found for a nft token contract
*TokenApi* | [**GetNFTMetadata**](docs/TokenApi.md#getnftmetadata) | **Get** /nft/{address}/metadata | Gets the global metadata for a given contract
*TokenApi* | [**GetNFTOwners**](docs/TokenApi.md#getnftowners) | **Get** /nft/{address}/owners | Gets the owners of the NFTs of a given contract
*TokenApi* | [**GetNFTTrades**](docs/TokenApi.md#getnfttrades) | **Get** /nft/{address}/trades | Get nft trades by marketplaces
*TokenApi* | [**GetNftTransfersFromToBlock**](docs/TokenApi.md#getnfttransfersfromtoblock) | **Get** /nft/transfers | Gets NFT transfers from a block number to a block number
*TokenApi* | [**GetTokenAddressTransfers**](docs/TokenApi.md#gettokenaddresstransfers) | **Get** /erc20/{address}/transfers | Gets erc20 transactions of a token contract
*TokenApi* | [**GetTokenAllowance**](docs/TokenApi.md#gettokenallowance) | **Get** /erc20/{address}/allowance | Gets the amount which the spender is allowed to withdraw from the owner.
*TokenApi* | [**GetTokenIdMetadata**](docs/TokenApi.md#gettokenidmetadata) | **Get** /nft/{address}/{token_id} | Gets the NFT with the given id of a given contract
*TokenApi* | [**GetTokenIdOwners**](docs/TokenApi.md#gettokenidowners) | **Get** /nft/{address}/{token_id}/owners | Gets the owners of NFTs for a given contract
*TokenApi* | [**GetTokenMetadata**](docs/TokenApi.md#gettokenmetadata) | **Get** /erc20/metadata | Gets token metadata
*TokenApi* | [**GetTokenMetadataBySymbol**](docs/TokenApi.md#gettokenmetadatabysymbol) | **Get** /erc20/metadata/symbols | Gets token metadata
*TokenApi* | [**GetTokenPrice**](docs/TokenApi.md#gettokenprice) | **Get** /erc20/{address}/price | Gets token price
*TokenApi* | [**GetWalletTokenIdTransfers**](docs/TokenApi.md#getwallettokenidtransfers) | **Get** /nft/{address}/{token_id}/transfers | Gets NFT transfers of a given contract
*TokenApi* | [**ReSyncMetadata**](docs/TokenApi.md#resyncmetadata) | **Get** /nft/{address}/{token_id}/metadata/resync | resync the metadata for a given token_id
*TokenApi* | [**SearchNFTs**](docs/TokenApi.md#searchnfts) | **Get** /nft/search | Retrieves the NFT data based on a metadata search
*TokenApi* | [**SyncNFTContract**](docs/TokenApi.md#syncnftcontract) | **Put** /nft/{address}/sync | Sync a Contract for NFT Index

## Documentation For Models

 - [Block](docs/Block.md)
 - [BlockDate](docs/BlockDate.md)
 - [BlockTransaction](docs/BlockTransaction.md)
 - [ChainList](docs/ChainList.md)
 - [EndpointWeights](docs/EndpointWeights.md)
 - [Ens](docs/Ens.md)
 - [Erc20Allowance](docs/Erc20Allowance.md)
 - [Erc20Metadata](docs/Erc20Metadata.md)
 - [Erc20Price](docs/Erc20Price.md)
 - [Erc20TokenBalance](docs/Erc20TokenBalance.md)
 - [Erc20Transaction](docs/Erc20Transaction.md)
 - [Erc20TransactionCollection](docs/Erc20TransactionCollection.md)
 - [Erc721Metadata](docs/Erc721Metadata.md)
 - [HistoricalNftTransfer](docs/HistoricalNftTransfer.md)
 - [IpfsFile](docs/IpfsFile.md)
 - [IpfsFileRequest](docs/IpfsFileRequest.md)
 - [Log](docs/Log.md)
 - [LogEvent](docs/LogEvent.md)
 - [LogEventByAddress](docs/LogEventByAddress.md)
 - [MetadataResync](docs/MetadataResync.md)
 - [NativeBalance](docs/NativeBalance.md)
 - [NativeErc20Price](docs/NativeErc20Price.md)
 - [Nft](docs/Nft.md)
 - [NftCollection](docs/NftCollection.md)
 - [NftContractMetadata](docs/NftContractMetadata.md)
 - [NftContractMetadataCollection](docs/NftContractMetadataCollection.md)
 - [NftMetadata](docs/NftMetadata.md)
 - [NftMetadataCollection](docs/NftMetadataCollection.md)
 - [NftOwner](docs/NftOwner.md)
 - [NftOwnerCollection](docs/NftOwnerCollection.md)
 - [NftTransfer](docs/NftTransfer.md)
 - [NftTransferCollection](docs/NftTransferCollection.md)
 - [ReservesCollection](docs/ReservesCollection.md)
 - [Resolve](docs/Resolve.md)
 - [RunContractDto](docs/RunContractDto.md)
 - [Trade](docs/Trade.md)
 - [TradeCollection](docs/TradeCollection.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionCollection](docs/TransactionCollection.md)
 - [Web3version](docs/Web3version.md)

## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author


