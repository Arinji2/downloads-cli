package logger

import (
	cli_log "github.com/Arinji2/cli-log"
)

var GLogger *cli_log.Logger

func InitLogger(logFile string) {
	logger := cli_log.NewLogger(logFile, 0, "Downloads CLI")
	GLogger = logger
}
