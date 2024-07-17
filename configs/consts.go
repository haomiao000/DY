package configs

const (
	DBUser     = "root"
	DBPassword = "LZXlzx251866"
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

	VideoURL = "http://192.168.1.33/assets/public/"
)

var MySecret = []byte("this is a very complex secret")
