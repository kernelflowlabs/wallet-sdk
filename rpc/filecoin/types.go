package filecoin

type BaseRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      uint64      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type BaseResponse struct {
	JsonRPC string `json:"jsonrpc"`
	ID      uint64 `json:"id"`
	Error   struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	}
}

type Message struct {
	Version    uint64 `json:"Version"`
	To         string `json:"To"`
	From       string `json:"From"`
	Nonce      uint64 `json:"Nonce"`
	Value      string `json:"Value"`
	GasLimit   int64  `json:"GasLimit"`
	GasFeeCap  string `json:"GasFeeCap"`
	GasPremium string `json:"GasPremium"`
	Method     uint64 `json:"Method"`
	Params     []byte `json:"Params"`
}

type SignedMessage struct {
	Message   *Message `json:"Message"`
	Signature struct {
		Type byte   `json:"Type"`
		Data []byte `json:"Data"`
	} `json:"Signature"`
}

type (
	GetBlockHeight struct {
		Height uint64
	}
	GetBlockHeightRes struct {
		BaseResponse
		Result *GetBlockHeight `json:"result"`
	}
	GetBalanceResponse struct {
		BaseResponse
		Result string `json:"result"`
	}
	GetTipSetByHeight struct {
		Blocks []struct {
			Timestamp int64 `json:"Timestamp"`
		}
		Height uint64 `json:"Height"`
	}
	GetTipSetByHeightRes struct {
		BaseResponse
		Result *GetTipSetByHeight `json:"result"`
	}
	GetMessageResponse struct {
		BaseResponse
		Result *Message `json:"result"`
	}
	GetMpoolGetNonceRes struct {
		BaseResponse
		Result uint64 `json:"result"`
	}
	MpoolPushResponse struct {
		BaseResponse
		Result struct {
			Cid string `json:"/"`
		} `json:"result"`
	}
	GasEstimateRes struct {
		BaseResponse
		Result struct {
			GasLimit   int64  `json:"GasLimit"`
			GasFeeCap  string `json:"GasFeeCap"`
			GasPremium string `json:"GasPremium"`
		} `json:"result"`
	}
	StateSearchMsgLimited struct {
		Receipt struct {
			ExitCode int64
			Return   string
			GasUsed  uint64
		}
		Height uint64
	}
	StateSearchMsgLimitedRes struct {
		BaseResponse
		Result *StateSearchMsgLimited `json:"result"`
	}
)
