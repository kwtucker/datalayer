package exampleQuery

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/kwtucker/datalayer"
)

const registryIdentifier = "exampleService"

type ExampleQuery struct {
	ctx context.Context
}

func (r *ExampleQuery) Run(object interface{}) error {
	if object != nil {
		// Make sure a pointer object is sent in.
		if reflect.ValueOf(object).Kind() != reflect.Ptr {
			return fmt.Errorf("must be a pointer of a type")
		}
	}

	// Lookup is a way to reference a datastore in the datalayer that was registered as a import in the main func.
	dataStore := datalayer.R.Lookup(registryIdentifier)
	if dataStore != nil {
		byt, err := json.Marshal(r)
		if err != nil {
			return err
		}

		err = dataStore.Value.Query(byt, object)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("ERROR")
}
