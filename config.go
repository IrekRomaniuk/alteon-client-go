package alteon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	//"crypto/tls"
)

// GetItem - Returns a specifc Item 
func (c *Client) GetItem(Table,Index string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", c.HostURL, Table, Index), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item  map[string]interface{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// CreateItem - Create new item
func (c *Client) CreateItem(Item []byte, Table string, Index string) (*Response, error) {
	//rb, err := json.Marshal(Items[0])
	/*rb, err :=json.MarshalIndent(Items, "", "    ")
	if err != nil {
		return nil, err
	}*/
    //fmt.Printf("Calling %s\n",string(c.HostURL))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s", c.HostURL, Table, Index), strings.NewReader(string(Item)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := Response{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) UpdateItem(Item []byte, Table string, Index string) (*Response, error) {
	//rb, err := json.Marshal(realServerItems[0])
	//rb, err :=json.MarshalIndent(Item, "", "    ")
	/*if err != nil {
		return nil, err
	}*/
    //fmt.Printf("%s\n",string(rb))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", c.HostURL, Table, Index), strings.NewReader(string(Item)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := Response{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}


func (c *Client) DeleteItem(Table string, Index string) error {
	
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", c.HostURL, Table, Index), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	r := Response{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	return nil
}