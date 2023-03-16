package domain

import "github.com/steve-care-software/libs/cryptography/hash"

// Builder represents an affiliate
type Builder interface {
	Create() Builder
	WithPubKey(hash hash.Hash) Builder
	WithParent(parent hash.Hash) Builder
	Now() (Affiliate, error)
}

// Affiliate represents an affiliate
type Affiliate interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	HasParent() bool
	Parent() hash.Hash
}
