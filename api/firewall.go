package api

//FirewallRuleVM VM firewall rule following iptables-based firewall service
type FirewallRuleVM struct {
	//Mandatory
	Action   string `json:"action"` //ACCEPT, DROP, REJECT or security group
	VMID     string `json:"vmid"`
	RuleType string `json:"type"` //in, out or group
	//Optional
	Comment          string `json:"comment"`
	Destination      string `json:"dest"`
	DestinationPort  string `json:"dport"`
	Enable           string `json:"enable"`
	NetworkInterface string `json:"iface"`
	LogLevel         string `json:"log"` // emerg, alert, crit, err, warning, notice, info, debug or nolog
	Position         int    `json:"pos"`
	Protocol         string `json:"prot"`
	Source           string `json:"source"`
	SourcePort       string `json:"sport"`
}
