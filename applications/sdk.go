package applications

import (
	"github.com/steve-care-software/affiliates/applications/affiliates"
	"github.com/steve-care-software/affiliates/applications/calculations"
	"github.com/steve-care-software/affiliates/applications/founders"
	"github.com/steve-care-software/affiliates/applications/lines"
	"github.com/steve-care-software/affiliates/applications/trees"
)

// Application represents an application
type Application interface {
	Affiliate() affiliates.Application
	Line() lines.Application
	Tree() trees.Application
	Founders() founders.Application
	Calculation() calculations.Application
}
