package config

type Configuration struct {
	BeaconURL        string `mapstructure:"beacon-url"`
	DatabaseUsername string `mapstructure:"database-username"`
	DatabasePassword string `mapstructure:"database-password"`
	DatabaseHost     string `mapstructure:"database-host"`
	DatabasePort     string `mapstructure:"database-port"`
	DatabaseName     string `mapstructure:"database-name"`
	JWTSecret        string `mapstructure:"jwt-secret"`
}

var Config = Configuration{
	BeaconURL:        "http://3.34.47.236:3500",
	DatabaseUsername: "test",
	DatabasePassword: "test",
	DatabaseHost:     "localhost",
	DatabasePort:     "5455",
	DatabaseName:     "tracker",
	JWTSecret:        "secret",
}
