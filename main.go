package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/Funkit/pve-go-api/api"
	"github.com/Funkit/pve-go-api/connection"
)

const secretsFilePath = "../secrets/secrets.yml"

func main() {
	tokenInfo, err := connection.ReadFile(secretsFilePath)
	if err != nil {
		panic(err)
	}

	pveClient := api.NewClient(*tokenInfo, &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2: true,
	})

	resources, err := pveClient.GetNodeNetwork("VMSRV01")
	if err != nil {
		panic(err)
	}

	fmt.Println(resources)

	res, err := pveClient.GetRawResponse("/cluster/resources")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
