package factory

import (
	"testing"

	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/config"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/core"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/testscommon"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/testscommon/middleware"
	"github.com/TerraDharitri/drt-go-sdk/authentication/native/mock"
	"github.com/stretchr/testify/assert"
)

func TestStartWebServer(t *testing.T) {
	t.Parallel()

	cfg := config.Configs{
		GeneralConfig: config.Config{
			Guardian: config.GuardianConfig{
				MnemonicFile:         "testdata/dharitri.mnemonic",
				RequestTimeInSeconds: 2,
			},
			Logs:      config.LogsConfig{},
			Antiflood: config.AntifloodConfig{},
		},
		ApiRoutesConfig: config.ApiRoutesConfig{},
		FlagsConfig: config.ContextFlagsConfig{
			RestApiInterface: core.WebServerOffString,
		},
	}

	webServer, err := StartWebServer(
		cfg,
		&testscommon.ServiceResolverStub{},
		&mock.AuthServerStub{},
		&mock.AuthTokenHandlerStub{},
		&middleware.NativeAuthWhitelistHandlerStub{},
		&testscommon.StatusMetricsStub{},
	)
	assert.Nil(t, err)
	assert.NotNil(t, webServer)

	err = webServer.Close()
	assert.Nil(t, err)
}
