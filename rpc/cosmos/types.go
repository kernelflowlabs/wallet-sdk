package cosmos

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type (
	GetAccInfoRes struct {
		BaseResponse
		Account struct {
			AccountNumber string `json:"account_number"`
			Sequence      string `json:"sequence"`
		} `json:"account"`
	}

	GetBalanceRes struct {
		BaseResponse
		Balances []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"balances"`
	}

	GetBlockRes struct {
		BaseResponse
		Block Block `json:"block"`
	}
	Block struct {
		Header struct {
			Height string `json:"height"`
		} `json:"header"`
	}
	RestTx struct {
		Height    string `json:"height"`
		Txhash    string `json:"txhash"`
		Code      int    `json:"code"`
		RawLog    string `json:"raw_log"`
		GasWanted string `json:"gas_wanted"`
		GasUsed   string `json:"gas_used"`
		Tx        struct {
			Body struct {
				Messages []struct {
					Type        string `json:"@type"`
					FromAddress string `json:"from_address"`
					ToAddress   string `json:"to_address"`
					Amount      []struct {
						Denom  string `json:"denom"`
						Amount string `json:"amount"`
					} `json:"amount"`
				} `json:"messages"`
				Memo string `json:"memo"`
			} `json:"body"`
		} `json:"tx"`
		Timestamp string `json:"timestamp,omitempty"`
	}

	SendTxReq struct {
		TxBytes []byte `json:"tx_bytes"`
		Mode    string `json:"mode"`
	}
	SendTxRes struct {
		BaseResponse
		TxResponse RestTx `json:"tx_response"`
	}
	GetTxRes struct {
		BaseResponse
		TxResponse RestTx `json:"tx_response"`
	}
)

type SeiStatus struct {
	SyncInfo struct {
		LatestBlockHeight string `json:"latest_block_height"`
	} `json:"sync_info"`
}
