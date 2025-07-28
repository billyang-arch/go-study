package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

const (
	maxClients = 1000 // 最大客户端数量
	maxNicklen = 32
)

var serverPort = flag.Int("p", 8972, "server port")

type Client struct {
	conn net.Conn
	nick string // 昵称
}

type ChatState struct {
	listener net.Listener

	clientsLock sync.RWMutex
	clients     map[net.Conn]*Client // 存储连接的客户端
	numClients  int                  // 当前连接的客户端数量
}

var chatState = &ChatState{
	clients: make(map[net.Conn]*Client),
}

func initChat() {
	var err error
	chatState.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", *serverPort))
	if err != nil {
		fmt.Println("listen failed:", err)
		os.Exit(1)
	}
}

func handleClient(client *Client) {
	// 发送欢迎信息
	welcomMsg := "Welcome Simple Chat! Use /nick to change nick name.\n"
	client.conn.Write([]byte(welcomMsg))

	buf := make([]byte, 256)
	for {
		n, err := client.conn.Read(buf)
		if err != nil {
			fmt.Printf("client left: %s\n", client.conn.RemoteAddr())
			chatState.clientsLock.Lock()
			delete(chatState.clients, client.conn)
			chatState.numClients--
			chatState.clientsLock.Unlock()
			return
		}

		msg := string(buf[:n])
		msg = strings.TrimSpace(msg)
		if msg[0] == '/' {
			//处理命令
			parts := strings.SplitN(msg, " ", 2)
			cmd := parts[0]
			if cmd == "/nick" && len(parts) > 1 {
				client.nick = parts[1]
			}
			continue
		}

		fmt.Printf("%s: %s\n", client.nick, msg)

		//将消息转发给其他客户端
		chatState.clientsLock.RLock()
		for conn, cli := range chatState.clients {
			if cli != client {
				conn.Write([]byte(client.nick + ":" + msg))
			}
		}
		chatState.clientsLock.RUnlock()
	}
}

func main() {
	flag.Parse()
	initChat()
	fmt.Printf("Chat server started on port %d\n", *serverPort)

	for {
		conn, err := chatState.listener.Accept()
		if err != nil {
			fmt.Println("accept failed:", err)
			continue
		}

		if chatState.numClients >= maxClients {
			conn.Write([]byte("Server is full. Try again later.\n"))
			conn.Close()
			continue
		}

		client := &Client{conn: conn}
		client.nick = fmt.Sprintf("user%d", conn.RemoteAddr().(*net.TCPAddr).Port)
		chatState.clientsLock.Lock()
		chatState.clients[conn] = client
		chatState.numClients++
		chatState.clientsLock.Unlock()
		go handleClient(client)
		fmt.Printf("new client: %s\n", conn.RemoteAddr())
	}
}
