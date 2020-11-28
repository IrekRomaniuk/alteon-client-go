package alteon

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	username string
	password string
	uri string
}{
	{"", "", "",},
	{"usernamenotempty", "", "",},
	{"", "passwordnotempty", "",},
	{"usernamenotempty", "passwordnotempty", "",},
}

func TestNewClient(t *testing.T) {
	//var username, password, uri string = "", "", ""
	// t.Log("TestNewClient")
	for _, tt := range tests {
		fmt.Printf("\t\tusername: %s password: %s uri: %s\n", tt.username, tt.password, tt.uri)
		//t.Log(tt)
		_, err := NewClient(&tt.uri, &tt.username, &tt.password)
		//t.Log(c.Token, err)
		if (tt.username != "") && (tt.password != "") {
			//t.Log("assert when username and password not empty")
			assert.Equal(t,err,nil)
		} else {
			//t.Log("assert when username or password are empty")
			assert.NotEqual(t,err,nil)
			//t.Errorf("error shouln't be nil when username %s password %s or uri %s are empty",tt.username, tt.password, tt.uri)
		}
	}
}