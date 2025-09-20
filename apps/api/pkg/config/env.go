package config

import (
	"os"
)

type EnvConfig struct {
	Mode           string
	FrontendDomain string
	Domain         string
}

type CorsConfig struct {
	AllowOrigins string
}

type PsqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type SqliteConfig struct {
	Path string
}

type DbConfig struct {
	Type     string
	Sqlite SqliteConfig
	Psql   PsqlConfig
}

type JwtConfig struct {
	Secret string
}

type Config struct {
	Db   DbConfig
	Jwt  JwtConfig
	Env  EnvConfig
	Cors CorsConfig
}

func LoadConfig() *Config {
	var (
		sqlitePath   string
		psqlHost     string
		psqlPort     string
		psqlUser     string
		psqlPassword string
		psqlName     string
	)
	dbType := os.Getenv("DB_TYPE")
	sqlitePath = os.Getenv("SQLITE_PATH")
	psqlHost = os.Getenv("PSQL_HOST")
	psqlPort = os.Getenv("PSQL_PORT")
	psqlUser = os.Getenv("PSQL_USER")
	psqlPassword = os.Getenv("PSQL_PASSWORD")
	psqlName = os.Getenv("PSQL_NAME")
	mode := os.Getenv("MODE")
	frontendDomain := os.Getenv("FRONTEND_DOMAIN")
	domain := os.Getenv("DOMAIN")
	if mode == "" {
		mode = "dev"
	}
	allow_origins := os.Getenv("ALLOW_ORIGINS")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret"
	}

	return &Config{
		Db: DbConfig{
			Sqlite: SqliteConfig{
				Path: sqlitePath,
			},
			Psql: PsqlConfig{
				Host:     psqlHost,
				Port:     psqlPort,
				User:     psqlUser,
				Password: psqlPassword,
				DbName:   psqlName,
			},
			Type: dbType,
		},
		Jwt: JwtConfig{
			Secret: jwtSecret,
		},
		Env: EnvConfig{
			Mode:           mode,
			FrontendDomain: frontendDomain,
			Domain:         domain,
		},
		Cors: CorsConfig{
			AllowOrigins: allow_origins,
		},
	}
}
