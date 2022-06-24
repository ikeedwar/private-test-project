package rwtest

type TestStruct struct {
	H string
}

type RwStruct struct {
	S string
	I int
}

func (rws *RwStruct) RwRead(p *string) (n int, err error) {
	p = &rws.S
	return rws.I, nil
}

func PointTest(str string, testStruct *TestStruct) (string, error) {
	testStruct.H = str
	return str, nil
}
