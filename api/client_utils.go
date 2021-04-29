package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Funkit/pve-go-api/connection"
)

//getData returns raw JSON data for structure-specific unmarshalling
func (c *Client) getData(url string) ([]*json.RawMessage, error) {
	respBody, err := c.get(url)
	if err != nil {
		return nil, err
	}

	return extractData(respBody)
}

//get returns the query answer body as a byte array
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

//newRequest build a request with the token-based Authorization header
func newRequest(c connection.Info, targetURL string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, c.Address+targetURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", tokenHeader(c))
	return req, nil
}

func tokenHeader(c connection.Info) string {
	return "PVEAPIToken=" + c.UserID.Username + "@" + c.UserID.IDRealm + "!" + c.APIToken.ID + "=" + c.APIToken.Token
}

//extractData extract the "data" field from the Proxmox API query response and return raw JSON field for structure-specific unmarshalling
func extractData(b []byte) ([]*json.RawMessage, error) {
	var rawContent map[string]*json.RawMessage
	var rawData []*json.RawMessage

	if err := json.Unmarshal(b, &rawContent); err != nil {
		return nil, err
	}
	for key, value := range rawContent {
		if key == "data" {
			if err := json.Unmarshal(*value, &rawData); err != nil {
				return nil, err
			}
		}
	}

	return rawData, nil
}
