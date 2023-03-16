package levels

import "github.com/steve-care-software/libs/cryptography/hash"

// Level represents a level
type Level interface {
	Hash() hash.Hash
	Power() uint
	TotalPower() uint
	HasParent() bool
	Parent() Level
}
