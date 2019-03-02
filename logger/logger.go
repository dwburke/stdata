package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"github.com/spf13/viper"
)

var Logfile *os.File

func init() {
	viper.SetDefault("logging.stdout.level", "info")
	viper.SetDefault("logging.file.maxsize", 5)
	viper.SetDefault("logging.file.maxbackups", 7)
	viper.SetDefault("logging.file.maxage", 7)
	viper.SetDefault("logging.file.enabled", false)
}

func interpretLogLevel(level string) (log_level log.Level) {

	switch level {
	case "trace":
		log_level = log.TraceLevel
	case "debug":
		log_level = log.DebugLevel
	case "info":
		log_level = log.InfoLevel
	case "warn":
		log_level = log.WarnLevel
	case "error":
		log_level = log.ErrorLevel
	case "fatal":
		log_level = log.FatalLevel
	case "panic":
		log_level = log.PanicLevel
	default:
		log_level = log.DebugLevel
	}

	return
}

func InitLogging() {
	log.SetOutput(os.Stdout)
	log.SetLevel(interpretLogLevel(viper.GetString("logging.stdout.level")))

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	setupFileLogging()
}

func setupFileLogging() {
	if !viper.GetBool("logging.file.enabled") {
		return
	}

	filename := viper.GetString("logging.file.name")

	log.Infof("logrus: opening logfile: %s", filename)

	if filename == "" {
		panic("logging.file.enabled == true && logging.file.name == \"\"\n")
	}

	var err error
	Logfile, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   filename,
		MaxSize:    viper.GetInt("logging.file.maxsize"),
		MaxBackups: viper.GetInt("logging.file.maxbackups"),
		MaxAge:     viper.GetInt("logging.file.maxage"),
		Level:      interpretLogLevel(viper.GetString("logging.file.level")),
		Formatter: &log.TextFormatter{
			FullTimestamp: true,
		},
	})

	log.AddHook(rotateFileHook)

	log.Infof("logrus: logging to file: %s", filename)
}

func Cleanup() {
	log.Info("logger: Cleanup()")
	if Logfile != nil {
		Logfile.Close()
	}
}
