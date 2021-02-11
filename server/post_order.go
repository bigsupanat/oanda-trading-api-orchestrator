package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/goanda"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/handler"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/service"
)

func postOrderFn(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			req goanda.OrderPayload
			res goanda.OrderResponse
		)

		byt, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(byt, &req)
		if err != nil {
			log.Println(err)
		}
		res, err = handler.CreateOrder(s, req)
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
