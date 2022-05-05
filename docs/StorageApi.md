# {{classname}}

All URIs are relative to *https://deep-index.moralis.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadFolder**](StorageApi.md#UploadFolder) | **Post** /ipfs/uploadFolder | Uploads multiple files and place them in a folder directory

# **UploadFolder**
> []IpfsFile UploadFolder(ctx, optional)
Uploads multiple files and place them in a folder directory

Uploads multiple files and place them in a folder directory 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageApiUploadFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiUploadFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []IpfsFileRequest**](ipfsFileRequest.md)| Array of JSON and Base64 Supported | 

### Return type

[**[]IpfsFile**](ipfsFile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

