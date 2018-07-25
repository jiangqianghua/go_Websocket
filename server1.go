/**
package main

import (
		"net/http"
		"github.com/gorilla/websocket"
		)
var(
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter , r *http.Request){
	//	w.Write([]byte("hello"))
	var(
		conn *websocket.Conn
		err error
		data []byte
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	if conn , err = upgrader.Upgrade(w,r,nil); err != nil{
		return 
	}
	// 开始长链接
	for{
		if _ , data , err = conn.ReadMessage();err != nil{
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage,data); err != nil{
			goto ERR
		}
	}

	ERR:
		conn.Close()

}

func main(){

	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe("0.0.0.0:7777",nil)
}
**/