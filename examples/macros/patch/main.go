package main

import (
	"os"

	"github.com/FlowingSPDG/kairos-go"
)

func main() {
	user := os.Getenv("KAIROS_USER")
	password := os.Getenv("KAIROS_PASSWORD")

	kc := kairos.NewKairosRestClient("192.168.10.10", user, password)

	macros, err := kc.GetMacros()
	if err != nil {
		panic(err)
	}

	macro := macros[0]
	if err := kc.PatchMacro(macro.UUID, "play"); err != nil {
		panic(err)
	}
}
