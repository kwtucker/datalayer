package datalayer

// R is for the top level object to store registered data stores.
var R = &Registry{}

// DataStorer is the interface for connecting to a third party data base or api
type DataStorer interface {
	Query(request []byte, object interface{}) error
}
