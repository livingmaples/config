## Install

```console
go get https://gitlab.com/livingmaples/packages/config
```


## What is packages/config (ConfigMan)?

ConfigMan is a configuration solution which uses [Viper](https://github.com/spf13/viper) under the hood. It is designed
to work within applications, and can handle all types of configuration needs
and formats. It supports:

* setting defaults
* reading from JSON, TOML, YAML, HCL, envfile
* live watching and re-reading of config files (optional)
* reading from environment variables
* setting explicit values

**Important:** configuration keys are case insensitive.


## Putting Values
A default value is not required for a key, but it’s useful in the event that a key hasn't been set via
config file, environment variable.

Examples:

```go
config.SetDefault("key2", "value2")
config.SetDefault("key3", map[string]string{"tag": "tags", "category": "categories"})
```

### Reading Config Files
ConfigMan supports JSON, TOML, YAML, HCL, INI, envfile files.

Here is an example of how to use ConfigMan to search for and read a configuration file.

```go
config.LoadFile("config", "yaml", "./configs") // name of config file (without extension), config file extension and config.LoadFile("config", "yaml", "./configs") // name of config file (without extension), config file extension and path to look for the config file in.
```

`LoadFile` raise a panic when no config file is found.

### Watching and re-reading config files

ConfigMan supports the ability to have your application live read a config file while running.

```go
config.WatchChanges()
```

### Setting Overrides

These could be from a command line flag, or from your own application logic.

```go
config.Set("OldKey", "NewValue")
```

## Getting Values From ConfigMan

In ConfigMan, there are a few ways to get a value depending on the value’s type.
The following functions and methods exist:

 * `IsSet(key string) : bool`
 * `Get(key string) : interface{}`
 * `GetString(key string) : string`
 * `GetBool(key string) : bool`
 * `GetInt(key string) : int`
 * `GetInt32(key string) : int32`
 * `GetInt64(key string) : int64`
 * `GetUint(key string) : uint`
 * `GetUint32(key string) : uint32`
 * `GetUint64(key string) : uint64`
 * `GetFloat64(key string) : float64`
 * `GetTime(key string) : time.Time`
 * `GetDuration(key string) : time.Duration`
 * `GetIntSlice(key string) : []int`
 * `GetStringSlice(key string) : []string`
 * `GetStringMap(key string) : map[string]interface{}`
 * `GetStringMapString(key string) : map[string]string`
 * `GetStringMapStringSlice(key string) : map[string][]string`
 * `GetSizeInBytes(key string) : uint`
 * `GetAll() : map[string]interface{}`
 * `GetNested() : *Config`

One important thing to recognize is that each Get function will return a zero
value if it’s not found. To check if a given key exists, the `IsSet()` method
has been provided.

### Accessing nested keys

The accessor methods also accept formatted paths to deeply nested keys. For
example, if the following JSON file is loaded:

```json
{
    "log": {
        "server": {
            "host": "127.0.0.1",
            "port": 4682
        }
    }
}

```

ConfigMan can access a nested field by passing a `.` delimited path of keys:

```go
config.GetString("log.server.port") // (returns "4682")
```

ConfigMan can access array indices by using numbers in the path. For example:

```json
{
    "host": {
        "address": "localhost",
        "ports": [
            5799,
            6029
        ]
    },
}

config.GetInt("host.ports.1") // returns 6029

```

Lastly, if there exists a key that matches the delimited key path, its value
will be returned instead. E.g.

```json
{
    "datastore.metric.host": "0.0.0.0",
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}

config.GetString("datastore.metric.host") // returns "0.0.0.0"
```

### Extracting a sub-tree

When developing reusable modules, it's often useful to extract a subset of the configuration
and pass it to a module. This way the module can be instantiated more than once, with different configurations.

For example, an application might use multiple different cache stores for different purposes:

```yaml
cache:
  cache1:
    max-items: 100
    item-size: 64
  cache2:
    max-items: 200
    item-size: 80
```

We could pass the cache name to a module (eg. `NewCache("cache1")`),
but it would require weird concatenation for accessing config keys and would be less separated from the global config.

So instead of doing that let's pass a ConfigMan instance to the constructor that represents a subset of the configuration:

```go
cache1Config := config.GetNested("cache.cache1")
if cache1Config == nil { // GetNested returns nil if the key cannot be found
    panic("cache configuration not found")
}

cache1 := NewCache(cache1Config)
```

**Note:** Always check the return value of `GetNested`. It returns `nil` if a key cannot be found.

Internally, the `NewCache` function can address `max-items` and `item-size` keys directly:

```go
func NewCache(c *Config) *Cache {
    return &Cache{
        MaxItems: c.GetInt("max-items"),
        ItemSize: c.GetInt("item-size"),
    }
}
```

The resulting code is easy to test, since it's decoupled from the main Config structure,
and easier to reuse (for the same reason).
