package alteon

// RealServer -
type RealServer struct {
	Index    int         `json:"index,omitempty"`
	Items []RealServerItem `json:"items,omitempty"`
}

// RealServerItem -
type RealServerItem struct {
	IpAddr  string `json:"IpAddr"`
	Weight int    `json:"Weight"`
	MaxConns  int `json:"MaxConns"`
	TimeOut int    `json:"TimeOut "`
	PingInterval  int `json:"PingInterval"`
	FailRetry int    `json:"FailRetry"`
	SuccRetry int    `json:"SuccRetry "`
}
