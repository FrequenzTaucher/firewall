package models

type Configuration struct {
	Server      ServerConfiguration
	Database    DatabaseConfiguration
	Auth        AuthConfiguration
	Environment EnvironmentConfiguration
}

type DatabaseConfiguration struct {
	ConnectionUri string
}

type ServerConfiguration struct {
	Port int
}

type AuthConfiguration struct {
	ApiKey string
}

type EnvironmentConfiguration struct {
	Mode string
}
