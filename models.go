package alteon

// RealServer -
type RealServer struct {
	Items []RealServerItem `json:"SlbNewCfgEnhRealServerTable,omitempty"`
}

// RealServerItem -
type RealServerItem struct {
	IpAddr  string `json:"IpAddr"`
	Weight int    `json:"Weight"`
	MaxConns  int `json:"MaxConns"`
	TimeOut int    `json:"TimeOut"`
	PingInterval  int `json:"PingInterval"`
	FailRetry int    `json:"FailRetry"`
	SuccRetry int    `json:"SuccRetry"`
	Name string    `json:"Name"`
	State int    `json:"State"`
}

type Response struct {
	Status  string `json:"status"`
	Message  string `json:"message"`
	TestErr  string `json:"testerr"`
	MibName  string `json:"mibName"`
}

