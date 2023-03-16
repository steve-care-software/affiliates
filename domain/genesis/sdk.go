package genesis

import (
	"github.com/steve-care-software/affiliates/domain/genesis/levels"
	fungible_unit_genesis "github.com/steve-care-software/fungible-unit-genesis/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Genesis represents a genesis
type Genesis interface {
	Hash() hash.Hash
	Parent() fungible_unit_genesis.Genesis
	ShareHoldersPercent() float64
	Level() levels.Level
}
