package utils

import "log"

// check error
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
