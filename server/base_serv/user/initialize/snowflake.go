package initialize

import (
	"github.com/bwmarrin/snowflake"
	configs "github.com/haomiao000/DY/server/base_serv/user/configs"
	globalconfigs "github.com/haomiao000/DY/server/common/configs"
)

func InitSnowFlake() {
	var err error
	configs.UserSnowFlakeNode, err = snowflake.NewNode(globalconfigs.UserSnowflakeNode)
	if err != nil {
		panic(err)
	}
}
