package api

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Client http.Client and connection information
type Client struct {
	httpClient *http.Client
	info       Info
}

//NewClient create new client with TLS check disabled and information to log with a token to the API
func NewClient(c Info) *Client {
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2: true,
	}
	httpClient := &http.Client{Transport: tr}

	client := &Client{
		httpClient: httpClient,
		info:       c,
	}

	return client
}

func tokenHeader(c Info) string {
	return "PVEAPIToken=" + c.UserID.Username + "@" + c.UserID.IDRealm + "!" + c.APIToken.ID + "=" + c.APIToken.Token
}

func newRequest(c Info, targetURL string) (*http.Request, error) {
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
	respBody, err := c.get("/nodes")
	if err != nil {
		return nil, err
	}

	nodeList, err := parseNodes(respBody)
	if err != nil {
		return nil, err
	}

	return nodeList, nil
}

//GetClusterResources query the /cluster/resources URL on the Proxmox API
func (c *Client) GetClusterResources() ([]NodeResource, []VMResource, error) {
	respBody, err := c.get("/cluster/resources")
	if err != nil {
		return nil, nil, err
	}

	nodeList, vmList, err := parseClusterResources(respBody)
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
