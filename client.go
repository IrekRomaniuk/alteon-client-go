package alteon

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"crypto/tls"
)

// HostURL - Default alteon URL
const HostURL string = "http://localhost:443"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}


// NewClient -
func NewClient(host, username, password *string) (*Client, error) {
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second, Transport: tr},
		// Default alteon URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if (username != nil) && (password != nil) {
		
		c.Token= "Basic " + b64.StdEncoding.EncodeToString([]byte(*username + ":" + *password))

	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
