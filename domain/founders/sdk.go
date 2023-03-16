package founders

import "github.com/steve-care-software/libs/cryptography/hash"

// Builder represents the founder builder
type Builder interface {
	Create() Builder
	WithPubKey(pubKey hash.Hash) Builder
	WithPower(power uint) Builder
	Now() (Founder, error)
}

// Founder represents the affiliate founders
type Founder interface {
	Hash() hash.Hash
	PubKey() hash.Hash
	Power() uint
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	WithKind(kind uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a founder repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Founder, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	WithKind(kind uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents a founder service
type Service interface {
	Insert(founder Founder) error
}
