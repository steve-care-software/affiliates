package domain

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type affiliate struct {
	hash      hash.Hash
	pubKey    hash.Hash
	createdOn time.Time
	parent    hash.Hash
}

func createAffiliate(
	hash hash.Hash,
	pubKey hash.Hash,
	createdOn time.Time,
) Affiliate {
	return createAffiliateInternally(hash, pubKey, createdOn, nil)
}

func createAffiliateWithParent(
	hash hash.Hash,
	pubKey hash.Hash,
	createdOn time.Time,
	parent hash.Hash,
) Affiliate {
	return createAffiliateInternally(hash, pubKey, createdOn, parent)
}

func createAffiliateInternally(
	hash hash.Hash,
	pubKey hash.Hash,
	createdOn time.Time,
	parent hash.Hash,
) Affiliate {
	out := affiliate{
		hash:      hash,
		pubKey:    pubKey,
		createdOn: createdOn,
		parent:    parent,
	}

	return &out
}

// Hash returns the hash
func (obj *affiliate) Hash() hash.Hash {
	return obj.hash
}

// PubKey returns the public key
func (obj *affiliate) PubKey() hash.Hash {
	return obj.hash
}

// CreatedOn returns the creation time
func (obj *affiliate) CreatedOn() time.Time {
	return obj.createdOn
}

// HasParent returns true if there is a parent, false otherwise
func (obj *affiliate) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *affiliate) Parent() hash.Hash {
	return obj.parent
}
