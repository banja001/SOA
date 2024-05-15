package config

import "os"

type Config struct {
	Address                   string
	StakeholderServiceAddress string
}

func GetConfig() Config {
	return Config{
		StakeholderServiceAddress: os.Getenv("STAKEHOLDERS_SERVICE_ADDRESS"),
		Address:                   os.Getenv("GATEWAY_ADDRESS"),
	}
}
