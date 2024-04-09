package top_level_domain

import (
	_ "embed"
	"encoding/json"
)

//go:embed top_level_domain.json
var data []byte

type domainTree map[string]*domainTree

var TopLevelDomainTree domainTree

func init() {
	err := json.Unmarshal(data, &TopLevelDomainTree)
	if err != nil {
		panic(err)
	}
}
