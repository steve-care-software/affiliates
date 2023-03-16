package calculations

import (
	"github.com/steve-care-software/affiliates/domain/calculations"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents the calculation
type Application interface {
	Execute(hash hash.Hash, amount uint) (calculations.Calculations, error)
}
