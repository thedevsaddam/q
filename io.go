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
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// prepareStdOutput prepares the standard output to show the result of the query
func prepareStdOutput() {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	if err := enc.Encode(output); err != nil {
		fmt.Println(err)
		return
	}
}
