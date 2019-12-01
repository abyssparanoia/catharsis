package psql

// Config ... psql config
type Config struct {
	ConnctionName string
	User          string
	Password      string
	Database      string
}

// NewConfig ... get new config
func NewConfig(connectionName string, user string, password string, database string) *Config {
	return &Config{
		ConnctionName: connectionName,
		User:          user,
		Password:      password,
		Database:      database,
	}
}
