package genesis

import (
	"github.com/steve-care-software/affiliates/domain/founders"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents the founder's application
type Application interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (founders.Founder, error)
	RetrieveList(hashes []hash.Hash) (founders.Founders, error)
}
