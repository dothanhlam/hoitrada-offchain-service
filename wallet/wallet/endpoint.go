package wallet

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type storeRequest struct {
	PassPhrase string `json:"passphrase"`
}

type storeResponse struct {
	Uuid string `json:"uuid,omitempty"`
	Err  string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeStoreEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(storeRequest)

		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ := bip39.NewMnemonic(entropy)

		// Generate a Bip32 HD wallet for the mnemonic and a user supplied password

		seed := bip39.NewSeed(mnemonic, req.PassPhrase)

		masterKey, _ := bip32.NewMasterKey(seed)
		publicKey := masterKey.PublicKey()

		v := Wallet{
			Mnemonic:  mnemonic,
			MasterKey: masterKey.String(),
			PublicKey: publicKey.String(),
		}
		_, err := svc.Store(ctx, v)
		if err != nil {
			return storeResponse{"", err.Error()}, err
		}

		return storeResponse{v.ID.String(), ""}, err
	}
}
