package main

import (
	"flag"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq"
)

var (
	defaultDataType string
	separator       string
	find            string
	command         string
	columns         string
	from            string
	where           string
	orWhere         string
	aggregateColumn string
	pretty          bool
	version         bool
)

func init() {
	flag.StringVar(&defaultDataType, "type", "json", "data type: json, yml, xml, yml, csv")
	flag.StringVar(&separator, "separator", ".", "separator can be: . / -> / --> / => etc. Default is DOT[.]")
	flag.StringVar(&from, "from", "", "from can be: items, users")
	flag.StringVar(&where, "where", "", "where would be the query where clause: name=macbook, price>1100")
	flag.StringVar(&orWhere, "orWhere", "", "orWhere would be the query orWhere clause: name=macbook, price>1100")
	flag.StringVar(&command, "command", "", "command can be: first, last, count, avg etc")
	flag.StringVar(&find, "find", "", `find works like: --find="items[0].price"`)
	flag.StringVar(&columns, "columns", "*", "columns can be: * or columnA,columnB")
	flag.BoolVar(&pretty, "pretty", false, "print formatted output")
	flag.BoolVar(&version, "version", false, "print version information")
	flag.Parse()
}

// checkFlags checks for each of the flag provided by the user if available
func checkFlags() {
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
				opt := getOptString(w)
				kk := strings.Split(w, opt) //todo: split with proper condition
				for _, dt := range dtypes {
					if strings.Contains(kk[1], dt) {
						v, err := strconv.Atoi(strings.TrimPrefix(kk[1], dt))
						panicOnError(err)
						jq.Where(kk[0], opt, v)
						break
					}
				}
			}
		}
	}

	if orWhere != "" {
		ww := strings.Split(orWhere, ",")
		for _, w := range ww {
			if w != "" {
				opt := getOptString(w)
				kk := strings.Split(w, opt) //todo: split with proper condition
				for _, dt := range dtypes {
					if strings.Contains(kk[1], dt) {
						v, err := strconv.Atoi(strings.TrimPrefix(kk[1], dt))
						panicOnError(err)
						jq.OrWhere(kk[0], opt, v)
						break
					}
				}
			}
		}
	}
}

// checkCommands checks for any command provided by the user during query & takes action accordingly
func checkCommands() {
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
}

// checkDataType checks for the type of the data to feed on
func checkDataType() {

	data := readFromStdin()

	switch defaultDataType {
	case "xml":
		jq = gojsonq.New(gojsonq.SetDecoder(&xmlDecoder{}), gojsonq.SetSeparator(separator)).JSONString(data)
	case "yml", "yaml":
		jq = gojsonq.New(gojsonq.SetDecoder(&yamlDecoder{}), gojsonq.SetSeparator(separator)).JSONString(data)
	case "csv":
		jq = gojsonq.New(gojsonq.SetDecoder(&csvDecoder{}), gojsonq.SetSeparator(separator)).JSONString(data)
	case "json":
		jq = gojsonq.New(gojsonq.SetSeparator(separator)).JSONString(data)
	}
}

func versionFlagProvided() bool {
	return version
}
