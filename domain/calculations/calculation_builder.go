package calculations

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type calculationBuilder struct {
	hashAdapter hash.Adapter
	pubKey      hash.Hash
	amount      uint
}

func createCalculationBuilder(
	hashAdapter hash.Adapter,
) CalculationBuilder {
	out := calculationBuilder{
		hashAdapter: hashAdapter,
		pubKey:      nil,
		amount:      0,
	}

	return &out
}

// Create initializes the builder
func (app *calculationBuilder) Create() CalculationBuilder {
	return createCalculationBuilder(app.hashAdapter)
}

// WithPubKey adds a public key to the builder
func (app *calculationBuilder) WithPubKey(pubKey hash.Hash) CalculationBuilder {
	app.pubKey = pubKey
	return app
}

// WithAmount adds an amount to the builder
func (app *calculationBuilder) WithAmount(amount uint) CalculationBuilder {
	app.amount = amount
	return app
}

// Now builds a new Calculation instance
func (app *calculationBuilder) Now() (Calculation, error) {
	if app.pubKey == nil {
		return nil, errors.New("the publicKey is mandatory in order to build a Calculation instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build a Calculation instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.pubKey.Bytes(),
		[]byte(fmt.Sprintf("%d", app.amount)),
	})

	if err != nil {
		return nil, err
	}

	return createCalculation(*pHash, app.pubKey, app.amount), nil
}
