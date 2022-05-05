/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type BlockTransaction struct {
	// The hash of the transaction
	Hash string `json:"hash"`
	// The nonce
	Nonce            string `json:"nonce"`
	TransactionIndex string `json:"transaction_index"`
	// The from address
	FromAddress string `json:"from_address"`
	// The to address
	ToAddress string `json:"to_address"`
	// The value sent
	Value string `json:"value"`
	Gas   string `json:"gas,omitempty"`
	// The gas price
	GasPrice                 string `json:"gas_price"`
	Input                    string `json:"input"`
	ReceiptCumulativeGasUsed string `json:"receipt_cumulative_gas_used"`
	ReceiptGasUsed           string `json:"receipt_gas_used"`
	ReceiptContractAddress   string `json:"receipt_contract_address,omitempty"`
	ReceiptRoot              string `json:"receipt_root,omitempty"`
	ReceiptStatus            string `json:"receipt_status"`
	// The block timestamp
	BlockTimestamp string `json:"block_timestamp"`
	// The block number
	BlockNumber string `json:"block_number"`
	// The hash of the block
	BlockHash string `json:"block_hash"`
	// The logs of the transaction
	Logs []Log `json:"logs"`
}
