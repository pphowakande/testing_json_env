package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Testmain() {
	creds := os.Getenv("GCP_KEY")
	fmt.Println("creds : ", creds)
	data := GCSCreds{}
	err := json.Unmarshal([]byte(creds), &data)
	fmt.Println("err : ", err)
	fmt.Println("data : ", data)
}
