package calculations

import (
	"github.com/steve-care-software/libs/cryptography/hash"
)

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
