package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const logo = `



                                        lllllll   iiii
                                        l:::::l  i::::i
                                        l:::::l   iiii
                                        l:::::l
   qqqqqqqqq   qqqqq    cccccccccccccccc l::::l iiiiiii
  q:::::::::qqq::::q  cc:::::::::::::::c l::::l i:::::i
 q:::::::::::::::::q c:::::::::::::::::c l::::l  i::::i
q::::::qqqqq::::::qqc:::::::cccccc:::::c l::::l  i::::i
q:::::q     q:::::q c::::::c     ccccccc l::::l  i::::i
q:::::q     q:::::q c:::::c              l::::l  i::::i
q:::::q     q:::::q c:::::c              l::::l  i::::i
q::::::q    q:::::q c::::::c     ccccccc l::::l  i::::i
q:::::::qqqqq:::::q c:::::::cccccc:::::cl::::::li::::::i
 q::::::::::::::::q  c:::::::::::::::::cl::::::li::::::i
  qq::::::::::::::q   cc:::::::::::::::cl::::::li::::::i
    qqqqqqqq::::::q     cccccccccccccccclllllllliiiiiiii
            q:::::q
            q:::::q
           q:::::::q
           q:::::::q
           q:::::::q
           qqqqqqqqq


Query JSON, CSV, YML, XML data from commandline
For more info visit: https://github.com/thedevsaddam/qcli
`

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
