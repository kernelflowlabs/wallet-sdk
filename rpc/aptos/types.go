package aptos

type (
	ErrMsg struct {
		Message     string `json:"message"`
		ErrorCode   string `json:"error_code"`
		VmErrorCode int    `json:"vm_error_code"`
	}
	AccountResourceRes struct {
		ErrMsg
		Type string `json:"type"`
		Data struct {
			Coin struct {
				Value string `json:"value"`
			} `json:"coin"`
			DepositEvents struct {
				Counter string `json:"counter"`
				Guid    struct {
					Id struct {
						Addr        string `json:"addr"`
						CreationNum string `json:"creation_num"`
					} `json:"id"`
				} `json:"guid"`
			} `json:"deposit_events"`
			Frozen         bool `json:"frozen"`
			WithdrawEvents struct {
				Counter string `json:"counter"`
				Guid    struct {
					Id struct {
						Addr        string `json:"addr"`
						CreationNum string `json:"creation_num"`
					} `json:"id"`
				} `json:"guid"`
			} `json:"withdraw_events"`
		} `json:"data"`
	}
	AccountCoreRes struct {
		ErrMsg
		SequenceNumber    string `json:"sequence_number"`
		AuthenticationKey string `json:"authentication_key"`
	}
	LedgerInfoRes struct {
		ErrMsg
		ChainId             uint64 `json:"chain_id"`
		LedgerVersion       string `json:"ledger_version" gencodec:"required"`
		LedgerTimestamp     string `json:"ledger_timestamp" gencodec:"required"`
		BlockHeight         string `json:"block_height" gencodec:"required"`
		Epoch               string `json:"epoch"`
		NodeRole            string `json:"node_role"`
		OldestBlockHeight   string `json:"oldest_block_height"`
		OldestLedgerVersion string `json:"oldest_ledger_version"`
	}
	GasPriceRes struct {
		ErrMsg
		DeprioritizedGasEstimate uint64 `json:"deprioritized_gas_estimate"`
		GasEstimate              uint64 `json:"gas_estimate"`
		PrioritizedGasEstimate   uint64 `json:"prioritized_gas_estimate"`
	}
	RpcTx struct {
		ErrMsg
		Version                 string `json:"version"`
		Hash                    string `json:"hash"`
		StateChangeHash         string `json:"state_change_hash"`
		EventRootHash           string `json:"event_root_hash"`
		GasUsed                 string `json:"gas_used"`
		Success                 bool   `json:"success"`
		VmStatus                string `json:"vm_status"`
		AccumulatorRootHash     string `json:"accumulator_root_hash"`
		Sender                  string `json:"sender"`
		SequenceNumber          string `json:"sequence_number"`
		MaxGasAmount            string `json:"max_gas_amount"`
		GasUnitPrice            string `json:"gas_unit_price"`
		ExpirationTimestampSecs string `json:"expiration_timestamp_secs"`
		Payload                 struct {
			Function      string        `json:"function"`
			TypeArguments []interface{} `json:"type_arguments"`
			Arguments     []interface{} `json:"arguments"`
			//TypeArguments []string `json:"type_arguments"`
			//Arguments     []string `json:"arguments"`
			Type string `json:"type"`
		} `json:"payload"`
		Signature struct {
			PublicKey interface{} `json:"public_key"`
			Signature interface{} `json:"signature"`
			Type      string      `json:"type"`
		} `json:"signature"`
		Timestamp string `json:"timestamp"`
		Type      string `json:"type"`
	}
	BlockTx struct {
		ErrMsg
		Transactions []RpcTx `json:"transactions"`
	}
	CoinInfoRes struct {
		ErrMsg
		Data CoinInfoData `json:"data"`
	}
	CoinInfoData struct {
		Decimals int    `json:"decimals"`
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
	}
)
