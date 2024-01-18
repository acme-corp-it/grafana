//go:build wireinject && oss
// +build wireinject,oss

package datasource

import (
	"github.com/google/wire"
)

func InitializeAPIServer(group string) (*DataSourceAPIBuilder, error) {
	wire.Build(wireExtsDataSourceApiServerSet)
	return &DataSourceAPIBuilder{}, nil
}
