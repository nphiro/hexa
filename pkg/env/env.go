package env

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v11"
)

func Parse[T any](obj T) (err error) {
	wd, err := os.Getwd()
	if err != nil {
		return ErrGetCurrWorkingDir
	}

	cfg, err := getEnvConfig(wd)
	if err != nil {
		return err
	}

	envFile := filepath.Join(wd, cfg.File)
	if err := loadEnvFile(envFile); err != nil {
		slog.Error("failed to load env file", slog.String("filename", envFile), slog.Any("error", err))
		return err
	}

	if err := env.Parse(obj); err != nil {
		return ErrParseEnv
	}
	return nil
}
