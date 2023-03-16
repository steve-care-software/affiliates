package levels

import "github.com/steve-care-software/libs/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a level builder
type Builder interface {
	Create() Builder
	WithPower(power uint) Builder
	WithParent(parent Level) Builder
	Now() (Level, error)
}

// Level represents a level
type Level interface {
	Hash() hash.Hash
	Power() uint
	TotalPower() uint
	HasParent() bool
	Parent() Level
}
