package types

type Config struct {
	Port           string
	JWT_SECRET_KEY string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_SSL_MODE    string
}