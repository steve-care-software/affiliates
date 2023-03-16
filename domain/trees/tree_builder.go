package trees

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type treeBuilder struct {
	hashAdapter hash.Adapter
	pubKey      hash.Hash
	children    Trees
}

func createTreeBuilder(
	hashAdapter hash.Adapter,
) TreeBuilder {
	out := treeBuilder{
		hashAdapter: hashAdapter,
		pubKey:      nil,
		children:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *treeBuilder) Create() TreeBuilder {
	return createTreeBuilder(app.hashAdapter)
}

// WithPubKey adds a pubKey to the builder
func (app *treeBuilder) WithPubKey(pubKey hash.Hash) TreeBuilder {
	app.pubKey = pubKey
	return app
}

// WithChildren adds a children to the builder
func (app *treeBuilder) WithChildren(children Trees) TreeBuilder {
	app.children = children
	return app
}

// Now builds a new Tree instance
func (app *treeBuilder) Now() (Tree, error) {
	if app.pubKey == nil {
		return nil, errors.New("the pubKey is mandatory in order to build a Tree instance")
	}

	data := [][]byte{
		app.pubKey.Bytes(),
	}

	if app.children != nil {
		data = append(data, app.children.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.children != nil {
		return createTreeWithChildren(*pHash, app.pubKey, app.children), nil
	}

	return createTree(*pHash, app.pubKey), nil
}
