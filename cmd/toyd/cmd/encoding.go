package cmd

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	chainparams "github.com/decento-labs/toy/app/params"
)

type EncodingConfig = chainparams.EncodingConfig

// makeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func makeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(codec, tx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

// MakeEncodingConfig creates an EncodingConfig for testing
func MakeEncodingConfig(moduleBasics module.BasicManager) EncodingConfig {
	encodingConfig := makeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	moduleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	moduleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	return encodingConfig
}
