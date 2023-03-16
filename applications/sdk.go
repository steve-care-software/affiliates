package applications

import (
	"github.com/steve-care-software/affiliates/applications/affiliates"
	"github.com/steve-care-software/affiliates/applications/calculations"
	"github.com/steve-care-software/affiliates/applications/genesis"
	"github.com/steve-care-software/affiliates/applications/lines"
	"github.com/steve-care-software/affiliates/applications/trees"
)

// Application represents an application
type Application interface {
	Genesis() genesis.Application
	Affiliate() affiliates.Application
	Line() lines.Application
	Tree() trees.Application
	Calculation() calculations.Application
}
