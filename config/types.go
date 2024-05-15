package config

type AppConfig struct {
	Postgres Postgres
	Fiber Fiber
}

type Fiber struct {
	Port string
}

type Postgres struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}