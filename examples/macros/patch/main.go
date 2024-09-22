package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/FlowingSPDG/kairos-go"
)

var (
	ip       = os.Getenv("KAIROS_IP")
	port     = os.Getenv("KAIROS_PORT")
	user     = os.Getenv("KAIROS_USER")
	password = os.Getenv("KAIROS_PASSWORD")
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	kc := kairos.NewKairosRestClient(ip, port, user, password)

	macros, err := kc.GetMacros(ctx)
	if err != nil {
		panic(err)
	}

	macro := macros[0]
	if err := kc.PatchMacro(ctx, macro.UUID, "play"); err != nil {
		panic(err)
	}
}
