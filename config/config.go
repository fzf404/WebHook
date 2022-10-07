package config

import (
	"io/ioutil"
	"log"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

// Config Struct
type Config struct {
	Host string `yaml:"host" validate:"required,hostname_port"`
	Hook Hook   `yaml:"hook" validate:"required,dive"`
	Mail Mail   `yaml:"mail" validate:"omitempty"`
}

// Hooks Struct
type Hook map[string]struct {
	Url    string            `yaml:"url" validate:"omitempty,uri"`
	Secret string            `yaml:"secret" validate:"omitempty"`
	Run    map[string]string `yaml:"run" validate:"omitempty"`
}

// Mail Struct
type Mail struct {
	Enable bool `yaml:"enable" validate:"omitempty"`

	Host     string `yaml:"host" validate:"required_if=Enable true,omitempty,hostname"`
	Port     int    `yaml:"port" validate:"required_if=Enable true,omitempty,min=1,max=65535"`
	Password string `yaml:"password" validate:"required_if=Enable true,omitempty"`

	From string   `yaml:"from" validate:"required_if=Enable true,omitempty,email"`
	To   []string `yaml:"to" validate:"required_if=Enable true,omitempty,dive,email"`
}

// Global Config
var Cfg Config

// Init Config File
func InitConfig() {

	// Read Config File
	file, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal("🔴 read config error: \n", err.Error())
	}

	// Unmarshal Config File
	if err := yaml.Unmarshal(file, &Cfg); err != nil {
		log.Fatal("🔴 parse config error: \n", err.Error())
	}

	// Complete Config Struct
	for k, v := range Cfg.Hook {
		// Complete Url
		if v.Url == "" {
			if entry, ok := Cfg.Hook[k]; ok {
				entry.Url = "/" + k
				Cfg.Hook[k] = entry
			}
		}
		// Complete Secret
		if v.Secret == "" {
			if entry, ok := Cfg.Hook[k]; ok {
				entry.Secret = k
				Cfg.Hook[k] = entry
			}
		}
		// Complete Run
		if v.Run == nil {
			if entry, ok := Cfg.Hook[k]; ok {
				entry.Run = map[string]string{
					"push": "./script/" + k + ".sh",
				}
				Cfg.Hook[k] = entry
			}
		}

	}

	// Validate Config File
	if err := validator.New().Struct(Cfg); err != nil {
		log.Fatal("🔴 validate config error: \n", err.Error())
	}
}
