package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/Funkit/pve-go-api/api"
	"github.com/Funkit/pve-go-api/common"
	"github.com/Funkit/pve-go-api/connection"
)

const secretsFilePath = "../secrets.yml"

func main() {
	var tokenInfo connection.Info

	if err := common.GetInfo(secretsFilePath, &tokenInfo); err != nil {
		panic(err)
	}

	pveClient := api.NewClient(tokenInfo, &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2: true,
	})

	rawResults, err := pveClient.GetRawResponse("/nodes")
	if err != nil {
		panic(err)
	}

	fmt.Println(rawResults)
}
