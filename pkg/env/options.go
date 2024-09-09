package env

type WriteOptions struct {
	EnvFile string
}

func NewWriteOptions() *WriteOptions {
	return &WriteOptions{
		EnvFile: ".env",
	}
}

func (o *WriteOptions) WithEnvFile(envFile string) *WriteOptions {
	o.EnvFile = envFile
	return o
}
