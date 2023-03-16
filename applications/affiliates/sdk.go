package affiliates

import (
	affiliates "github.com/steve-care-software/affiliates/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents an affiliate application
type Application interface {
	Insert(afiliate affiliates.Affiliate) error
	Retrieve(hash hash.Hash) (affiliates.Affiliate, error)
}
