package main

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

var cdc *codec.Codec

func init() {
	cdc = codec.New()

	// bank.RegisterCodec(cdc)
	// staking.RegisterCodec(cdc)
	// distr.RegisterCodec(cdc)
	// slashing.RegisterCodec(cdc)
	// gov.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
}
