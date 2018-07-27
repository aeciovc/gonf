package gonf

import (
	"log"
)

// Load the config file
func Load(filePath string, struc interface{}) error{

	//Build OS reader
	reader := new(ReaderOS)
	
	//Reading File
	contentBytes, err := reader.Read(filePath)
	if err != nil{
		log.Println(TAG, err)
		return err
	}

	//Get Parser
	parser := GetParser()

	//Parsing
	if err := parser.Parse(contentBytes, struc); err != nil {
		log.Println(TAG, err)
		return ErrParserFile
	}

	return nil
}