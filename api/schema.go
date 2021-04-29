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
	AllocatedCPU      float64 `json:"maxcpu"`
	AllocatedRAMBytes int     `json:"maxmem"` // bytes
	CPU               float64 `json:"cpu"`    // %
	ID                string  `json:"id"`
	Name              string  `json:"name,omitempty"`
	Node              string  `json:"node"`
	Pool              string  `json:"pool,omitempty"`
	RAM               int     `json:"mem"` // bytes
	Status            string  `json:"status"`
	Template          int     `json:"template,omitempty"`
	Type              string  `json:"type,omitempty"`
	Uptime            int     `json:"uptime"` // sec
	VMID              int     `json:"vmid,omitempty"`
}
