package utility

import "fmt"

type CDR struct {
	Number, Code, MessageId string
}

func (cdr *CDR) Log() string {
	s := fmt.Sprintf("%s,%s,%s", cdr.Number, cdr.Code, cdr.MessageId)
	return s
}
