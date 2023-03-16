package levels

import "github.com/steve-care-software/libs/cryptography/hash"

type level struct {
	hash   hash.Hash
	power  uint
	parent Level
}

func createLevel(
	hash hash.Hash,
	power uint,
) Level {
	return createLevelInternally(hash, power, nil)
}

func createLevelWithParent(
	hash hash.Hash,
	power uint,
	parent Level,
) Level {
	return createLevelInternally(hash, power, parent)
}

func createLevelInternally(
	hash hash.Hash,
	power uint,
	parent Level,
) Level {
	out := level{
		hash:   hash,
		power:  power,
		parent: parent,
	}

	return &out
}

// Hash returns the hash
func (obj *level) Hash() hash.Hash {
	return obj.hash
}

// Power returns the power
func (obj *level) Power() uint {
	return obj.power
}

// TotalPower returns the total power
func (obj *level) TotalPower() uint {
	power := uint(0)
	if obj.HasParent() {
		power += obj.Parent().TotalPower()
	}

	return power
}

// HasParent returns true if there is a parent, false otherwise
func (obj *level) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *level) Parent() Level {
	return obj.parent
}
