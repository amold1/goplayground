package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"sync"

	edgedns "github.com/akamai/AkamaiOPEN-edgegrid-golang/configdns-v2"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

var (
	libraryVersion = "v1.2.2"
	// UserAgent is the User-Agent value sent for all requests
	UserAgent = "Akamai-Open-Edgegrid-golang/" + libraryVersion + " golang/" + strings.TrimPrefix(runtime.Version(), "go")
	// Client is the *http.Client to use
	Client = http.DefaultClient
	// Mutex Lock
	reqLock sync.Mutex
	// Zone
	zone = "akafn.com"
	// Zones URL
	zonesURL = "config-dns/v2/zones"
)

type RecordSet struct {
	Name  string   `json:"name"`
	Type  string   `json:"type"`
	TTL   int      `json:"ttl"`
	RData []string `json:"rdata"`
}

type Metadata struct {
	ShowAll       bool `json:"showAll"`
	LastPage      int  `json:"lastPage"`
	PageSize      int  `json:"pageSize"`
	Page          int  `json:"page"`
	TotalElements int  `json:"totalElements"`
}

type GetResponseInterface struct {
	Metadata   Metadata    `json:"metadata"`
	Recordsets []RecordSet `json:"recordsets"`
}

func getRecordSets() {

	// Set config using edgerc or env var
	config, err := edgegrid.InitEnv("")
	if err != nil {
		log.Fatal(err)
	}

	// create new request for given URL
	req, err := NewRequest(config, "GET", fmt.Sprintf("https://%s/%s/%s/recordsets", config.Host, zonesURL, zone), nil)
	if err != nil {
		log.Fatal(err)
	}
	// Add auth header
	req = edgegrid.AddRequestHeader(config, req)

	edgedns.Config = edgegrid.Config{
		Host:         "akab-pvi2kxobnjna6cf2-7to2refrsunbzavl.luna.akamaiapis.net",
		AccessToken:  "akab-ea7bv2een2syz5cc-yxhc2vd4e4dmykcx",
		ClientToken:  "akab-gepihiqucit4p2ls-7l4vpvu2mrbgtiky",
		ClientSecret: "WULnDE9zAFsOunIgSIDdUn10C7Tuwom9P8Q04/ovjp0=",
	}
	zones, _ := edgedns.ListZones()
	fmt.Printf("abir \n\n %s", zones)

	// // Send the request to the API
	// resp, err := Client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Convert response to byte and print response
	// byt, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(byt))

	// // Use Akamai EdgeGrid Client method BodyJSON to unmarshal response to passed stuct and print the struct
	// jsonStruct := GetResponseInterface{}
	// client.BodyJSON(resp, jsonStruct)
	// fmt.Println(&jsonStruct)

}

// NewRequest creates an HTTP request that can be sent to Akamai APIs. A relative URL can be provided in path, which will be resolved to the
// Host specified in Config. If body is specified, it will be sent as the request body.
func NewRequest(config edgegrid.Config, method, path string, body io.Reader) (*http.Request, error) {
	var (
		baseURL *url.URL
		err     error
	)

	reqLock.Lock()
	defer reqLock.Unlock()

	if strings.HasPrefix(config.Host, "https://") {
		baseURL, err = url.Parse(config.Host)
	} else {
		baseURL, err = url.Parse("https://" + config.Host)
	}

	if err != nil {
		return nil, err
	}

	rel, err := url.Parse(strings.TrimPrefix(path, "/"))
	if err != nil {
		return nil, err
	}

	u := baseURL.ResolveReference(rel)
	if config.AccountKey != "" {
		q := u.Query()
		q.Add("accountSwitchKey", config.AccountKey)
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", UserAgent)
	// req.Header.Add("Content-Type", "application/json")

	return req, nil
}
