package gonf

import "errors"

var (
	// ErrFileNotFound is returned when there is no file with path 
	ErrFileNotFound = errors.New("No such file")

	// ErrReadingFile is returned when a invalid file has read
	ErrReadingFile = errors.New("Error trying read the file")

	//
	ErrParserFile = errors.New("Parse failed, bad format file")

	//
	ErrInvalidStruct = errors.New("Invalid Struct reference")
)