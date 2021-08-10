package dstore

import (
	"os"

	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var traceEnabled = os.Getenv("TRACE") == "true"
var zlog *zap.Logger

func init() {
	logging.Register("github.com/streamingfast/dstore", &zlog)
}
