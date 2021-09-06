package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kogisin/maphack/config"
)

func TestReadConfigFile(t *testing.T) {
	configFilePath := "../config.toml"

	cfg, err := config.Read(configFilePath)
	require.NoError(t, err)

	require.Equal(t, "http://localhost:26657", cfg.RPC.Address)
	require.Equal(t, "", cfg.RPC.Token)
	require.Equal(t, "localhost:9090", cfg.GRPC.Address)
	require.Equal(t, "", cfg.GRPC.Token)
	require.Equal(t, true, cfg.GRPC.Insecure)
	require.Equal(t, "http://localhost:1317", cfg.LCD.Address)
}

func TestParseConfigString(t *testing.T) {
	var sampleConfig = `
[rpc]
address = "http://localhost:26657"
token = ""

[grpc]
address = "localhost:9090"
token = ""
insecure = true

[lcd]
address = "http://localhost:1317"

[database]
host = "http://localhost"
port = "5432"
user = "root"
password = "admin123"
name = "persistence"
schema = "persistence-testnet"
`
	cfg, err := config.ParseString([]byte(sampleConfig))
	require.NoError(t, err)

	require.Equal(t, "http://localhost:26657", cfg.RPC.Address)
	require.Equal(t, "", cfg.RPC.Token)
	require.Equal(t, "localhost:9090", cfg.GRPC.Address)
	require.Equal(t, "", cfg.GRPC.Token)
	require.Equal(t, true, cfg.GRPC.Insecure)
	require.Equal(t, "http://localhost:1317", cfg.LCD.Address)
	require.Equal(t, "http://localhost:5432", cfg.DB.Host+":"+cfg.DB.Port)
}
