package utility

import "fmt"

type CDR struct {
	From, To string
}

func (cdr *CDR) Log() string {
	s := fmt.Sprintf("%s,%s", cdr.From, cdr.To)
	return s
}
