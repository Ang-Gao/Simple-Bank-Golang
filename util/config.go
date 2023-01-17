package util

import "github.com/spf13/viper"

//Config contains all configuration of the application

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	ServerAddr string `mapstructure:"SERVER_ADDR"`
}

//The name of retval has been defined, so we don't need to 'return ..'.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //yaml xml env...

	//AutomaticEnv makes Viper check if 'environment variables' match any of the existing keys (config, default or flags).
	//If matching env vars are found, they are loaded into Viper.
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
