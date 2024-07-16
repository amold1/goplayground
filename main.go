package main

import "log"

func main() {
	// v8GetRecord()
	if err := setRecord("test5.akafn.com", "A", "1.1.1.3"); err != nil {
		log.Fatal(err)
	}
	// getRecord("test.akafn.com")
	// GetKey()
	// getAllRecordSets()
	// interactWithLinodeGo()
}
