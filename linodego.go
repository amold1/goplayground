package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/linode/linodego"
	"golang.org/x/oauth2"
)

func interactWithLinodeGo() {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "805eb2a8b9cb9929f8a92c326f058d12e9acb9b77269532be1a463febe990aca"})

	oauth2Client := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
		},
	}

	linodeClient := linodego.NewClient(oauth2Client)
	linodeClient.SetDebug(true)

	domain := "lkedevs.net"
	var domainID int
	filter, err := json.Marshal(map[string]string{"domain": domain})

	domains, err := linodeClient.ListDomains(context.Background(), linodego.NewListOptions(0, string(filter)))
	if err != nil {
		log.Fatal(err)
	}
	if len(domains) == 0 {
		return
	}

	domainID = domains[0].ID
	fmt.Println("domains %s", domains[0])
	domainHostname := "test-a1b2c3"

	recordReq := linodego.DomainRecordCreateOptions{
		Type:   "A",
		Name:   domainHostname,
		Target: "10.10.10.10",
		TTLSec: 30,
	}

	if _, err := linodeClient.CreateDomainRecord(context.Background(), domainID, recordReq); err != nil {
		log.Fatal(err)
	}

	filter1, err := json.Marshal(map[string]interface{}{"name": domainHostname})
	if err != nil {
		log.Fatal(err)
	}

	domainRecords1, err := linodeClient.ListDomainRecords(context.Background(), domainID, linodego.NewListOptions(0, string(filter1)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("domain record is", domainRecords1)

	recordReqUpdate := linodego.DomainRecordUpdateOptions{
		Type:   "A",
		Name:   domainHostname,
		Target: "10.10.10.11",
		TTLSec: 30,
	}

	if _, err := linodeClient.UpdateDomainRecord(context.Background(), domainID, domainRecords1[0].ID, recordReqUpdate); err != nil {
		log.Fatal(err)
	}

	filter2, err := json.Marshal(map[string]interface{}{"name": domainHostname, "type": "TXT", "target": "owner:amol"})
	if err != nil {
		log.Fatal(err)
	}

	domainRecords2, err := linodeClient.ListDomainRecords(context.Background(), domainID, linodego.NewListOptions(0, string(filter2)))
	if err != nil {
		log.Fatal(err)
	}
	if domainRecords2 == nil {
		fmt.Println("goddammit")
	} else {
		fmt.Println("domain record is", domainRecords2[0])
	}

}
