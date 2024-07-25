package main

func main() {
	// v8GetRecord()
	// if err := setRecord("test.akafn.com", "A", "1.1.1.4"); err != nil {
	// 	log.Fatal(err)
	// }
	// result, err := getRecord("dodo.akafn.com", "A")
	// if strings.Contains(err.Error(), "sundar") {
	// 	fmt.Println("Jai Gopal")
	// } else {
	// 	fmt.Println(result)
	// }
	// GetKey()
	// getAllRecordSets()
	// interactWithLinodeGo()
	var data = Cluster{
		Name:   "abir",
		IPList: []string{"ip1", "ip2", "ip3"},
	}

	fileIO(data)
}
