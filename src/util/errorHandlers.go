package util

import (
	"fmt"
)

func HandleWarning(e error) {
	if e != nil {
		fmt.Println("Warning: " + e.Error())
	}
}

func HandleFatal(e error) {
	if e != nil {
		fmt.Println("Fatal error encountered: " + e.Error())
		panic(e)
	}
}
