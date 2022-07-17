package wtf

import (
	"errors"
	"strings"
)

var (
	ErrRowAlreadyExists = errors.New("duplicate key, row already exists")
	ErrNoResults        = errors.New("no results found")
)

var (
	cassandraErrors = map[string]error{
		"duplicate": ErrRowAlreadyExists,
		"not exist": ErrNoResults,
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
