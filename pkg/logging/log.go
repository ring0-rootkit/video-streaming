package logging

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strings"
)

var logLevel int8

type Log struct {
	logger log.Logger
}

func loadLogLvl() {
	lvl := strings.ToLower(os.Getenv("log_level"))
	switch lvl {
	case "info":
		logLevel = 1
		log.Default().Println("log level set to info")
	case "warn":
		logLevel = 2
		log.Default().Println("log level set to warn")
	case "error":
		logLevel = 3
		log.Default().Println("log level set to error")
	case "":
		log.Default().Println("log_level was not spcified in .env file, defaulting to 'INFO'")
	default:
		log.Default().Println("please check your log_level in .env file,",
			"allowed values 'INFO' 'WARN' and 'ERROR',",
			"provided", lvl,
			"defaulting to INFO")
		logLevel = 1
	}

}

func New(prefix string) *Log {
	if logLevel == 0 {
		loadLogLvl()
	}
	return &Log{*log.New(os.Stdout, prefix+" ", log.Default().Flags())}
}

func (l *Log) Error(msg string) {
	if logLevel <= 3 {
		l.logger.Print("[ERROR]", msg)
	}
}

func (l *Log) Warn(msg string) {
	if logLevel <= 2 {
		l.logger.Print("[WARN]", msg)
	}
}

func (l *Log) Info(msg string) {
	if logLevel <= 1 {
		l.logger.Print("[INFO]", msg)
	}
}
