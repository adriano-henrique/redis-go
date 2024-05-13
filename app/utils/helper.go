package utils

import (
	"errors"
	"strings"
)

func GetMasterHostAddress(replicaOf string, args []string) (string, error) {
	if len(args) >= 1 {
		return strings.Join([]string{replicaOf, args[0]}, ":"), nil
	}

	splitString := strings.Split(replicaOf, " ")
	if len(splitString) == 2 {
		return strings.Join(splitString, ":"), nil
	}

	return "", errors.New("unable to parse master host")
}
