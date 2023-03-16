package founders

import "github.com/steve-care-software/libs/cryptography/hash"

// Founder represents the affiliate founders
type Founder interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	Power() uint
}
