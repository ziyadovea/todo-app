package apiserver

import "github.com/spf13/viper"

func InitConfig() error {
	viper.SetConfigName("apiserver") // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("configs")   // path to look for the config file in
	return viper.ReadInConfig()      // Find and read the config file
}
