package domain

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

// Builder represents an affiliate
type Builder interface {
	Create() Builder
	WithPubKey(hash hash.Hash) Builder
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
