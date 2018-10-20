package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/ilovelili/dongfeng/sharedlib"
)

var once sync.Once
var instance *Config

// GetConfig get config defined in config.json
func GetConfig() *Config {
	once.Do(func() {
		env := os.Getenv("DF_ENVIROMENT")
		if env == "" {
			env = "dev"
		}

		var config *Config
		var filepath string
		pwd, _ := os.Getwd()

		if flag.Lookup("test.v") == nil {
			// normal run
			filepath = path.Join(pwd, fmt.Sprintf("config.%s.json", strings.ToLower(env)))
		} else {
			// under go test
			filepath = path.Join(pwd, "testdata", "config.unit.test.json")
		}

		configFile, err := os.Open(filepath)
		defer configFile.Close()
		if err != nil {
			sharedlib.NewStdOutLogger().Panic("Failed to load config")
			return
		}

		jsonParser := json.NewDecoder(configFile)
		err = jsonParser.Decode(&config)
		if err != nil {
			sharedlib.NewStdOutLogger().Panic("Failed to load config")
			return
		}

		instance = config
	})

	return instance
}

// Auth auth0 fields
type Auth struct {
	// https://auth0.com/docs/jwks
	JWKS string `json:"jwks"`
}

// Redis redis config
type Redis struct {
	Host     string `json:"host"`
	Password string `json:"password,omitempty"`
	Size     int    `json:"maxconnectioncount"`
}

// Services external services like Mysql
type Services struct {
	Redis `json:"redis"`
}

// ServiceNames servicename config
type ServiceNames struct {
	CoreProxy  string `json:"core_proxy"`
	CoreServer string `json:"core_server"`
}

// ServiceMeta service meta data including service discovery specs
type ServiceMeta struct {
	RegistryTTL       int    `json:"registry_ttl"`
	RegistryHeartbeat int    `json:"registry_heartbeat"`
	Version           string `json:"api_version"`
}

// Config config entry
type Config struct {
	Auth         `json:"auth"`
	Services     `json:"services"`
	ServiceNames `json:"servicenames"`
	ServiceMeta  `json:"servicemeta"`
}

// GetMaxConnectionCount get redis max connection count
func (r *Redis) GetMaxConnectionCount() int {
	if r.Size == 0 {
		return 100
	}
	return r.Size
}

// GetRegistryTTL get registry ttl
func (s *ServiceMeta) GetRegistryTTL() time.Duration {
	if s.RegistryTTL == 0 {
		return 30 * time.Second
	}

	return time.Duration(s.RegistryTTL) * time.Second
}

// GetRegistryHeartbeat get registry heartbeat
func (s *ServiceMeta) GetRegistryHeartbeat() time.Duration {
	if s.RegistryHeartbeat == 0 {
		return 10 * time.Second
	}

	return time.Duration(s.RegistryHeartbeat) * time.Second
}

// GetVersion get api version
func (s *ServiceMeta) GetVersion() string {
	return s.Version
}
