package utxo

import "time"

type (
	BaseRequest struct {
		JsonRPC string      `json:"jsonrpc"`
		ID      string      `json:"id"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
	}
	BaseResponse struct {
		ID    string `json:"id"`
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
	}
	GetBlockCountRes struct {
		BaseResponse
		Result uint64 `json:"result"`
	}
	SendRawTransactionRes struct {
		BaseResponse
		Result string `json:"result"`
	}
	GetBlockHashRes struct {
		BaseResponse
		Result string `json:"result"`
	}
	GetBlockRes struct {
		BaseResponse
		Result string `json:"result"`
	}
	GetBlockVerboseRes struct {
		BaseResponse
		Result RpcBlock `json:"result"`
	}
	RpcBlock struct {
		Hash          string  `json:"hash"`
		Confirmations int64   `json:"confirmations"`
		Height        int64   `json:"height"`
		Version       int32   `json:"version"`
		MerkleRoot    string  `json:"merkleroot"`
		Time          int64   `json:"time"`
		Nonce         uint32  `json:"nonce"`
		Bits          string  `json:"bits"`
		Difficulty    float64 `json:"difficulty"`
		PreviousHash  string  `json:"previousblockhash"`
		NextHash      string  `json:"nextblockhash,omitempty"`
		Tx            []RpcTx `json:"tx"`
	}
	RpcTx struct {
		Txid     string    `json:"txid"`
		Version  int32     `json:"version"`
		Locktime uint32    `json:"locktime"`
		Vin      []RpcVin  `json:"vin"`
		Vout     []RpcVout `json:"vout"`
	}
	RpcVin struct {
		Coinbase  string `json:"coinbase,omitempty"`
		Txid      string `json:"txid"`
		Vout      uint32 `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		TxinWitness []string `json:"txinwitness,omitempty"` //only for btc
		Sequence    uint32   `json:"sequence"`
	}
	RpcVout struct {
		Value        float64         `json:"value"`
		N            int             `json:"n"`
		ScriptPubKey RpcScriptPubKey `json:"scriptPubKey"`
	}
	RpcScriptPubKey struct {
		Asm       string   `json:"asm"`
		Hex       string   `json:"hex"`
		ReqSigs   int      `json:"reqSigs,omitempty"`
		Type      string   `json:"type"`
		Address   string   `json:"address,omitempty"`   // for BTC new version
		Addresses []string `json:"addresses,omitempty"` // for BTC old version and DOGE
	}
)

// electrs
type (
	ElectrsAddressRes struct {
		ChainStats struct {
			FundedTxoCount int    `json:"funded_txo_count"`
			FundedTxoSum   uint64 `json:"funded_txo_sum"`
			SpentTxoCount  int    `json:"spent_txo_count"`
			SpentTxoSum    uint64 `json:"spent_txo_sum"`
			TxCount        int    `json:"tx_count"`
		} `json:"chain_stats"`
		MempoolStats struct {
			FundedTxoCount int    `json:"funded_txo_count"`
			FundedTxoSum   uint64 `json:"funded_txo_sum"`
			SpentTxoCount  int    `json:"spent_txo_count"`
			SpentTxoSum    uint64 `json:"spent_txo_sum"`
			TxCount        int    `json:"tx_count"`
		} `json:"mempool_stats"`
	}
	ElectrsUtxo struct {
		TxID   string `json:"txid"`
		Vout   uint64 `json:"vout"`
		Status struct {
			Confirmed bool `json:"confirmed"`
		} `json:"status"`
		Value uint64 `json:"value"`
	}
	ElectrsUtxoRes []ElectrsUtxo

	ElectrsTxRes struct {
		Fee      uint64 `json:"fee"`
		Weight   int    `json:"weight"`
		Size     int    `json:"size"`
		Version  int    `json:"version"`
		Locktime int    `json:"locktime"`
		Txid     string `json:"txid"`
		Vin      []struct {
			Txid    string `json:"txid"`
			Vout    int    `json:"vout"`
			Prevout struct {
				Scriptpubkey        string `json:"scriptpubkey"`
				ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
				ScriptpubkeyType    string `json:"scriptpubkey_type"`
				ScriptpubkeyAddress string `json:"scriptpubkey_address"`
				Value               int64  `json:"value"`
			} `json:"prevout"`
			Scriptsig    string   `json:"scriptsig"`
			ScriptsigAsm string   `json:"scriptsig_asm"`
			Witness      []string `json:"witness"`
			IsCoinbase   bool     `json:"is_coinbase"`
			Sequence     int64    `json:"sequence"`
		} `json:"vin"`
		Vout []struct {
			Scriptpubkey        string `json:"scriptpubkey"`
			ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
			ScriptpubkeyType    string `json:"scriptpubkey_type"`
			ScriptpubkeyAddress string `json:"scriptpubkey_address"`
			Value               int64  `json:"value"`
		} `json:"vout"`
		Status struct {
			Confirmed   bool  `json:"confirmed"`
			BlockHeight int64 `json:"block_height"`
			BlockTime   int64 `json:"block_time"`
		} `json:"status"`
	}

	ElectrsFeeRes struct {
		One   float64 `json:"1"`
		Two   float64 `json:"2"`
		Three float64 `json:"3"`
	}
)

