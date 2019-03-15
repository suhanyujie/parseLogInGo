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

var wg sync.WaitGroup

func main() {
	file := parser.File
	wg.Add(2)
	go func(file *parser.LogFile,wg *sync.WaitGroup) {
		log.Println("文件读取开始！")
		err := file.ReadLine("/home/www/go/src/parseLogInGo/log1.log")
		if err!= nil {
			wg.Done()
			log.Fatalln(err)
		}
		log.Println("文件读取完成！")
		wg.Done()
	}(file,&wg)
	go func(file *parser.LogFile,wg *sync.WaitGroup) {
		// for循环中的超时，需要对timer进行重用，否则造成CPU的浪费 http://xargin.com/go-timer/
		timer := time.NewTimer(time.Second)
		defer timer.Stop()
		var tmpCon []byte
		for {
			timer.Reset(time.Second*2)
			select {
			case tmpCon = <- file.ConCh:
				log.Println(string(tmpCon))
			case <-timer.C:
				wg.Done()
				log.Println("timeout.May be finished...")
			}
		}
	}(file,&wg)
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
