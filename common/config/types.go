package config

type (
	Config struct {
		Server  ServerConfig
		MongoDB MongoDBConfig
	}

	ServerConfig struct {
		Host string
		Port string
	}

	MongoDBConfig struct {
		Uri          string
		DatabaseName string
		Collection   string
	}
)
