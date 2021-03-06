package main

import (
	"flag"
	"fmt"
	"log"
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
	sort            string
	sortBy          string
	groupBy         string
	distinct        string
	offset          string
	limit           string
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
	flag.StringVar(&find, "find", "", `find works like: -find="items[0].price"`)
	flag.StringVar(&columns, "columns", "*", "columns can be: * or columnA,columnB")
	flag.StringVar(&sort, "sort", "", `sort accept argument: -sort="asc/desc"`)
	flag.StringVar(&sortBy, "sortBy", "", `sort accept argument: -sortBy="price:desc"`)
	flag.StringVar(&groupBy, "groupBy", "", `groupBy accept argument: -groupBy="category"`)
	flag.StringVar(&offset, "offset", "", `offset accept argument: -offset="5"`)
	flag.StringVar(&limit, "limit", "", `limit accept argument: -limit="5"`)
	flag.StringVar(&distinct, "distinct", "", `distinct accept argument: -distinct="category"`)
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

	if where != "" { // FIXME: handle null query; -where="id!=null"
		ww := strings.Split(where, ",")
		for _, w := range ww {
			if w != "" {
				opt := getOptString(w)
				kk := strings.Split(w, opt) //todo: split with proper condition
				isStringVal := true
				for _, dt := range dtypes {
					if strings.Contains(kk[1], dt) {
						v, err := strconv.Atoi(strings.TrimPrefix(kk[1], dt))
						handleError(err)
						jq.Where(kk[0], opt, v)
						isStringVal = false
						break
					}
				}
				if isStringVal {
					sw := strings.HasSuffix(kk[1], "%")
					ew := strings.HasPrefix(kk[1], "%")
					kk[1] = trimLeadingTrailingPercents(kk[1])
					if sw && ew {
						jq.WhereContains(kk[0], kk[1])
					} else if sw && !ew {
						jq.WhereStartsWith(kk[0], kk[1])
					} else if !sw && ew {
						jq.WhereEndsWith(kk[0], kk[1])
					} else {
						jq.Where(kk[0], opt, kk[1])

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
						handleError(err)
						jq.OrWhere(kk[0], opt, v)
						break
					}
				}
			}
		}
	}

	if sort != "" {
		if sort == "asc" {
			jq.Sort(sort)
		} else {
			jq.Sort("desc")
		}
	}

	if sortBy != "" {
		if strings.Contains(sortBy, ":") {
			ss := strings.Split(sortBy, ":")
			jq.SortBy(ss[0], ss[1])
		} else {
			jq.SortBy(sortBy)
		}
	}

	if groupBy != "" {
		jq.GroupBy(groupBy)
	}

	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			log.Println("Invalid offset value:", err)
			return
		}
		jq.Offset(o)
	}

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			log.Println("Invalid limit value:", err)
			return
		}
		jq.Limit(l)
	}

	if distinct != "" {
		jq.Distinct(distinct)
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
	case "pluck":
		output = jq.Pluck(aggregateColumn)
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

// setDecoder checks for the type of the data to feed on
func setDecoder() bool {

	data := readFromStdin()

	if data == "" {
		fmt.Println("Empty input!")
		return false
	}

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
	return true
}

func versionFlagProvided() bool {
	return version
}
