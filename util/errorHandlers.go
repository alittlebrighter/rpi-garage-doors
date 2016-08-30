package util

import "log"

func HandleWarning(e error) {
	if e != nil {
		log.Println("Warning: " + e.Error())
	}
}

func HandleFatal(e error) {
	if e != nil {
		log.Println("Fatal error encountered: " + e.Error())
		panic(e)
	}
}
