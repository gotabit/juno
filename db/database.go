package db

import (
	"github.com/angelorc/desmos-parser/config"
	"github.com/angelorc/desmos-parser/types"
	"github.com/cosmos/cosmos-sdk/codec"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// Database represents an abstract database that can be used to save data inside it
type Database interface {
	// HasBlock tells whether or not the database has already stored the block having the given height.
	// An error is returned if the operation fails.
	HasBlock(height int64) (bool, error)

	// SaveBlock will be called when a new block is parsed, passing the block itself
	// and the transactions contained inside that block.
	// An error is returned if the operation fails.
	// NOTE. For each transaction inside txs, SaveTx will be called as well.
	SaveBlock(block *tmctypes.ResultBlock, totalGas, preCommits uint64) error

	// SaveTx will be called to save each transaction contained inside a block.
	// An error is returned if the operation fails.
	SaveTx(tx types.Tx) error

	// HasValidator returns true if a given validator by HEX address exists.
	// An error is returned if the operation fails.
	HasValidator(address string) (bool, error)

	// SetValidator stores a validator if it does not already exist.
	// An error is returned if the operation fails.
	SaveValidator(address, publicKey string) error

	// SetPreCommit stores a validator's pre-commit.
	// An error is returned if the operation fails.
	SavePreCommit(pc *tmtypes.CommitSig, votingPower, proposerPriority int64) error
}

// Builder represents a method that allows to build any database from a given codec and configuration
type Builder func(config.Config, *codec.Codec) (*Database, error)
