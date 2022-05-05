/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package moralisAPI

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/DevilsTear/moralis-go-client/models"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type DefiApiService service

/*
DefiApiService Get pair address based on token0 and token1 address
Fetches and returns pair data of the provided token0+token1 combination. The token0 and token1 options are interchangable (ie. there is no different outcome in \&quot;token0&#x3D;WETH and token1&#x3D;USDT\&quot; or \&quot;token0&#x3D;USDT and token1&#x3D;WETH\&quot;)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param exchange The factory name or address of the token exchange
 * @param token0Address Token0 address
 * @param token1Address Token1 address
 * @param optional nil or *DefiApiGetPairAddressOpts - Optional Parameters:
     * @param "Chain" (optional.Interface of ChainList) -  The chain to query
     * @param "ToBlock" (optional.String) -  To get the reserves at this block number
     * @param "ToDate" (optional.String) -  Get the reserves to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.
@return ReservesCollection
*/

type DefiApiGetPairAddressOpts struct {
	Chain   optional.Interface
	ToBlock optional.String
	ToDate  optional.String
}

func (a *DefiApiService) GetPairAddress(ctx context.Context, exchange string, token0Address string, token1Address string, localVarOptionals *DefiApiGetPairAddressOpts) (models.ReservesCollection, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.ReservesCollection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{token0_address}/{token1_address}/pairAddress"
	localVarPath = strings.Replace(localVarPath, "{"+"token0_address"+"}", fmt.Sprintf("%v", token0Address), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"token1_address"+"}", fmt.Sprintf("%v", token1Address), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Chain.IsSet() {
		localVarQueryParams.Add("chain", parameterToString(localVarOptionals.Chain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ToBlock.IsSet() {
		localVarQueryParams.Add("to_block", parameterToString(localVarOptionals.ToBlock.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ToDate.IsSet() {
		localVarQueryParams.Add("to_date", parameterToString(localVarOptionals.ToDate.Value(), ""))
	}
	localVarQueryParams.Add("exchange", parameterToString(exchange, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-Key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v models.ReservesCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DefiApiService Get liquidity pair reserves for an Uniswap V2 based Exchange.
Get the liquidity reserves for a given pair address. Only Uniswap V2 based exchanges supported at the moment.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pairAddress Liquidity pair address
 * @param optional nil or *DefiApiGetPairReservesOpts - Optional Parameters:
     * @param "Chain" (optional.Interface of ChainList) -  The chain to query
     * @param "ToBlock" (optional.String) -  To get the reserves at this block number
     * @param "ToDate" (optional.String) -  Get the reserves to this date (any format that is accepted by momentjs) * Provide the param &#x27;to_block&#x27; or &#x27;to_date&#x27; * If &#x27;to_date&#x27; and &#x27;to_block&#x27; are provided, &#x27;to_block&#x27; will be used.
     * @param "ProviderUrl" (optional.String) -  web3 provider url to user when using local dev chain
@return ReservesCollection
*/

type DefiApiGetPairReservesOpts struct {
	Chain       optional.Interface
	ToBlock     optional.String
	ToDate      optional.String
	ProviderUrl optional.String
}

func (a *DefiApiService) GetPairReserves(ctx context.Context, pairAddress string, localVarOptionals *DefiApiGetPairReservesOpts) (models.ReservesCollection, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.ReservesCollection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{pair_address}/reserves"
	localVarPath = strings.Replace(localVarPath, "{"+"pair_address"+"}", fmt.Sprintf("%v", pairAddress), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Chain.IsSet() {
		localVarQueryParams.Add("chain", parameterToString(localVarOptionals.Chain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ToBlock.IsSet() {
		localVarQueryParams.Add("to_block", parameterToString(localVarOptionals.ToBlock.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ToDate.IsSet() {
		localVarQueryParams.Add("to_date", parameterToString(localVarOptionals.ToDate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProviderUrl.IsSet() {
		localVarQueryParams.Add("provider_url", parameterToString(localVarOptionals.ProviderUrl.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-Key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v models.ReservesCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
