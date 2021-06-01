package utility

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var FileWriters = map[string]*FileWriter{}
var writer *FileWriter

var CreditFile = "credit.txt"
var counter int

type FileWriter struct {
	mu   sync.Mutex
	File string
}

func IniitalizeCredit() {
	writer = NewFileWriter(CreditFile)
	counter = ReadCredit()
}
func ReadCredit() int {
	f, _ := os.Open(CreditFile)
	r4 := bufio.NewReader(f)
	b4, _ := r4.Peek(10)
	ii, _ := strconv.Atoi(string(b4))
	return ii
}
func GetCredit() int {
	return counter
}

func HasCredit() bool {
	if ReadCredit() > 0 {
		return true
	}
	return false
}
func DecreaseCredit() {
	if counter > 0 {
		bs := []byte(strconv.Itoa(counter - 1))
		counter--
		writer.Write(bs)
		Log("INFO", "Decrease credit : ", counter)
	}
}

func IncreaseCredit() {
	bs := []byte(strconv.Itoa(counter - 1))
	counter++
	writer.Write(bs)
	Log("INFO", "increment credit : ", counter)
}

func SetCredit(c int) {
	bs := []byte(strconv.Itoa(c))
	counter = c
	writer.Write(bs)
	Log("INFO", "Set Credit : ", counter)
}

func NewFileWriter(file string) *FileWriter {
	path, err := filepath.Abs(file)
	if err != nil {
		return nil
	}

	writer, exists := FileWriters[path]
	if !exists {
		writer = &FileWriter{File: path}
		FileWriters[path] = writer
	}

	return writer
}

func (w *FileWriter) Write(content []byte) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	file, err := os.OpenFile(w.File, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	file.Write(content)
	file.Close()

	return nil
}
