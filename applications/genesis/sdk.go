package genesis

import (
	"github.com/steve-care-software/affiliates/domain/genesis"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents the genesis application
type Application interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (genesis.Genesis, error)
	Insert(genesis genesis.Genesis) error
	Delete(hash hash.Hash) error
}
