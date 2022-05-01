package utils

import "log"

// CheckErr The function return panic if error != nil
func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// InListString THe function check is entry exists on list
func InListString(entry string, list []string) bool {
	for _, listEntry := range list {
		if entry == listEntry {
			return true
		}
	}
	return false
}
