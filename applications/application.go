package applications

import (
	"github.com/steve-care-software/affiliates/applications/affiliates"
	"github.com/steve-care-software/affiliates/applications/calculations"
	"github.com/steve-care-software/affiliates/applications/genesis"
	"github.com/steve-care-software/affiliates/applications/lines"
	"github.com/steve-care-software/affiliates/applications/trees"
)

type application struct {
	genesis     genesis.Application
	affiliate   affiliates.Application
	line        lines.Application
	tree        trees.Application
	calculation calculations.Application
}

func createApplication(
	genesis genesis.Application,
	affiliate affiliates.Application,
	line lines.Application,
	tree trees.Application,
	calculation calculations.Application,
) Application {
	out := application{
		genesis:     genesis,
		affiliate:   affiliate,
		line:        line,
		tree:        tree,
		calculation: calculation,
	}

	return &out
}

// Genesis returns the genesis application
func (app *application) Genesis() genesis.Application {
	return app.genesis
}

// Affiliate returns the affiliate application
func (app *application) Affiliate() affiliates.Application {
	return app.affiliate
}

// Line returns the line application
func (app *application) Line() lines.Application {
	return app.line
}

// Tree returns the tree application
func (app *application) Tree() trees.Application {
	return app.tree
}

// Calculation returns the calculation application
func (app *application) Calculation() calculations.Application {
	return app.calculation
}
