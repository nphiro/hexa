package env

import (
	"os"
	"path/filepath"
	"strings"
)

func loadEnvFile(envFile string) error {
	wd, err := os.Getwd()
	if err != nil {
		return ErrGetCurrWorkingDir
	}
	dat, err := os.ReadFile(filepath.Join(wd, envFile))
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
