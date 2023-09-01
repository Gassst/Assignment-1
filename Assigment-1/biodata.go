package main

import (
	"Assigment-1/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func biodata(id string) {
	jsonPath := filepath.Join("data", "participants.json")

	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	var participants model.Participants

	if err := json.Unmarshal(data, &participants); err != nil {
		log.Fatal(err)
	}

	for _, person := range participants.Participants {
		if person.ID == id {
			fmt.Printf("Id: %s\n", person.ID)
			fmt.Printf("Student_Code: %s\n", person.StudentCode)
			fmt.Printf("Student_Name: %s\n", person.StudentName)
			fmt.Printf("Student_Address: %s\n", person.StudentAddress)
			fmt.Printf("Student_Occupation: %s\n", person.StudentOccupation)
			fmt.Printf("Joining_Reason: %s\n", person.JoiningReason)
			return
		}
	}

	fmt.Printf("No data found for ID: %s\n", id)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go <id>")
		return
	}

	id := os.Args[1]
	biodata(id)
}
