package service

import (
	"github.com/bigsupanat/oanda-trading-api-orchestrator/goanda"
)

type oandaRequestFunc func(endpoint string) []byte

type oandaSendFunc func(endpoint string, data []byte) []byte

type oandaUpdateFunc func(endpoint string, data []byte) []byte

type oandaGetOrderDetailsFunc func(instrument string, units string) goanda.OrderDetails

type oandaGetGetAccountSummaryFunc func(subAccountID string) goanda.AccountSummary

type oandaCreateOrderfunc func(body goanda.OrderPayload) goanda.OrderResponse

type getAccountSummaryFunc func(subAccountID string) goanda.AccountSummary

//type OandaRequestAccountSummaryFn func(subAccountID string, c )
