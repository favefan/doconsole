package websocket

import (
	"fmt"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"gitee.com/favefan/doconsole/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	//ReadBufferSize:  1024,
	//WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func TerminalHandle(c *gin.Context, execID string) {
	appG := app.Gin{C: c}

	// websocket握手
	websocketConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer websocketConn.Close()

	// 执行exec，获取到容器终端的连接
	hijackedResponse, err := service.ContainerExecAttach(execID)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, err)
		return
	}
	defer hijackedResponse.Close()
	// 退出进程
	defer func() {
		hijackedResponse.Conn.Write([]byte("exit\r"))
	}()

	// 转发输入/输出至websocket
	go func() {
		wsWriterCopy(hijackedResponse.Conn, websocketConn)
	}()
	wsReaderCopy(websocketConn, hijackedResponse.Conn)
}

func wsWriterCopy(reader io.Reader, writer *websocket.Conn) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}
}

func wsReaderCopy(reader *websocket.Conn, writer io.Writer) {
	for {
		messageType, p, err := reader.ReadMessage()
		if err != nil {
			return
		}
		if messageType == websocket.TextMessage {
			writer.Write(p)
		}
	}
}