package utils

import (
	"fmt"
	"strconv"
)

func ParseString(s string, dest interface{}) error {
	d, ok := dest.(*string)
	if !ok {
		return fmt.Errorf("wrong type for ParseString: %T", dest)
	}
	// assume error = false
	*d = s
	return nil
}

func ParseBool(s string, dest interface{}) error {
	d, ok := dest.(*bool)
	if !ok {
		return fmt.Errorf("wrong type for ParseBool: %T", dest)
	}
	// assume error = false
	*d, _ = strconv.ParseBool(s)
	return nil
}

func ParseInt32(s string, dest interface{}) error {
	d, ok := dest.(*int32)
	if !ok {
		return fmt.Errorf("wrong type for ParseInt: %T", dest)
	}
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	*d = int32(n)
	return nil
}

func ParseInt(s string, dest interface{}) error {
	d, ok := dest.(*int)
	if !ok {
		return fmt.Errorf("wrong type for ParseInt: %T", dest)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*d = n
	return nil
}
