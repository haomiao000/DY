package configs
var (
    MySqlDSN       = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
    UserDBUser     = "root"
    UserDBPassword = "11111111"
    UserDBIP       = "127.0.0.1"
    UserDBPort     = 3306
    UserDBName     = "DY"

	RedisIP        = "127.0.0.1"
    RedisPort      = ":50051"
	UserServerAddr = "127.0.0.1:8081"
)
