package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func testmain(t *testing.T) {
	fmt.Println("Inside tests")
	creds := os.Getenv("GCP_KEY")
	fmt.Println("creds : ", creds)
	data := GCSCreds{}
	err := json.Unmarshal([]byte(creds), &data)
	fmt.Println("err : ", err)
	fmt.Println("data : ", data)
}
