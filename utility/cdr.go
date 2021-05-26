package utility

import "fmt"

type CDR struct {
	Number, Code, MessageId, MyMessageId string
}

func (cdr *CDR) Log() string {
	s := fmt.Sprintf("%s,%s,%s,%s", cdr.Number, cdr.Code, cdr.MyMessageId, cdr.MessageId)
	return s
}
