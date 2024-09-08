package flags

import (
	"fmt"
	"os"
)

var (
	serverAddress    string
	postgresConn     string
	postgresJdbcUrl  string
	postgresUsername string
	postgresPassword string
	postgresHost     string
	postgresPort     string
	postgresDBName   string
)

func init() {
	serverAddress = os.Getenv("SERVER_ADDRESS")
	postgresConn = os.Getenv("POSTGRES_CONN")
	postgresJdbcUrl = os.Getenv("POSTGRES_JDBC_URL")
	postgresUsername = os.Getenv("POSTGRES_USERNAME")
	postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	postgresHost = os.Getenv("POSTGRES_HOST")
	postgresPort = os.Getenv("POSTGRES_PORT")
	postgresDBName = os.Getenv("POSTGRES_DATABASE")
}

func GetPostgresConn() string {
	if postgresConn != "" {
		return postgresDBName
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		postgresUsername,
		postgresPassword,
		postgresHost,
		postgresPort,
		postgresDBName)
}

func GetPostgresJdbcUrl() string {
	if postgresJdbcUrl != "" {
		return postgresJdbcUrl
	}

	return fmt.Sprintf("jdbc:postgresql://%s:%s/%s",
		postgresHost,
		postgresPort,
		postgresDBName)
}

func GetServerEndPoint() string {
	return serverAddress
}
