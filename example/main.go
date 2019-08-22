package main

import (
	"context"
	"fmt"

	"github.com/kwtucker/datalayer/example/query"
	_ "github.com/kwtucker/datalayer/example/services/exampleService/init" // Only calls the init func in the scte service
)

// Test would be the object to pass to the query to be filled.
type Test struct {
	Text string `json:"text,omitempty"`
}

func main() {
	// How you would start a new query to the datalayer.
	q := query.NewExampleQuery(context.TODO())
	obj := &Test{} // Needs to be a pointer so you can reference obj after the Run func returns without a error
	err := q.Run(obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj.Text)
}
