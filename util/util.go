package util

import "log"

func ErrorHandle(err error) {
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
}
