package domain

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents an affiliate
type Builder interface {
	Create() Builder
	WithPubKey(pubKey hash.Hash) Builder
	WithParent(parent hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Affiliate, error)
}

// Affiliate represents an affiliate
type Affiliate interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	CreatedOn() time.Time
	HasParent() bool
	Parent() hash.Hash
}
