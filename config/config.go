package config

import "github.com/spf13/viper"

type Configuration struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
	EXAMPLE_PATH string
	EXAMPLE_VAR  string
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func GetConfig() (*Configuration, error) {
	// Provide the config.yml file to Viper
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	// Add environment variables to Viper
	viper.AutomaticEnv()

	// Allows Viper to reload the config.yml at run time
	viper.WatchConfig()

	// Read config.yml
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Set defaults
	viper.SetDefault("database.server.port", "8080")

	// Bind config.yml to model
	configuration := &Configuration{}
	err := viper.Unmarshal(configuration)
	return configuration, err
}
