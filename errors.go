package gonf

import "errors"


//Readers Errors
var (
	// ErrFileNotFound is returned when there is no file with path 
	ErrFileNotFound = errors.New("No such file in path")

	// ErrReadingFile is returned when a invalid file has read
	ErrReadingFile = errors.New("Error trying read the file")
)

//Parser Errrors
var (
	//
	ErrParserFile = errors.New("Parse failed, bad format file")

	//
	ErrInvalidStruct = errors.New("Invalid Struct reference")
)