package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        // シンプルにするためにエラーを無視
        conn, _ := upgrader.Upgrade(w, r, nil)

        for {
            // ブラウザからのメッセージを読み込む
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }

            // メッセージをコンソールに表示する
            fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

            // 書き込んだメッセージをブラウザに戻す
            if err = conn.WriteMessage(msgType, msg); err != nil {
                return
            }
        }
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "websockets.html")
    })

    http.ListenAndServe(":80", nil)
}
