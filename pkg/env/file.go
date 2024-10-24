package env

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type envConfig struct {
	File string `json:"file"`
}

func defaultEnvConfig() *envConfig {
	return &envConfig{
		File: ".env",
	}
}

func getEnvConfig(wd string) (*envConfig, error) {
	configFile := filepath.Join(wd, "env.config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		slog.Info("env.config.json not found, using default env config")
		return defaultEnvConfig(), nil
	}
	dat, err := os.ReadFile(configFile)
	if err != nil {
		slog.Error("failed to read env config file", slog.String("filename", configFile), slog.Any("error", err))
		return nil, ErrParseConfig
	}

	cfg := defaultEnvConfig()
	if len(dat) == 0 {
		slog.Info("env.config.json is empty, using default env config")
		return cfg, nil
	}
	if err := json.Unmarshal(dat, cfg); err != nil {
		slog.Error("failed to parse env config", slog.String("filename", configFile), slog.Any("error", err))
		return nil, ErrParseConfig
	}
	return cfg, nil
}

func loadEnvFile(envFile string) error {
	dat, err := os.ReadFile(envFile)
	if err != nil {
		return ErrReadFileEnv
	}
	lines := strings.Split(string(dat), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		os.Setenv(parts[0], parts[1])
	}
	return nil
}
