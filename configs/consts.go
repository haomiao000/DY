package configs

const (
	DBUser     = "root"
	DBPassword = "11111111"
	DBIP       = "127.0.0.1"
	DBPort     = "3306"
	DBName     = "DY"

	IsLike   = 1
	Like     = 1
	IsUnLike = 2
	UnLike   = 2

	AddComment    = 1
	DeleteComment = 2

	UserSnowflakeNode    = 1
	NacosSnowflakeNode   = 2
	CommentSnowFlakeNode = 3
	VideoSnowFlakeNode   = 4
	MinioSnowFlakeNode   = 5

	VideoURL = "http://localhost:8080/assets/public/"
)

var MySecret = []byte("this is a very complex secret")
