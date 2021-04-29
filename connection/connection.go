package connection

import "github.com/Funkit/pve-go-api/common"

//Info token-based API access information
type Info struct {
	Address  string   `yaml:"apiaddress"`
	UserID   UserID   `yaml:"userid"`
	APIToken APIToken `yaml:"apitoken"`
}

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

//ReadFile returns API token-based connection info
func ReadFile(filePath string) (*Info, error) {
	var info Info

	if err := common.GetInfo(filePath, &info); err != nil {
		return nil, err
	}

	return &info, nil
}
