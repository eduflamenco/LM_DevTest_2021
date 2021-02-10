package TraslateImplementation

import (
	"fmt"
)

func NewTraslator() *translator {
	return &translator{
	}
}

type translator struct {
}

func(t *translator) StringToBinary(text string)(string, error){
	var binString string
	for _, c := range text {
		binString = fmt.Sprintf("%s%b",binString, c)
	}
	return binString, nil
}

func(t *translator) BinaryToString(binary []byte)(string, error){
	s := string(binary)
	//t.logger.Println("Binary converted to string: ", s)
	return s, nil
}
