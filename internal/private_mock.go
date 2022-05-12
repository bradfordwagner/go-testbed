package internal

type PrivateMock struct {
}

func (p *PrivateMock) SayHello() string {
	return p.getMessage()
}

func (p *PrivateMock) getMessage() string {
	return "hi friends"
}
