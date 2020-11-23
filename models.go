package alteon

// RealServer -
type RealServer struct {
	Items []RealServerItem `json:"SlbNewCfgEnhRealServerTable,omitempty"`
}

// RealServerItem -
type RealServerItem struct {
	IpAddr  string `json:"IpAddr"`
	IpAddr  string `json:"IpAddr"`
	Weight int    `json:"Weight"`
	MaxConns  int `json:"MaxConns"`
	TimeOut int    `json:"TimeOut"`
	PingInterval  int `json:"PingInterval"`
	FailRetry int    `json:"FailRetry"`
	SuccRetry int    `json:"SuccRetry"`
	Name string    `json:"Name"`
}

type Response struct {
	Status  string `json:"status"`
}
