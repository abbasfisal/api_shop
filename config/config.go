package config

type AppConfig struct {
	Name  string
	Env   string
	Debug bool
	Port  int
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type Config struct {
	App      AppConfig
	Database Mysql
}
