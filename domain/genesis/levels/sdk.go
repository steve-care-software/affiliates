package levels

import "github.com/steve-care-software/libs/cryptography/hash"

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
