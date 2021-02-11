package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/handler"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/service"
)

func cancelOrderFn(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		orderSpecifier := "" //Todo: receive from request
		res, err := handler.CancelOrder(s, orderSpecifier)
		if err != nil {
			log.Println(err)
		}
		bytRet, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bytRet)
	}
}
