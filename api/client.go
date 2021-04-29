package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Funkit/pve-go-api/common"
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

func tokenHeader(c connection.Info) string {
	return "PVEAPIToken=" + c.UserID.Username + "@" + c.UserID.IDRealm + "!" + c.APIToken.ID + "=" + c.APIToken.Token
}

func newRequest(c connection.Info, targetURL string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, c.Address+targetURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", tokenHeader(c))
	return req, nil
}

func (c *Client) get(url string) (responseBody []byte, err error) {
	req, err := newRequest(c.info, url)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *Client) getData(url string) ([]*json.RawMessage, error) {
	req, err := newRequest(c.info, url)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return common.GetRawData(respBody)
}

//GetRawResponse query the API endpoint and return the response body before serialization
func (c *Client) GetRawResponse(url string) (*Results, error) {
	respBody, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var buffer Results

	if err := json.Unmarshal([]byte(respBody), &buffer); err != nil {
		return nil, err
	}

	return &buffer, nil
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
func (c *Client) GetClusterResources() ([]NodeResource, []VMResource, error) {
	rawData, err := c.getData("/cluster/resources")
	if err != nil {
		return nil, nil, err
	}

	nodeList, vmList, err := parseClusterResources(rawData)
	if err != nil {
		log.Fatal(err)
	}

	return nodeList, vmList, nil
}

//GetNodeNetwork query the /nodes/<node name>/network URL on the Proxmox API
func (c *Client) GetNodeNetwork(nodeName string) ([]NodeNetworkInterface, error) {
	respBody, err := c.get("/nodes/" + nodeName + "/network")
	if err != nil {
		return nil, err
	}

	networkList, err := parseNodeNetwork(respBody)
	if err != nil {
		log.Fatal(err)
	}

	return networkList, nil
}
