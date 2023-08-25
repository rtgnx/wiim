package wiim

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/av1"
)

const WiiMPort = 49152

type WiiM struct {
	target    *url.URL
	transport *av1.AVTransport1
	device    goupnp.Device
}

func Connect(ctx context.Context, addr string) (*WiiM, error) {
	wiim := &WiiM{
		target: &url.URL{
			Host:   fmt.Sprintf("%s:%d", addr, WiiMPort),
			Scheme: "http",
			Path:   "/description.xml",
		},
	}

	transports, err := av1.NewAVTransport1ClientsByURLCtx(ctx, wiim.target)
	if err != nil {
		return nil, err
	}

	services, _ := goupnp.NewServiceClientsByURL(wiim.target, "*")

	for _, svc := range services {
		log.Print(svc.Location)
	}

	if len(transports) > 0 {
		wiim.transport = transports[0]
		return wiim, nil
	}
	return nil, errors.New("unable t o connect to device")
}

func (wiim *WiiM) GetMediaInfo(ctx context.Context) (string, error) {

	_, _, _, metadata, _, _, _, _, _, err := wiim.transport.GetMediaInfoCtx(ctx, 0)
	return metadata, err
}

func (wiim *WiiM) Subscribe(ctx context.Context, interval time.Duration, ch chan string) error {
	for {
		select {
		case <-time.After(interval):
			metadata, err := wiim.GetMediaInfo(ctx)
			if err != nil {
				log.Print(err)
				continue
			}

			ch <- metadata
		case <-ctx.Done():
			return nil
		}
	}
}
