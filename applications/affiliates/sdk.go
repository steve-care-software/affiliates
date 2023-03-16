package affiliates

import (
	affiliates "github.com/steve-care-software/affiliates/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents an affiliate application
type Application interface {
	List() ([]hash.Hash, error)
	ListByParent(parent hash.Hash) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (affiliates.Affiliate, error)
	RetrieveList(hashes []hash.Hash) ([]affiliates.Affiliate, error)
	Insert(affiliate affiliates.Affiliate) error
}
