package d2s

import (
	"log"
	"os"
	"testing"
)

/**
* Parse NokkaSorc.d2s from examples and verify the chars name
**/
func TestParse(t *testing.T) {
	path := "examples/NokkaSorc.d2s"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening .d2s file", err)
	}

	defer file.Close()

	char, err := Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	if char.Header.Name.String() != "NokkaSorc" {
		t.Errorf("Expected char name to be %v. Got %v", "NokkaSorc", char.Header.Name)
	}
}
