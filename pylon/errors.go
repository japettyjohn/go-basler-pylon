package pylon

import "fmt"

func newError(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

/* EOF */
