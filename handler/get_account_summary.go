package handler

import (
	_ "context"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/goanda"
)

type accountSummaryService interface {
	OandaRequest(endpoint string) []byte
	// OandaSend(endpoint string, data []byte) []byte
	// OandaUpdate(endpoint string, data []byte) []byte
	// OandaGetOrderDetails(instrument string, units string) goanda.OrderDetails
	// OandaGetAccountSummary(subAccountID string) goanda.AccountSummary
	// OandaCreateOrder(body goanda.OrderPayload) goanda.OrderResponse
}

func GetAccountSummary(svc accountSummaryService, accountID string) goanda.AccountSummary {
	f := newGetAccountSummaryFlow(svc, accountID)
	res := f.doAccountSummaryRequest()
	return res
}
