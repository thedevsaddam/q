package main

import (
	"fmt"
	"regexp"
)

const (
	eq             = "equals"
	gt             = "greater_than"
	gte            = "greater_than_or_equals"
	lt             = "less_than"
	lte            = "less_than_or_equals"
	ne             = "not_equals"
	dataTypeInt    = "int:"
	dataTypeFloat  = "float:"
	dataTypeString = "string:" //for string
)

var (
	dtypes     []string
	operations = make(map[string]string)
)

func init() {
	dtypes = []string{dataTypeInt, dataTypeFloat}
	operations[eq] = "="
	operations[gt] = ">"
	operations[gte] = ">="
	operations[lt] = "<"
	operations[lte] = "<="
	operations[ne] = "!="
}

func getOptString(s string) string {
	opt := ""
	for _, o := range operations {
		if matched, err := regexp.MatchString(fmt.Sprintf(".%s.", o), s); matched {
			panicOnError(err)
			opt = o
		}
	}
	return opt
}
