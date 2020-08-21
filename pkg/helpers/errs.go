package helpers

import (
	"errors"
	"strings"
)

func JoinErrs(errs []error) error {
	n := len(errs)
	if n == 0 {
		return nil
	}

	slice := make([]string, len(errs))

	for i, err := range errs {
		slice[i] = err.Error()
	}

	return errors.New(strings.Join(slice, "; "))
}
