package main

import (
	"strings"
)

func Handler(text string) (command string, key string, value string) {
	command, key, value = "", "", ""
	arr := strings.Fields(text)
	switch len(arr) {
	case 1:
		command = arr[0]

	case 2:
		command = arr[0]
		key = arr[1]

	case 3:
		command = arr[0]
		key = arr[1]
		value = arr[2]
	}

	return command, key, value
}
