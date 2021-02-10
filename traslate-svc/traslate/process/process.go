package process

import (
	"fmt"
	"traslate-svc/traslate/TraslateImplementation"
)

type IProcess interface {
	Process(text, fmtOrig, fmtDest string) (string, error)
}

func NewProcessTraslator(traslator TraslateImplementation.ITraslator) *processTraslator{
	return &processTraslator{
		traslator: traslator,
	}
}

type processTraslator struct {
	traslator TraslateImplementation.ITraslator
}

func (p *processTraslator) Process(input, fmtOrig, fmtDest string) (result string, err error){
	if fmtOrig == "text" && fmtDest == "byte" {
		 p.traslator.StringToBinary(input)
	}else if fmtOrig == "text" && fmtDest == "morse"{
		result, err = p.traslator.StringToMorse(input)
	} else {
		result, err = p.traslator.BinaryToString([]byte(input))
	}

	if err != nil{
		fmt.Println("Error:", err.Error())
	}
	return result, nil
}


