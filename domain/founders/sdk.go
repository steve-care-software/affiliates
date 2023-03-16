package founders

import "github.com/steve-care-software/libs/cryptography/hash"

// Builder represents the founder's builder
type Builder interface {
	Create() Builder
	WithList(list []Founder) Builder
	Now() (Founders, error)
}

// Founders represents the founders
type Founders interface {
	Hash() hash.Hash
	List() []Founder
}

// FounderBuilder represents the founder builder
type FounderBuilder interface {
	Create() FounderBuilder
	WithPubKey(pubKey hash.Hash) FounderBuilder
	WithPower(power uint) FounderBuilder
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
	InsertList(founders Founders) error
}
