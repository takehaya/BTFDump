package main

import(
	"log"
	"os"
	"github.com/takehaya/btfdump/internal"
)

func main() {
	app := internal.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("%+v", err)
	}
}
