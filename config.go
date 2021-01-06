package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	configs *viper.Viper
}

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

// Flush resets configuration and load default values
func Flush() {
	viper.Reset()
}

// WatchChanges watching and re-reading changes from config files
func WatchChanges() {
	viper.WatchConfig()
}

// Set sets default values
func Default(key string, value interface{}) {
	viper.SetDefault(key, value)
}

// Set sets value for the key
func Set(key string, value interface{}) {
	viper.Set(key, value)
}

// IsSet checks to see if the key has been set in any of the data locations.
func IsSet(key string) bool             { return viper.IsSet(key) }
func (c *Config) IsSet(key string) bool { return c.configs.IsSet(key) }

// Get returns an interface
func Get(key string) interface{}             { return viper.Get(key) }
func (c *Config) Get(key string) interface{} { return c.configs.Get(key) }

// GetString returns the value associated with the key as a string.
func GetString(key string) string             { return viper.GetString(key) }
func (c *Config) GetString(key string) string { return c.configs.GetString(key) }

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool             { return viper.GetBool(key) }
func (c *Config) GetBool(key string) bool { return c.configs.GetBool(key) }

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int             { return viper.GetInt(key) }
func (c *Config) GetInt(key string) int { return c.configs.GetInt(key) }

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) int32             { return viper.GetInt32(key) }
func (c *Config) GetInt32(key string) int32 { return c.configs.GetInt32(key) }

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64             { return viper.GetInt64(key) }
func (c *Config) GetInt64(key string) int64 { return c.configs.GetInt64(key) }

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) uint             { return viper.GetUint(key) }
func (c *Config) GetUint(key string) uint { return c.configs.GetUint(key) }

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) uint32             { return viper.GetUint32(key) }
func (c *Config) GetUint32(key string) uint32 { return c.configs.GetUint32(key) }

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) uint64             { return viper.GetUint64(key) }
func (c *Config) GetUint64(key string) uint64 { return c.configs.GetUint64(key) }

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64             { return viper.GetFloat64(key) }
func (c *Config) GetFloat64(key string) float64 { return c.configs.GetFloat64(key) }

// GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time             { return viper.GetTime(key) }
func (c *Config) GetTime(key string) time.Time { return c.configs.GetTime(key) }

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration             { return viper.GetDuration(key) }
func (c *Config) GetDuration(key string) time.Duration { return c.configs.GetDuration(key) }

// GetIntSlice returns the value associated with the key as a slice of int values.
func GetIntSlice(key string) []int             { return viper.GetIntSlice(key) }
func (c *Config) GetIntSlice(key string) []int { return c.configs.GetIntSlice(key) }

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key string) []string             { return viper.GetStringSlice(key) }
func (c *Config) GetStringSlice(key string) []string { return c.configs.GetStringSlice(key) }

// GetStringMap returns the value associated with the key as a map of interfaces.
func GetStringMap(key string) map[string]interface{}             { return viper.GetStringMap(key) }
func (c *Config) GetStringMap(key string) map[string]interface{} { return c.configs.GetStringMap(key) }

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) map[string]string             { return viper.GetStringMapString(key) }
func (c *Config) GetStringMapString(key string) map[string]string { return c.configs.GetStringMapString(key) }

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) map[string][]string             { return viper.GetStringMapStringSlice(key) }
func (c *Config) GetStringMapStringSlice(key string) map[string][]string { return c.configs.GetStringMapStringSlice(key) }

// GetSizeInBytes returns the size of the value associated with the given key in bytes.
func GetSizeInBytes(key string) uint             { return viper.GetSizeInBytes(key) }
func (c *Config) GetSizeInBytes(key string) uint { return c.configs.GetSizeInBytes(key) }

// GetAll returns all keys defined in config
func GetAll() interface{}                       { return viper.AllKeys() }
func (c *Config) GetAll(key string) interface{} { return c.configs.AllKeys() }

// GetNested returns all keys under specific key
func GetNested(key string) *Config {
	nested := viper.Sub(key)
	if nested == nil {
		return nil
	}

	return &Config{configs: nested}
}
func (c *Config) GetNested(key string) *Config {
	nested := c.configs.Sub(key)
	if nested == nil {
		return nil
	}

	return &Config{configs: nested}
}
