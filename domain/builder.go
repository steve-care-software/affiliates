package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pubKey      hash.Hash
	pCreatedOn  *time.Time
	parent      hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pubKey:      nil,
		pCreatedOn:  nil,
		parent:      nil,
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

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent hash.Hash) Builder {
	app.parent = parent
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Affiliate instance
func (app *builder) Now() (Affiliate, error) {
	if app.pubKey == nil {
		return nil, errors.New("the pubKey hash is mandatory in order to build an Affiliate instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build an Affiliate instance")
	}

	data := [][]byte{
		app.pubKey.Bytes(),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.Unix())),
	}

	if app.parent != nil {
		data = append(data, app.parent.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createAffiliateWithParent(*pHash, app.pubKey, *app.pCreatedOn, app.parent), nil
	}

	return createAffiliate(*pHash, app.pubKey, *app.pCreatedOn), nil
}
