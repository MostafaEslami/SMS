package utility

import (
	"fmt"
	"strconv"
	"sync"
)

//var FileWriters = map[string]*FileWriter{}
//var writer *FileWriter

//var CreditFile = "credit.txt"
var counter int

//type FileWriter struct {
//	mu   sync.Mutex
//	File string
//}

var mu sync.Mutex

func InitializeCredit(credit string) {
	//writer = NewFileWriter(CreditFile)
	x, _ := strconv.Atoi(credit)
	counter = x
	//NewFileWriter(credit)
}

//func ReadCredit() int {
//	f, _ := os.Open(CreditFile)
//	r4 := bufio.NewReader(f)
//	b4, _ := r4.Peek(5)
//	ii, _ := strconv.Atoi(string(b4))
//	return ii
//}
func GetCredit() int {
	if counter == 0 {
		Log("WARNING", "credit is zero")
	}
	return counter
}

func HasCredit() bool {
	return GetCredit() > 0

}

func DecreaseCreditAsync() {

	if counter > 0 {
		mu.Lock()
		defer mu.Unlock()
		//bs := []byte(strconv.Itoa(counter - 1))
		counter = counter - 1
		//writer.Write(bs)
		Log("DEBUG", "Decrease credit : ", counter)
		s := fmt.Sprintf("%d", counter)
		LogCredit(s)
	}

}

//func NewFileWriter(file string) *FileWriter {
//	path, err := filepath.Abs(file)
//	if err != nil {
//		return nil
//	}
//
//	writer, exists := FileWriters[path]
//	if !exists {
//		writer = &FileWriter{File: path}
//		FileWriters[path] = writer
//	}
//
//	return writer
//}
//
//func (w *FileWriter) Write(content []byte) error {
//	w.mu.Lock()
//	defer w.mu.Unlock()
//
//	file, err := os.OpenFile(w.File, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
//	if err != nil {
//		return err
//	}
//
//	file.Write(content)
//	file.Close()
//
//	return nil
//}
