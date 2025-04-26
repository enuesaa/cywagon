package model

import (
	"maps"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

// The structs in this file are for configuration parsing only.
// DO NOT use them directly.

type PartialConstsConfig struct {
	Consts []Const  `hcl:"const,block"`
	Remain hcl.Body `hcl:",remain"`
}

func (p *PartialConstsConfig) FlattenConsts() cty.Value {
	merged := make(map[string]cty.Value)
	for _, co := range p.Consts {
		maps.Copy(merged, co.Attrs)
	}
	return cty.ObjectVal(merged)
}

type LogicNameOnly struct {
	Name   string   `hcl:"name,label"`
	Remain hcl.Body `hcl:",remain"`
}

type PartialLogicNameOnlyConfig struct {
	Logics []LogicNameOnly `hcl:"logic,block"`
	Remain hcl.Body        `hcl:",remain"`
}

func (p *PartialLogicNameOnlyConfig) FlattenLogicNames() cty.Value {
	merged := make(map[string]cty.Value)
	for _, co := range p.Logics {
		merged[co.Name] = cty.StringVal(co.Name)
	}
	return cty.ObjectVal(merged)
}
