package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	output interface{}
)

// readFromStdin reads the data to feed on
func readFromStdin() string {
	var data []byte
	var err error
	stat, _ := os.Stdin.Stat() // for checking if data provided from stdin

	if stat.Size() > 0 { // if data provided on stdin
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	}
	return string(data)
}

// prepareStdOutput prepares the standard output to show the result of the query
func prepareStdOutput() {
	enc := json.NewEncoder(os.Stdout)
	if pretty {
		enc.SetIndent("", "\t")
	}
	if err := enc.Encode(output); err != nil {
		fmt.Println(err)
		return
	}
}
