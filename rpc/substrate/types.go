package substrate

type (
	SubscanExtrinsicReq struct {
		Hash string `json:"hash"`
	}
	SubscanExtrinsicRes struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		GeneratedAt int    `json:"generated_at"`
		Data        struct {
			BlockTimestamp     int    `json:"block_timestamp"`
			BlockNum           int    `json:"block_num"`
			ExtrinsicIndex     string `json:"extrinsic_index"`
			CallModuleFunction string `json:"call_module_function"`
			CallModule         string `json:"call_module"`
			AccountId          string `json:"account_id"`
			Signature          string `json:"signature"`
			Nonce              int    `json:"nonce"`
			ExtrinsicHash      string `json:"extrinsic_hash"`
			Success            bool   `json:"success"`
			Params             []struct {
				Name     string      `json:"name"`
				Type     string      `json:"type"`
				TypeName string      `json:"type_name"`
				Value    interface{} `json:"value"`
			} `json:"params"`
			Transfer struct {
				From             string `json:"from"`
				To               string `json:"to"`
				Module           string `json:"module"`
				Amount           string `json:"amount"`
				Hash             string `json:"hash"`
				Success          bool   `json:"success"`
				AssetSymbol      string `json:"asset_symbol"`
				ToAccountDisplay struct {
					Address string `json:"address"`
				} `json:"to_account_display"`
			} `json:"transfer"`
			Event []struct {
				EventIndex     string `json:"event_index"`
				BlockNum       int    `json:"block_num"`
				ExtrinsicIdx   int    `json:"extrinsic_idx"`
				ModuleId       string `json:"module_id"`
				EventId        string `json:"event_id"`
				Params         string `json:"params"`
				Phase          int    `json:"phase"`
				EventIdx       int    `json:"event_idx"`
				ExtrinsicHash  string `json:"extrinsic_hash"`
				Finalized      bool   `json:"finalized"`
				BlockTimestamp int    `json:"block_timestamp"`
			} `json:"event"`
			EventCount int    `json:"event_count"`
			Fee        string `json:"fee"`
			FeeUsed    string `json:"fee_used"`
			Error      struct {
				Module     string `json:"module"`
				Name       string `json:"name"`
				Doc        string `json:"doc"`
				Value      string `json:"value"`
				BatchIndex int    `json:"batch_index"`
			} `json:"error"`
			Finalized bool `json:"finalized"`
			Lifetime  struct {
				Birth int `json:"birth"`
				Death int `json:"death"`
			} `json:"lifetime"`
			Tip            string `json:"tip"`
			AccountDisplay struct {
				Address string `json:"address"`
			} `json:"account_display"`
			CrosschainOp interface{} `json:"crosschain_op"`
			BlockHash    string      `json:"block_hash"`
			Pending      bool        `json:"pending"`
		} `json:"data"`
	}
)
