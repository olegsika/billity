package utils

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func InListString(entry string, list []string) bool {
	for _, listEntry := range list {
		if entry == listEntry {
			return true
		}
	}
	return false
}
