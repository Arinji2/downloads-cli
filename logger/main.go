package logger

import (
	cli_log "github.com/Arinji2/cli-log"
)

var GLogger *cli_log.Logger

func InitLogger() {
	logger := cli_log.NewLogger("log.txt", 0, "Downloads CLI")
	GLogger = logger
}
