package applications

import (
	"github.com/steve-care-software/affiliates/applications/affiliates"
	"github.com/steve-care-software/affiliates/applications/calculations"
	"github.com/steve-care-software/affiliates/applications/lines"
	"github.com/steve-care-software/affiliates/applications/trees"
	"github.com/steve-care-software/affiliates/domain/genesis"
)

// Application represents an application
type Application interface {
	Genesis() (genesis.Genesis, error)
	Affiliate() affiliates.Application
	Line() lines.Application
	Tree() trees.Application
	Calculation() calculations.Application
}
