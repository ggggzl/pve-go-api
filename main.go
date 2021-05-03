package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/Funkit/pve-go-api/api"
	"github.com/Funkit/pve-go-api/connection"
	"github.com/Funkit/pve-go-api/utils"
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

	//resp, err := pveClient.GetRawResponse("/nodes/VMSRV02/qemu/280/firewall/rules")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(resp)

	resList, err := pveClient.GetClusterResources()
	if err != nil {
		panic(err)
	}

	vm, err := utils.GetVM(resList, 280)
	if err != nil {
		panic(err)
	}

	fmt.Println(vm)

	server1, err := utils.GetNode(resList, "VMSRV09")
	if err != nil {
		panic(err)
	}

	fmt.Println(server1)
}
