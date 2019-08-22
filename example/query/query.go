package query

import (
	"context"

	"github.com/kwtucker/datalayer/example/query/exampleQuery"
)

// NewExampleQuery ...
func NewExampleQuery(ctx context.Context) *exampleQuery.ExampleQuery {
	s := &exampleQuery.ExampleQuery{}
	return s
}
