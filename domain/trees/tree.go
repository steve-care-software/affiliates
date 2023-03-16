package trees

import "github.com/steve-care-software/libs/cryptography/hash"

type tree struct {
	hash     hash.Hash
	pubKey   hash.Hash
	children Trees
}

func createTree(
	hash hash.Hash,
	pubKey hash.Hash,
) Tree {
	return createTreeInternally(hash, pubKey, nil)
}

func createTreeWithChildren(
	hash hash.Hash,
	pubKey hash.Hash,
	children Trees,
) Tree {
	return createTreeInternally(hash, pubKey, children)
}

func createTreeInternally(
	hash hash.Hash,
	pubKey hash.Hash,
	children Trees,
) Tree {
	out := tree{
		hash:     hash,
		pubKey:   pubKey,
		children: children,
	}

	return &out
}

// Hash returns the hash
func (obj *tree) Hash() hash.Hash {
	return obj.hash

}

// PubKey returns the pubKey
func (obj *tree) PubKey() hash.Hash {
	return obj.pubKey

}

// HasChildren returns true if there is a children, false otherwise
func (obj *tree) HasChildren() bool {
	return obj.children != nil
}

// Children returns the children, if any
func (obj *tree) Children() Trees {
	return obj.children
}
