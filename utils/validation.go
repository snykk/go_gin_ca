package utils

import (
	"fmt"

	"github.com/snykk/go_gin_ca/constants"
)

func IsPriorityValid(priority string) error {
	if !isArrayContains(constants.ListOfPriority, priority) {
		var option string
		for index, priority := range constants.ListOfPriority {
			option += priority
			if index != len(constants.ListOfPriority)-1 {
				option += ", "
			}
		}

		return fmt.Errorf("priority must be one of [%s]", option)
	}

	return nil
}

func isArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}
