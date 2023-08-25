package api

import (
	"encoding/json"
	"log"
	"time"

	"github.com/rtgnx/wiim/internal/wiim"
	"golang.org/x/net/websocket"
)

type WebSocketHandler func(*websocket.Conn)

func MediaInfoUpdater(dev *wiim.WiiM) WebSocketHandler {
	return func(conn *websocket.Conn) {
		defer conn.Close()

		log.Print("connected client: ", conn.RemoteAddr())

		metadataCh := make(chan string)

		go dev.Subscribe(conn.Request().Context(), time.Second*4, metadataCh)

		for {
			select {
			case metadata := <-metadataCh:
				mediaInfo, err := wiim.ParseMediaInfo(metadata)
				log.Print(mediaInfo.Item.Album)
				if err != nil {
					log.Println(err)
					continue
				}
				json.NewEncoder(conn).Encode(mediaInfo.MediaInfo())
			case <-conn.Request().Context().Done():
				return
			}
		}

	}
}
