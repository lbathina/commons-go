package conf

import (
	"fmt"
	"github.com/caarlos0/env"
)

//Registry represents type of used service discovery server
type Registry string

const (
	//Consul service discovery
	Consul Registry = "consul"
	//Eureka service discovery
	Eureka Registry = "eureka"
)

//ServerConfig represents Main service configuration
type ServerConfig struct {
	Hostname string `env:"HOSTNAME" envDefault:"localhost"`
	Port     int    `env:"RP_SERVER_PORT" envDefault:"8080"`
}

//EurekaConfig represents Eureka Discovery service configuration
type EurekaConfig struct {
	URL          string `env:"RP_EUREKA_URL" envDefault:"http://localhost:8761/eureka"`
	PollInterval int    `env:"RP_EUREKA_POLL_INTERVAL" envDefault:"5"`
}

//ConsulConfig represents Consul Discovery service configuration
type ConsulConfig struct {
	Address      string   `env:"RP_CONSUL_ADDRESS" envDefault:"registry:8500"`
	Scheme       string   `env:"RP_CONSUL_SCHEME" envDefault:"http"`
	Token        string   `env:"RP_CONSUL_TOKEN"`
	PollInterval int      `env:"RP_CONSUL_POLL_INTERVAL" envDefault:"5"`
	Tags         []string `env:"RP_CONSUL_TAGS"`
}

//AddTags parses tags to string array. Extremely slow implementation - simplicity over speed
func (c *ConsulConfig) AddTags(tags ...string) {
	c.Tags = append(c.Tags, tags...)
}

//RpConfig represents Composite of all app configs
type RpConfig struct {
	AppName  string   `env:"RP_APP_NAME" envDefault:"goRP"`
	Registry Registry `env:"RP_REGISTRY" envDefault:"consul"`
	Server   *ServerConfig
	Eureka   *EurekaConfig
	Consul   *ConsulConfig
}

//LoadConfig loads configuration from provided file and serializes it into RpConfig struct
func LoadConfig(cfg interface{}) error {
	err := env.Parse(cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}

	return nil
}

//EmptyConfig creates empty config
func EmptyConfig() *RpConfig {
	return &RpConfig{
		Consul: &ConsulConfig{},
		Eureka: &EurekaConfig{},
		Server: &ServerConfig{},
	}
}
