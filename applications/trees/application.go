package trees

import (
	application_affiliates "github.com/steve-care-software/affiliates/applications/affiliates"
	"github.com/steve-care-software/affiliates/domain/trees"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type application struct {
	affiliateApp application_affiliates.Application
	builder      trees.Builder
	treeBuilder  trees.TreeBuilder
}

func createApplication(
	affiliateApp application_affiliates.Application,
	builder trees.Builder,
	treeBuilder trees.TreeBuilder,
) Application {
	out := application{
		affiliateApp: affiliateApp,
		builder:      builder,
		treeBuilder:  treeBuilder,
	}

	return &out
}

// Retrieve retrieves a tree by affiliate hash
func (app *application) Retrieve(hash hash.Hash) (trees.Tree, error) {
	// retrieve the affiliate:
	affiliate, err := app.affiliateApp.Retrieve(hash)
	if err != nil {
		return nil, err
	}

	// retrieve the children:
	hashList, err := app.affiliateApp.ListByParent(hash)
	if err != nil {
		return nil, err
	}

	pubKey := affiliate.PubKey()
	builder := app.treeBuilder.Create().WithPubKey(pubKey)
	if len(hashList) > 0 {
		children, err := app.RetrieveList(hashList)
		if err != nil {
			return nil, err
		}

		builder.WithChildren(children)
	}

	return builder.Now()
}

// RetrieveList retrieves trees by hashes
func (app *application) RetrieveList(hashes []hash.Hash) (trees.Trees, error) {
	list := []trees.Tree{}
	for _, oneHash := range hashes {
		ins, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.builder.Create().
		WithList(list).
		Now()
}
