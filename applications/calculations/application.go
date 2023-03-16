package calculations

import (
	"math"

	"github.com/steve-care-software/affiliates/applications/genesis"
	application_line "github.com/steve-care-software/affiliates/applications/lines"
	"github.com/steve-care-software/affiliates/domain/calculations"
	"github.com/steve-care-software/affiliates/domain/genesis/levels"
	"github.com/steve-care-software/affiliates/domain/lines"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type application struct {
	genesisApp         genesis.Application
	lineApp            application_line.Application
	builder            calculations.Builder
	calculationBuilder calculations.CalculationBuilder
}

func createApplication(
	genesisApp genesis.Application,
	lineApp application_line.Application,
	builder calculations.Builder,
	calculationBuilder calculations.CalculationBuilder,
) Application {
	out := application{
		genesisApp:         genesisApp,
		lineApp:            lineApp,
		builder:            builder,
		calculationBuilder: calculationBuilder,
	}

	return &out
}

// Execute executes a calculation
func (app *application) Execute(affiliateHash hash.Hash, genesisHash hash.Hash, amount uint) (calculations.Calculations, error) {
	// retrieve the genesis instance:
	genesisIns, err := app.genesisApp.Retrieve(genesisHash)
	if err != nil {
		return nil, err
	}

	// retrieve the distribution's instance:
	distributions := genesisIns.Parent().Distributions()

	// create the calculation's list
	calculationList := []calculations.Calculation{}

	// create the founder's calculations:
	amountUsed := uint(0)
	totalPower := distributions.Power()
	founderPercent := genesisIns.ShareHoldersPercent()
	foundersAmount := math.Floor(float64(amount) * founderPercent)
	foundersList := distributions.List()
	for _, oneFounder := range foundersList {
		percent := float64(totalPower / oneFounder.Power())
		founderAmount := math.Ceil(foundersAmount * percent)
		pubKey := oneFounder.PubKey()
		calculation, err := app.calculationBuilder.Create().WithPubKey(pubKey).WithAmount(uint(founderAmount)).Now()
		if err != nil {
			return nil, err
		}

		calculationList = append(calculationList, calculation)
		amountUsed += uint(founderAmount)
	}

	// retrieve the line of the affiliate:
	line, err := app.lineApp.Retrieve(affiliateHash)
	if err != nil {
		return nil, err
	}

	// fetch the level from the genesis:
	level := genesisIns.Level()
	remainingAmount := amount - amountUsed
	calculations, err := app.calculateByLevel(line, level, remainingAmount)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithList(calculations).Now()
}

func (app *application) calculateByLevel(line lines.Line, level levels.Level, totalAmount uint) ([]calculations.Calculation, error) {
	levelPower := level.Power()
	totalPower := level.TotalPower()
	percent := totalPower / levelPower
	amount := uint(math.Ceil(float64(totalAmount * percent)))
	pubKey := line.PubKey()
	calculation, err := app.calculationBuilder.Create().WithPubKey(pubKey).WithAmount(amount).Now()
	if err != nil {
		return nil, err
	}

	list := []calculations.Calculation{
		calculation,
	}

	if line.HasChild() && level.HasParent() {
		child := line.Child()
		parentLine := level.Parent()
		subAmount := totalAmount - amount
		subCalculations, err := app.calculateByLevel(child, parentLine, subAmount)
		if err != nil {
			return nil, err
		}

		list = append(list, subCalculations...)
	}

	return list, nil
}
