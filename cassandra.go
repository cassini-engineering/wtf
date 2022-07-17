package wtf

import (
	"errors"
	"strings"
)

var (
	errRowAlreadyExists = errors.New("duplicate key, row already exists")
	errNoResults        = errors.New("no results found")
)

var (
	cassandraErrors = map[string]error{
		"duplicate": errRowAlreadyExists,
		"not exist": errNoResults,
	}
)

func IsAnyCassandraError(err error) (bool, error) {
	for k, v := range cassandraErrors {
		if strings.Contains(err.Error(), k) {
			return true, Wrap(v.Error(), err)
		}
	}
	return false, nil
}
