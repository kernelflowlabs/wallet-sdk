package starknet

const ethContractAddress = "0x049d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7"

type (
	BaseRequest struct {
		JsonRPC string      `json:"jsonrpc"`
		ID      string      `json:"id"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params,omitempty"`
	}
	ErrorResponse struct {
		Code    int    `json:"code"`
		Name    string `json:"name"`
		Message string `json:"message"`
	}
	BaseResponse struct {
		JsonRPC string         `json:"jsonrpc"`
		ID      string         `json:"id"`
		Error   *ErrorResponse `json:"error"`
	}

	GetHeightRes struct {
		BaseResponse
		Result uint64 `json:"result"`
	}
	GetNonceRes struct {
		BaseResponse
		Result string `json:"result"`
	}
	GetBalanceRes struct {
		BaseResponse
		Result []string `json:"result"`
	}
	GetDecimalRes struct {
		BaseResponse
		Result []string `json:"result"`
	}
	FeeEstimate struct {
		L1GasConsumed     string `json:"l1_gas_consumed"`
		L1GasPrice        string `json:"l1_gas_price"`
		L2GasConsumed     string `json:"l2_gas_consumed"`
		L2GasPrice        string `json:"l2_gas_price"`
		L1DataGasConsumed string `json:"l1_data_gas_consumed"`
		L1DataGasPrice    string `json:"l1_data_gas_price"`
		OverallFee        string `json:"overall_fee"`
	}
	EstimateFeeRes struct {
		BaseResponse
		Result []FeeEstimate `json:"result"`
	}
	SendTxRes struct {
		BaseResponse
		Result struct {
			TransactionHash string `json:"transaction_hash"`
		} `json:"result"`
	}
	GetTransactionByHash struct {
		BaseResponse
		Result struct {
			Calldata        []string `json:"calldata"`
			SenderAddress   string   `json:"sender_address"`
			TransactionHash string   `json:"transaction_hash"`
			Type            string   `json:"type"`
			Version         string   `json:"version"`
		} `json:"result"`
	}
	GetTransaction struct {
		BaseResponse
		Result struct {
			BlockNumber     int            `json:"block_number"`
			ExecutionStatus string         `json:"execution_status"`
			FinalityStatus  string         `json:"finality_status"`
			TransactionHash string         `json:"transaction_hash"`
			Events          []ReceiptEvent `json:"events"`
		} `json:"result"`
	}
	ReceiptEvent struct {
		FromAddress string   `json:"from_address"`
		Keys        []string `json:"keys"`
		Data        []string `json:"data"`
	}
)
