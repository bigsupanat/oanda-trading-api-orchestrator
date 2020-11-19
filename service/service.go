package service

type Service struct {
	OandaRequestFunc           oandaRequestFunc
	OandaSendFunc              oandaSendFunc
	OandaUpdateFunc            oandaUpdateFunc
	OandaGetOrderDetailsFunc   oandaGetOrderDetailsFunc
	OandaGetAccountSummaryFunc oandaGetGetAccountSummaryFunc
	OandaCreateOrderFunc       oandaCreateOrderfunc

	GetAccountSummaryFunc getAccountSummaryFunc
}

func (s Service) OandaRequest(endpoint string) []byte {
	return s.OandaRequestFunc(endpoint)
}
