package config

import (
	"encoding/json"
	"log"
	"os"
)

type SqlConfig struct {
	Driver       string `json:"driver"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database"`
	Encoding     string `json:"encoding"`
	Host         string `json:"host"`
	Port         string `json:"port"`
}

type ErrorLoggerConfig struct {
	OutputPath string `json:"output_path"`
}
type RequestLoggerConfig struct {
	OutputPath string `json:"output_path"`
}

type Config struct {
	Env        string              `json:"env"`
	Sql        SqlConfig           `json:"sql"`
	ErrorLog   ErrorLoggerConfig   `json:"error_logger"`
	RequestLog RequestLoggerConfig `json:"request_logger"`
	Port       int                 `json:"port"`
}

// New creates a new config by reading a json file that matches the types above
func Load(path string) (Config, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	cfg := Config{}

	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}
