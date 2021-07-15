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

	//resources, err := pveClient.GetNodeNetwork("VMSRV01")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resources)

	//resList, err := pveClient.GetClusterResources()
	//if err != nil {
	//	panic(err)
	//}
	//for _, res := range resList {
	//	if res.Type == "qemu" {
	//		fmt.Println(res.Name)
	//	}
	//}

	result, err := pveClient.GetRawResponse("/nodes/VMSRV05/qemu/580/firewall/rules")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
