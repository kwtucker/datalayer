package datalayer

import "context"

var (
	// DataStores List of datastores that are registered.
	DataStores = make(map[string]DataStorer)

	// DataLayerRegistry is for the top level object to store registered data stores.
	DataLayerRegistry = &Registry{}
)

type Requester interface {
	// GetContext Gets the context of the request
	GetContext() context.Context

	// SetContext Sets the context to the object that implements it.
	SetContext(ctx context.Context)

	// GetInfo gives the flexiblity to unmarshal the byte slice to an object
	GetInfo() ([]byte, error)

	// SetInfo takes a byte slice to unmarshal and update the requester info.
	SetInfo([]byte) error
}

// DataStorer is the interface for connecting to a third party data base or api
type DataStorer interface {
	Query(requester Requester, object interface{}) error
}
