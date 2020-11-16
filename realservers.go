package alteon

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"crypto/tls"
)

// GetRealServer - Returns a specifc RealServer 
func (c *Client) GetRealServer(RealServerID string) (*RealServerItem, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/SlbNewCfgEnhRealServerTable/%s", c.HostURL, RealServerID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	realServerItem := RealServerItem{}
	err = json.Unmarshal(body, &realServerItem)
	if err != nil {
		return nil, err
	}

	return &realServerItem, nil
}