package trees

import "github.com/steve-care-software/libs/cryptography/hash"

type trees struct {
	hash hash.Hash
	list []Tree
}

func createTrees(
	hash hash.Hash,
	list []Tree,
) Trees {
	out := trees{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *trees) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *trees) List() []Tree {
	return obj.list
}
