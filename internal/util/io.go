package util

import (
	"io/ioutil"
	"log"
)

func WriteStringToFile(filename string, s string) error {
	if err := ioutil.WriteFile(filename, []byte(s), 0644); err != nil {
		return err
	}
	log.Printf("[WRITE] %s", filename)
	return nil
}
