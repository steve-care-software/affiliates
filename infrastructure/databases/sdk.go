package databases

import (
	application_affiliates "github.com/steve-care-software/affiliates/applications/affiliates"
	application_founders "github.com/steve-care-software/affiliates/applications/founders"
	affiliates "github.com/steve-care-software/affiliates/domain"
	"github.com/steve-care-software/affiliates/domain/founders"
	databases "github.com/steve-care-software/databases/applications"
)

// NewFounderApplication creates a new founder application
func NewFounderApplication(
	databaseApp databases.Application,
	repositoryBuilder founders.RepositoryBuilder,
	serviceBuilder founders.ServiceBuilder,
	dbName string,
	kind uint,
) application_founders.Application {
	return createFounderApplication(
		databaseApp,
		repositoryBuilder,
		serviceBuilder,
		dbName,
		kind,
	)
}

// NewAffiliateApplication creates a new affiliate application
func NewAffiliateApplication(
	databaseApp databases.Application,
	repositoryBuilder affiliates.RepositoryBuilder,
	serviceBuilder affiliates.ServiceBuilder,
	dbName string,
	kind uint,
) application_affiliates.Application {
	return createAffiliateApplication(
		databaseApp,
		repositoryBuilder,
		serviceBuilder,
		dbName,
		kind,
	)
}
