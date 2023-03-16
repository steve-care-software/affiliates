package trees

import "github.com/steve-care-software/libs/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewTreeBuilder creates a new tree builder
func NewTreeBuilder() TreeBuilder {
	hashAdapter := hash.NewAdapter()
	return createTreeBuilder(hashAdapter)
}

// Builder represents the trees builder
type Builder interface {
	Create() Builder
	WithList(list []Tree) Builder
	Now() (Trees, error)
}

// Trees represents trees
type Trees interface {
	Hash() hash.Hash
	List() []Tree
}

// TreeBuilder represents the tree builder
type TreeBuilder interface {
	Create() TreeBuilder
	WithPubKey(pubKey hash.Hash) TreeBuilder
	WithChildren(children Trees) TreeBuilder
	Now() (Tree, error)
}

// Tree represents a referral tree
type Tree interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	HasChildren() bool
	Children() Trees
}
