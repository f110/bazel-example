package main

import (
	"encoding/json"
	"errors"
	"os"
)

type confFile []string

func testSuite(args []string) error {
	if len(args) == 0 {
		return errors.New("requires input file")
	}
	f, err := os.Open(args[0])
	if err != nil {
		return err
	}
	var conf confFile
	if err := json.NewDecoder(f).Decode(&conf); err != nil {
		return err
	}

	if len(conf) == 0 {
		return errors.New("empty file")
	}

	return nil
}

func main() {
	if err := testSuite(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
