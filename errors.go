package wtf

import "errors"

func Wrap(message string, err error) error {
	return errors.New(message + "::" + err.Error())
}
