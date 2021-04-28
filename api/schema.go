package api

import "encoding/json"

//Node Node resource type at /nodes
type Node struct {
	Name           string `json:"node,omitempty"`
	Status         string `json:"status"`
	SslFingerprint string `json:"ssl_fingerprint,omitempty"`
}

//ParseMap parse generic map into the object
func (node *Node) ParseMap(element map[string]interface{}) error {
	jsonbody, err := json.Marshal(element)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(jsonbody), &node); err != nil {
		return err
	}
	return nil
}

//VMResource VM resource type at /cluster/resources
type VMResource struct {
	Name              string  `json:"name"`
	VMID              int     `json:"vmid"`
	Pool              string  `json:"pool,omitempty"`
	Node              string  `json:"node"`
	Status            string  `json:"status,omitempty"`
	Uptime            int     `json:"uptime"`
	AllocatedCPUCores int     `json:"maxcpu"`
	CPU               float64 `json:"cpu"`    // %
	AllocatedRAMBytes int     `json:"maxmem"` // in bytes
	RAM               int     `json:"mem"`    // in bytes
	Template          int     `json:"template"`
}

//ParseMap parse generic map into the object
func (node *VMResource) ParseMap(element map[string]interface{}) error {
	jsonbody, err := json.Marshal(element)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(jsonbody), &node); err != nil {
		return err
	}
	return nil
}

//NodeResource Node resource type at /cluster/resources
type NodeResource struct {
	Node              string  `json:"node"`
	Status            string  `json:"status"`
	CPU               float64 `json:"cpu"`    // %
	AllocatedRAMBytes int     `json:"maxmem"` // in bytes
	RAM               int     `json:"mem"`    // in bytes
}

//ParseMap parse generic map into the object
func (node *NodeResource) ParseMap(element map[string]interface{}) error {
	jsonbody, err := json.Marshal(element)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(jsonbody), &node); err != nil {
		return err
	}
	return nil
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
