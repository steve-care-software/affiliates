package databases

import (
	application_founders "github.com/steve-care-software/affiliates/applications/founders"
	"github.com/steve-care-software/affiliates/domain/founders"
	databases "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type founderApplication struct {
	databaseApp       databases.Application
	repositoryBuilder founders.RepositoryBuilder
	serviceBuilder    founders.ServiceBuilder
	dbName            string
	kind              uint
}

func createFounderApplication(
	databaseApp databases.Application,
	repositoryBuilder founders.RepositoryBuilder,
	serviceBuilder founders.ServiceBuilder,
	dbName string,
	kind uint,
) application_founders.Application {
	out := founderApplication{
		databaseApp:       databaseApp,
		repositoryBuilder: repositoryBuilder,
		serviceBuilder:    serviceBuilder,
		dbName:            dbName,
		kind:              kind,
	}

	return &out
}

// List lists the founders
func (app *founderApplication) List() ([]hash.Hash, error) {
	pContext, err := app.databaseApp.Open(app.dbName)
	if err != nil {
		return nil, err
	}

	defer app.databaseApp.Close(*pContext)
	repository, err := app.repositoryBuilder.Create().WithContext(*pContext).WithKind(app.kind).Now()
	if err != nil {
		return nil, err
	}

	return repository.List()
}

// Retrieve retrieves a founder by hash
func (app *founderApplication) Retrieve(hash hash.Hash) (founders.Founder, error) {
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

// RetrieveList retrieves a founder list by hashes
func (app *founderApplication) RetrieveList(hashes []hash.Hash) (founders.Founders, error) {
	pContext, err := app.databaseApp.Open(app.dbName)
	if err != nil {
		return nil, err
	}

	defer app.databaseApp.Close(*pContext)
	repository, err := app.repositoryBuilder.Create().WithContext(*pContext).WithKind(app.kind).Now()
	if err != nil {
		return nil, err
	}

	return repository.RetrieveList(hashes)
}
