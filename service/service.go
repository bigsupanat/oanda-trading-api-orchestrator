package service

type Service struct {
	OandaGetFunc  func(endpoint string) []byte
	OandaPostFunc func(endpoint string, data []byte) []byte
	OandaPutFunc  func(endpoint string, data []byte) []byte
}

func (s Service) Request(endpoint string) []byte {
	return s.OandaGetFunc(endpoint)
}

func (s Service) Send(endpoint string, data []byte) []byte {
	return s.OandaPostFunc(endpoint, data)
}

func (s Service) Update(endpoint string, data []byte) []byte {
	return s.OandaPutFunc(endpoint, data)
}
