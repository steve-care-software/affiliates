package genesis

import (
	"github.com/steve-care-software/affiliates/domain/genesis/levels"
	fungible_unit_genesis "github.com/steve-care-software/fungible-unit-genesis/domain"
	genesis "github.com/steve-care-software/fungible-unit-genesis/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithParent(parent fungible_unit_genesis.Genesis) Builder
	WithShareHoldersPercent(shareHoldersPercent float64) Builder
	WithLevel(level levels.Level) Builder
	Now() (Genesis, error)
}

// Genesis represents a genesis
type Genesis interface {
	Hash() hash.Hash
	Parent() fungible_unit_genesis.Genesis
	ShareHoldersPercent() float64
	Level() levels.Level
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	WithKind(kind uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a genesis repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Genesis, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	WithKind(kind uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents a genesis service
type Service interface {
	Insert(genesis genesis.Genesis) error
	Delete(hash hash.Hash) error
}
