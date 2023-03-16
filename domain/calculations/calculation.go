package calculations

import "github.com/steve-care-software/libs/cryptography/hash"

type calculation struct {
	hash   hash.Hash
	pubKey hash.Hash
	amount uint
}

func createCalculation(
	hash hash.Hash,
	pubKey hash.Hash,
	amount uint,
) Calculation {
	out := calculation{
		hash:   hash,
		pubKey: pubKey,
		amount: amount,
	}

	return &out
}

// Hash returns the hash
func (obj *calculation) Hash() hash.Hash {
	return obj.hash
}

// PubKey returns the pubKey
func (obj *calculation) PubKey() hash.Hash {
	return obj.pubKey
}

// Amount returns the amount
func (obj *calculation) Amount() uint {
	return obj.amount
}
