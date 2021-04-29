package api

import "encoding/json"

//Node Node resource type at /nodes
type Node struct {
	Name           string `json:"node,omitempty"`
	Status         string `json:"status"`
	SslFingerprint string `json:"ssl_fingerprint,omitempty"`
}

//NodeNetworkInterface Network interface for a given Node at /nodes/<node name>/network
type NodeNetworkInterface struct {
	Name          string   `json:"iface"`
	InterfaceType string   `json:"type"`
	Active        int      `json:"active"`
	IPAddress     string   `json:"address"`
	Gateway       string   `json:"gateway"`
	Autostart     int      `json:"autostart"`
	BridgePorts   string   `json:"bridge_ports"`
	CIDR          string   `json:"cidr"`
	Families      []string `json:"families"`
	Options       []string `json:"options"`
}

//ParseMap parse generic map into the object
func (node *NodeNetworkInterface) ParseMap(element map[string]interface{}) error {
	jsonbody, err := json.Marshal(element)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(jsonbody), &node); err != nil {
		return err
	}
	return nil
}

//Resource resource at /cluster/resources
type Resource struct {
	Name              string  `json:"name,omitempty"`
	Node              string  `json:"node"`
	ID                string  `json:"id"`
	Status            string  `json:"status"`
	CPU               float64 `json:"cpu"` // %
	AllocatedCPU      float64 `json:"maxcpu"`
	AllocatedRAMBytes int     `json:"maxmem"` // bytes
	RAM               int     `json:"mem"`    // bytes
	Uptime            int     `json:"uptime"` // sec
	VMID              int     `json:"vmid,omitempty"`
	Pool              string  `json:"pool,omitempty"`
	Template          int     `json:"template,omitempty"`
}
