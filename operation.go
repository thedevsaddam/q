package main

import (
	"fmt"
	"regexp"
)

const (
	eq  = "="
	gt  = ">"
	gte = ">="
	lt  = "<"
	lte = "<="
	ne  = "!="
)

func getOptString(s string) string {
	opt := ""
	if matched, err := regexp.MatchString(fmt.Sprintf(".%s.", eq), s); matched {
		panicOnError(err)
		opt = eq
	}
	if matched, err := regexp.MatchString(fmt.Sprintf(".%s.", gt), s); matched {
		panicOnError(err)
		opt = gt
	}
	if matched, err := regexp.MatchString(fmt.Sprintf(".%s.", gte), s); matched {
		panicOnError(err)
		opt = gte
	}
	if matched, err := regexp.MatchString(fmt.Sprintf(".%s.", lt), s); matched {
		panicOnError(err)
		opt = lt
	}
	if matched, err := regexp.MatchString(fmt.Sprintf(".%s.", lte), s); matched {
		panicOnError(err)
		opt = lte
	}
	if matched, err := regexp.MatchString(fmt.Sprintf(".%s.", ne), s); matched {
		panicOnError(err)
		opt = ne
	}
	return opt
}
