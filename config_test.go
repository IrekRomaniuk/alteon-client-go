package alteon

import (
    "fmt"
	"testing"
	"encoding/json"
	"os"
	"math/rand"
	"time"
	"net/http/httptest"
	"net/http"
	"flag"
)

const (
	ok = `{"status":"ok"}`
	ipaddr  = "10.2.3.4"
	ipaddrUpdate  = "10.222.111.123"
	table  = "SlbNewCfgEnhRealServerTable"
	weight = 1
	timeout = 2
	state = 3
)

var (
	username,password, uri,name string
)

var online = flag.Bool("online", false, "only perform mocktests")

func init() {
    username  = os.Getenv("ALTEON_USERNAME")
	password = os.Getenv("ALTEON_PASSWORD")
	uri  = os.Getenv("ALTEON_URI")
	rand.Seed(time.Now().UnixNano())
	name  = randomString(10)	
}

func TestMain(m *testing.M) {
	//Do stuff BEFORE the tests!
	if username == "" || password=="" || uri=="" {
		fmt.Printf("env vraibles not set: uri=%s, username=%s, password=%s\n",uri, username,password)
		os.Exit(1)
	}
	exitVal := m.Run()
	//Do stuff AFTER the tests!
    os.Exit(exitVal)
}

func TestCreateItem(t *testing.T) {
	srv, url := alteonMock(table, name)
	if srv != nil {
		fmt.Println("Starting mock server")
		defer srv.Close()
	}
	c1, err := NewClient(&url, &username, &password)
	if err != nil {
		t.Error(err)
	}
	c1.HostURL= url
	items := &RealServerItem{IpAddr:ipaddr, Name:name, Weight:weight, TimeOut:timeout, State:state}
	helper, err :=json.MarshalIndent(items, "", "    ")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(string(helper))
	response, err := c1.CreateItem(helper, table, name)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("Server %s created: %s\n", name, string(response.Status))
	}
}

func TestGetItem(t *testing.T) {
	srv, url := alteonMock(table, name)
	if srv != nil {
		fmt.Println("Starting mock server")
		defer srv.Close()
	}
	c1, err := NewClient(&url, &username, &password)
	if err != nil {
		t.Error(err)
	}
	realServer, err := c1.GetItem(table, name)
	if err != nil {
		t.Error(err)
	}
	items := realServer[table]
	helper, err  := json.Marshal(items)
	if err != nil {
		t.Error(err)
	}	
	var item []RealServerItem
	json.Unmarshal(helper, &item)
	if err != nil {
		t.Error(err)
	}	else {
		for _, r := range item {
			if r.Index != name {
				t.Errorf("%s  is NOT %s",r.Index, name)
			}
		}
		
	}
}

func TestUpdateItem(t *testing.T) {
	srv, url := alteonMock(table, name)
	if srv != nil {
		fmt.Println("Starting mock server")
		defer srv.Close()
	}
	c1, err := NewClient(&url, &username, &password)
	if err != nil {
		t.Error(err)
	}
	items := &RealServerItem{IpAddr:ipaddrUpdate, Name:name, Weight:weight, TimeOut:timeout, State:state}
	helper, err :=json.MarshalIndent(items, "", "    ")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(string(helper))
	response, err := c1.UpdateItem(helper, table, name)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("Server %s updated: %s\n", name, string(response.Status))
	}
}

func TestDeleteItem(t *testing.T) {
	srv, url := alteonMock(table, name)
	if srv != nil {
		fmt.Println("Starting mock server")
		defer srv.Close()
	}
	c1, err := NewClient(&url, &username, &password)
	if err != nil {
		t.Error(err)
	}
	err = c1.DeleteItem(table, name)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("Server %s deleted\n", name)
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}

func alteonMock(t,n string) (srv *httptest.Server, url string) {
	if !*online {
		handler := http.NewServeMux()
	//fmt.Printf("%s/%s\n",t,n)
	handler.HandleFunc(fmt.Sprintf("/%s/%s", t, n), responseMock)
	srv = httptest.NewServer(handler)
	url = srv.URL
	} else {
		srv = nil
		url = uri
	}
	return srv, url
}

func responseMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(ok))
}