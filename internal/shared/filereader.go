package shared

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return ioutil.ReadAll(file)
}
