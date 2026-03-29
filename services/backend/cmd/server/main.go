package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kamil-abbasi/TicTacToe.git/internal"
)

func main() {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	queue := internal.NewQueue()
	go queue.Run()

	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			w.Write([]byte("name cannot be empty"))
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(fmt.Errorf("failed to upgrade connection: %v", err))
			return
		}

		log.Printf("client %v connected", name)

		conn.SetCloseHandler(func(code int, text string) error {
			log.Printf("client %v disconnected", name)
			return nil
		})

		client := internal.NewClient(name, conn)

		go client.WritePump()
		go client.ReadPump()

		queue.Enqueue(client)
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(fmt.Errorf("failed to start http server: %v", err))
	}
}
