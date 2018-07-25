package gonf

import (
	"encoding/json"
	"log"
)

type JSONParser struct {
}

func (jsonParser *JSONParser) Parse(file []byte, struc interface{}) error{
	if len(file) == 0{
		log.Println(TAG, ErrParserFile)
		return ErrParserFile
	}

	if struc == nil{
		log.Println(TAG, ErrInvalidStruct)
		return ErrInvalidStruct
	}

	err := json.Unmarshal(file, struc)
	if err != nil{
		log.Println(TAG, "Unmarshal Error: ", err)
		return ErrParserFile
	}

	return nil
}