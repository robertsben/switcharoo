package main 

import (
	"fmt"
)

type DataError struct {
	Err string
	Data []byte
}

func (e DataError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}