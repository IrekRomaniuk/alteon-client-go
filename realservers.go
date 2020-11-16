package alteon

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"crypto/tls"
)

// GetRealServer - Returns a specifc RealServer 
func (c *Client) GetRealServer(RealServerID string) (*RealServer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/SlbNewCfgEnhRealServerTable/%s", c.HostURL, RealServerID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	realServer:= RealServer{}
	err = json.Unmarshal(body, &realServer)
	if err != nil {
		return nil, err
	}

	return &realServer, nil
}