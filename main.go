package main

import "log"

func main() {
	// v8GetRecord()
	if err := setRecord("test.akafn.com", "A", "1.1.1.4"); err != nil {
		log.Fatal(err)
	}
	// getRecord("test.akafn.com")
	// GetKey()
	// getAllRecordSets()
	// interactWithLinodeGo()
}
