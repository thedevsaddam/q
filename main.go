package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

const logo = `              
                    
     QQQQQQQQQ      
   QQ:::::::::QQ    
 QQ:::::::::::::QQ  
Q:::::::QQQ:::::::Q 
Q::::::O   Q::::::Q 
Q:::::O     Q:::::Q 
Q:::::O     Q:::::Q 
Q:::::O     Q:::::Q 
Q:::::O     Q:::::Q 
Q:::::O     Q:::::Q 
Q:::::O  QQQQ:::::Q 
Q::::::O Q::::::::Q 
Q:::::::QQ::::::::Q 
 QQ::::::::::::::Q  
   QQ:::::::::::Q   
     QQQQQQQQ::::QQ 
             Q:::::Q
              QQQQQQ
					

Query JSON, CSV, YML, XML data from commandline
For more info visit: https://github.com/thedevsaddam/qcli
`

var (
	jq *gojsonq.JSONQ
)

func main() {

	if versionFlagProvided() {
		fmt.Println(logo)
		return
	}

	if !setDecoder() { // Checking for the datatype feeded, i.e: json/xml/yaml/csv, returns true if data provided
		return
	}

	checkFlags()       // Checking for the flags provided by the user
	checkCommands()    // checking for the commands provided by the user to query over
	prepareStdOutput() // preparing the std output for the query result
}
