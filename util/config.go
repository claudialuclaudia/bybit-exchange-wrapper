package util

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config type stores all the configuration needed to run the app
// the values are read by Viper from a config file or .env variables
type Config struct {

	// Binance
	BinanceApiKey    string `mapstructure:"BINANCE_API_KEY"`
	BinanceSecretKey string `mapstructure:"BINANCE_SECRET_KEY"`

	// Ftx
	FtxApiKey    string `mapstructure:"FTX_API_KEY"`
	FtxSecretKey string `mapstructure:"FTX_SECRET_KEY"`

	// ADDRESSES
	PawoVaultAddress string `mapstructure:"PAWO_VAULT_ADDRESS"`
}

// Runtime config variable
var runtimeConfig_ *Config

// LoadConfig reads configuration from file or environment variables
func LoadConfig(env, path string) (config *Config, err error) {
	viper.AddConfigPath(path)  // Set config path
	viper.SetConfigType("env") // Look for specific type
	viper.SetConfigName(env)   // Register config file name (no extension)
	viper.AutomaticEnv()       // override vars if found in env file
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	runtimeConfig_ = config
	return
}

// Get configs
func GetConfigs() *Config {
	return runtimeConfig_
}
