package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type GCSCreds struct {
	Type            string `json:"type"`
	ProjectID       string `json:"project_id"`
	ProjectKeyID    string `json:"private_key_id"`
	PrivateKey      string `json:"private_key"`
	ClientEmail     string `json:"client_email"`
	ClientID        string `json:"client_id"`
	AuthURI         string `json:"auth_uri"`
	TokenURI        string `json:"token_uri"`
	AuthProviderURI string `json:"auth_provider_x509_cert_url"`
	ClientCertURI   string `json:"client_x509_cert_url"`
	UniverseDomain  string `json:"universe_domain"`
}

func testfunc() bool {
	creds := os.Getenv("GCP_KEY")
	fmt.Println("creds : ", creds)
	data := GCSCreds{}
	err := json.Unmarshal([]byte(creds), &data)
	if err != nil {
		return false
	}
	fmt.Println("err : ", err)
	fmt.Println("data : ", data)
	return true
}
