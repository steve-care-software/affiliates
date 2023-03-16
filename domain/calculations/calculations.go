package calculations

import "github.com/steve-care-software/libs/cryptography/hash"

type calculations struct {
	hash hash.Hash
	list []Calculation
}

func createCalculations(
	hash hash.Hash,
	list []Calculation,
) Calculations {
	out := calculations{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *calculations) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *calculations) List() []Calculation {
	return obj.list
}
