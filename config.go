package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var supportedFiles = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "dotenv", "env", "ini"}

// LoadFile loads configurations from file, if app has more than one config file, we can call this function multiple times.
// Supported types: json, toml, yaml, yml, properties, props, prop, hcl, dotenv, env, ini
func LoadFile(configName string, configType string, configPath string) {
	if !stringInSlice(configType, supportedFiles) {
		panic(fmt.Errorf("Configuration file not supported: %s \n", configType))
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// WatchChanges watching and re-reading changes from config files
func WatchChanges() {
	viper.WatchConfig()
}

// Set sets value for the key
func Set(key string, value interface{}) {
	viper.Set(key, value)
}

// Get returns an interface
func Get(key string) interface{} { return viper.Get(key) }

// GetString returns the value associated with the key as a string.
func GetString(key string) string { return viper.GetString(key) }

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool { return viper.GetBool(key) }

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int { return viper.GetInt(key) }

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) int32 { return viper.GetInt32(key) }

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64 { return viper.GetInt64(key) }

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) uint { return viper.GetUint(key) }

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) uint32 { return viper.GetUint32(key) }

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) uint64 { return viper.GetUint64(key) }

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64 { return viper.GetFloat64(key) }

// GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time { return viper.GetTime(key) }

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration { return viper.GetDuration(key) }

// GetIntSlice returns the value associated with the key as a slice of int values.
func GetIntSlice(key string) []int { return viper.GetIntSlice(key) }

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key string) []string { return viper.GetStringSlice(key) }

// GetStringMap returns the value associated with the key as a map of interfaces.
func GetStringMap(key string) map[string]interface{} { return viper.GetStringMap(key) }

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) map[string]string { return viper.GetStringMapString(key) }

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) map[string][]string { return viper.GetStringMapStringSlice(key) }

// GetSizeInBytes returns the size of the value associated with the given key in bytes.
func GetSizeInBytes(key string) uint { return viper.GetSizeInBytes(key) }
