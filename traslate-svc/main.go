package main

import (
	"fmt"
	"traslate-svc/traslate/TraslateImplementation"
	processTrasl "traslate-svc/traslate/process"
)

func main(){


	traslate := TraslateImplementation.NewTraslator()
	process := processTrasl.NewProcessTraslator(traslate)
	fmt.Println(process.Process("CC", "text","binary"))
	//fmt.Println(binTras.BinaryToString([]byte{65, 66, 67, 226, 130, 172}))
	fmt.Println(traslate.StringToMorse("Convert this to morse"))
}
