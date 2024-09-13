package main

import (
	"context"
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

	macro := macros[0]
	if err := kc.PatchMacro(macro.UUID, "play"); err != nil {
		panic(err)
	}
}
