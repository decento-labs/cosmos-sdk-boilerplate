package cmd

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	chain "github.com/decento-labs/toy/app"
)

var (
	// AddressVerifier address verifier
	AddressVerifier = func(bz []byte) error {
		if n := len(bz); n != 20 && n != 32 {
			return fmt.Errorf("incorrect address length %d", n)
		}

		return nil
	}
)

func SetPrefixes() {
	config := sdk.GetConfig()
	config.SetPurpose(sdk.Purpose) // BIP-44 path "44'/118'/0'/0/0"
	config.SetCoinType(sdk.CoinType)
	config.SetBech32PrefixForAccount(chain.Bech32PrefixAccAddr, chain.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(chain.Bech32PrefixValAddr, chain.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(chain.Bech32PrefixConsAddr, chain.Bech32PrefixConsPub)
	config.SetAddressVerifier(AddressVerifier)
	config.Seal()
}
