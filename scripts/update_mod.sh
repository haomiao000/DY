z#!/bin/bash

directories=(
    "../comm/discovery"
    "../comm/redis"
    "../comm/trace"
    "../server/gateway_serv/gateway"
    "../server/base_serv/interact"
    "../server/base_serv/relation"
    "../server/base_serv/user"
    "../server/base_serv/video"
    "../server/common"
    "../server/redis_svr"
)

for dir in "${directories[@]}"; do
    echo "进入目录: $dir"
    cd "$dir" || exit
    go mod tidy
    cd - || exit
done

echo "所有目录的 go mod tidy 操作完成。"