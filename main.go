package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

var (
	jq *gojsonq.JSONQ
)

func main() {

	if versionFlagProvided() {
		fmt.Println(logo)
		return
	}

	checkDataType() // Checking for the datatype feeded, i.e: json/xml/yaml/csv

	checkFlags() // Checking for the flags provided by the user

	checkCommands() // checking for the commands provided by the user to query over

	prepareStdOutput() // preparing the std output for the query result
}
