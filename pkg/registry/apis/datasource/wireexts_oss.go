//go:build wireinject && oss
// +build wireinject,oss

// This file should contain wiresets which contain OSS-specific implementations.
package datasource

import (
	"github.com/google/wire"
)

var wireExtsDataSourceApiServerSet = wire.NewSet(
	wire.Bind(new(APIServer), new(*TestDataAPIServer)),
	NewTestDataAPIServer,
	NewDataSourceAPIServer,
)
