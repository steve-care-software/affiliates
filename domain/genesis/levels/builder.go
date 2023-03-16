package levels

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	power       uint
	parent      Level
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		power:       0,
		parent:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithPower adds a power to the builder
func (app *builder) WithPower(power uint) Builder {
	app.power = power
	return app
}

// WithParent adds a parent level to the builder
func (app *builder) WithParent(parent Level) Builder {
	app.parent = parent
	return app
}

// Now builds a new Level instance
func (app *builder) Now() (Level, error) {
	if app.power <= 0 {
		return nil, errors.New("the power must be greater than zero (0) in order to build a Level instance")
	}

	data := [][]byte{
		[]byte(fmt.Sprintf("%d", app.power)),
	}

	if app.parent != nil {
		data = append(data, app.parent.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createLevelWithParent(*pHash, app.power, app.parent), nil
	}

	return createLevel(*pHash, app.power), nil
}
