package gonf

import (
)

// Parser must implement Parse
type Parser interface {
	Parse([]byte, interface{}) error
}

func GetParser() Parser{
	return &JSONParser{}
}