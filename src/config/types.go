package config

type Configuration struct {
	ServerPort       uint32 `envconfig:"server_port" default:"1337"`
	JWTEncryptionKey string `envconfig:"jwt_encryption_key" default:"G3RFhwr4WMWy2XLDi73V5ir8yFTSthjT"`

	DatabaseHost          string `envconfig:"db_host" default:"127.0.0.1"`
	DatabasePort          uint32 `envconfig:"db_port" default:"3306"`
	DatabaseName          string `envconfig:"db_name" default:"kubesmith"`
	DatabaseUser          string `envconfig:"db_user" default:"root"`
	DatabasePass          string `envconfig:"db_pass" default:"root"`
	DatabaseDebug         bool   `envconfig:"db_debug" default:"false"`
	DatabaseMaxOpen       int    `envconfig:"db_max_open" default:"100"`
	DatabaseMaxIdle       int    `envconfig:"db_max_idle" default:"10"`
	DatabaseEncryptionKey string `envconfig:"db_encryption_key" default:"Leq2yHQPRdp8NKWcNrfpZLJWAfWbDzih"`
}
