package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq"
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

var (
	defaultDataType string
	find            string
	command         string
	columns         string
	from            string
	where           string
	orWhere         string
	output          interface{}
	aggregateColumn string
	version         bool
)

func init() {
	flag.StringVar(&defaultDataType, "type", "json", "data type: json, yml, xml, yml, csv")
	flag.StringVar(&from, "from", "", "from can be: items, users")
	flag.StringVar(&where, "where", "", "where would be the query where clause: name=macbook, price>1100")
	flag.StringVar(&orWhere, "orWhere", "", "orWhere would be the query orWhere clause: name=macbook, price>1100")
	flag.StringVar(&command, "command", "", "command can be: first, last, count, avg etc")
	flag.StringVar(&find, "find", "", `find works like: --find="items[0].price"`)
	flag.StringVar(&columns, "columns", "*", "columns can be: * or columnA,columnB")
	flag.BoolVar(&version, "version", false, "print version information")
}

func main() {
	flag.Parse()
	if version {
		fmt.Println(logo)
		return
	}
	data := readFromStdin()
	jq := gojsonq.New()

	switch defaultDataType {
	case "xml":
		fmt.Println("implement xml decoder")
	case "yml":
		fmt.Println("implement yml decoder")
	case "csv":
		fmt.Println("implement csv decoder")
	case "json":
		jq.JSONString(data)
	}

	if command != "" {
		cc := strings.Split(command, ":")
		command = strings.ToLower(cc[0])
		if len(cc) > 1 {
			aggregateColumn = cc[1]
		}
	}

	if find != "" {
		command = "find"
	} else if from != "" {
		jq.From(from)
	}

	if columns != "*" {
		a := strings.Split(columns, ",")
		for i := 0; i < len(a); i++ {
			a[i] = strings.TrimSpace(a[i])
		}
		jq.Select(a...)
	}

	if where != "" {
		ww := strings.Split(where, ",")
		for _, w := range ww {
			if w != "" {
				kk := strings.Split(w, "=") //todo: split with proper condition
				if strings.Contains(kk[1], "int:") {
					v, err := strconv.Atoi(strings.TrimPrefix(kk[1], "int:"))
					if err != nil {
						panic(err)
					}
					jq.Where(kk[0], "=", v)
				}
			}
		}
	}

	if orWhere != "" {
		ww := strings.Split(orWhere, ",")
		for _, w := range ww {
			if w != "" {
				kk := strings.Split(w, "=") //todo: split with proper condition
				if strings.Contains(kk[1], "int:") {
					v, err := strconv.Atoi(strings.TrimPrefix(kk[1], "int:"))
					if err != nil {
						panic(err)
					}
					jq.OrWhere(kk[0], "=", v)
				}
			}
		}
	}

	switch command {
	case "find":
		output = jq.Find(find)
	case "first":
		output = jq.First()
	case "last":
		output = jq.Last()
	case "count":
		output = jq.Count()
	case "avg":
		if aggregateColumn != "" {
			output = jq.Avg(aggregateColumn)
		} else {
			output = jq.Avg()
		}
	case "sum":
		if aggregateColumn != "" {
			output = jq.Sum(aggregateColumn)
		} else {
			output = jq.Sum()
		}
	case "min":
		if aggregateColumn != "" {
			output = jq.Min(aggregateColumn)
		} else {
			output = jq.Min()
		}
	case "max":
		if aggregateColumn != "" {
			output = jq.Max(aggregateColumn)
		} else {
			output = jq.Max()
		}
	default:
		output = jq.Get()
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	if err := enc.Encode(output); err != nil {
		fmt.Println(err)
		return
	}
}

func readFromStdin() string {
	var data []byte
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return string(data)
}
