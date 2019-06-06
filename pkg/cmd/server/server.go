package server

// PostgresConfig contains options for the Postgres connections pool
type PostgresConfig struct {
	DSN string
}

// Config is configuration for the Server
type Config struct {
	Postgres PostgresConfig
}
