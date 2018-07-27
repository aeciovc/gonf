package gonf

import (
	"os"
	"io/ioutil"
	"log"
)

type ReaderOS struct {
}

func (r *ReaderOS) Read(path string) ([]byte, error){

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println(TAG, err)
		return nil, ErrFileNotFound
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(TAG, err)
		return nil, ErrReadingFile
	}

	return data, nil
}