package logger

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	LogLevels = iota * 2
	INFO
	DEBUG
	DIEBUG
	WARNING
	ERROR
)

var (
	logLevelsMap = map[int]string{
		INFO:    "INFO",
		DEBUG:   "DEBUG",
		DIEBUG:  "DIEBUG",
		WARNING: "WARNING",
		ERROR:   "ERROR",
	}
)

func Info(message string) {
	generateLog(INFO, message, nil)
}

func Infof(message string, data ...interface{}) {
	generateLog(INFO, message, data...)
}

func Debug(message string) {
	generateLog(DEBUG, message, nil)
}

func Debugf(message string, data ...interface{}) {
	generateLog(DEBUG, message, data...)
}

func Diebug(message string) {
	generateLog(DIEBUG, message, nil)
}

func Diebugf(message string, data ...interface{}) {
	generateLog(DIEBUG, message, data...)
}

func Warning(message string) {
	generateLog(WARNING, message, nil)
}

func Warningf(message string, data ...interface{}) {
	generateLog(WARNING, message, data...)
}

func Error(message string) {
	generateLog(ERROR, message, nil)
}

func Errorf(message string, data ...interface{}) {
	generateLog(ERROR, message, data...)
}

func generateLog(level int, message string, data ...interface{}) {
	var (
		msgStrWithData string
		logPrint       = fmt.Sprintf(`\e%s[%s]\e[0m message: %s`, colorByLabel(level), logLevelsMap[level], message)
	)

	for index, v := range data {
		if v != nil {
			msgStrWithData = fmt.Sprintf("%s | %v", msgStrWithData, v)
			if index == len(data)-1 {
				msgStrWithData = msgStrWithData + " |"
			}
		}
	}

	logPrint = logPrint + msgStrWithData

	out, err := exec.Command("echo", "-e", logPrint).Output()

	if err != nil {
		Error(err.Error())
	}

	if level == DIEBUG {
		log.Fatal(string(out))
	}

	log.Print(string(out))
}

func colorByLabel(level int) string {
	switch level {
	case INFO:
		return "[32m"
	case DEBUG:
		return "[36m"
	case DIEBUG:
		return "[36m"
	case WARNING:
		return "[33m"
	case ERROR:
		return "[31m"
	default:
		return "[37m"
	}
}
