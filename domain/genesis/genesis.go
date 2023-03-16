package genesis

import (
	"github.com/steve-care-software/affiliates/domain/genesis/levels"
	fungible_unit_genesis "github.com/steve-care-software/fungible-unit-genesis/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type genesis struct {
	hash                hash.Hash
	parent              fungible_unit_genesis.Genesis
	shareHoldersPercent float64
	level               levels.Level
}

func createGenesis(
	hash hash.Hash,
	parent fungible_unit_genesis.Genesis,
	shareHoldersPercent float64,
	level levels.Level,
) Genesis {
	out := genesis{
		hash:                hash,
		parent:              parent,
		shareHoldersPercent: shareHoldersPercent,
		level:               level,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// Parent returns the parent
func (obj *genesis) Parent() fungible_unit_genesis.Genesis {
	return obj.parent
}

// ShareHoldersPercent returns the shareHolder's percent
func (obj *genesis) ShareHoldersPercent() float64 {
	return obj.shareHoldersPercent
}

// Level returns the level
func (obj *genesis) Level() levels.Level {
	return obj.level
}
