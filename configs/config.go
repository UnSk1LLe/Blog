package configs

type Config struct {
	App    App
	Server Server
	Db     Db
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type Db struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}
