package gonf

import (
	"os"
	"testing"
)

type Other struct {
	Count int
}

type Config struct{
	Name string
	Other Other
}

func TestMain(m *testing.M) {
	result := m.Run()
	os.Exit(result)
}