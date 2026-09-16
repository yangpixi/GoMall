package id

import (
	"github.com/bwmarrin/snowflake"
	"github.com/yangpixi/GoMall/services/user-service/internal/domain/address"
)

type SnowflakeGenerator struct {
	node *snowflake.Node
}

func NewGenerator(nodeID int64) (address.IDGenerator, error) {
	node, err := snowflake.NewNode(nodeID)

	return &SnowflakeGenerator{node: node}, err
}

func (g *SnowflakeGenerator) NextID() int64 {
	return int64(g.node.Generate())
}
