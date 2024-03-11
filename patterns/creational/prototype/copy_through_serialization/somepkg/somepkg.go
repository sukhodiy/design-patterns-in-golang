package somepkg

import (
	"bytes"
	"encoding/gob"
)

type SomeStruct struct {
	SomeField1 int
	someField2 int
}

func (p *SomeStruct) DeepCopy() *SomeStruct {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	//fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	result := SomeStruct{}
	_ = d.Decode(&result)

	return &result
}

func NewSomeStruct() *SomeStruct {
	return &SomeStruct{5, 10}
}

func (somestruct *SomeStruct) GetSomeField1() int {
	return somestruct.SomeField1
}

func (somestruct *SomeStruct) GetSomeField2() int {
	return somestruct.someField2
}

func (somestruct *SomeStruct) SetSomeField1(SomeField1 int) *SomeStruct {
	somestruct.SomeField1 = SomeField1
	return somestruct
}

func (somestruct *SomeStruct) SetSomeField2(someField2 int) *SomeStruct {
	somestruct.someField2 = someField2
	return somestruct
}