// ltc blockcypher
type (
	BlockcypherFeeRes struct {
		Name             string    `json:"name"`
		Height           int64     `json:"height"`
		Hash             string    `json:"hash"`
		Time             time.Time `json:"time"`
		LatestUrl        string    `json:"latest_url"`
		PreviousHash     string    `json:"previous_hash"`
		PreviousUrl      string    `json:"previous_url"`
		PeerCount        int64     `json:"peer_count"`
		UnconfirmedCount int64     `json:"unconfirmed_count"`
		HighFeePerKb     float64   `json:"high_fee_per_kb"`
		MediumFeePerKb   float64   `json:"medium_fee_per_kb"`
		LowFeePerKb      float64   `json:"low_fee_per_kb"`
		LastForkHeight   int64     `json:"last_fork_height"`
		LastForkHash     string    `json:"last_fork_hash"`
	}
	BlockcypherUtxoRes struct {
		Address            string `json:"address"`
		TotalReceived      int64  `json:"total_received"`
		TotalSent          int64  `json:"total_sent"`
		Balance            int64  `json:"balance"`
		UnconfirmedBalance int64  `json:"unconfirmed_balance"`
		FinalBalance       int64  `json:"final_balance"`
		NTx                int64  `json:"n_tx"`
		UnconfirmedNTx     int64  `json:"unconfirmed_n_tx"`
		FinalNTx           int64  `json:"final_n_tx"`
		Txrefs             []struct {
			TxHash        string `json:"tx_hash"`
			BlockHeight   uint64 `json:"block_height"`
			TxInputN      int64  `json:"tx_input_n"`
			TxOutputN     int64  `json:"tx_output_n"`
			Value         int64  `json:"value"`
			RefBalance    int64  `json:"ref_balance"`
			Spent         bool   `json:"spent"`
			Confirmations int64  `json:"confirmations"`
			Confirmed     string `json:"confirmed"`
			DoubleSpend   bool   `json:"double_spend"`
		} `json:"txrefs"`
		TxUrl string `json:"tx_url"`
	}
	BlockcypherTxRes struct {
		BlockHash     string   `json:"block_hash"`
		BlockHeight   int64    `json:"block_height"`
		BlockIndex    int64    `json:"block_index"`
		Hash          string   `json:"hash"`
		Addresses     []string `json:"addresses"`
		Total         int64    `json:"total"`
		Fees          int64    `json:"fees"`
		Size          int64    `json:"size"`
		Vsize         int64    `json:"vsize"`
		Preference    string   `json:"preference"`
		RelayedBy     string   `json:"relayed_by"`
		Confirmed     string   `json:"confirmed"`
		Received      string   `json:"received"`
		Ver           int64    `json:"ver"`
		DoubleSpend   bool     `json:"double_spend"`
		VinSz         int64    `json:"vin_sz"`
		VoutSz        int64    `json:"vout_sz"`
		OptInRbf      bool     `json:"opt_in_rbf"`
		Confirmations int64    `json:"confirmations"`
		Confidence    int64    `json:"confidence"`
		Inputs        []struct {
			PrevHash    string   `json:"prev_hash"`
			OutputIndex int      `json:"output_index"`
			Script      string   `json:"script"`
			OutputValue int      `json:"output_value"`
			Sequence    int64    `json:"sequence"`
			Addresses   []string `json:"addresses"`
			ScriptType  string   `json:"script_type"`
			Age         int      `json:"age"`
		} `json:"inputs"`
		Outputs []struct {
			Value      int64    `json:"value"`
			Script     string   `json:"script"`
			Addresses  []string `json:"addresses"`
			ScriptType string   `json:"script_type"`
		} `json:"outputs"`
	}
)

