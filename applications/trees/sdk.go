package trees

import (
	"github.com/steve-care-software/affiliates/domain/trees"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents the tree application
type Application interface {
	Retrieve(hash hash.Hash) (trees.Tree, error)
}
