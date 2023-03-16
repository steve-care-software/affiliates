package lines

import "github.com/steve-care-software/libs/cryptography/hash"

type line struct {
	hash   hash.Hash
	pubKey hash.Hash
	child  Line
}

func createLine(
	hash hash.Hash,
	pubKey hash.Hash,
) Line {
	return createLineInternally(hash, pubKey, nil)
}

func createLineWithChild(
	hash hash.Hash,
	pubKey hash.Hash,
	child Line,
) Line {
	return createLineInternally(hash, pubKey, child)
}

func createLineInternally(
	hash hash.Hash,
	pubKey hash.Hash,
	child Line,
) Line {
	out := line{
		hash:   hash,
		pubKey: pubKey,
		child:  child,
	}

	return &out
}

// Hash returns the hash
func (obj *line) Hash() hash.Hash {
	return obj.hash
}

// PubKey returns the pubKey
func (obj *line) PubKey() hash.Hash {
	return obj.pubKey
}

// HasChild returns true if there is a child, false otherwise
func (obj *line) HasChild() bool {
	return obj.child != nil
}

// Child returns the child, if any
func (obj *line) Child() Line {
	return obj.child
}