// dogeinfo
type (
	DogeInfoAddressRes struct {
		Status string `json:"status"`
		Data   struct {
			Confirmed   string `json:"confirmed"`
			Unconfirmed string `json:"unconfirmed"`
			Error       string `json:"error_message"`
		} `json:"data"`
	}
	DogeInfoUtxoRes struct {
		Status string `json:"status"`
		Data   struct {
			Outputs []struct {
				Hash    string `json:"hash"`
				Index   uint64 `json:"index"`
				Script  string `json:"script"`
				Address string `json:"address"`
				Value   string `json:"value"`
				Block   int    `json:"block"`
				TxHex   string `json:"tx_hex"`
			} `json:"outputs"`
			Error string `json:"error_message"`
		} `json:"data"`
	}

	DogeTransaction struct {
		Hash          string `json:"hash"`
		Confirmations int    `json:"confirmations"`
		Size          int    `json:"size"`
		Version       int    `json:"version"`
		Locktime      int    `json:"locktime"`
		BlockHash     string `json:"block_hash"`
		Time          int    `json:"time"`
		InputsN       int    `json:"inputs_n"`
		Inputs        []struct {
			Pos       int    `json:"pos"`
			Value     string `json:"value"`
			Type      string `json:"type"`
			Address   string `json:"address"`
			ScriptSig struct {
				Hex string `json:"hex"`
			} `json:"scriptSig"`
			PreviousOutput struct {
				Hash string `json:"hash"`
				Pos  int    `json:"pos"`
			} `json:"previous_output"`
		} `json:"inputs"`
		InputsValue string `json:"inputs_value"`
		OutputsN    int    `json:"outputs_n"`
		Outputs     []struct {
			Pos     int    `json:"pos"`
			Value   string `json:"value"`
			Type    string `json:"type"`
			Address string `json:"address"`
		} `json:"outputs"`
		OutputsValue string `json:"outputs_value"`
		Fee          string `json:"fee"`
	}
	DogeTransactionRes struct {
		Status string `json:"status"`
		Data   struct {
			Error         string      `json:"error_message"`
			Hash          string      `json:"hash"`
			Confirmations uint64      `json:"confirmations"`
			Size          int         `json:"size"`
			Vsize         int         `json:"vsize"`
			Weight        interface{} `json:"weight"`
			Version       int         `json:"version"`
			Time          uint64      `json:"time"`
			BlockHash     string      `json:"block_hash"`
			BlockHeight   int         `json:"block_height"`
			Fee           string      `json:"fee"`
			Inputs        []struct {
				Index     int    `json:"index"`
				Value     string `json:"value"`
				Address   string `json:"address"`
				ScriptSig struct {
					Hex string `json:"hex"`
					Asm string `json:"asm"`
				} `json:"scriptSig"`
				Witness        interface{} `json:"witness"`
				PreviousOutput struct {
					Hash  string `json:"hash"`
					Index int    `json:"index"`
				} `json:"previous_output"`
			} `json:"inputs"`
			Outputs []struct {
				Index   int    `json:"index"`
				Value   string `json:"value"`
				Type    string `json:"type"`
				Address string `json:"address"`
				Script  struct {
					Hex string `json:"hex"`
					Asm string `json:"asm"`
				} `json:"script"`
				Spent interface{} `json:"spent"`
			} `json:"outputs"`
		} `json:"data"`
	}
)

// syscoin
type (
	BaseResponseSysCoin struct {
		ID    string `json:"id"`
		Error string `json:"error"`
	}
	GetBlockCountResSysCoin struct {
		BaseResponseSysCoin
		Blockbook struct {
			Coin       string `json:"coin"`
			Host       string `json:"host"`
			Version    string `json:"version"`
			BestHeight uint64 `json:"bestHeight"`
		} `json:"blockbook"`
	}

	GetBalanceResSysCoin struct {
		BaseResponseSysCoin
		Address            string `json:"address"`
		Balance            string `json:"balance"`
		TotalReceived      string `json:"totalReceived"`
		TotalSent          string `json:"totalSent"`
		UnconfirmedBalance string `json:"unconfirmedBalance"`
	}

	GetUtxoResSysCoin struct {
		BaseResponseSysCoin
		Utxos []struct {
			Txid          string `json:"txid"`
			Vout          uint64 `json:"vout"`
			Value         string `json:"value"`
			Height        int    `json:"height"`
			Confirmations int    `json:"confirmations"`
		} `json:"utxos"`
	}

	SendRawTransactionResSysCoin struct {
		BaseResponseSysCoin
		Result string `json:"result"`
	}

	GetTransactionRes struct {
		BaseResponseSysCoin
		Txid    string `json:"txid"`
		Version int    `json:"version"`
		Vin     []struct {
			Txid      string   `json:"txid"`
			Vout      int      `json:"vout"`
			Sequence  int64    `json:"sequence"`
			N         int      `json:"n"`
			Addresses []string `json:"addresses"`
			IsAddress bool     `json:"isAddress"`
			Value     string   `json:"value"`
		} `json:"vin"`
		Vout []struct {
			Value     string   `json:"value"`
			N         int      `json:"n"`
			Hex       string   `json:"hex"`
			Addresses []string `json:"addresses"`
			IsAddress bool     `json:"isAddress"`
		} `json:"vout"`
		BlockHeight   int    `json:"blockHeight"`
		Confirmations int    `json:"confirmations"`
		BlockTime     int    `json:"blockTime"`
		Value         string `json:"value"`
		ValueIn       string `json:"valueIn"`
		Fees          string `json:"fees"`
		Hex           string `json:"hex"`
		Rbf           bool   `json:"rbf"`
	}
)
