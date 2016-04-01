package yai

import (
	log "github.com/cihub/seelog"
	"os"
	"fmt"
)

func InitLog() log.LoggerInterface {
	logfile := "seelog.xml"

	minLogLevel := os.Getenv("MIN_LOG_LEVEL")
	fmt.Printf("minLogLevel:'%v'\n", minLogLevel)
	if minLogLevel == "TRACE" {
		logfile = "seelog_trace.xml"
	}
	if minLogLevel == "INFO" {
		logfile = "seelog_info.xml"
	}
	if minLogLevel == "ERROR" {
		logfile = "seelog_error.xml"
	}

	if _, err := os.Stat(logfile); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("No log file %v", logfile))
		return log.Default
	} else {
		logger, _ := log.LoggerFromConfigAsFile(logfile)
		log.ReplaceLogger(logger)

		return logger
	}
}
