package main

import "fmt"

type CDR struct {
	from, to string
}

func (cdr CDR) Log() string {
	s := fmt.Sprintf("%s,%s", cdr.from, cdr.to)
	return s
}
