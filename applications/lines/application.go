package lines

import (
	application_affiliates "github.com/steve-care-software/affiliates/applications/affiliates"
	"github.com/steve-care-software/affiliates/domain/lines"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type application struct {
	affiliateApp application_affiliates.Application
	builder      lines.Builder
}

func createApplication(
	affiliateApp application_affiliates.Application,
	builder lines.Builder,
) Application {
	out := application{
		affiliateApp: affiliateApp,
		builder:      builder,
	}

	return &out
}

// Retrieve retrieves a line by affiliate hash
func (app *application) Retrieve(hash hash.Hash) (lines.Line, error) {
	// retrieve the affiliate:
	affiliate, err := app.affiliateApp.Retrieve(hash)
	if err != nil {
		return nil, err
	}

	pubKey := affiliate.PubKey()
	builder := app.builder.Create().WithPubKey(pubKey)
	if affiliate.HasParent() {
		parent := affiliate.Parent()
		child, err := app.Retrieve(parent)
		if err != nil {
			return nil, err
		}

		builder.WithChild(child)
	}

	return builder.Now()
}
