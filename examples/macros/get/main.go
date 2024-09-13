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

	macros, err := kc.GetMacros(ctx)
	if err != nil {
		panic(err)
	}

	for i, macro := range macros {
		fmt.Printf("Got Macro[%d] - %s\n", i, macro.Name)
	}
}
