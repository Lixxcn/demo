package main

import "fmt"

type ParameterCodec interface {
	DecodeParameters(s string) (string, error)
	EncodeParameters(s string) (string, error)
}

type ParameterCodecImpl struct {
	decodestring string
	encodestring string
}

func (p *ParameterCodecImpl) DecodeParameters(s string) (string, error) {
	return p.decodestring, nil
}

func (p *ParameterCodecImpl) EncodeParameters(s string) (string, error) {
	return p.encodestring, nil
}

func main() {
	p := ParameterCodecImpl{
		decodestring: "decode string",
		encodestring: "encode string",
	}

	s, error := (&p).DecodeParameters("123")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(s)

	var pc ParameterCodec = &p
	var pc2 *ParameterCodecImpl = &p
	s, error = pc.DecodeParameters("123")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(s)

	s, error = pc2.DecodeParameters("123")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(s)

}
