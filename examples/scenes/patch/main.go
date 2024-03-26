package main

import (
	"fmt"
	"os"
	"time"

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

	// get Main(PGM) Scene
	pgmScene := scenes[0]
	pgmLayer := pgmScene.Layers[0] // Main Layer[0] as a Background
	pgmSources := pgmLayer.Sources // Main Layer[0] Sources(scenes)

	fmt.Println("Switching PGM")
	for i, source := range pgmSources {
		// get preview source
		prv := pgmSources[i]
		if i != len(pgmSources) && i != 0 {
			prv = pgmSources[i-1]
		}

		// patch scene
		if err := kc.PatchScene(pgmScene.UUID, pgmLayer.UUID, prv, source, pgmSources); err != nil {
			panic(err)
		}

		// sleep 1 second
		time.Sleep(1 * time.Second)
	}
}
