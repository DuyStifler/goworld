package main

import (
	"../../../goworld"
	"../../../goworld/engine/entity"
)

// AOITester type
type AOITester struct {
	goworld.Entity // Entity type should always inherit entity.Entity
}

func (m *AOITester) DescribeEntityType(desc *entity.EntityTypeDesc) {
	desc.SetUseAOI(true, 100)
	desc.DefineAttr("name", "AllClients")
}
