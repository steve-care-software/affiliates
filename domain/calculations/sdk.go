package calculations

import (
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewCalculationBuilder creates a new calculation builder
func NewCalculationBuilder() CalculationBuilder {
	hashAdapter := hash.NewAdapter()
	return createCalculationBuilder(hashAdapter)
}

// Builder represents the calculations builder
type Builder interface {
	Create() Builder
	WithList(list []Calculation) Builder
	Now() (Calculations, error)
}

// Calculations represents calculations
type Calculations interface {
	Hash() hash.Hash
	List() []Calculation
}

// CalculationBuilder represents the calculation builder
type CalculationBuilder interface {
	Create() CalculationBuilder
	WithPubKey(pubKey hash.Hash) CalculationBuilder
	WithAmount(amount uint) CalculationBuilder
	Now() (Calculation, error)
}

// Calculation represents a calculation
type Calculation interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	Amount() uint
}
