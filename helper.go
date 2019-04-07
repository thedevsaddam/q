package main

import "strings"

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

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func trimLeadingTrailingPercents(s string) string {
	s = strings.TrimLeft(s, "%")
	s = strings.TrimRight(s, "%")
	return s
}
