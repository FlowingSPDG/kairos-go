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

	macros, err := kc.GetMacros()
	if err != nil {
		panic(err)
	}

	for i, macro := range macros {
		fmt.Printf("Got Macro[%d] - %s\n", i, macro.Name)
	}
}
