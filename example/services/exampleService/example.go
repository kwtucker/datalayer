package exampleService

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/kwtucker/datalayer"
)

// ExampleService ...
type ExampleService struct {
	ctx  context.Context // NOTE: required
	Text string
	// And what ever else you want to add here.
}

// Register the sctevs package.
func Register(reg *datalayer.Registry) {
	reg.Register(&datalayer.Object{
		Name:    "exampleService",
		Value:   &ExampleService{},
		Enabled: true,
	})
}

// Query implements the DataStorer interface and is the entry point to hte sctevs service.
func (s *ExampleService) Query(byt []byte, object interface{}) error {
	var err error
	exp := &ExampleService{}

	if object != nil {
		// Make sure a pointer object is sent in.
		if reflect.ValueOf(object).Kind() != reflect.Ptr {
			return fmt.Errorf("must be a pointer of a type")
		}
	}

	err = json.Unmarshal(byt, exp)
	if err != nil {
		return err
	}

	// NOTE: Do all un exported logic here. Api calls etc.

	exp.Text = "ExampleService Object Filled."

	// Sets the filled scte to the requester interface.
	setByt, _ := json.Marshal(exp)

	return json.Unmarshal(setByt, object)
}
