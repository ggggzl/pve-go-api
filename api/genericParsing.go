package api

import "encoding/json"

type GenericResult interface{}
type GenericResultBuilder interface {
	NewResultList() []GenericResult
}

func ParseResultList(responseBody []byte, builder GenericResultBuilder) (interface{}, error) {
	Results := builder.NewResultList()

	err := json.Unmarshal([]byte(responseBody), &Results)

	return Results, err
}

type NodeBuilder struct{}
type testNode struct {
	CPU string `json:"cpu"`
}

func (NodeBuilder) NewResultList() interface{} {
	var nodeList []testNode
	return nodeList
}
