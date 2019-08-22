package init

import (
	"github.com/kwtucker/datalayer"
	"github.com/kwtucker/datalayer/example/services/exampleService"
)

func init() {
	exampleService.Register(datalayer.R)
}
