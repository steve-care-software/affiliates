package genesis

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/affiliates/domain/genesis/levels"
	fungible_unit_genesis "github.com/steve-care-software/fungible-unit-genesis/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter         hash.Adapter
	parent              fungible_unit_genesis.Genesis
	shareHoldersPercent float64
	level               levels.Level
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:         hashAdapter,
		parent:              nil,
		shareHoldersPercent: 0,
		level:               nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent fungible_unit_genesis.Genesis) Builder {
	app.parent = parent
	return app
}

// WithShareHoldersPercent adds a shareHolder's percent to the builder
func (app *builder) WithShareHoldersPercent(shareHoldersPercent float64) Builder {
	app.shareHoldersPercent = shareHoldersPercent
	return app
}

// WithLevel adds a level to the builder
func (app *builder) WithLevel(level levels.Level) Builder {
	app.level = level
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.parent == nil {
		return nil, errors.New("the parent is mandatory in order to build a Genesis instance")
	}

	if app.shareHoldersPercent <= 0 {
		return nil, errors.New("the shareHolder's percent is mandatory in order to build a Genesis instance")
	}

	if app.level == nil {
		return nil, errors.New("the level is mandatory in order to build a Genesis instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.parent.Hash().Bytes(),
		[]byte(fmt.Sprintf("%f", app.shareHoldersPercent)),
		app.level.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createGenesis(*pHash, app.parent, app.shareHoldersPercent, app.level), nil
}
