package main

import "strings"

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func trimLeadingTrailingPercents(s string) string {
	s = strings.TrimLeft(s, "%")
	s = strings.TrimRight(s, "%")
	return s
}
