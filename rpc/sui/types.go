package sui

import "encoding/json"

type signedTx struct {
	Tx        string
	Signature string
}

type (
	BaseRequest struct {
		JsonRPC string      `json:"jsonrpc"`
		ID      int         `json:"id"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
	}
	BaseResponse struct {
		ID    int `json:"id"`
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
	}
	GetBlockCountRes struct {
		BaseResponse
		Result string `json:"result"`
	}
	GetBalanceRes struct {
		BaseResponse
		Result struct {
			CoinType        string      `json:"coinType"`
			CoinObjectCount int         `json:"coinObjectCount"`
			TotalBalance    string      `json:"totalBalance"`
			LockedBalance   interface{} `json:"lockedBalance"`
		}
	}

	MetaData struct {
		Decimals    int64  `json:"decimals"`
		Name        string `json:"name"`
		Symbol      string `json:"symbol"`
		Description string `json:"description"`
		IconUrl     string `json:"iconUrl"`
		Id          string `json:"id"`
	}
	GetCoinMetadataRes struct {
		BaseResponse
		Result MetaData `json:"result"`
	}
	GetGasPriceRes struct {
		BaseResponse
		Result string `json:"result"`
	}
	CoinItem struct {
		CoinType            string `json:"coinType"`
		CoinObjectId        string `json:"coinObjectId"`
		Version             string `json:"version"`
		Digest              string `json:"digest"`
		Balance             string `json:"balance"`
		PreviousTransaction string `json:"previousTransaction"`
	}
	Coins       []CoinItem
	GetCoinsRes struct {
		BaseResponse
		Result struct {
			Data        Coins  `json:"data"`
			NextCursor  string `json:"nextCursor"`
			HasNextPage bool   `json:"hasNextPage"`
		} `json:"result"`
	}
	GetTransactionsRes struct {
		BaseResponse
		Result struct {
			Digest  string `json:"digest"`
			Effects struct {
				MessageVersion string `json:"messageVersion"`
				Status         struct {
					Status string `json:"status"`
				} `json:"status"`
				ExecutedEpoch string `json:"executedEpoch"`
				GasUsed       struct {
					ComputationCost         string `json:"computationCost"`
					StorageCost             string `json:"storageCost"`
					StorageRebate           string `json:"storageRebate"`
					NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
				} `json:"gasUsed"`
				TransactionDigest string `json:"transactionDigest"`
			} `json:"effects"`
			TimestampMs string  `json:"timestampMs"`
			Checkpoint  *string `json:"checkpoint"`
		} `json:"result"`
	}

	SendRawTransactionRes struct {
		BaseResponse
		Result struct {
			Digest      string `json:"digest"`
			Transaction struct {
				Data struct {
					MessageVersion string `json:"messageVersion"`
					Transaction    struct {
						Kind   string `json:"kind"`
						Inputs []struct {
							Type       string          `json:"type"`
							ValueType  string          `json:"valueType,omitempty"`
							Value      json.RawMessage `json:"value,omitempty"`
							ObjectType string          `json:"objectType,omitempty"`
							ObjectId   string          `json:"objectId,omitempty"`
							Version    string          `json:"version,omitempty"`
							Digest     string          `json:"digest,omitempty"`
						} `json:"inputs"`
						Transactions []struct {
							TransferObjects []interface{} `json:"TransferObjects"`
						} `json:"transactions"`
					} `json:"transaction"`
					Sender  string `json:"sender"`
					GasData struct {
						Payment []struct {
							ObjectId string `json:"objectId"`
							Version  int    `json:"version"`
							Digest   string `json:"digest"`
						} `json:"payment"`
						Owner  string `json:"owner"`
						Price  string `json:"price"`
						Budget string `json:"budget"`
					} `json:"gasData"`
				} `json:"data"`
				TxSignatures []string `json:"txSignatures"`
			} `json:"transaction"`
			RawTransaction string `json:"rawTransaction"`
			Effects        struct {
				MessageVersion string `json:"messageVersion"`
				Status         struct {
					Status string `json:"status"`
				} `json:"status"`
				ExecutedEpoch string `json:"executedEpoch"`
				GasUsed       struct {
					ComputationCost         string `json:"computationCost"`
					StorageCost             string `json:"storageCost"`
					StorageRebate           string `json:"storageRebate"`
					NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
				} `json:"gasUsed"`
				TransactionDigest string `json:"transactionDigest"`
				Mutated           []struct {
					Owner struct {
						AddressOwner string `json:"AddressOwner"`
					} `json:"owner"`
					Reference struct {
						ObjectId string `json:"objectId"`
						Version  int    `json:"version"`
						Digest   string `json:"digest"`
					} `json:"reference"`
				} `json:"mutated"`
				GasObject struct {
					Owner struct {
						ObjectOwner string `json:"ObjectOwner"`
					} `json:"owner"`
					Reference struct {
						ObjectId string `json:"objectId"`
						Version  int    `json:"version"`
						Digest   string `json:"digest"`
					} `json:"reference"`
				} `json:"gasObject"`
				EventsDigest string `json:"eventsDigest"`
			} `json:"effects"`
			ObjectChanges []struct {
				Type      string `json:"type"`
				Sender    string `json:"sender"`
				Recipient struct {
					AddressOwner string `json:"AddressOwner"`
				} `json:"recipient"`
				ObjectType string `json:"objectType"`
				ObjectId   string `json:"objectId"`
				Version    string `json:"version"`
				Digest     string `json:"digest"`
			} `json:"objectChanges"`
		} `json:"result"`
	}
)
