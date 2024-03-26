package main

import (
	"fmt"
	"os"

	"github.com/FlowingSPDG/kairos-go"
)

func main() {
	user := os.Getenv("KAIROS_USER")
	password := os.Getenv("KAIROS_PASSWORD")

	kc := kairos.NewKairosRestClient("192.168.10.10", user, password)

	scenes, err := kc.GetScene()
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
