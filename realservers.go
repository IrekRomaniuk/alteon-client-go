package alteon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

// CreateRealServer - Create new server
func (c *Client) CreateRealServer(realServerItems []RealServerItem, RealServerID string) (*RealServer, error) {
	rb, err := json.Marshal(realServerItems)
	if err != nil {
		return nil, err
	}
    fmt.Printf("%s/SlbNewCfgEnhRealServerTable/%s", c.HostURL, RealServerID)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/SlbNewCfgEnhRealServerTable/%s", c.HostURL, RealServerID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	server := RealServer{}
	err = json.Unmarshal(body, &server)
	if err != nil {
		return nil, err
	}

	return &server, nil
}