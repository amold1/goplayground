package main

import (
	"context"
	"log"
	"os"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v8/pkg/dns"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v8/pkg/edgegrid"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v8/pkg/session"
)

var domain = "akafn.com"

func setUpEdgeRCConfig() *edgegrid.Config {
	return &edgegrid.Config{
		Host:         os.Getenv("AKAMAI_HOST"),
		AccessToken:  os.Getenv("AKAMAI_ACCESS_TOKEN"),
		ClientToken:  os.Getenv("AKAMAI_CLIENT_TOKEN"),
		ClientSecret: os.Getenv("AKAMAI_CLIENT_SECRET"),
		MaxBody:      131072,
	}
}

func setUpDNSInterface() dns.DNS {
	var err error
	config := setUpEdgeRCConfig()
	sess, err := session.New(session.WithSigner(config))
	if err != nil {
		log.Fatal(err)
	}
	return dns.Client(sess)
}

func getRecord(recordName, recordType string) (*dns.RecordBody, error) {
	d := setUpDNSInterface()
	recordBody, err := d.GetRecord(context.Background(), domain, recordName, recordType)
	if err != nil {
		return nil, err
	}
	return recordBody, nil
}

func setRecord(recordName, recordType, ip string) error {
	d := setUpDNSInterface()
	record, err := getRecord(recordName, recordType)
	if err == nil {
		record.Target = append(record.Target, ip)
		return d.UpdateRecord(context.Background(), record, domain)
	}
	body := dns.RecordBody{
		Name:       recordName,
		RecordType: recordType,
		TTL:        30,
		Target:     []string{ip},
	}
	return d.CreateRecord(context.Background(), &body, domain)
}
