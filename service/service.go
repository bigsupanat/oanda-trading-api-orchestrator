package service

type Service struct {
	OandaGetFunc  func(endpoint string) []byte
	OandaPostFunc func(endpoint string, data []byte) []byte
	// OandaSendFunc              oandaSendFunc
	// OandaUpdateFunc            oandaUpdateFunc
	// OandaGetOrderDetailsFunc   oandaGetOrderDetailsFunc
	// OandaGetAccountSummaryFunc oandaGetGetAccountSummaryFunc
	// OandaCreateOrderFunc       oandaCreateOrderfunc

	// GetAccountSummaryFunc getAccountSummaryFunc
}

func (s Service) Request(endpoint string) []byte {
	return s.OandaGetFunc(endpoint)
}

func (s Service) Send(endpoint string, data []byte) []byte {
	return s.OandaPostFunc(endpoint, data)
}
