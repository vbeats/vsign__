package cleanup

import (
	"os"
	"vsign/logger"
)

func ExitService(s os.Signal) {
	logger.Info("exit signal: %s received, clean up working...", s.String())

	logger.Info("clea up done, exit...")
}
