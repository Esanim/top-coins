package main // import "github.com/esanim/top-coins"
import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/esanim/top-coins/pkg/app"
	"github.com/esanim/top-coins/pkg/coins"
	midl "github.com/esanim/top-coins/pkg/middleware"
)

func main() {
	// Configuration
	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":8080"
	}
	// Initialize
	app, err := app.NewApp(addr)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := recover(); err != nil {
			app.Logger.
				Fatal().
				Msgf("Exception: %s", err)
		}
	}()

	// GET /
	coinsHandler := coins.NewCoinsHandler(app)
	app.Handler.HandleFunc("/", midl.WrapMiddleware(coinsHandler.GetCoins))

	// Process
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go app.Serve()

	// Stop
	<-stop
	app.Shutdown()
}
