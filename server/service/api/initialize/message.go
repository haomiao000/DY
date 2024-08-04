package initialize

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
	model "github.com/haomiao000/DY/server/service/api/model"
)

// chatConnMap 是一个并发安全的映射，用于存储用户之间的连接。
var chatConnMap = sync.Map{}

// RunMessageServer 启动消息服务器，监听指定地址和端口接收客户端连接。
func RunMessageServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}

	// 持续接收新的客户端连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept conn failed: %v\n", err)
			continue
		}

		// 启动一个 goroutine 处理每个连接
		go process(conn)
	}
}

// process 处理单个客户端连接的函数。
func process(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		// 读取客户端发送的数据
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break // 客户端关闭连接
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue
		}

		// 解析收到的消息
		var event = model.MessageSendEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		fmt.Printf("Receive Message：%+v\n", event)

		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
		if len(event.MsgContent) == 0 {
			// 如果消息内容为空，则将连接存储在 chatConnMap 中
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			// 如果接收者不在线，则打印消息并继续等待下一条消息
			fmt.Printf("User %d offline\n", event.ToUserId)
			continue
		}

		// 准备要推送的消息事件
		pushEvent := model.MessagePushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent)

		// 将消息推送给接收者的连接
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			fmt.Printf("Push message failed: %v\n", err)
		}
	}
}