package parser

type Parser interface {
	LogReader()
	LogOutput()
}

type LogParser struct {
	FilePath string
}

// 日志外层的数据结构
type LogDataWrap struct {

}

func (this LogParser) LogReader()  {

}

func (this LogParser) LogOutput()  {

}


