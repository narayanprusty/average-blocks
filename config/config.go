package config

import (
	"os"
)

type Configuration struct {
	BeaconURL string `mapstructure:"beacon-url"`
}

var Config = Configuration{
	BeaconURL: "http://3.34.47.236:3500",
}

func init() {
	beaconURL := os.Getenv("BEACON_URL")
	if beaconURL != "" {
		Config.BeaconURL = beaconURL
	}
}
