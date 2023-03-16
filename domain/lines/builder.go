package lines

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pubKey      hash.Hash
	child       Line
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pubKey:      nil,
		child:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithPubKey adds a pubKey to the builder
func (app *builder) WithPubKey(pubKey hash.Hash) Builder {
	app.pubKey = pubKey
	return app
}

// WithChild adds a child to the builder
func (app *builder) WithChild(child Line) Builder {
	app.child = child
	return app
}

// Now builds a new Line instance
func (app *builder) Now() (Line, error) {
	if app.pubKey == nil {
		return nil, errors.New("the pubKey is mandatory in order to build a Line instance")
	}

	data := [][]byte{
		app.pubKey.Bytes(),
	}

	if app.child != nil {
		data = append(data, app.pubKey.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.child != nil {
		return createLineWithChild(*pHash, app.pubKey, app.child), nil
	}

	return createLine(*pHash, app.pubKey), nil
}
