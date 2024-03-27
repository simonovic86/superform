package model

// Transaction represents a single transaction
type Transaction struct {
	TransactionHash   string `json:"transaction_hash"`
	TxType            int    `json:"tx_type"`
	Sender            string `json:"sender"`
	FromChain         int    `json:"from_chain"`
	ToChain           int    `json:"to_chain"`
	CreationTimestamp int64  `json:"creation_timestamp"`
	SuperformID       string `json:"superform_id"`
	Vault             string `json:"vault"`
	Protocol          string `json:"protocol"`
	Status            int    `json:"status"`
}

// DetailedTransaction represents the detailed view of a transaction with additional information
type DetailedTransaction struct {
	Type                string             `json:"type"`
	ID                  int                `json:"id"`
	Hash                string             `json:"hash"`
	Chain               string             `json:"chain"`
	Status              string             `json:"status"`
	SuperformID         string             `json:"superform_ID"`
	Phases              []TransactionPhase `json:"phases"`
	Rows                []Row              `json:"rows"`
	CreationTimestamp   int64              `json:"creation_timestamp"`    // artificially added
	ProcessedTimestamp  int64              `json:"processed_timestamp"`   // artificially added
	Amount              string             `json:"amount"`                // artificially added
	InitialOutputAmount string             `json:"initial_output_amount"` // artificially added
	DelayedOutputAmount string             `json:"delayed_output_amount"` // artificially added
	// Add other fields as necessary
}

// Row represents a row within a transaction
type Row struct {
	Component string      `json:"component"`
	Label     string      `json:"label"`
	Status    string      `json:"status"`
	Name      string      `json:"name"`
	ChainID   int         `json:"chainId"`
	Hash      string      `json:"hash"`
	Value     interface{} `json:"value"`
	Address   string      `json:"address"`
	// Include other fields as necessary
}

type ProcessedTransaction struct {
	Type                string `json:"type"`
	ID                  int    `json:"id"`
	Hash                string `json:"hash"`
	Chain               string `json:"chain"`
	Status              string `json:"status"`
	SuperformId         string `json:"superform_id"`
	Amount              string `json:"amount"`                // artificially added
	InitialTimestamp    int64  `json:"initial_timestamp"`     // artificially added
	DelayedTimestamp    int64  `json:"delayed_timestamp"`     // artificially added
	InitialOutputAmount string `json:"initial_output_amount"` // artificially added
	DelayedOutputAmount string `json:"delayed_output_amount"` // artificially added
	// Add other fields as necessary
}

// TransactionPhase represents a phase in a transaction's lifecycle
type TransactionPhase struct {
	Status string     `json:"status"`
	Name   string     `json:"name"`
	Rows   []PhaseRow `json:"rows"`
}

// PhaseRow represents a row within a transaction phase
type PhaseRow struct {
	Component string      `json:"component"`
	Label     string      `json:"label"`
	Status    string      `json:"status"` // For AMBs
	Name      string      `json:"name"`   // For AMBs
	ChainID   int         `json:"chainId"`
	Hash      string      `json:"hash"`
	Value     interface{} `json:"value"`
	Amount    string      `json:"amount"`
	// Include other fields as necessary
}

type Chain struct {
	ChainID   int      `json:"chainId"`
	Name      string   `json:"name"`
	Rpcs      []string `json:"rpcs"`
	Explorers []string `json:"explorers"`
}

type Currency struct {
	Icon                    string `json:"icon"`
	Name                    string `json:"name"`
	Symbol                  string `json:"symbol"`
	Decimals                int    `json:"decimals"`
	MinNativeCurrencyForGas string `json:"minNativeCurrencyForGas"`
}
