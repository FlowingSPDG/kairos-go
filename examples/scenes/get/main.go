package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/FlowingSPDG/kairos-go"
)

var (
	user     = os.Getenv("KAIROS_USER")
	password = os.Getenv("KAIROS_PASSWORD")
	ip       = os.Getenv("KAIROS_IP")
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	kc := kairos.NewKairosRestClient(ip, user, password)

	scenes, err := kc.GetScenes(ctx)
	if err != nil {
		panic(err)
	}

	for i, scene := range scenes {
		fmt.Printf("Got Scene[%d] - %s\n", i, scene.Name)
		for j, layer := range scene.Layers {
			fmt.Printf("	- Layer[%d] %s\n", j, layer.Name)
		}
	}
}
