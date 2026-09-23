package algorand

type (
	SendTxOut struct {
		Message string `json:"message"`
		TxId    string `json:"txId"`
	}
	TransactionParams struct {
		Message          string `json:"message"`
		ConsensusVersion string `json:"consensus-version"`
		Fee              uint64 `json:"fee"`
		GenesisHash      string `json:"genesis-hash"`
		GenesisId        string `json:"genesis-id"`
		LastRound        uint64 `json:"last-round"`
		MinFee           uint64 `json:"min-fee"`
	}
	NodeStatusResponse struct {
		Message                     string `json:"message"`
		Catchpoint                  string `json:"catchpoint,omitempty"`
		CatchpointAcquiredBlocks    uint64 `json:"catchpoint-acquired-blocks,omitempty"`
		CatchpointProcessedAccounts uint64 `json:"catchpoint-processed-accounts,omitempty"`
		CatchpointTotalAccounts     uint64 `json:"catchpoint-total-accounts,omitempty"`
		CatchpointTotalBlocks       uint64 `json:"catchpoint-total-blocks,omitempty"`
		CatchpointVerifiedAccounts  uint64 `json:"catchpoint-verified-accounts,omitempty"`
		CatchupTime                 uint64 `json:"catchup-time"`
		LastCatchpoint              string `json:"last-catchpoint,omitempty"`
		LastRound                   uint64 `json:"last-round"`
		LastVersion                 string `json:"last-version"`
		NextVersion                 string `json:"next-version"`
		NextVersionRound            uint64 `json:"next-version-round"`
		NextVersionSupported        bool   `json:"next-version-supported"`
		StoppedAtUnsupportedRound   bool   `json:"stopped-at-unsupported-round"`
		TimeSinceLastRound          uint64 `json:"time-since-last-round"`
	}
	ApplicationStateSchema struct {
		NumByteSlice uint64 `json:"num-byte-slice"`
		NumUint      uint64 `json:"num-uint"`
	}
	AssetHolding struct {
		Amount          uint64 `json:"amount"`
		AssetId         uint64 `json:"asset-id"`
		Creator         string `json:"creator"`
		Deleted         bool   `json:"deleted,omitempty"`
		IsFrozen        bool   `json:"is-frozen"`
		OptedInAtRound  uint64 `json:"opted-in-at-round,omitempty"`
		OptedOutAtRound uint64 `json:"opted-out-at-round,omitempty"`
	}
	TealValue struct {
		Bytes string `json:"bytes"`
		Type  uint64 `json:"type"`
		Uint  uint64 `json:"uint"`
	}

	TealKeyValue struct {
		Key   string    `json:"key"`
		Value TealValue `json:"value"`
	}

	ApplicationParams struct {
		ApprovalProgram   []byte                 `json:"approval-program"`
		ClearStateProgram []byte                 `json:"clear-state-program"`
		Creator           string                 `json:"creator,omitempty"`
		ExtraProgramPages uint64                 `json:"extra-program-pages,omitempty"`
		GlobalState       []TealKeyValue         `json:"global-state,omitempty"`
		GlobalStateSchema ApplicationStateSchema `json:"global-state-schema,omitempty"`
		LocalStateSchema  ApplicationStateSchema `json:"local-state-schema,omitempty"`
	}

	Application struct {
		CreatedAtRound uint64            `json:"created-at-round,omitempty"`
		Deleted        bool              `json:"deleted,omitempty"`
		DeletedAtRound uint64            `json:"deleted-at-round,omitempty"`
		Id             uint64            `json:"id"`
		Params         ApplicationParams `json:"params"`
	}
	ApplicationLocalState struct {
		ClosedOutAtRound uint64                 `json:"closed-out-at-round,omitempty"`
		Deleted          bool                   `json:"deleted,omitempty"`
		Id               uint64                 `json:"id"`
		KeyValue         []TealKeyValue         `json:"key-value,omitempty"`
		OptedInAtRound   uint64                 `json:"opted-in-at-round,omitempty"`
		Schema           ApplicationStateSchema `json:"schema"`
	}
	AssetParams struct {
		Clawback      string `json:"clawback,omitempty"`
		Creator       string `json:"creator"`
		Decimals      uint64 `json:"decimals"`
		DefaultFrozen bool   `json:"default-frozen,omitempty"`
		Freeze        string `json:"freeze,omitempty"`
		Manager       string `json:"manager,omitempty"`
		MetadataHash  []byte `json:"metadata-hash,omitempty"`
		Name          string `json:"name,omitempty"`
		NameB64       []byte `json:"name-b64,omitempty"`
		Reserve       string `json:"reserve,omitempty"`
		Total         uint64 `json:"total"`
		UnitName      string `json:"unit-name,omitempty"`
		UnitNameB64   []byte `json:"unit-name-b64,omitempty"`
		Url           string `json:"url,omitempty"`
		UrlB64        []byte `json:"url-b64,omitempty"`
	}

	Asset struct {
		CreatedAtRound   uint64      `json:"created-at-round,omitempty"`
		Deleted          bool        `json:"deleted,omitempty"`
		DestroyedAtRound uint64      `json:"destroyed-at-round,omitempty"`
		Index            uint64      `json:"index"`
		Params           AssetParams `json:"params"`
	}

	AccountParticipation struct {
		SelectionParticipationKey []byte `json:"selection-participation-key"`
		VoteFirstValid            uint64 `json:"vote-first-valid"`
		VoteKeyDilution           uint64 `json:"vote-key-dilution"`
		VoteLastValid             uint64 `json:"vote-last-valid"`
		VoteParticipationKey      []byte `json:"vote-participation-key"`
	}
	AccountAddress struct {
		Message                     string                  `json:"message"`
		Address                     string                  `json:"address"`
		Amount                      uint64                  `json:"amount"`
		AmountWithoutPendingRewards uint64                  `json:"amount-without-pending-rewards"`
		AppsLocalState              []ApplicationLocalState `json:"apps-local-state,omitempty"`
		AppsTotalExtraPages         uint64                  `json:"apps-total-extra-pages,omitempty"`
		AppsTotalSchema             ApplicationStateSchema  `json:"apps-total-schema,omitempty"`
		Assets                      []AssetHolding          `json:"assets,omitempty"`
		AuthAddr                    string                  `json:"auth-addr,omitempty"`
		ClosedAtRound               uint64                  `json:"closed-at-round,omitempty"`
		CreatedApps                 []Application           `json:"created-apps,omitempty"`
		CreatedAssets               []Asset                 `json:"created-assets,omitempty"`
		CreatedAtRound              uint64                  `json:"created-at-round,omitempty"`
		Deleted                     bool                    `json:"deleted,omitempty"`
		Participation               AccountParticipation    `json:"participation,omitempty"`
		PendingRewards              uint64                  `json:"pending-rewards"`
		RewardBase                  uint64                  `json:"reward-base,omitempty"`
		Rewards                     uint64                  `json:"rewards"`
		Round                       uint64                  `json:"round"`
		SigType                     string                  `json:"sig-type,omitempty"`
		Status                      string                  `json:"status"`
	}

	PendingTransactionResponse struct {
		Message            string `json:"message,omitempty"`
		ApplicationIndex   uint64 `json:"application-index,omitempty"`
		AssetClosingAmount uint64 `json:"asset-closing-amount,omitempty"`
		AssetIndex         uint64 `json:"asset-index,omitempty"`
		CloseRewards       uint64 `json:"close-rewards,omitempty"`
		ClosingAmount      uint64 `json:"closing-amount,omitempty"`
		ConfirmedRound     uint64 `json:"confirmed-round,omitempty"`
		PoolError          string `json:"pool-error"`
		Txn                struct {
			Sig string `json:"sig"`
			Txn struct {
				Type string `json:"type"`
				Snd  string `json:"snd"`
				Rcv  string `json:"rcv"`
				Arcv string `json:"arcv"`
				Aamt uint64 `json:"aamt"`
				Amt  uint64 `json:"amt"`
				Xaid uint64 `json:"xaid"`
				Fee  uint64 `json:"fee"`
				Note string `json:"note"`
			} `json:"txn"`
		} `json:"txn"`
	}
)

type IndexerTransactionRes struct {
	Message     string `json:"message,omitempty"`
	Transaction struct {
		Sender             string `json:"sender"`
		TxType             string `json:"tx-type"`
		ConfirmedRound     uint64 `json:"confirmed-round"`
		Fee                uint64 `json:"fee"`
		PaymentTransaction struct {
			Amount   uint64 `json:"amount"`
			Receiver string `json:"receiver"`
		} `json:"payment-transaction"`
		AssetTransferTransaction struct {
			Amount   uint64 `json:"amount"`
			AssetID  uint64 `json:"asset-id"`
			Receiver string `json:"receiver"`
		} `json:"asset-transfer-transaction"`
	} `json:"transaction"`
}
