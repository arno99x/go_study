package main

import (
	"fmt"
	"time"
)

type Reader interface {
	Read(readChan chan string)
}
type Processer interface {
	Process(readChan <-chan string, writeChan chan<- string)
}
type Writer interface {
	Write(writeChan <-chan string)
}

type LogFileReader struct {
	Path string
}
type LogProcesser struct {
}
type EsWriter struct {
	Address string
	Path    string
}

func (reader *LogFileReader) Read(readChan chan string) {
	readChan <- reader.Path
}

func (processer *LogProcesser) Process(readChan <-chan string, writeChan chan<- string) {
	ms := <-readChan
	ms = "processed \"" + ms + "\""
	writeChan <- ms
}

func (writer *EsWriter) Write(writeChan <-chan string) {
	ms := <-writeChan
	fmt.Println("write ", ms, " to ", writer.Address, writer.Path)
}

func main() {
	readChan := make(chan string)
	writeChan := make(chan string)
	reader := LogFileReader{Path: "/usr/local/log/logfile.log"}
	processer := LogProcesser{}
	writer := EsWriter{Address: "http://localhost:9200", Path: "/log/serverlog"}

	go reader.Read(readChan)
	go processer.Process(readChan, writeChan)
	go writer.Write(writeChan)

	time.Sleep(1 * time.Second)
}
