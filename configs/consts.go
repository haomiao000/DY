package configs

const (
	DBUser     = "root"
	DBPassword = "11111111"
	DBIP       = "127.0.0.1"
	DBPort     = "3306"
	DBName     = "DY"

	IsLike   = true
	Like     = 1
	IsUnLike = false
	UnLike   = 2

	Minus_like = -1
	Plus_like = 1

	AddComment    = 1
	DeleteComment = 2

	Minus_comment = -1
	Plus_comment = 1

	UserSnowflakeNode    = 1
	NacosSnowflakeNode   = 2
	CommentSnowFlakeNode = 3
	VideoSnowFlakeNode   = 4
	MinioSnowFlakeNode   = 5

	//action
	Follow   = 1
	UnFollow = 2
	//status
	IsFollow  = true
	NotFollow = false

	If_Delete_All_Tables_Startup = false

	VideoURL = "http://localhost:8080/assets/public/"
)

var MySecret = []byte("this is a very complex secret")
