package founders

import "github.com/steve-care-software/libs/cryptography/hash"

// Builder represents the founder builder
type Builder interface {
	Create() Builder
	WithPubKey(pubKey hash.Hash) Builder
	WithPower(power uint) Builder
	Now() (Founder, error)
}

// Founder represents the affiliate founders
type Founder interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	Power() uint
}
