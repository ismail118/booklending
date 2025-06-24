package util

type Config struct {
	DbSource   string `mapstructure:"DB_SOURCE"`
	AddrServer string `mapstructure:"ADDR_SERVER"`
	SecretKey  string `mapstructure:"SECRET_KEY"`
}
