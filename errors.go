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
	// ErrParserFile is returned when JSON has a bad format content
	ErrParserFile = errors.New("Parse failed, bad format file")

	// ErrInvalidStruct is returned when a nil struct is passed as param to load configs
	ErrInvalidStruct = errors.New("Invalid Struct reference")
)