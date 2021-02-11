package goanda

// type Headers struct {
// 	contentType    string
// 	agent          string
// 	DatetimeFormat string
// 	auth           string
// }

// type Connection interface {
// 	Get(endpoint string) []byte
// 	Post(endpoint string, data []byte) []byte
// 	// Update(endpoint string, data []byte) []byte
// }

// type OandaConnection struct {
// 	hostname       string
// 	port           int
// 	ssl            bool
// 	token          string
// 	accountID      string
// 	DatetimeFormat string
// 	headers        *Headers
// }

// const OANDA_AGENT string = "v20-golang/0.0.1"

// func NewConnection(accountID string, token string, env string) (*OandaConnection, error) {
// 	hostname := ""
// 	// should we use the live API?
// 	if env == "live" {
// 		hostname = "https://api-fxtrade.oanda.com/v3"
// 	} else if env == "practice" {
// 		hostname = "https://api-fxpractice.oanda.com/v3"
// 	} else {
// 		return nil, errors.New("ENVIRONMENT in env.sh should be \"live\" or \"practice\"")
// 	}

// 	var buffer bytes.Buffer
// 	// Generate the auth header
// 	buffer.WriteString("Bearer ")
// 	buffer.WriteString(token)

// 	authHeader := buffer.String()
// 	// Create headers for oanda to be used in requests
// 	headers := &Headers{
// 		contentType:    "application/json",
// 		agent:          OANDA_AGENT,
// 		DatetimeFormat: "RFC3339",
// 		auth:           authHeader,
// 	}
// 	// Create the connection object
// 	connection := &OandaConnection{
// 		hostname:  hostname,
// 		port:      443,
// 		ssl:       true,
// 		token:     token,
// 		headers:   headers,
// 		accountID: accountID,
// 	}

// 	return connection, nil
// }

// // TODO: include params as a second option
// func (c *OandaConnection) Get(endpoint string) []byte {
// 	client := http.Client{
// 		Timeout: time.Second * 5, // 5 sec timeout
// 	}

// 	url := createUrl(c.hostname, endpoint)

// 	// New request object
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	checkErr(err)

// 	body := makeRequest(c, endpoint, client, req)

// 	return body
// }

// func (c *OandaConnection) Post(endpoint string, data []byte) []byte {
// 	client := http.Client{
// 		Timeout: time.Second * 5, // 5 sec timeout
// 	}

// 	url := createUrl(c.hostname, endpoint)

// 	// New request object
// 	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
// 	checkErr(err)

// 	body := makeRequest(c, endpoint, client, req)

// 	return body
// }

// func (c *OandaConnection) Update(endpoint string, data []byte) []byte {
// 	client := http.Client{
// 		Timeout: time.Second * 5,
// 	}

// 	url := createUrl(c.hostname, endpoint)

// 	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
// 	checkErr(err)
// 	body := makeRequest(c, endpoint, client, req)
// 	return body
// }
