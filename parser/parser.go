package parser

type Parser interface {
	LogReader()
	LogOutput()
}

type LogParser struct {
	FilePath string
}

func (this LogParser) LogReader()  {

}

func (this LogParser) LogOutput()  {

}


