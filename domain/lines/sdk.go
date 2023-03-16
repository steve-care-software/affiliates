package lines

import "github.com/steve-care-software/libs/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a line builder
type Builder interface {
	Create() Builder
	WithPubKey(pubKey hash.Hash) Builder
	WithChild(child Line) Builder
	Now() (Line, error)
}

// Line represents a referral line
type Line interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	HasChild() bool
	Child() Line
}
