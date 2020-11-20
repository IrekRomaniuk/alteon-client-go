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
	/*{
		"status": "ok"
	}*/
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
func (c *Client) CreateRealServer(realServerItems []RealServerItem, RealServerID string) (*Status, error) {
	//rb, err := json.Marshal(realServerItems[0])
	rb, err :=json.MarshalIndent(realServerItems[0], "", "    ")
	if err != nil {
		return nil, err
	}
    //fmt.Printf("%s\n",string(rb))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/SlbNewCfgEnhRealServerTable/%s", c.HostURL, RealServerID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	status := Status{}
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}