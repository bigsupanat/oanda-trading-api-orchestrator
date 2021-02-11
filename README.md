# oanda-trading-api-orchestrator


env.sh

SERVICE_PORT="0.0.0.0:8080"
ENVIRONMENT="practice"
ACCOUNT_ID="" //from oanda
ACCESS_TOKEN="" //from oanda


Endpoint: 

 /get_account_summary : getsummary of the account using http.Get

 /createOrder : json request using http.Post
                {
                     "order": 
                     {      "type": "LIMIT",
                            "price": "1.25000",
                            "units": 1,
                            "instrument": "EUR_USD",
                            "timeInForce": "GTC"
                    }
                }

 /cancelOrder : cancel order by order ID using http.Put
                (not complete)

 This project is not complete. I will try to finished it asap.