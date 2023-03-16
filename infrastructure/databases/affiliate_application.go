package databases

import (
	application_affiliates "github.com/steve-care-software/affiliates/applications/affiliates"
	affiliates "github.com/steve-care-software/affiliates/domain"
	databases "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type affiliateApplication struct {
	databaseApp       databases.Application
	repositoryBuilder affiliates.RepositoryBuilder
	serviceBuilder    affiliates.ServiceBuilder
	dbName            string
	kind              uint
}

func createAffiliateApplication(
	databaseApp databases.Application,
	repositoryBuilder affiliates.RepositoryBuilder,
	serviceBuilder affiliates.ServiceBuilder,
	dbName string,
	kind uint,
) application_affiliates.Application {
	out := affiliateApplication{
		databaseApp:       databaseApp,
		repositoryBuilder: repositoryBuilder,
		serviceBuilder:    serviceBuilder,
		dbName:            dbName,
		kind:              kind,
	}

	return &out
}

// List returns the list of affiliate hashes that have no parent
func (app *affiliateApplication) List() ([]hash.Hash, error) {
	return nil, nil
}

// ListByParent returns the list of affiliate hashes by parent hash
func (app *affiliateApplication) ListByParent(parent hash.Hash) ([]hash.Hash, error) {
	return nil, nil
}

// Retrieve retrieves an affiliate by hash
func (app *affiliateApplication) Retrieve(hash hash.Hash) (affiliates.Affiliate, error) {
	pContext, err := app.databaseApp.Open(app.dbName)
	if err != nil {
		return nil, err
	}

	defer app.databaseApp.Close(*pContext)
	repository, err := app.repositoryBuilder.Create().WithContext(*pContext).WithKind(app.kind).Now()
	if err != nil {
		return nil, err
	}

	return repository.Retrieve(hash)
}

// RetrieveList retrieves a list of affiliates by hashes
func (app *affiliateApplication) RetrieveList(hashes []hash.Hash) ([]affiliates.Affiliate, error) {
	return nil, nil
}

// Insert inserts an affiliate
func (app *affiliateApplication) Insert(affiliate affiliates.Affiliate) error {
	pContext, err := app.databaseApp.Open(app.dbName)
	if err != nil {
		return err
	}

	defer app.databaseApp.Close(*pContext)
	service, err := app.serviceBuilder.Create().WithContext(*pContext).WithKind(app.kind).Now()
	if err != nil {
		return err
	}

	err = service.Insert(affiliate)
	if err != nil {
		return app.databaseApp.Cancel(*pContext)
	}

	return app.databaseApp.Commit(*pContext)
}
