package main

import (
	"fmt"
	"log"
	"os"
	"parseLogInGo/parser"
	"sync"
	"time"
)

type ParseLog struct {
	parser parser.Parser
}

func main() {
	wg := sync.WaitGroup{}
	file := parser.File
	go func(file *parser.LogFile) {
		wg.Add(1)
		err := file.ReadLine("/home/www/go/src/parseLogInGo/log1.log")
		if err!= nil {
			log.Fatalln(err)
		}
		log.Println("文件读取完成！")
		wg.Done()
	}(file)
	go func(file *parser.LogFile) {
		wg.Add(1)
		var tmpCon []byte
		for {
			select {
			case tmpCon = <- file.ConCh:
				log.Println(string(tmpCon))
			case <-time.NewTicker(time.Duration(1*time.Second)).C:
				log.Println("timeout for 1s")
			}
		}
		wg.Done()
	}(file)

	wg.Wait()
	os.Exit(0)

	var parser = &parser.LogParser{
		"./log1.log",
	}
	//var v1 Parser = *parser
	//fmt.Println(v1)
	fmt.Println(parser)
}

//func main() {
//	var a animal
//	var c cat
//	a=c
//	a.printInfo()
//	//使用另外一个类型赋值
//	var d dog
//	a=d
//	a.printInfo()
//}
//type animal interface {
//	printInfo()
//}
//type cat int
//type dog int
//func (c cat) printInfo(){
//	fmt.Println("a cat")
//}
//func (d dog) printInfo(){
//	fmt.Println("a dog")
//}
