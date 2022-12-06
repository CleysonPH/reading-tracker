package config

var (
	Port string = "8000"
	Env  string = "dev"
	Dsn  string = "root:root@tcp(localhost:3306)/reading-tracker?parseTime=true"
	Host string = ""
)

func Addr() string {
	return Host + ":" + Port
}
