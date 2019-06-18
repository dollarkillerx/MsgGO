/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-1
* Time: 上午10:40
* */
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sync"
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/chat/:id", middlewareAuth(chat))

	return router
}

func main() {
	router := RegisterRouter()

	fmt.Println("server run http://127.0.0.1:9001 ...")
	if err := http.ListenAndServe(":9001", router); err != nil {
		panic(err.Error())
	}
}

func middlewareAuth(route httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token := r.Header.Get("token")
		// 当用户校验通过设置用户的value
		if token == "token" {
			r.ParseForm()
			r.PostForm.Set("userid", "dollarkiller")
			route(w, r, p)
			return
		}
		w.WriteHeader(401)
		w.Write([]byte("401"))
	})
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clientMap sync.Map
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
}

func chat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	userId := r.PostForm.Get("userid")
	conn, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Println(e.Error())
		return
	}
	// new Node
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
	}

	// userid 和 node 绑定
	clientMap.Store(userId, node)

	go sendproc(node)
	go recvproc(node)

}

// 发送协程
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

// 接收协程
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}

		fmt.Printf("%s \n", data)
	}
}

// 发送消息
func sendMsg(userid int64, msg []byte) {
	node, ok := clientMap.Load(userid)
	if ok {
		i, ok := node.(*Node)
		if ok {
			i.DataQueue <- msg
			return
		}
	}
}
