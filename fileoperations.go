package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Cluster struct {
	Name   string   `json:"name"`
	IPList []string `json:"ipList"`
}

var (
	allClusters []Cluster
	foundData   = false
)

func fileIO(data Cluster) {
	byteValue, _ := os.ReadFile("log.json")
	if len(byteValue) != 0 {
		if err := json.Unmarshal(byteValue, &allClusters); err != nil {
			fmt.Println(err.Error())
		}
	}
	for index, cluster := range allClusters {
		if cluster.Name == data.Name {
			foundData = true
			allClusters[index] = data
		}
	}
	if !foundData {
		allClusters = append(allClusters, data)
	}
	dataToWrite, err := json.Marshal(allClusters)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("log.json", dataToWrite, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
