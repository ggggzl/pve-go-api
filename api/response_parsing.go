package api

import (
	"encoding/json"
)

//Results generic object to store API answers
type Results struct {
	Rows []map[string]interface{} `json:"data"`
}

func parseClusterResources(rawData []*json.RawMessage) ([]NodeResource, []VMResource, error) {

	var nodeList []NodeResource
	var vmList []VMResource

	//for _, element := range rawData {
	//	var node Node
	//	if err := json.Unmarshal(*element, &node); err != nil {
	//		return nil, nil, err
	//	}
	//	nodeList = append(nodeList, node)
	//}
	//
	//for _, row := range buffer.Rows {
	//	if row["type"] == "node" {
	//		var buffer NodeResource
	//		if err := buffer.ParseMap(row); err != nil {
	//			return nil, nil, err
	//		}
	//		nodeList = append(nodeList, buffer)
	//	}
	//	if row["type"] == "qemu" {
	//		var buffer VMResource
	//		if err := buffer.ParseMap(row); err != nil {
	//			return nil, nil, err
	//		}
	//		vmList = append(vmList, buffer)
	//	}
	//}
	return nodeList, vmList, nil
}

func parseNodeNetwork(responseBody []byte) ([]NodeNetworkInterface, error) {
	var buffer Results

	if err := json.Unmarshal([]byte(responseBody), &buffer); err != nil {
		return nil, err
	}

	var nodeList []NodeNetworkInterface
	for _, row := range buffer.Rows {
		var buffer NodeNetworkInterface
		if err := buffer.ParseMap(row); err != nil {
			return nil, err
		}
		nodeList = append(nodeList, buffer)
	}
	return nodeList, nil
}
