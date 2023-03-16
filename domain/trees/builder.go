package trees

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Tree
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Tree) Builder {
	app.list = list
	return app
}

// Now builds a new Trees instance
func (app *builder) Now() (Trees, error) {
	if app.list == nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Tree in order to build a Trees instance")
	}

	data := [][]byte{}
	for _, oneTree := range app.list {
		data = append(data, oneTree.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createTrees(*pHash, app.list), nil
}
