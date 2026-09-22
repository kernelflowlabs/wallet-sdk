package multiversx

type (
	NetWorkConfigRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			Config *RpcNetworkConfig `json:"config"`
		}
	}
	RpcNetworkConfig struct {
		ChainID                  string  `json:"erd_chain_id"`
		Denomination             int     `json:"erd_denomination"`
		GasPerDataByte           uint64  `json:"erd_gas_per_data_byte"`
		LatestTagSoftwareVersion string  `json:"erd_latest_tag_software_version"`
		MetaConsensusGroup       uint32  `json:"erd_meta_consensus_group_size"`
		MinGasLimit              uint64  `json:"erd_min_gas_limit"`
		MinGasPrice              uint64  `json:"erd_min_gas_price"`
		MinTransactionVersion    uint32  `json:"erd_min_transaction_version"`
		NumMetachainNodes        uint32  `json:"erd_num_metachain_nodes"`
		NumNodesInShard          uint32  `json:"erd_num_nodes_in_shard"`
		NumShardsWithoutMeta     uint32  `json:"erd_num_shards_without_meta"`
		RoundDuration            int64   `json:"erd_round_duration"`
		ShardConsensusGroupSize  uint64  `json:"erd_shard_consensus_group_size"`
		StartTime                int64   `json:"erd_start_time"`
		Adaptivity               bool    `json:"erd_adaptivity,string"`
		Hysteresys               float32 `json:"erd_hysteresis,string"`
	}
	StatusRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			Status struct {
				ErdCurrentRound               uint64 `json:"erd_current_round"`
				ErdEpochNumber                uint64 `json:"erd_epoch_number"`
				ErdHighestFinalNonce          uint64 `json:"erd_highest_final_nonce"`
				ErdNonce                      uint64 `json:"erd_nonce"`
				ErdNonceAtEpochStart          uint64 `json:"erd_nonce_at_epoch_start"`
				ErdNoncesPassedInCurrentEpoch uint64 `json:"erd_nonces_passed_in_current_epoch"`
				ErdRoundAtEpochStart          uint64 `json:"erd_round_at_epoch_start"`
				ErdRoundsPassedInCurrentEpoch uint64 `json:"erd_rounds_passed_in_current_epoch"`
				ErdRoundsPerEpoch             uint64 `json:"erd_rounds_per_epoch"`
			} `json:"status"`
		} `json:"data"`
	}
	AddressRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			Account struct {
				Address string `json:"address"`
				Nonce   uint64 `json:"nonce"`
				Balance string `json:"balance"`
			} `json:"account"`
		} `json:"data"`
	}
	TopicsArr     []string
	EgldEventItem struct {
		Address    string    `json:"address"`
		Identifier string    `json:"identifier"`
		Data       string    `json:"data"`
		Topics     TopicsArr `json:"topics"`
	}
	EgldLogsItem struct {
		Address string           `json:"address"`
		Events  []*EgldEventItem `json:"events"`
	}
	EgldSmartContractResultsItem struct {
		Hash          string        `json:"hash"`
		Sender        string        `json:"sender"`
		Receiver      string        `json:"receiver"`
		Nonce         uint64        `json:"nonce"`
		Data          string        `json:"data"`
		CallType      int           `json:"callType"`
		ReturnMessage string        `json:"returnMessage"`
		Logs          *EgldLogsItem `json:"logs"`
	}
	EgldTransaction struct {
		Type                              string `json:"type"`
		ProcessingTypeOnSource            string `json:"processingTypeOnSource"`
		ProcessingTypeOnDestination       string `json:"processingTypeOnDestination"`
		Hash                              string `json:"hash"`
		Nonce                             int    `json:"nonce"`
		Round                             int    `json:"round"`
		Epoch                             int    `json:"epoch"`
		Value                             string `json:"value"`
		Receiver                          string `json:"receiver"`
		Sender                            string `json:"sender"`
		GasPrice                          int    `json:"gasPrice"`
		GasLimit                          int    `json:"gasLimit"`
		Data                              string `json:"data"`
		Signature                         string `json:"signature"`
		SourceShard                       int    `json:"sourceShard"`
		DestinationShard                  int    `json:"destinationShard"`
		BlockNonce                        int    `json:"blockNonce"`
		BlockHash                         string `json:"blockHash"`
		NotarizedAtSourceInMetaNonce      int    `json:"notarizedAtSourceInMetaNonce"`
		NotarizedAtSourceInMetaHash       string `json:"NotarizedAtSourceInMetaHash"`
		NotarizedAtDestinationInMetaNonce int    `json:"notarizedAtDestinationInMetaNonce"`
		NotarizedAtDestinationInMetaHash  string `json:"notarizedAtDestinationInMetaHash"`
		MiniblockType                     string `json:"miniblockType"`
		MiniblockHash                     string `json:"miniblockHash"`
		Timestamp                         int    `json:"timestamp"`
		SmartContractResults              []struct {
			Hash           string        `json:"hash"`
			Nonce          int           `json:"nonce"`
			Value          int64         `json:"value"`
			Receiver       string        `json:"receiver"`
			Sender         string        `json:"sender"`
			Data           string        `json:"data"`
			PrevTxHash     string        `json:"prevTxHash"`
			OriginalTxHash string        `json:"originalTxHash"`
			GasLimit       int           `json:"gasLimit"`
			GasPrice       int           `json:"gasPrice"`
			CallType       int           `json:"callType"`
			Operation      string        `json:"operation"`
			IsRefund       bool          `json:"isRefund"`
			ReturnMessage  string        `json:"returnMessage"`
			Logs           *EgldLogsItem `json:"logs"`
		} `json:"smartContractResults"`
		Logs struct {
			Address string `json:"address"`
			Events  []struct {
				Address    string   `json:"address"`
				Identifier string   `json:"identifier"`
				Topics     []string `json:"topics"`
				Data       *string  `json:"data"`
			} `json:"events"`
		} `json:"logs"`
		Status           string   `json:"status"`
		Tokens           []string `json:"tokens"`
		EsdtValues       []string `json:"esdtValues"`
		Operation        string   `json:"operation"`
		InitiallyPaidFee string   `json:"initiallyPaidFee"`
	}
	EgldShard struct {
		Hash  string `json:"hash"`
		Nonce uint64 `json:"nonce"`
		Shard int    `json:"shard"`
	}
	Block struct {
		Nonce        uint64             `json:"nonce"`
		Round        uint64             `json:"round"`
		Hash         string             `json:"hash"`
		Epoch        uint64             `json:"epoch"`
		NumTxs       uint64             `json:"numTxs"`
		Timestamp    int64              `json:"timestamp"`
		Status       string             `jons:"status"`
		ShardBlocks  []*EgldShard       `json:"shardBlocks"`
		Transactions []*EgldTransaction `json:"transactions"`
	}
	HyperblockRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			Hyperblock *Block `json:"hyperblock"`
		} `json:"data"`
	}
	TransactionRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			Transaction *EgldTransaction `json:"transaction"`
		} `json:"data"`
	}
	TransactionStatusRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			Status string `json:"status"`
		} `json:"data"`
	}
	TransactionSendRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			TxHash string `json:"txHash"`
		} `json:"data"`
	}
	TransactionSimulateRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			Result struct {
				Status string `json:"status"`
				Hash   string `json:"hash"`
			} `json:"result"`
		} `json:"data"`
	}
	EgldTokenData struct {
		Symbol          string `json:"symbol,omitempty"`
		Balance         string `json:"balance"`
		Properties      string `json:"properties"`
		TokenIdentifier string `json:"tokenIdentifier"`
	}
	EsdtBalanceRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			TokenData *EgldTokenData `json:"tokenData"`
		} `json:"data"`
	}
	TokenInfoRes struct {
		Code  string `json:"code"`
		Error string `json:"error"`
		Data  struct {
			Decimals int64  `json:"decimals"`
			Name     string `json:"name"`
			Ticker   string `json:"ticker"`
		} `json:"data"`
	}
)

type nativeTx struct {
	Nonce     uint64 `json:"nonce"`
	Value     string `json:"value"`
	RcvAddr   string `json:"receiver"`
	SndAddr   string `json:"sender"`
	GasPrice  uint64 `json:"gasPrice,omitempty"`
	GasLimit  uint64 `json:"gasLimit,omitempty"`
	Data      string `json:"data,omitempty"`
	Signature string `json:"signature,omitempty"`
	ChainID   string `json:"chainID"`
	Version   uint32 `json:"version"`
	Options   uint32 `json:"options,omitempty"`
}
