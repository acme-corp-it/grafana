package datasource

import (
	"context"
	"fmt"

	"github.com/grafana/grafana/pkg/services/accesscontrol/actest"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	common "github.com/grafana/grafana/pkg/apis/common/v0alpha1"
	"github.com/grafana/grafana/pkg/plugins"
	"github.com/grafana/grafana/pkg/setting"
	testdatasource "github.com/grafana/grafana/pkg/tsdb/grafana-testdata-datasource"
)

func NewDataSourceAPIServer(s APIServer) (*DataSourceAPIBuilder, error) {
	return NewDataSourceAPIBuilder(
		s.PluginJSON(),
		s.QuerierProvider(),
		s.PluginContextProvider(),
		&actest.FakeAccessControl{ExpectedEvaluate: true},
	)
}

type APIServer interface {
	QuerierProvider() QuerierProvider
	PluginContextProvider() PluginContextProvider
	PluginJSON() plugins.JSONData
}

type TestDataAPIServer struct {
	querierProvider       QuerierProvider
	pluginContextProvider PluginContextProvider
	pluginJSON            plugins.JSONData
}

func NewTestDataAPIServer(group string) (*TestDataAPIServer, error) {
	pluginID := "grafana-testdata-datasource"

	if group != "testdata.datasource.grafana.app" {
		return nil, fmt.Errorf("only %s is currently supported", pluginID)
	}

	cfg, err := setting.NewCfgFromArgs(setting.CommandLineArgs{
		// TODO: Add support for args?
	})
	if err != nil {
		return nil, err
	}

	_, pluginStore, dsService, dsCache, err := apiBuilderServices(cfg, pluginID)
	if err != nil {
		return nil, err
	}

	td, exists := pluginStore.Plugin(context.Background(), pluginID)
	if !exists {
		return nil, fmt.Errorf("plugin %s not found", pluginID)
	}

	var testsDataQuerierFactory QuerierFactoryFunc = func(ctx context.Context, ri common.ResourceInfo, pj plugins.JSONData) (Querier, error) {
		return NewDefaultQuerier(ri, td.JSONData, testdatasource.ProvideService(), dsService, dsCache), nil
	}

	return &TestDataAPIServer{
		querierProvider:       NewQuerierProvider(testsDataQuerierFactory),
		pluginContextProvider: &TestDataPluginContextProvider{},
		pluginJSON:            td.JSONData,
	}, nil
}

func (b *TestDataAPIServer) QuerierProvider() QuerierProvider {
	return b.querierProvider
}

func (b *TestDataAPIServer) PluginContextProvider() PluginContextProvider {
	return b.pluginContextProvider
}

func (b *TestDataAPIServer) PluginJSON() plugins.JSONData {
	return b.pluginJSON
}

type TestDataPluginContextProvider struct{}

func (p *TestDataPluginContextProvider) PluginContextForDataSource(_ context.Context, _, _ string) (backend.PluginContext, error) {
	return backend.PluginContext{}, nil
}
