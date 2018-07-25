package gonf

import (
	"io"
	"io/ioutil"
	"os"
	"log"
)

// Load the config file
func Load(configFile string, struc interface{}) error{

	var err error
	var input = io.ReadCloser(os.Stdin)
	
	log.Println(configFile)

	//Trying open file
	if input, err = os.Open(configFile); err != nil {
		log.Fatalln(TAG, err)
		return ErrFileNotFound
	}

	// Read the config file
	contentBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(TAG, err)
		return ErrReadingFile
	}

	//Get Parser
	parser := GetParser()

	if err := parser.Parse(contentBytes, struc); err != nil {
		log.Fatalln(TAG, "Could not parse file:", err)
		return ErrParserFile
	}

	return nil
}