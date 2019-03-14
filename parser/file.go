package parser

import (
	"bufio"
	"io"
	"os"
)

type LogFile struct {
	ConCh chan []byte
}

var File *LogFile

func init()  {
	File = &LogFile{
		make(chan []byte, 512),
	}
}

func (this *LogFile) ReadLine(path string) (error) {
	file,err := os.Open(path)
	if err!= nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		conByte,_,err := reader.ReadLine()
		if err!=nil {
			if err == io.EOF {
				break
			}
			return err
		}
		File.ConCh <- conByte
	}
	return nil
}
