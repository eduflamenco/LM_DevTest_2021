package TraslateImplementation

type ITraslator interface {
	StringToBinary(text string)(string, error)
	BinaryToString(binary []byte)(string, error)
	StringToMorse(text string)(string, error)
}
