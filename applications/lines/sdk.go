package lines

import (
	"github.com/steve-care-software/affiliates/domain/lines"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents the line application
type Application interface {
	Retrieve(hash hash.Hash) (lines.Line, error)
}
