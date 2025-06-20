package logging

/*
	Currently just placeholder functions to ease adding a proper logger later on
*/

import (
	"log"
)

func LogDebug(args ...any) {
	log.Println(args...)
}

func LogVerbose(args ...any) {
	log.Println(args...)
}

func LogWarning(args ...any) {
	log.Println(args...)
}

func LogError(args ...any) {
	log.Println(args...)
}

func LogFatal(args ...any) {
	log.Println(args...)
}
