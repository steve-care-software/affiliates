package trees

import "github.com/steve-care-software/libs/cryptography/hash"

// Trees represents trees
type Trees interface {
	Hash() hash.Hash
	List() []Tree
}

// Tree represents a referral tree
type Tree interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	HasChildren() bool
	Children() []Trees
}
