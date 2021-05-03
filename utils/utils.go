package utils

import (
	"errors"

	"github.com/Funkit/pve-go-api/api"
)

//GetVM From a resource list obtained through GetClusterResource, return the pointer to the VM resource with a matching VMID
func GetVM(resList []api.Resource, vmid int) (*api.Resource, error) {
	for i := range resList {
		if resList[i].VMID == vmid {
			return &resList[i], nil
		}
	}
	return nil, errors.New("VM not found in the resource list")
}

//GetNode From a resource list obtained through GetClusterResource, return the pointer to the node resource with a matching name
func GetNode(resList []api.Resource, nodeName string) (*api.Resource, error) {
	for i := range resList {
		if resList[i].Type == api.TypeNode && resList[i].Node == nodeName {
			return &resList[i], nil
		}
	}
	return nil, errors.New("Node not found in the resource list")
}
