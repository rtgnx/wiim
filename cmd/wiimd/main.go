package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rtgnx/wiim/internal/api"
	"github.com/rtgnx/wiim/internal/res"
	"github.com/rtgnx/wiim/internal/wiim"
	"golang.org/x/net/websocket"
)

const (
	addr = "127.0.0.1:9090"
)

func main() {
	e := echo.New()

	dev, err := wiim.Connect(context.Background(), "10.1.1.159")
	if err != nil {
		log.Fatal(err)
	}

	assetHandler := http.FileServer(res.GetFileSystem())
	e.GET("/", echo.WrapHandler(assetHandler))

	e.GET("/ws", func(c echo.Context) error {
		websocket.Handler(api.MediaInfoUpdater(dev)).ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Debug(e.Start(addr))
}

// func main() {

// 	dev, err := wiim.Connect(context.Background(), "10.1.1.159")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	metadataCh := make(chan string)

// 	go dev.Subscribe(context.Background(), time.Second, metadataCh)

// 	for {
// 		log.Print(<-metadataCh)
// 	}

// }
