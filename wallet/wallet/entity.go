package wallet

import "github.com/google/uuid"

type Wallet struct {
	ID         uuid.UUID
	Mnemonic   string `json:"mnemonic,omitempty"`
	MasterKey  string `json:"masterKey,omitempty"`
	PublicKey  string `json:"publicKey,omitempty"`
	PassPharse string `json:"passPharse,omitempty"`
}
