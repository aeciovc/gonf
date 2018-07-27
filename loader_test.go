package gonf

import (
	"testing"
	test "github.com/aeciovc/gonf/test"
)

func TestLoad(t *testing.T) {

	config := &Config{}
	configExpected := &Config{Name:"myconfig", Other: other{Count:80}}

	err := Load("./etc/config.json", config)

	test.Equals(t, nil, err)
	test.Equals(t, configExpected.Name, config.Name)
	test.Equals(t, configExpected.Other.Count, config.Other.Count)
}

func TestLoad_FileNotFound(t *testing.T) {

	config := &Config{}

	err := Load("./etc/any_file.json", config)

	test.Equals(t, ErrFileNotFound, err)
	test.Equals(t, "", config.Name)
	test.Equals(t, 0, config.Other.Count)
}