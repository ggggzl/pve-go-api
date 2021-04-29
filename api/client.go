package api

import (
	"encoding/json"
	"net/http"

	"github.com/Funkit/pve-go-api/connection"
)

// Client http.Client and connection information
type Client struct {
	httpClient *http.Client
	info       connection.Info
}

//NewClient create new client with TLS check disabled and information to log with a token to the API
func NewClient(i connection.Info, transportSettings *http.Transport) *Client {
	httpClient := &http.Client{Transport: transportSettings}

	client := &Client{
		httpClient: httpClient,
		info:       i,
	}

	return client
}

//GetRawResponse query the API endpoint and return a map containing the response body
func (c *Client) GetRawResponse(url string) (interface{}, error) {
	respBody, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

//GetNodes query the /nodes URL on the Proxmox API
func (c *Client) GetNodes() ([]Node, error) {
	rawData, err := c.getData("/nodes")
	if err != nil {
		return nil, err
	}

	var nodeList []Node
	for _, element := range rawData {
		var node Node
		if err = json.Unmarshal(*element, &node); err != nil {
			return nil, err
		}
		nodeList = append(nodeList, node)
	}
	return nodeList, err
}

//GetClusterResources query the /cluster/resources URL on the Proxmox API
func (c *Client) GetClusterResources() ([]Resource, error) {
	rawData, err := c.getData("/cluster/resources")
	if err != nil {
		return nil, err
	}

	var resList []Resource
	for _, element := range rawData {
		var res Resource
		if err = json.Unmarshal(*element, &res); err != nil {
			return nil, err
		}
		resList = append(resList, res)
	}
	return resList, err
}

//GetNodeNetwork query the /nodes/<node name>/network URL on the Proxmox API
func (c *Client) GetNodeNetwork(nodeName string) ([]NodeNetworkInterface, error) {

	rawData, err := c.getData("/nodes/" + nodeName + "/network")
	if err != nil {
		return nil, err
	}

	var resList []NodeNetworkInterface
	for _, element := range rawData {
		var iface NodeNetworkInterface
		if err = json.Unmarshal(*element, &iface); err != nil {
			return nil, err
		}
		resList = append(resList, iface)
	}

	return resList, err
}
