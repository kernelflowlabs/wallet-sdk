package kaspa

// types
type (
	GetTransactionRes struct {
		SubnetworkID            string    `json:"subnetwork_id"`
		TransactionID           string    `json:"transaction_id"`
		Hash                    string    `json:"hash"`
		Mass                    string    `json:"mass"`
		BlockHash               []string  `json:"block_hash"`
		BlockTime               int64     `json:"block_time"`
		IsAccepted              bool      `json:"is_accepted"`
		AcceptingBlockHash      string    `json:"accepting_block_hash"`
		AcceptingBlockBlueScore uint64    `json:"accepting_block_blue_score"`
		Inputs                  []Inputs  `json:"inputs"`
		Outputs                 []Outputs `json:"outputs"`
		Detail                  string    `json:"detail"`
	}

	Inputs struct {
		ID                    int    `json:"id"`
		TransactionID         string `json:"transaction_id"`
		Index                 int    `json:"index"`
		PreviousOutpointHash  string `json:"previous_outpoint_hash"`
		PreviousOutpointIndex string `json:"previous_outpoint_index"`
		SignatureScript       string `json:"signature_script"`
		SigOpCount            string `json:"sig_op_count"`
	}

	Outputs struct {
		ID                     int         `json:"id"`
		TransactionID          string      `json:"transaction_id"`
		Index                  int         `json:"index"`
		Amount                 int         `json:"amount"`
		ScriptPublicKey        string      `json:"script_public_key"`
		ScriptPublicKeyAddress string      `json:"script_public_key_address"`
		ScriptPublicKeyType    string      `json:"script_public_key_type"`
		AcceptingBlockHash     interface{} `json:"accepting_block_hash"`
	}
)

type (
	KasScanGetTransactionRes struct {
		Message                 string      `json:"message"`
		Detail                  string      `json:"detail"`
		SubnetworkId            string      `json:"subnetwork_id"`
		TransactionId           string      `json:"transaction_id"`
		Hash                    string      `json:"hash"`
		Mass                    string      `json:"mass"`
		Payload                 interface{} `json:"payload"`
		BlockHash               []string    `json:"block_hash"`
		BlockTime               int64       `json:"block_time"`
		IsAccepted              bool        `json:"is_accepted"`
		AcceptingBlockHash      string      `json:"accepting_block_hash"`
		AcceptingBlockBlueScore int         `json:"accepting_block_blue_score"`
		AcceptingBlockTime      int64       `json:"accepting_block_time"`
		Inputs                  []struct {
			TransactionId           string      `json:"transaction_id"`
			Index                   int         `json:"index"`
			PreviousOutpointHash    string      `json:"previous_outpoint_hash"`
			PreviousOutpointIndex   string      `json:"previous_outpoint_index"`
			PreviousOutpointAddress interface{} `json:"previous_outpoint_address"`
			PreviousOutpointAmount  interface{} `json:"previous_outpoint_amount"`
			SignatureScript         string      `json:"signature_script"`
			SigOpCount              string      `json:"sig_op_count"`
		} `json:"inputs"`
		Outputs []struct {
			TransactionId          string `json:"transaction_id"`
			Index                  int    `json:"index"`
			Amount                 int    `json:"amount"`
			ScriptPublicKey        string `json:"script_public_key"`
			ScriptPublicKeyAddress string `json:"script_public_key_address"`
			ScriptPublicKeyType    string `json:"script_public_key_type"`
		} `json:"outputs"`
	}
)

type (
	KasplexGetBalanceRes struct {
		Message string `json:"message"`
		Result  []struct {
			Tick       string `json:"tick"`
			Balance    string `json:"balance"`
			Locked     string `json:"locked"`
			Dec        string `json:"dec"`
			OpScoreMod string `json:"opScoreMod"`
		} `json:"result"`
	}
	KasplexGetTransactionRes struct {
		Message string `json:"message"`
		Result  []struct {
			P          string `json:"p"`
			Op         string `json:"op"`
			Tick       string `json:"tick"`
			Amt        string `json:"amt"`
			From       string `json:"from"`
			To         string `json:"to"`
			OpScore    string `json:"opScore"`
			HashRev    string `json:"hashRev"`
			FeeRev     string `json:"feeRev"`
			TxAccept   string `json:"txAccept"`
			OpAccept   string `json:"opAccept"`
			OpError    string `json:"opError"`
			Checkpoint string `json:"checkpoint"`
			MtsAdd     string `json:"mtsAdd"`
			MtsMod     string `json:"mtsMod"`
		} `json:"result"`
	}
	KASSendRequest struct {
		Recipient string `json:"recipient"`
		Amount    string `json:"amount"`
		OrderId   string `json:"orderId"`
	}
	KasplexSendResponse struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			OrderInfo struct {
				CommitHash string `json:"commitHash"`
				Amount     string `json:"amount"`
				Tick       string `json:"tick"`
				Recipient  string `json:"recipient"`
				Error      string `json:"error"`
				RevealHash string `json:"revealHash"`
				OrderId    string `json:"orderId"`
			} `json:"orderInfo"`
			Timestamp string `json:"timestamp"`
		} `json:"data"`
	}
	KRC20SendRequest struct {
		Tick      string `json:"tick"`
		Recipient string `json:"recipient"`
		Amount    string `json:"amount"`
		Minter    string `json:"minter"`
		OrderId   string `json:"orderId"`
	}
	KRC20SubmitTransactionResponse struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			OrderInfo struct {
				CommitHash string `json:"commitHash"`
				Amount     string `json:"amount"`
				Tick       string `json:"tick"`
				Recipient  string `json:"recipient"`
				Error      string `json:"error"`
				RevealHash string `json:"revealHash"`
				OrderId    string `json:"orderId"`
			} `json:"orderInfo"`
			Timestamp string `json:"timestamp"`
		} `json:"data"`
	}
	KRC20EstimateFee struct {
		PriorityBucket struct {
			Feerate          int `json:"feerate"`
			EstimatedSeconds int `json:"estimatedSeconds"`
		} `json:"priorityBucket"`
		NormalBuckets []struct {
			Feerate          int     `json:"feerate"`
			EstimatedSeconds float64 `json:"estimatedSeconds"`
		} `json:"normalBuckets"`
		LowBuckets []struct {
			Feerate          int     `json:"feerate"`
			EstimatedSeconds float64 `json:"estimatedSeconds"`
		} `json:"lowBuckets"`
	}
	KRC20DelOrderResponse struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Timestamp string `json:"timestamp"`
			OrderId   string `json:"orderId"`
		} `json:"data"`
	}
)
