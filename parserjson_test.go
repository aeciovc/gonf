package gonf

import (
	"testing"
	"io/ioutil"
	_"log"
	test "github.com/aeciovc/gonf/test"
)

func TestParse(t *testing.T) {

	//Inputs
	config := &Config{}
	fileBytes := readFile("./etc/config.json")

	//Output
	configExpected := &Config{Name:"myconfig", Other: Other{Count:80}}
	
	//Get JSON Parser (Default)
	parser := GetParser()
	err := parser.Parse(fileBytes, config)

	test.Equals(t, err, nil)
	test.Equals(t, configExpected.Name, config.Name)
	test.Equals(t, configExpected.Other.Count, config.Other.Count)
}
	

func TestParse_InvalidStructReference(t *testing.T) {

	//Inputs
	fileBytes := readFile("./etc/config.json")

	//Get JSON Parser (Default)
	parser := GetParser()
	err := parser.Parse(fileBytes, nil)

	test.Equals(t, err, ErrInvalidStruct)
}

func TestParse_BadFormatFile(t *testing.T) {

	//Inputs
	config := &Config{}
	fileBytes := readFile("./etc/config_bad.json")

	//Get JSON Parser (Default)
	parser := GetParser()
	err := parser.Parse(fileBytes, config)

	test.Equals(t, err, ErrParserFile)
}


/*************************** Private ***************************/
func interfaceToStructConfig(i interface{}) *Config{
	var result *Config
	if i != nil{
		result = i.(*Config)
	}

	return result
}

func readFile(path string) []byte{
	contentBytes, err := ioutil.ReadFile(path)
	check(err)

	return contentBytes
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}