package gonf

import (
	"testing"

	"io/ioutil"
	_"log"
	test "github.com/aeciovc/gonf/test"
)

type other struct {
	Count int
}

type config struct{
	Name string
	Other other
}

func TestParser(t *testing.T) {

	//Inputs
	conf := &config{}
	fileBytes := readFile("./etc/config.json")
	fileBytesBadFormated := readFile("./etc/config_bad.json")

	type inputs struct {
		File []byte
		Struc interface{}
	}

	//Outputs
	confLoaded := &config{Name:"myconfig", Other: other{Count:80}}

	type outputs struct{
		Struc interface{}
		Err   error
	}

	tests := []struct {
		name      string
		inputs    inputs
		outputs   outputs
		
	}{
		{"TestParser - Invalid bytes file", inputs{File:nil, Struc:conf}, outputs{Struc:conf, Err:ErrParserFile}},
		{"TestParser - Invalid struct reference", inputs{File:fileBytes, Struc:nil}, outputs{Struc:nil, Err:ErrInvalidStruct}},
		{"TestParser - JSON bad formated", inputs{File:fileBytesBadFormated, Struc:conf}, outputs{Struc:conf, Err:ErrParserFile}},
		{"TestParser - Success", inputs{File:fileBytes, Struc:conf}, outputs{Struc:confLoaded, Err:nil}},
	}

	for _, tt := range tests {

		// Run tests
		t.Run(tt.name, func(t *testing.T) {
			
			//Get JSON Parser (Default)
			parser := GetParser()
			err := parser.Parse(tt.inputs.File, tt.inputs.Struc)

			//get result
			resultConfig := interfaceToStructConfig(tt.inputs.Struc)

			//get expected result
			expectedConfig := interfaceToStructConfig(tt.outputs.Struc)
			
			//Assert error
			test.Equals(t, tt.outputs.Err, err)

			//Assert Config loaded
			if tt.outputs.Err == nil{
				test.Equals(t, expectedConfig.Name, resultConfig.Name)
				test.Equals(t, expectedConfig.Other.Count, resultConfig.Other.Count)
			}
		})
	}
}

func interfaceToStructConfig(i interface{}) *config{
	var result *config
	if i != nil{
		result = i.(*config)
	}

	return result
}

func readFile(path string) []byte{
	contentBytes, err := ioutil.ReadFile(path)
	check(err)

	return contentBytes
}

//Check error
func check(e error) {
    if e != nil {
        panic(e)
    }
}