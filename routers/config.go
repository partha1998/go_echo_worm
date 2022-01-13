package routers

type configDatabase struct {
	AppName  string `env:"APP_NAME" env-default:"GO_ECHO_GORM"`
	AppEnv   string `env:"APP_ENV" env-default:"DEV"`
	Port     string `env:"MY_APP_PORT" env-default:"8080"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLevel string `env:"LOG_LEVEL" env-default:"ERROR"`
}

var cfg configDatabase
