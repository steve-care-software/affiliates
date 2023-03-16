package databases

import (
	application_affiliates "github.com/steve-care-software/affiliates/applications/affiliates"
	affiliates "github.com/steve-care-software/affiliates/domain"
	databases "github.com/steve-care-software/databases/applications"
)

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
