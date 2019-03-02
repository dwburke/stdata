package util

import (
	log "github.com/sirupsen/logrus"
)

func PanicIf(err error) {
	if err == nil {
		return
	}

	log.Panic(err)
}

func FatalIf(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)
}
