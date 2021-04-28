package connection

// Example yml file //
// address: 192.0.2.12
// userid:
//   username: toto
//   idrealm: pve
// apitoken:
//   id: prometheus
//   token: AAAAABBBBBCCCCCDDDDD

//UserID User name and realm
type UserID struct {
	Username string `yaml:"username"`
	IDRealm  string `yaml:"idrealm"`
}

//APIToken Token ID and actual token
type APIToken struct {
	ID    string `yaml:"id"`
	Token string `yaml:"token"`
}

//Info token-based API access information
type Info struct {
	Address  string   `yaml:"apiaddress"`
	UserID   UserID   `yaml:"userid"`
	APIToken APIToken `yaml:"apitoken"`
}
