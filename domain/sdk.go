package domain

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents an affiliate
type Builder interface {
	Create() Builder
	WithPubKey(pubKey hash.Hash) Builder
	WithParent(parent hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Affiliate, error)
}

// Affiliate represents an affiliate
type Affiliate interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	CreatedOn() time.Time
	HasParent() bool
	Parent() hash.Hash
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	WithKind(kind uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents an afiliate repository
type Repository interface {
	List() ([]hash.Hash, error)
	ListByParent(parent hash.Hash) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Affiliate, error)
	RetrieveList(hashes []hash.Hash) ([]Affiliate, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	WithKind(kind uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents an affiliate service
type Service interface {
	Insert(afiliate Affiliate) error
}
