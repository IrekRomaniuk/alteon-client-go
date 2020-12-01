package alteon

// RealServer -
type RealServer struct {
	Items []RealServerItem `json:"SlbNewCfgEnhRealServerTable,omitempty"`
}

// RealServerItem -
type RealServerItem struct {
	Index  string `json:"Index"`
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

// ServerGroup -
type ServerGroup struct {
	Items [ServerGroupItem `json:"SlbNewCfgEnhGroupTable,omitempty"`
}

// ServerGroupItem -
type ServerGroupItem struct {
	Index  string `json:"Index"`
	AddServer  string `json:"AddServer"`
	RemoveServer string  `json:"RemoveServer"`
	Name  string `json:"Name"`
}

type Response struct {
	Status  string `json:"status"`
	Message  string `json:"message"`
	TestErr  string `json:"testerr"`
	MibName  string `json:"mibName"`
}

